package sys

import (
	"errors"
	"github.com/fatih/structs"
	"github.com/lbemi/lbemi/pkg/bootstrap/log"
	"github.com/lbemi/lbemi/pkg/model/form"
	"github.com/lbemi/lbemi/pkg/model/sys"
	"gorm.io/gorm"
)

// IMenu 菜单操作接口
type IMenu interface {
	Create(*sys.Menu) (*sys.Menu, error)
	Update(*form.UpdateMenusReq, uint64) error
	Delete(uint64) error
	Get(uint64) (*sys.Menu, error)
	List(page, limit int, menuType []int8) (res *form.PageMenu, err error)

	GetByIds([]uint64) (*[]sys.Menu, error)
	GetMenuByMenuNameUrl(string, string) (*sys.Menu, error)
	UpdateStatus(menuId, status uint64) error
}

type menu struct {
	db *gorm.DB
}

func NewMenu(db *gorm.DB) *menu {
	return &menu{db}
}

func (m *menu) Create(obj *sys.Menu) (*sys.Menu, error) {
	if err := m.db.Create(obj).Error; err != nil {
		return nil, err
	}
	return obj, nil
}

func (m *menu) Update(obj *form.UpdateMenusReq, mId uint64) error {
	objMap := structs.Map(obj)
	tx := m.db.Model(&sys.Menu{}).Where("id = ?  ", mId).Updates(objMap)
	if tx.RowsAffected == 0 {
		return errors.New("update failed")
	}
	return tx.Error
}

func (m *menu) Delete(mId uint64) error {
	tx := m.db.Begin()
	defer func() {
		if err := recover(); err != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		tx.Rollback()
		log.Logger.Error(err)
		return err
	}

	// 清除role_menus
	if err := tx.Where("menu_id = ?", mId).Delete(&sys.RoleMenu{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	// 清除menus
	if err := tx.Where("id =  ?", mId).
		Or("parent_id = ?", mId).
		Delete(&sys.Menu{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func (m *menu) Get(mId uint64) (menu *sys.Menu, err error) {
	if err = m.db.Model(&sys.Menu{}).Where("id = ?", mId).First(&menu).Error; err != nil {
		return nil, err
	}
	return
}

func (m *menu) List(page, limit int, menuType []int8) (res *form.PageMenu, err error) {

	var (
		menuList []sys.Menu
		total    int64
	)

	// 全量查询
	if page == 0 && limit == 0 {
		if tx := m.db.Order("sequence DESC").Where("menu_type in (?)", menuType).Find(&menuList); tx.Error != nil {
			return nil, tx.Error
		}
		treeMenu := GetTreeMenus(menuList, 0)

		if err := m.db.Model(&sys.Menu{}).Where("menu_type in (?)", menuType).Count(&total).Error; err != nil {
			return nil, err
		}

		res = &form.PageMenu{
			Menus: treeMenu,
			Total: total,
		}
		return res, err
	}

	//分页数据
	if err := m.db.Order("sequence DESC").Where("parent_id = 0").Where("menu_type in (?)", menuType).Limit(limit).Offset((page - 1) * limit).
		Find(&menuList).Error; err != nil {
		return nil, err
	}

	var menuIds []uint64
	for _, menuInfo := range menuList {
		menuIds = append(menuIds, menuInfo.ID)
	}

	// 查询子角色
	if len(menuIds) != 0 {
		var menus []sys.Menu
		if err := m.db.Where("parent_id in ?", menuIds).Where("menu_type in (?)", menuType).Find(&menus).Error; err != nil {
			return nil, err
		}
		menuList = append(menuList, menus...)

		// 查询子角色的按钮及API
		var ids []uint64
		for _, menuInfo := range menus {
			ids = append(ids, menuInfo.ID)
		}
		if len(ids) != 0 {
			var ms []sys.Menu
			if err := m.db.Where("parent_id in ?", ids).Where("menu_type in (?)", menuType).Find(&ms).Error; err != nil {
				return nil, err
			}
			menuList = append(menuList, ms...)
		}

	}

	if err := m.db.Model(&sys.Menu{}).Where("parent_id = 0").Where("menu_type in (?)", menuType).Count(&total).Error; err != nil {
		return nil, err
	}

	treeMenus := GetTreeMenus(menuList, 0)
	res = &form.PageMenu{
		Menus: treeMenus,
		Total: total,
	}

	return

}

func (m *menu) GetByIds(mIds []uint64) (menus *[]sys.Menu, err error) {
	if err = m.db.Where("id in ?", mIds).Find(&menus).Error; err != nil {
		return nil, err
	}
	return
}

func (m *menu) GetMenuByMenuNameUrl(url, method string) (menu *sys.Menu, err error) {
	err = m.db.Where("url = ? and method = ?", url, method).First(&menu).Error
	return
}

func (m *menu) UpdateStatus(menuId, status uint64) error {
	return m.db.Model(&sys.Menu{}).Where("id = ?", menuId).Update("status", status).Error
}

func GetTreeMenus(menusList []sys.Menu, pid uint64) (treeMenusList []sys.Menu) {
	for _, node := range menusList {
		if node.ParentID == pid {
			child := GetTreeMenus(menusList, node.ID)
			node.Children = child
			treeMenusList = append(treeMenusList, node)
		}
	}
	return treeMenusList
}