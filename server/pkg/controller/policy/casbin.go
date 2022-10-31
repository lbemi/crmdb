package policy

import (
	"context"
	"github.com/lbemi/lbemi/pkg/bootstrap/log"
	"github.com/lbemi/lbemi/pkg/model/sys"
	"github.com/lbemi/lbemi/pkg/services"

	"github.com/casbin/casbin/v2"
)

type PolicyGetter interface {
	Policy() PolicyInterface
}

type PolicyInterface interface {
	GetEnforce() *casbin.Enforcer
	AddRoleForUser(ctx context.Context, userid uint64, roleIds []uint64) (err error)
	SetRolePermission(ctx context.Context, roleId uint64, menus *[]sys.Menu) (bool, error)
	DeleteRole(ctx context.Context, roleId uint64) error
	DeleteRolePermission(ctx context.Context, resource ...string) error
}

type policy struct {
	factory services.IDbFactory
}

func NewPolicy(f services.IDbFactory) PolicyInterface {
	return &policy{
		factory: f,
	}
}

// GetEnforce 获取全局enforcer
func (c *policy) GetEnforce() *casbin.Enforcer {
	return c.factory.Authentication().GetEnforce()
}

// AddRoleForUser 添加用户角色权限
func (c *policy) AddRoleForUser(ctx context.Context, userid uint64, roleIds []uint64) (err error) {
	err = c.factory.Authentication().AddRoleForUser(userid, roleIds)
	if err != nil {
		log.Logger.Error(err)
		return
	}
	return
}

// SetRolePermission 设置角色权限
func (c *policy) SetRolePermission(ctx context.Context, roleId uint64, menus *[]sys.Menu) (bool, error) {
	ok, err := c.factory.Authentication().SetRolePermission(roleId, menus)
	if err != nil {
		log.Logger.Error(err)
		return ok, err
	}
	return ok, err
}

// DeleteRole 删除角色
func (c *policy) DeleteRole(ctx context.Context, roleId uint64) error {
	err := c.factory.Authentication().DeleteRole(roleId)
	if err != nil {
		log.Logger.Error(err)
		return err
	}
	return err
}

// DeleteRolePermission 删除角色权限
func (c *policy) DeleteRolePermission(ctx context.Context, resource ...string) error {
	err := c.factory.Authentication().DeleteRolePermission(resource...)
	if err != nil {
		log.Logger.Error(err)
		return err
	}
	return err
}