package sys

import (
	"github.com/fatih/structs"
	"github.com/lbemi/lbemi/pkg/model"
	"github.com/lbemi/lbemi/pkg/model/form"
	"github.com/lbemi/lbemi/pkg/model/sys"
	"gorm.io/gorm"
)

// IRole 角色操作接口
type IRole interface {
	Create(role *sys.Role) error
	Update(role *form.UpdateRoleReq, roleID uint64) error
	Delete(roleID uint64) (*gorm.DB, error)
	Get(uint64) (*[]sys.Role, error)
	List(*model.PageParam, *sys.Role) (*form.PageResult, error)
	GetMenusByRoleID(roleID uint64, menuType []int8) (*[]sys.Menu, error)
	SetRole(roleID uint64, menuIDs []uint64) error
	GetRolesByMenuID(menuID uint64) (*[]uint64, error)
	GetRoleByRoleName(roleName string) (*sys.Role, error)
	UpdateStatus(roleID, status uint64) error
}

type role struct {
	db *gorm.DB
}

func NewRole(db *gorm.DB) IRole {
	return &role{db}
}

func (r *role) Create(role *sys.Role) error {
	return r.db.Create(role).Error
}

func (r *role) Update(role *form.UpdateRoleReq, rid uint64) error {
	roleMap := structs.Map(role)
	return r.db.Model(&sys.Role{}).Where("id = ? ", rid).Updates(roleMap).Error
}

func (r *role) Delete(roleID uint64) (*gorm.DB, error) {
	tx := r.db.Begin()
	defer func() {
		if err := recover(); err != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		tx.Rollback()
		return tx, err
	}

	//删除角色相关的菜单
	if err := tx.Where("role_id = ?", roleID).Delete(&sys.RoleMenu{}).Error; err != nil {
		tx.Rollback()
		return tx, err
	}

	// 删除角色及其子角色
	if err := tx.Where("id  = ?", roleID).
		Or("parent_id  = ?", roleID).
		Delete(&sys.Role{}).Error; err != nil {
		tx.Rollback()
		return tx, err
	}

	// 删除用户绑定的角色信息(用户需要重新绑定角色)
	if err := tx.Where("role_id = ?", roleID).
		Delete(&sys.UserRole{}).Error; err != nil {
		tx.Rollback()
		return tx, err
	}
	return tx, tx.Commit().Error
}

func (r *role) Get(roleID uint64) (roles *[]sys.Role, err error) {
	err = r.db.Where("id = ?", roleID).
		Or("parent_id = ?", roleID).
		Order("sequence DESC").
		First(&roles).Error
	if err != nil {
		return nil, err
	}

	res := GetTreeRoles(*roles, 0)
	return &res, err
}

func (r *role) List(query *model.PageParam, condition *sys.Role) (*form.PageResult, error) {
	var (
		roleList []sys.Role
		total    int64
	)
	db := r.db
	offset := (query.Page - 1) * query.Limit
	if condition.Name != "" {
		db = db.Where("name like ?", condition.Name)
	}
	res := &form.PageResult{}
	// 全量查询
	if query.Page == 0 && query.Limit == 0 {
		err := db.Order("sequence DESC").Find(&roleList).Error
		if err != nil {
			return nil, err
		}
		treeRole := GetTreeRoles(roleList, 0)
		err = db.Model(&sys.Role{}).Count(&total).Error
		if err != nil {
			return nil, err
		}

		res.Total = total
		res.Data = treeRole
		return res, nil
	}

	//分页数据
	err := db.Order("sequence DESC").Where("parent_id = 0").Limit(query.Limit).Offset(offset).
		Find(&roleList).Error
	if err != nil {
		return nil, err
	}

	var roleIds []uint64
	for _, role := range roleList {
		roleIds = append(roleIds, role.ID)
	}

	// 查询子角色
	if len(roleIds) != 0 {
		var roles []sys.Role
		err = db.Where("parent_id in ?", roleIds).Find(&roles).Error
		if err != nil {
			return nil, err
		}
		roleList = append(roleList, roles...)
	}
	err = db.Model(&sys.Role{}).Where("parent_id = 0").Count(&total).Error
	if err != nil {
		return nil, err
	}
	treeRoles := GetTreeRoles(roleList, 0)
	res.Total = total
	res.Data = treeRoles
	return res, nil
}

func (r *role) GetMenusByRoleID(roleID uint64, menuType []int8) (*[]sys.Menu, error) {
	var menus []sys.Menu
	err := r.db.Table("menus").Select("menus.id, menus.parentID,menus.name,menus.memo, menus.path, menus.icon,menus.sequence,+"+
		"menus.method, menus.menuType, menus.status, menus.component, menus.title, menus.isLink,menus.isHide,menus.isAffix,menus.isKeepAlive,menus.isIframe").
		Joins("left join role_menus on menus.id = role_menus.menuID", roleID).
		Where("role_menus.roleID = ?", roleID).
		Where("menus.menuType in ?", menuType).
		Order("parentId ASC").
		Order("sequence ASC").
		Scan(&menus).Error

	if err != nil {
		return nil, err
	}
	return &menus, nil
}

// SetRole 设置角色菜单权限
func (r *role) SetRole(roleID uint64, menuIDs []uint64) error {
	tx := r.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Where(&sys.RoleMenu{RoleID: roleID}).Delete(&sys.RoleMenu{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	if len(menuIDs) > 0 {
		roleMens := make([]*sys.RoleMenu, len(menuIDs))
		for _, mid := range menuIDs {
			rm := new(sys.RoleMenu)
			rm.RoleID = roleID
			rm.MenuID = mid
			roleMens = append(roleMens, rm)
		}
		if err := tx.Create(&roleMens).Error; err != nil {
			tx.Rollback()
			return err
		}
	}
	return tx.Commit().Error
}

func (r *role) GetRolesByMenuID(menuID uint64) (roleIds *[]uint64, err error) {
	err = r.db.Where("menuID = ?", menuID).Table("role_menus").Pluck("roleID", &roleIds).Error
	return
}

func (r *role) GetRoleByRoleName(roleName string) (role *sys.Role, err error) {
	err = r.db.Where("name = ?", roleName).First(&role).Error
	return
}

func (r *role) UpdateStatus(roleID, status uint64) error {
	return r.db.Model(&sys.Role{}).Where("id = ?", roleID).Update("status", status).Error
}

func GetTreeRoles(rolesList []sys.Role, parentID uint64) (treeRolesList []sys.Role) {
	for _, node := range rolesList {
		if node.ParentID == parentID {
			child := GetTreeRoles(rolesList, node.ID)
			node.Children = child
			treeRolesList = append(treeRolesList, node)
		}
	}
	return treeRolesList
}
