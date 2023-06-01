package sys

import (
	"github.com/fatih/structs"
	"github.com/lbemi/lbemi/pkg/model"
	"github.com/lbemi/lbemi/pkg/model/form"
	"github.com/lbemi/lbemi/pkg/model/sys"
	"github.com/lbemi/lbemi/pkg/restfulx"
	"gorm.io/gorm"
)

// IRole 角色操作接口
type IRole interface {
	Create(role *sys.Role)
	Update(role *form.UpdateRoleReq, roleID uint64)
	Delete(roleID uint64)
	Get(uint64) *[]sys.Role
	List(*model.PageParam, *sys.Role) *form.PageResult
	GetMenusByRoleID(roleID uint64, menuType []int8) *[]sys.Menu
	SetRole(roleID uint64, menuIDs []uint64)
	GetRolesByMenuID(menuID uint64) *[]uint64
	GetRoleByRoleName(roleName string) *sys.Role
	UpdateStatus(roleID, status uint64)
}

type role struct {
	db *gorm.DB
}

func NewRole(db *gorm.DB) IRole {
	return &role{db}
}

func (r *role) Create(role *sys.Role) {
	restfulx.ErrIsNilRes(r.db.Create(role).Error, restfulx.OperatorErr)
}

func (r *role) Update(role *form.UpdateRoleReq, rid uint64) {
	roleMap := structs.Map(role)
	restfulx.ErrIsNilRes(r.db.Model(&sys.Role{}).Where("id = ? ", rid).Updates(roleMap).Error, restfulx.OperatorErr)
}

func (r *role) Delete(roleID uint64) {
	tx := r.db.Begin()
	defer func() {
		if err := recover(); err != nil {
			tx.Rollback()
			restfulx.ErrIsNilRes(err.(error), restfulx.OperatorErr)
		}

		restfulx.ErrIsNilRes(tx.Commit().Error, restfulx.OperatorErr)
	}()

	if err := tx.Error; err != nil {
		tx.Rollback()
		panic(err)
	}

	//删除角色相关的菜单
	if err := tx.Where("role_id = ?", roleID).Delete(&sys.RoleMenu{}).Error; err != nil {
		tx.Rollback()
		panic(err)
	}

	// 删除角色及其子角色
	if err := tx.Where("id  = ?", roleID).
		Or("parent_id  = ?", roleID).
		Delete(&sys.Role{}).Error; err != nil {
		tx.Rollback()
		panic(err)
	}

	// 删除用户绑定的角色信息(用户需要重新绑定角色)
	if err := tx.Where("role_id = ?", roleID).
		Delete(&sys.UserRole{}).Error; err != nil {
		tx.Rollback()
		panic(err)
	}

}

func (r *role) Get(roleID uint64) (roles *[]sys.Role) {
	err := r.db.Where("id = ?", roleID).
		Or("parent_id = ?", roleID).
		Order("sequence DESC").
		First(&roles).Error

	restfulx.ErrIsNilRes(err, restfulx.OperatorErr)
	res := GetTreeRoles(*roles, 0)
	return &res
}

func (r *role) List(query *model.PageParam, condition *sys.Role) *form.PageResult {
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
		restfulx.ErrIsNilRes(db.Order("sequence DESC").Find(&roleList).Error, restfulx.OperatorErr)
		treeRole := GetTreeRoles(roleList, 0)
		restfulx.ErrIsNilRes(db.Model(&sys.Role{}).Count(&total).Error, restfulx.OperatorErr)
		res.Total = total
		res.Data = treeRole
		return res
	}

	//分页数据
	err := db.Order("sequence DESC").Where("parent_id = 0").Limit(query.Limit).Offset(offset).
		Find(&roleList).Error
	restfulx.ErrIsNilRes(err, restfulx.OperatorErr)

	var roleIds []uint64
	for _, role := range roleList {
		roleIds = append(roleIds, role.ID)
	}

	// 查询子角色
	if len(roleIds) != 0 {
		var roles []sys.Role
		restfulx.ErrIsNilRes(db.Where("parent_id in ?", roleIds).Find(&roles).Error, restfulx.OperatorErr)
		roleList = append(roleList, roles...)
	}
	restfulx.ErrIsNilRes(db.Model(&sys.Role{}).Where("parent_id = 0").Count(&total).Error, restfulx.OperatorErr)
	treeRoles := GetTreeRoles(roleList, 0)
	res.Total = total
	res.Data = treeRoles
	return res
}

func (r *role) GetMenusByRoleID(roleID uint64, menuType []int8) *[]sys.Menu {
	var menus []sys.Menu
	err := r.db.Table("menus").Select("menus.id, menus.parentID,menus.name,menus.memo, menus.path, menus.icon,menus.sequence,+"+
		"menus.method, menus.menuType, menus.status, menus.component, menus.title, menus.isLink,menus.isHide,menus.isAffix,menus.isKeepAlive,menus.isIframe").
		Joins("left join role_menus on menus.id = role_menus.menuID", roleID).
		Where("role_menus.roleID = ?", roleID).
		Where("menus.menuType in ?", menuType).
		Order("parentId ASC").
		Order("sequence ASC").
		Scan(&menus).Error

	restfulx.ErrIsNilRes(err, restfulx.OperatorErr)
	return &menus
}

// SetRole 设置角色菜单权限
func (r *role) SetRole(roleID uint64, menuIDs []uint64) {
	tx := r.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			restfulx.ErrIsNilRes(r.(error), restfulx.OperatorErr)
		}
		restfulx.ErrIsNilRes(tx.Commit().Error, restfulx.OperatorErr)
	}()

	if err := tx.Error; err != nil {
		tx.Rollback()
		panic(err)
	}

	if err := tx.Where(&sys.RoleMenu{RoleID: roleID}).Delete(&sys.RoleMenu{}).Error; err != nil {
		tx.Rollback()
		panic(err)
	}

	if len(menuIDs) > 0 {
		for _, mid := range menuIDs {
			rm := new(sys.RoleMenu)
			rm.RoleID = roleID
			rm.MenuID = mid
			if err := tx.Create(rm).Error; err != nil {
				tx.Rollback()
				panic(err)
			}
		}
	}

}

func (r *role) GetRolesByMenuID(menuID uint64) (roleIds *[]uint64) {
	err := r.db.Where("menuID = ?", menuID).Table("role_menus").Pluck("roleID", &roleIds).Error
	restfulx.ErrIsNilRes(err, restfulx.OperatorErr)
	return
}

func (r *role) GetRoleByRoleName(roleName string) (role *sys.Role) {
	restfulx.ErrIsNilRes(r.db.Where("name = ?", roleName).First(&role).Error, restfulx.OperatorErr)
	return
}

func (r *role) UpdateStatus(roleID, status uint64) {
	restfulx.ErrIsNilRes(
		r.db.Model(&sys.Role{}).Where("id = ?", roleID).Update("status", status).Error,
		restfulx.OperatorErr)
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
