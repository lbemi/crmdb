package services

import (
	"context"
	"github.com/fatih/structs"
	form2 "github.com/lbemi/lbemi/apps/system/api/form"
	"github.com/lbemi/lbemi/apps/system/entity"
	"github.com/lbemi/lbemi/pkg/bootstrap/policy"
	entity2 "github.com/lbemi/lbemi/pkg/common/entity"
	"github.com/lbemi/lbemi/pkg/restfulx"
	"gorm.io/gorm"
)

type RoleGetter interface {
	Role() IRole
}

// IRole 角色操作接口
type IRole interface {
	Create(ctx context.Context, obj *form2.RoleReq)
	Update(ctx context.Context, role *form2.UpdateRoleReq, roleID uint64)
	Delete(ctx context.Context, roleID uint64)
	Get(ctx context.Context, roleID uint64) (roles []*entity.Role)
	List(context.Context, *entity2.PageParam, *entity.Role) *entity2.PageResult
	GetMenusByRoleID(ctx context.Context, roleID uint64, menuType []int8) *[]entity.Menu
	SetRole(ctx context.Context, roleID uint64, menuIDs []uint64)
	GetRolesByMenuID(ctx context.Context, menuID uint64) (roleIDs *[]uint64)
	GetRoleByRoleName(ctx context.Context, roleName string) *entity.Role
	CheckRoleIsExist(ctx context.Context, name string) bool
	UpdateStatus(ctx context.Context, roleID, status uint64)
}

type Role struct {
	db     *gorm.DB
	policy policy.IPolicy
	menu   IMenu
}

func NewRole(db *gorm.DB, policy policy.IPolicy, menu IMenu) IRole {
	return &Role{
		db:     db,
		policy: policy,
		menu:   menu,
	}
}

func (r *Role) Create(ctx context.Context, obj *form2.RoleReq) {
	r.CheckRoleIsExist(ctx, obj.Name)
	role := &entity.Role{
		Name:     obj.Name,
		Memo:     obj.Memo,
		ParentID: obj.ParentID,
		Sequence: obj.Sequence,
		Status:   obj.Status}
	restfulx.ErrNotNilDebug(r.db.Create(role).Error, restfulx.OperatorErr)
}

func (r *Role) Update(ctx context.Context, role *form2.UpdateRoleReq, roleID uint64) {
	roleMap := structs.Map(role)
	err := r.db.Model(&entity.Role{}).Where("id = ? ", roleID).Updates(roleMap).Error
	restfulx.ErrNotNilDebug(err, restfulx.OperatorErr)
}

func (r *Role) Delete(ctx context.Context, roleID uint64) {
	// 1.删除user_role
	tx := r.db.Begin()
	defer func() {
		if err := recover(); err != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		tx.Rollback()
		restfulx.ErrNotNilDebug(err, restfulx.OperatorErr)
	}

	//删除角色相关的菜单
	if err := tx.Where("roleID = ?", roleID).Delete(&entity.RoleMenu{}).Error; err != nil {
		tx.Rollback()
		restfulx.ErrNotNilDebug(err, restfulx.OperatorErr)
	}

	// 删除角色及其子角色
	if err := tx.Where("id  = ?", roleID).
		Or("parent_id  = ?", roleID).
		Delete(&entity.Role{}).Error; err != nil {
		tx.Rollback()
		restfulx.ErrNotNilDebug(err, restfulx.OperatorErr)
	}

	// 删除用户绑定的角色信息(用户需要重新绑定角色)
	if err := tx.Where("role_id = ?", roleID).
		Delete(&entity.UserRole{}).Error; err != nil {
		tx.Rollback()
		restfulx.ErrNotNilDebug(err, restfulx.OperatorErr)
	}
	restfulx.ErrNotNilDebug(tx.Commit().Error, restfulx.OperatorErr)

	// 2.先清除rule
	err := r.policy.DeleteRole(ctx, roleID)
	if err != nil {
		tx.Rollback()
		restfulx.ErrNotNilDebug(err, restfulx.OperatorErr)
	}
}

func (r *Role) Get(ctx context.Context, roleID uint64) []*entity.Role {
	var roles []*entity.Role
	err := r.db.Where("id = ?", roleID).
		Or("parent_id = ?", roleID).
		Order("sequence DESC").
		First(&roles).Error
	restfulx.ErrNotNilDebug(err, restfulx.OperatorErr)

	return GetTreeRoles(roles, 0)
}

func (r *Role) List(ctx context.Context, query *entity2.PageParam, condition *entity.Role) *entity2.PageResult {
	var (
		roleList []*entity.Role
		total    int64
	)
	db := r.db
	offset := (query.Page - 1) * query.Limit
	if condition.Name != "" {
		db = db.Where("name like ?", "%"+condition.Name+"%")
	}
	if condition.Status != 0 {
		db = db.Where("status  = ?", condition.Status)
	}
	res := &entity2.PageResult{}
	// 全量查询
	if query.Page == 0 && query.Limit == 0 {
		err := db.Order("sequence DESC").Find(&roleList).Error
		restfulx.ErrNotNilDebug(err, restfulx.OperatorErr)
		treeRole := GetTreeRoles(roleList, 0)
		err = db.Model(&entity.Role{}).Count(&total).Error
		restfulx.ErrNotNilDebug(err, restfulx.OperatorErr)
		res.Total = total
		res.Data = treeRole
		return res
	}

	//分页数据
	err := db.Order("sequence DESC").Where("parent_id = 0").Limit(query.Limit).Offset(offset).
		Find(&roleList).Error
	restfulx.ErrNotNilDebug(err, restfulx.OperatorErr)

	var roleIds []uint64
	for _, role := range roleList {
		roleIds = append(roleIds, role.ID)
	}

	// 查询子角色
	if len(roleIds) != 0 {
		var roles []*entity.Role
		err = db.Where("parent_id in ?", roleIds).Find(&roles).Error
		restfulx.ErrNotNilDebug(err, restfulx.OperatorErr)
		roleList = append(roleList, roles...)
	}
	err = db.Model(&entity.Role{}).Where("parent_id = 0").Count(&total).Error
	restfulx.ErrNotNilDebug(err, restfulx.OperatorErr)
	treeRoles := GetTreeRoles(roleList, 0)
	res.Total = total
	res.Data = treeRoles
	return res
}

func (r *Role) GetMenusByRoleID(ctx context.Context, roleID uint64, menuType []int8) *[]entity.Menu {
	var menus []entity.Menu
	err := r.db.Table("menus").Select("menus.id, menus.parentID,menus.name,menus.memo, menus.path, menus.icon,menus.sequence,+"+
		"menus.method, menus.menuType, menus.status, menus.component, menus.title, menus.isLink,menus.isHide,menus.isAffix,menus.isKeepAlive,menus.isIframe").
		Joins("left join role_menus on menus.id = role_menus.menuID", roleID).
		Where("role_menus.roleID = ?", roleID).
		Where("menus.menuType in ?", menuType).
		Order("parentId ASC").
		Order("sequence ASC").
		Scan(&menus).Error

	restfulx.ErrNotNilDebug(err, restfulx.OperatorErr)
	return &menus
}

// SetRole 设置角色菜单权限
func (r *Role) SetRole(ctx context.Context, roleID uint64, menuIDs []uint64) {
	// 查询menus信息
	menus := r.menu.GetByIds(ctx, menuIDs)

	tx := r.db.Begin()
	if err := tx.Where(&entity.RoleMenu{RoleID: roleID}).Delete(&entity.RoleMenu{}).Error; err != nil {
		tx.Rollback()
		restfulx.ErrNotNilDebug(err, restfulx.OperatorErr)
	}

	if len(menuIDs) > 0 {
		roleMens := make([]entity.RoleMenu, len(menuIDs))
		for _, mid := range menuIDs {
			rm := entity.RoleMenu{}
			rm.RoleID = roleID
			rm.MenuID = mid
			roleMens = append(roleMens, rm)
		}
		if err := tx.Create(&roleMens).Error; err != nil {
			tx.Rollback()
			restfulx.ErrNotNilDebug(err, restfulx.OperatorErr)
		}
	}

	// 添加rule规则
	ok, err := r.policy.SetRolePermission(ctx, roleID, menus)
	if err != nil || !ok {
		tx.Rollback()
		restfulx.ErrNotNilDebug(err, restfulx.OperatorErr)
	}

	restfulx.ErrNotNilDebug(tx.Commit().Error, restfulx.OperatorErr)
}

func (r *Role) GetRolesByMenuID(ctx context.Context, menuID uint64) *[]uint64 {
	var roleIds *[]uint64
	err := r.db.Where("menuID = ?", menuID).Table("role_menus").Pluck("roleID", &roleIds).Error
	restfulx.ErrNotNilDebug(err, restfulx.OperatorErr)
	return roleIds
}

func (r *Role) GetRoleByRoleName(ctx context.Context, roleName string) *entity.Role {
	var role *entity.Role
	err := r.db.Where("name = ?", roleName).First(&role).Error
	restfulx.ErrNotNilDebug(err, restfulx.OperatorErr)
	return role
}

func (r *Role) UpdateStatus(ctx context.Context, roleID, status uint64) {
	restfulx.ErrNotNilDebug(r.db.Model(&entity.Role{}).Where("id = ?", roleID).Update("status", status).Error, restfulx.OperatorErr)
}

// CheckRoleIsExist 判断角色是否存在
func (r *Role) CheckRoleIsExist(ctx context.Context, name string) bool {
	err := r.db.Where("name = ?", name).First(&entity.Role{}).Error
	if err != nil {
		return false
	}
	return true
}

func GetTreeRoles(rolesList []*entity.Role, parentID uint64) (treeRolesList []*entity.Role) {
	for _, node := range rolesList {
		if node.ParentID == parentID {
			child := GetTreeRoles(rolesList, node.ID)
			node.Children = child
			treeRolesList = append(treeRolesList, node)
		}
	}
	return treeRolesList
}
