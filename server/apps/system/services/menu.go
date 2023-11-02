package services

import (
	"context"
	"github.com/fatih/structs"
	"github.com/lbemi/lbemi/apps/system/api/form"
	"github.com/lbemi/lbemi/apps/system/entity"
	"github.com/lbemi/lbemi/pkg/common/commService"
	"github.com/lbemi/lbemi/pkg/restfulx"
	"gorm.io/gorm"
)

type MenuGetter interface {
	Menu() IMenu
}

// IMenu 菜单操作接口
type IMenu interface {
	Create(ctx context.Context, obj *form.MenusReq) *entity.Menu
	Update(ctx context.Context, menu *form.UpdateMenusReq, menuID uint64)
	Delete(ctx context.Context, menuID uint64)
	Get(ctx context.Context, menuID uint64) *entity.Menu
	List(ctx context.Context, page, limit int, menuType []int8, isTree bool, condition *entity.Menu) *form.PageMenu

	GetByIds(ctx context.Context, menuIDs []uint64) *[]entity.Menu
	GetMenuByMenuNameUrl(context.Context, string, string) *entity.Menu
	CheckMenusIsExist(ctx context.Context, menuID uint64) bool
	UpdateStatus(ctx context.Context, menuID, status uint64)
}

type Menu struct {
	//factory commService.Interface
	db     *gorm.DB
	policy commService.IPolicy
}

func NewMenu(db *gorm.DB, policy commService.IPolicy) IMenu {
	return &Menu{
		db:     db,
		policy: policy,
	}
}

func (m *Menu) Create(ctx context.Context, obj *form.MenusReq) *entity.Menu {
	isUnique := m.CheckUniqueMenuNameUrl(ctx, obj.Path, obj.Method)
	restfulx.ErrNotTrue(isUnique, restfulx.ResourceExist)

	newMenu := &entity.Menu{
		Name:     obj.Name,
		Memo:     obj.Memo,
		ParentID: obj.ParentID,
		Status:   obj.Status,
		Path:     obj.Path,
		Group:    obj.Group,
		Meta: entity.Meta{
			Icon:        obj.Meta.Icon,
			Title:       obj.Meta.Title,
			IsLink:      obj.Meta.IsLink,
			IsAffix:     obj.Meta.IsAffix,
			IsHide:      obj.Meta.IsHide,
			IsIframe:    obj.Meta.IsIframe,
			IsKeepAlive: obj.Meta.IsKeepAlive,
			IsK8s:       obj.Meta.IsK8s,
		},
		Redirect:  obj.Redirect,
		Component: obj.Component,
		Sequence:  obj.Sequence,
		MenuType:  obj.MenuType,
		Method:    obj.Method,
		Code:      obj.Code,
	}
	err := m.db.Create(obj).Error
	restfulx.ErrNotNilDebug(err, restfulx.OperatorErr)
	return newMenu
}

func (m *Menu) Update(ctx context.Context, newMenu *form.UpdateMenusReq, menuID uint64) {
	restfulx.ErrNotTrue(m.CheckMenusIsExist(ctx, menuID), restfulx.ResourceExist)

	oldMenu := m.Get(ctx, menuID)

	objMap := structs.Map(newMenu)
	delete(objMap, "Meta")
	metaMap := structs.Map(newMenu.Meta)
	for k, v := range metaMap {
		objMap[k] = v
	}

	err := m.db.Model(&entity.Menu{}).Where("id = ?  ", menuID).Updates(objMap).Error
	restfulx.ErrNotNilDebug(err, restfulx.OperatorErr)

	if oldMenu.Path != newMenu.Path || oldMenu.Method != newMenu.Method {
		err = m.policy.UpdatePermissions(ctx, oldMenu.Path, oldMenu.Method, newMenu.Path, newMenu.Method)
		if err != nil {
			restfulx.ErrNotNilDebug(err, restfulx.OperatorErr)
		}
	}
}

func (m *Menu) Delete(ctx context.Context, menuID uint64) {
	restfulx.ErrNotTrue(m.CheckMenusIsExist(ctx, menuID), restfulx.ResourceExist)
	menuInfo := m.Get(ctx, menuID)

	// 清除rules
	err := m.policy.DeleteRolePermission(ctx, menuInfo.Path, menuInfo.Method)
	restfulx.ErrNotNilDebug(err, restfulx.OperatorErr)

	// 清除menus
	tx := m.db.Begin()
	defer func() {
		if err := recover(); err != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		tx.Rollback()
		restfulx.ErrNotNilDebug(err, restfulx.OperatorErr)
	}

	// 清除role_menus
	if err := tx.Where("menuID= ?", menuID).Delete(&entity.RoleMenu{}).Error; err != nil {
		tx.Rollback()
		restfulx.ErrNotNilDebug(err, restfulx.OperatorErr)
	}

	// 清除menus
	if err := tx.Where("id =  ?", menuID).
		Or("parentID = ?", menuID).
		Delete(&entity.Menu{}).Error; err != nil {
		tx.Rollback()
		restfulx.ErrNotNilDebug(err, restfulx.OperatorErr)
	}

	restfulx.ErrNotNilDebug(err, restfulx.OperatorErr)
}

func (m *Menu) Get(ctx context.Context, menuID uint64) *entity.Menu {
	menuResult := &entity.Menu{}
	err := m.db.Model(&entity.Menu{}).Where("id = ?", menuID).First(&menuResult).Error
	restfulx.ErrNotNilDebug(err, restfulx.OperatorErr)

	return menuResult
}

func (m *Menu) List(ctx context.Context, page, limit int, menuType []int8, isTree bool, condition *entity.Menu) *form.PageMenu {
	db := m.db
	res := &form.PageMenu{}
	var (
		menuList []entity.Menu
		total    int64
		err      error
	)
	if condition.Memo != "" {
		db = db.Where("memo like ?", "%"+condition.Memo+"%")
	}
	if condition.Group != "" {
		db = db.Where("`group` like ?", "%"+condition.Group+"%")
	}
	if condition.Status != 0 {
		db = db.Where("status = ?", condition.Status)
	}

	// 全量查询
	if page == 0 && limit == 0 {
		err = db.Order("sequence DESC").Where("menuType in (?)", menuType).Find(&menuList).Error
		restfulx.ErrNotNilDebug(err, restfulx.OperatorErr)

		err = db.Model(&entity.Menu{}).Count(&total).Error
		restfulx.ErrNotNilDebug(err, restfulx.OperatorErr)
		if isTree {
			treeMenu := GetTreeMenus(menuList, 0)
			res = &form.PageMenu{
				Menus: treeMenu,
				Total: total,
			}
			return res
		}
		res = &form.PageMenu{
			Menus: menuList,
			Total: total,
		}
		return res
	}

	//分页数据
	err = db.Order("sequence DESC").Where("menuType in (?)", menuType).Limit(limit).Offset((page - 1) * limit).
		Find(&menuList).Error
	restfulx.ErrNotNilDebug(err, restfulx.OperatorErr)

	//查询 total 数量
	err = db.Model(&entity.Menu{}).Where("menuType in (?)", menuType).Count(&total).Error
	restfulx.ErrNotNilDebug(err, restfulx.OperatorErr)
	var menuIds []uint64
	for _, menuInfo := range menuList {
		menuIds = append(menuIds, menuInfo.ID)
	}

	// 查询子角色
	if len(menuIds) != 0 {
		var menus []entity.Menu
		err = db.Where("parentID in ?", menuIds).Where("menuType in (?)", menuType).Find(&menus).Error
		restfulx.ErrNotNilDebug(err, restfulx.OperatorErr)
		if len(menus) != 0 {
			menuList = append(menuList, menus...)
			// 查询子角色的按钮及API
			var ids []uint64
			for _, menuInfo := range menus {
				ids = append(ids, menuInfo.ID)
			}
			var ms []entity.Menu
			err = db.Where("parentID in ?", ids).Where("menuType in (?)", menuType).Find(&ms).Error
			restfulx.ErrNotNilDebug(err, restfulx.OperatorErr)
			menuList = append(menuList, ms...)
		}

	}

	if isTree {
		treeMenus := GetTreeMenus(menuList, 0)
		res = &form.PageMenu{
			Menus: treeMenus,
			Total: total,
		}
		return res
	}

	res = &form.PageMenu{
		Menus: menuList,
		Total: total,
	}

	restfulx.ErrNotNilDebug(err, restfulx.OperatorErr)
	return res
}

func (m *Menu) GetByIds(ctx context.Context, menuIDs []uint64) *[]entity.Menu {
	var menus *[]entity.Menu
	err := m.db.Where("id in ?", menuIDs).Find(&menus).Error
	restfulx.ErrNotNilDebug(err, restfulx.OperatorErr)
	return menus
}

func (m *Menu) GetMenuByMenuNameUrl(ctx context.Context, url, method string) *entity.Menu {
	men := &entity.Menu{}
	err := m.db.Where("path = ? and method = ?", url, method).First(men).Error

	restfulx.ErrNotNilDebug(err, restfulx.OperatorErr)
	return men
}
func (m *Menu) CheckUniqueMenuNameUrl(ctx context.Context, url, method string) bool {
	men := &entity.Menu{}
	err := m.db.Where("path = ? and method = ?", url, method).First(men).Error
	if err != nil {
		return false
	}
	return true
}

func (m *Menu) CheckMenusIsExist(ctx context.Context, menuID uint64) bool {
	var menu *entity.Menu
	if err := m.db.Model(&entity.Menu{}).Where("id = ?", menuID).First(&menu).Error; err != nil {
		return false
	}
	return true
}

func (m *Menu) UpdateStatus(ctx context.Context, menuID, status uint64) {
	restfulx.ErrNotNilDebug(m.db.Model(&entity.Menu{}).Where("id = ?", menuID).Update("status", status).Error, restfulx.OperatorErr)
}

func GetTreeMenus(menusList []entity.Menu, pid uint64) (treeMenusList []entity.Menu) {
	for _, node := range menusList {
		if node.ParentID == pid {
			child := GetTreeMenus(menusList, node.ID)
			node.Children = child
			treeMenusList = append(treeMenusList, node)
		}
	}
	return treeMenusList
}
