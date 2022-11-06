package sys

import (
	"errors"
	"github.com/fatih/structs"
	"github.com/lbemi/lbemi/pkg/model/form"
	"github.com/lbemi/lbemi/pkg/model/sys"
	"gorm.io/gorm"
)

// IRole 角色操作接口
type IRole interface {
	Create(role *sys.Role) (*sys.Role, error)
	Update(role *form.UpdateRoleReq, roleID uint64) error
	Delete(roleID uint64) error
	Get(uint64) (*[]sys.Role, error)
	List(page, limit int) (res *form.PageRole, err error)

	GetMenusByRoleID(roleID uint64) (*[]sys.Menu, error)
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

func (r *role) Create(role *sys.Role) (*sys.Role, error) {
	if err := r.db.Create(role).Error; err != nil {
		return nil, err
	}

	return role, nil
}

func (r *role) Update(role *form.UpdateRoleReq, rid uint64) error {

	roleMap := structs.Map(role)
	tx := r.db.Model(&sys.Role{}).Where("id = ? ", rid).Updates(roleMap)
	if tx.RowsAffected == 0 {
		return errors.New("update failed")
	}
	return tx.Error
}

func (r *role) Delete(roleID uint64) error {
	tx := r.db.Begin()
	defer func() {
		if err := recover(); err != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		tx.Rollback()
		return err
	}

	//删除角色相关的菜单
	if err := tx.Where("role_id = ?", roleID).Delete(&sys.RoleMenu{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	// 删除角色及其子角色
	if err := tx.Where("id  = ?", roleID).
		Or("parent_id  = ?", roleID).
		Delete(&sys.Role{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	// 删除用户绑定的角色信息(用户需要重新绑定角色)
	if err := tx.Where("role_id = ?", roleID).
		Delete(&sys.UserRole{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
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

func (r *role) List(page, limit int) (res *form.PageRole, err error) {
	var (
		roleList []sys.Role
		total    int64
	)
	// 全量查询
	if page == 0 && limit == 0 {
		if tx := r.db.Order("sequence DESC").Find(&roleList); tx.Error != nil {
			return nil, tx.Error
		}
		treeRole := GetTreeRoles(roleList, 0)

		if err := r.db.Model(&sys.Role{}).Count(&total).Error; err != nil {
			return nil, err
		}

		res = &form.PageRole{
			Roles: treeRole,
			Total: total,
		}
		return res, err
	}

	//分页数据
	if err := r.db.Order("sequence DESC").Where("parent_id = 0").Limit(limit).Offset((page - 1) * limit).
		Find(&roleList).Error; err != nil {
		return nil, err
	}

	var roleIds []uint64
	for _, role := range roleList {
		roleIds = append(roleIds, role.ID)
	}

	// 查询子角色
	if len(roleIds) != 0 {
		var roles []sys.Role
		if err := r.db.Where("parent_id in ?", roleIds).Find(&roles).Error; err != nil {
			return nil, err
		}
		roleList = append(roleList, roles...)
	}

	if err := r.db.Model(&sys.Role{}).Where("parent_id = 0").Count(&total).Error; err != nil {
		return nil, err
	}

	treeRoles := GetTreeRoles(roleList, 0)
	res = &form.PageRole{
		Roles: treeRoles,
		Total: total,
	}

	return
}

func (r *role) GetMenusByRoleID(roleID uint64) (*[]sys.Menu, error) {
	var menus []sys.Menu
	err := r.db.Table("menus").Select(" menus.id, menus.parent_id,menus.name, menus.url, menus.icon,menus.menu_type,menus.sequence,menus.code,menus.method").
		Joins("left join role_menus on menus.id = role_menus.menu_id", roleID).
		Where("role_menus.role_id = ?", roleID).
		Where("menus.menu_type != 1").
		Order("parent_id ASC").
		Order("sequence ASC").
		Scan(&menus).Error
	if err != nil {
		return nil, err
	}

	//res := getTreeMenus(menus, 0)
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

	if err := tx.Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Where(&sys.RoleMenu{RoleID: roleID}).Delete(&sys.RoleMenu{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	if len(menuIDs) > 0 {
		for _, mid := range menuIDs {
			rm := new(sys.RoleMenu)
			rm.RoleID = roleID
			rm.MenuID = mid
			if err := tx.Create(rm).Error; err != nil {
				tx.Rollback()
				return err
			}
		}
	}

	return tx.Commit().Error
}

func (r *role) GetRolesByMenuID(menuID uint64) (roleIds *[]uint64, err error) {
	err = r.db.Where("menu_id = ?", menuID).Table("role_menus").Pluck("role_id", &roleIds).Error
	if err != nil {
		return
	}
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
