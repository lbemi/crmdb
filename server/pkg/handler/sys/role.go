package sys

import (
	"context"
	"github.com/lbemi/lbemi/pkg/bootstrap/log"
	"github.com/lbemi/lbemi/pkg/model/form"
	"github.com/lbemi/lbemi/pkg/model/sys"
	"github.com/lbemi/lbemi/pkg/services"
)

type RoleGetter interface {
	Role() IRole
}

// IRole 角色操作接口
type IRole interface {
	Create(c context.Context, obj *form.RoleReq) (role *sys.Role, err error)
	Update(c context.Context, role *form.UpdateRoleReq, roleID uint64) error
	Delete(c context.Context, roleID uint64) error
	Get(c context.Context, roleID uint64) (roles *[]sys.Role, err error)
	List(c context.Context, page, limit int) (res *form.PageRole, err error)

	GetMenusByRoleID(c context.Context, roleID uint64) (*[]sys.Menu, error)
	SetRole(c context.Context, roleID uint64, menuIDs []uint64) error
	GetRolesByMenuID(c context.Context, menuID uint64) (roleIDs *[]uint64, err error)
	GetRoleByRoleName(c context.Context, roleName string) (*sys.Role, error)
	CheckRoleIsExist(c context.Context, name string) bool
	UpdateStatus(c context.Context, roleID, status uint64) error
}

type role struct {
	factory services.FactoryImp
}

func NewRole(f services.FactoryImp) *role {
	return &role{
		factory: f,
	}
}

func (r *role) Create(c context.Context, obj *form.RoleReq) (role *sys.Role, err error) {
	if role, err = r.factory.Role().Create(&sys.Role{
		Name:     obj.Name,
		Memo:     obj.Memo,
		ParentID: obj.ParentID,
		Sequence: obj.Sequence,
		Status:   obj.Status,
	}); err != nil {
		log.Logger.Error(err)
		return
	}
	return
}

func (r *role) Update(c context.Context, role *form.UpdateRoleReq, roleID uint64) error {
	if err := r.factory.Role().Update(role, roleID); err != nil {
		log.Logger.Error(err)
		return err
	}
	return nil
}

func (r *role) Delete(c context.Context, roleID uint64) error {
	// 1.先清除rule
	err := r.factory.Authentication().DeleteRole(roleID)
	if err != nil {
		log.Logger.Error(err)
		return err
	}

	// 2.删除user_role
	err = r.factory.Role().Delete(roleID)
	if err != nil {
		log.Logger.Error(err)
		return err
	}
	return nil
}

func (r *role) Get(c context.Context, roleID uint64) (roles *[]sys.Role, err error) {
	if roles, err = r.factory.Role().Get(roleID); err != nil {
		log.Logger.Error(err)
		return
	}
	return
}

func (r *role) List(c context.Context, page, limit int) (res *form.PageRole, err error) {
	if res, err = r.factory.Role().List(page, limit); err != nil {
		log.Logger.Error(err)
		return
	}
	return
}

func (r *role) GetMenusByRoleID(c context.Context, roleID uint64) (*[]sys.Menu, error) {
	menus, err := r.factory.Role().GetMenusByRoleID(roleID)
	if err != nil {
		log.Logger.Error(err)
		return menus, err
	}
	return menus, err
}

// SetRole 设置角色菜单权限
func (r *role) SetRole(c context.Context, roleID uint64, menuIDs []uint64) error {
	// 查询menus信息
	menus, err := r.factory.Menu().GetByIds(menuIDs)
	if err != nil {
		log.Logger.Error(err)
		return err
	}

	// 添加rule规则
	ok, err := r.factory.Authentication().SetRolePermission(roleID, menus)
	if !ok || err != nil {
		log.Logger.Error(err)
		return err
	}

	// 配置role_menus, 如果操作失败，则将rule表中规则清除
	if err = r.factory.Role().SetRole(roleID, menuIDs); err != nil {
		log.Logger.Error(err)
		//清除rule表中规则
		for _, menu := range *menus {
			err := r.factory.Authentication().DeleteRolePermissionWithRole(roleID, menu.URL, menu.Method)
			if err != nil {
				log.Logger.Error(err)
				break
			}
		}
		return err
	}

	return nil
}

func (r *role) GetRolesByMenuID(c context.Context, menuID uint64) (roleIDs *[]uint64, err error) {
	roleIDs, err = r.factory.Role().GetRolesByMenuID(menuID)
	if err != nil {
		log.Logger.Error(err)
		return
	}
	return
}

func (r *role) GetRoleByRoleName(c context.Context, roleName string) (role *sys.Role, err error) {
	role, err = r.factory.Role().GetRoleByRoleName(roleName)
	if err != nil {
		log.Logger.Error(err)
		return
	}
	return
}

func (r *role) UpdateStatus(c context.Context, roleID, status uint64) (err error) {
	err = r.factory.Role().UpdateStatus(roleID, status)
	if err != nil {
		log.Logger.Error(err)
	}
	return
}

// CheckRoleIsExist 判断角色是否存在
func (r *role) CheckRoleIsExist(c context.Context, name string) bool {
	_, err := r.factory.Role().GetRoleByRoleName(name)
	if err != nil {
		log.Logger.Error(err)
		return false
	}

	return true
}
