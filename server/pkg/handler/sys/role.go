package sys

import (
	"github.com/lbemi/lbemi/pkg/bootstrap/log"
	"github.com/lbemi/lbemi/pkg/model"
	"github.com/lbemi/lbemi/pkg/model/form"
	"github.com/lbemi/lbemi/pkg/model/sys"
	"github.com/lbemi/lbemi/pkg/services"
)

type RoleGetter interface {
	Role() IRole
}

// IRole 角色操作接口
type IRole interface {
	Create(obj *form.RoleReq)
	Update(role *form.UpdateRoleReq, roleID uint64)
	Delete(roleID uint64)
	Get(roleID uint64) (roles *[]sys.Role)
	List(*model.PageParam, *sys.Role) *form.PageResult
	GetMenusByRoleID(roleID uint64, menuType []int8) *[]sys.Menu
	SetRole(roleID uint64, menuIDs []uint64)
	GetRolesByMenuID(menuID uint64) (roleIDs *[]uint64)
	GetRoleByRoleName(roleName string) *sys.Role
	CheckRoleIsExist(name string) bool
	UpdateStatus(roleID, status uint64)
}

type role struct {
	factory services.FactoryImp
}

func NewRole(f services.FactoryImp) *role {
	return &role{
		factory: f,
	}
}

func (r *role) Create(obj *form.RoleReq) {
	r.CheckRoleIsExist(obj.Name)
	r.factory.Role().Create(&sys.Role{
		Name:     obj.Name,
		Memo:     obj.Memo,
		ParentID: obj.ParentID,
		Sequence: obj.Sequence,
		Status:   obj.Status})
}

func (r *role) Update(role *form.UpdateRoleReq, roleID uint64) {
	r.factory.Role().Update(role, roleID)
}

func (r *role) Delete(roleID uint64) {
	// 1.先清除rule
	r.factory.Authentication().DeleteRole(roleID)

	// 2.删除user_role
	r.factory.Role().Delete(roleID)
}

func (r *role) Get(roleID uint64) *[]sys.Role {
	return r.factory.Role().Get(roleID)
}

func (r *role) List(query *model.PageParam, condition *sys.Role) *form.PageResult {
	return r.factory.Role().List(query, condition)
}

func (r *role) GetMenusByRoleID(roleID uint64, menuType []int8) *[]sys.Menu {
	return r.factory.Role().GetMenusByRoleID(roleID, menuType)
}

// SetRole 设置角色菜单权限
func (r *role) SetRole(roleID uint64, menuIDs []uint64) {
	// 查询menus信息
	menus, err := r.factory.Menu().GetByIds(menuIDs)
	if err != nil {
		log.Logger.Error(err)
		panic(err)
	}

	// 添加rule规则
	ok, err := r.factory.Authentication().SetRolePermission(roleID, menus)
	if !ok || err != nil {
		log.Logger.Error(err)
		panic(err)
	}

	// 配置role_menus, 如果操作失败，则将rule表中规则清除
	r.factory.Role().SetRole(roleID, menuIDs)
	//清除rule表中规则
	for _, menu := range *menus {
		err := r.factory.Authentication().DeleteRolePermissionWithRole(roleID, menu.Path, menu.Method)
		if err != nil {
			log.Logger.Error(err)
			break
		}
	}

}

func (r *role) GetRolesByMenuID(menuID uint64) *[]uint64 {
	return r.factory.Role().GetRolesByMenuID(menuID)
}

func (r *role) GetRoleByRoleName(roleName string) *sys.Role {
	return r.factory.Role().GetRoleByRoleName(roleName)
}

func (r *role) UpdateStatus(roleID, status uint64) {
	r.factory.Role().UpdateStatus(roleID, status)
}

// CheckRoleIsExist 判断角色是否存在
func (r *role) CheckRoleIsExist(name string) bool {
	if r := recover(); r != nil {
		return false
	}
	r.factory.Role().GetRoleByRoleName(name)
	return true
}
