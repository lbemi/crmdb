package sys

import (
	"github.com/lbemi/lbemi/pkg/model"
	"github.com/lbemi/lbemi/pkg/model/form"
	"github.com/lbemi/lbemi/pkg/model/sys"
	"github.com/lbemi/lbemi/pkg/restfulx"
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
	// 1.删除user_role
	tx, err := r.factory.Role().Delete(roleID)
	restfulx.ErrNotNilDebug(err, restfulx.OperatorErr)

	// 2.先清除rule
	err = r.factory.Authentication().DeleteRole(roleID)
	if err != nil {
		tx.Rollback()
		restfulx.ErrNotNilDebug(err, restfulx.OperatorErr)
	}
}

func (r *role) Get(roleID uint64) *[]sys.Role {
	res, err := r.factory.Role().Get(roleID)
	restfulx.ErrNotNilDebug(err, restfulx.OperatorErr)
	return res
}

func (r *role) List(query *model.PageParam, condition *sys.Role) *form.PageResult {
	res, err := r.factory.Role().List(query, condition)
	restfulx.ErrNotNilDebug(err, restfulx.OperatorErr)
	return res
}

func (r *role) GetMenusByRoleID(roleID uint64, menuType []int8) *[]sys.Menu {
	res, err := r.factory.Role().GetMenusByRoleID(roleID, menuType)
	restfulx.ErrNotNilDebug(err, restfulx.OperatorErr)
	return res
}

// SetRole 设置角色菜单权限
func (r *role) SetRole(roleID uint64, menuIDs []uint64) {
	// 查询menus信息
	menus, err := r.factory.Menu().GetByIds(menuIDs)
	restfulx.ErrNotNilDebug(err, restfulx.OperatorErr)

	// 添加rule规则
	ok, err := r.factory.Authentication().SetRolePermission(roleID, menus)
	restfulx.ErrNotNilDebug(err, restfulx.OperatorErr)
	restfulx.ErrNotTrue(ok, restfulx.OperatorErr)

	// 配置role_menus, 如果操作失败，则将rule表中规则清除
	err = r.factory.Role().SetRole(roleID, menuIDs)
	restfulx.ErrNotNilDebug(err, restfulx.OperatorErr)

}

func (r *role) GetRolesByMenuID(menuID uint64) *[]uint64 {
	res, err := r.factory.Role().GetRolesByMenuID(menuID)
	restfulx.ErrNotNilDebug(err, restfulx.OperatorErr)
	return res
}

func (r *role) GetRoleByRoleName(roleName string) *sys.Role {
	res, err := r.factory.Role().GetRoleByRoleName(roleName)
	restfulx.ErrNotNilDebug(err, restfulx.OperatorErr)
	return res
}

func (r *role) UpdateStatus(roleID, status uint64) {
	restfulx.ErrNotNilDebug(r.factory.Role().UpdateStatus(roleID, status), restfulx.OperatorErr)
}

// CheckRoleIsExist 判断角色是否存在
func (r *role) CheckRoleIsExist(name string) bool {
	_, err := r.factory.Role().GetRoleByRoleName(name)
	if err != nil {
		return false
	}
	return true
}
