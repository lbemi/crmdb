package policy

import (
	"context"
	"fmt"
	"github.com/lbemi/lbemi/pkg/model/rules"
	"github.com/lbemi/lbemi/pkg/restfulx"
	"gorm.io/gorm"
	"strconv"

	"github.com/casbin/casbin/v2"
	"github.com/lbemi/lbemi/pkg/model/sys"
)

type PolicyGetter interface {
	Policy() IPolicy
}

type IPolicy interface {
	GetEnforce() *casbin.SyncedEnforcer
	AddRoleForUser(ctx context.Context, userid uint64, roleIds []uint64) (err error)
	SetRolePermission(ctx context.Context, roleId uint64, menus *[]sys.Menu) (bool, error)
	DeleteRole(ctx context.Context, roleId uint64) error
	DeleteRolePermission(ctx context.Context, resource ...string) error
	UpdatePermissions(ctx context.Context, oldPath, oldMethod, newPath, newMethod string) error
	DeleteUser(userID uint64)
}

type Policy struct {
	db       *gorm.DB
	enforcer *casbin.SyncedEnforcer
}

func NewPolicy(db *gorm.DB, enforcer *casbin.SyncedEnforcer) IPolicy {
	return &Policy{
		db:       db,
		enforcer: enforcer,
	}
}

// GetEnforce 获取全局enforcer
func (c *Policy) GetEnforce() *casbin.SyncedEnforcer {
	return c.enforcer
}

// AddRoleForUser 添加用户角色权限
func (c *Policy) AddRoleForUser(ctx context.Context, userID uint64, roleIDs []uint64) (err error) {
	c.DeleteUser(userID)

	uidStr := strconv.FormatUint(userID, 10)
	for _, roleId := range roleIDs {
		ok, err := c.enforcer.AddRoleForUser(uidStr, strconv.FormatUint(roleId, 10))
		if err != nil || !ok {
			break
		}
	}
	return
}

// SetRolePermission 设置角色权限
func (c *Policy) SetRolePermission(ctx context.Context, roleId uint64, menus *[]sys.Menu) (bool, error) {
	_, err := c.clearCasbin(0, strconv.FormatUint(roleId, 10))
	if err != nil {
		return false, err
	}
	_, err = c.setRolePermission(roleId, menus)
	if err != nil {
		return false, err
	}
	return true, nil
}

// DeleteRole 删除角色
func (c *Policy) DeleteRole(ctx context.Context, roleId uint64) error {
	ok, err := c.enforcer.DeletePermissionsForUser(strconv.FormatUint(roleId, 10))
	if err != nil || !ok {
		return err
	}
	_, err = c.enforcer.DeleteRole(strconv.FormatUint(roleId, 10))
	if err != nil {
		return err
	}

	return nil
}

// DeleteRolePermission 删除角色权限
func (c *Policy) DeleteRolePermission(ctx context.Context, resource ...string) error {
	_, err := c.enforcer.DeletePermission(resource...)
	if err != nil {
		return err
	}
	return nil
}

// UpdatePermissions 更新权限
func (c *Policy) UpdatePermissions(ctx context.Context, oldPath, oldMethod, newPath, newMethod string) error {
	err := c.db.Model(&rules.Rule{}).
		Where("v1 =? and v2 =?", oldPath, oldMethod).
		Updates(&rules.Rule{Path: newPath, Method: newMethod}).Error

	err = c.enforcer.LoadPolicy()
	return err
}

func (c *Policy) DeleteUser(userID uint64) {
	_, err := c.enforcer.DeleteUser(strconv.FormatUint(userID, 10))
	restfulx.ErrNotNilDebug(err, restfulx.OperatorErr)
}

func (c *Policy) clearCasbin(v int, p ...string) (bool, error) {
	success, err := c.enforcer.RemoveFilteredPolicy(v, p...)
	return success, err
}

// 设置角色权限
func (c *Policy) setRolePermission(roleId uint64, menus *[]sys.Menu) (bool, error) {
	rules := [][]string{}
	for _, menu := range *menus {
		if menu.MenuType == 2 || menu.MenuType == 3 {
			rules = append(rules, []string{strconv.FormatUint(roleId, 10), menu.Path, menu.Method})
		}
	}

	if len(rules) == 0 {
		return true, nil
	}

	ok, err := c.enforcer.AddPolicies(rules)
	if !ok {
		return false, fmt.Errorf("操作失败")
	}
	if err != nil {
		return false, err
	}
	return true, nil
}
