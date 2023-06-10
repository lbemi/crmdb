package auth

import (
	"fmt"
	"github.com/casbin/casbin/v2"
	"github.com/lbemi/lbemi/pkg/model/rules"
	"github.com/lbemi/lbemi/pkg/model/sys"
	"github.com/lbemi/lbemi/pkg/restfulx"
	"gorm.io/gorm"

	"strconv"
)

// TODO: 整体优化

type AuthenticationInterface interface {
	GetEnforce() *casbin.SyncedEnforcer
	AddRoleForUser(userID uint64, roleIDs []uint64) (err error)
	SetRolePermission(roleId uint64, menus *[]sys.Menu) (bool, error)
	DeleteRole(roleId uint64) error
	DeleteRolePermission(resource ...string) error
	DeleteRoleWithUser(uid, roleId uint64) error
	DeleteRolePermissionWithRole(roleId uint64, resource ...string) error
	UpdatePermissions(oldPath, oldMethod, newPath, newMethod string) error
	DeleteUser(userID uint64)
}

type authentication struct {
	db       *gorm.DB
	enforcer *casbin.SyncedEnforcer
}

func NewAuthentication(db *gorm.DB, e *casbin.SyncedEnforcer) *authentication {
	return &authentication{db, e}
}

func (c *authentication) GetEnforce() *casbin.SyncedEnforcer {
	return c.enforcer
}

// AddRoleForUser 分配用户角色
func (c *authentication) AddRoleForUser(userID uint64, roleIDs []uint64) (err error) {
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
func (c *authentication) SetRolePermission(roleId uint64, menus *[]sys.Menu) (bool, error) {
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

// 设置角色权限
func (c *authentication) setRolePermission(roleId uint64, menus *[]sys.Menu) (bool, error) {
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

// DeleteRole 删除角色
func (c *authentication) DeleteRole(roleId uint64) error {

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

// DeleteRoleWithUser 删除角色
func (c *authentication) DeleteRoleWithUser(uid, roleId uint64) error {
	ok, err := c.enforcer.DeleteRoleForUser(strconv.FormatUint(uid, 10), strconv.FormatUint(roleId, 10))
	if err != nil || !ok {
		return err
	}

	return nil
}

// DeleteRolePermission 删除角色权限
func (c *authentication) DeleteRolePermission(resource ...string) error {
	_, err := c.enforcer.DeletePermission(resource...)
	if err != nil {
		return err
	}
	return nil
}

// DeleteRolePermissionWithRole 删除角色的权限
func (c *authentication) DeleteRolePermissionWithRole(roleId uint64, resource ...string) error {
	_, err := c.enforcer.DeletePermissionForUser(strconv.FormatUint(roleId, 10), resource...)
	if err != nil {
		return err
	}
	return nil
}

func (c *authentication) UpdatePermissions(oldPath, oldMethod, newPath, newMethod string) error {
	err := c.db.Model(&rules.Rule{}).Where("v1 =? and v2 =?", oldPath, oldMethod).Updates(&rules.Rule{Path: newPath, Method: newMethod}).Error
	c.enforcer.LoadPolicy()
	return err
	//_, err := c.enforcer.UpdatePolicy(oldPolicy, newPolicy)
	//if err != nil {
	//	return err
	//}
	//return nil
}

func (c *authentication) DeleteUser(userID uint64) {
	_, err := c.enforcer.DeleteUser(strconv.FormatUint(userID, 10))
	restfulx.ErrNotNilDebug(err, restfulx.OperatorErr)
}

func (c *authentication) clearCasbin(v int, p ...string) (bool, error) {
	success, err := c.enforcer.RemoveFilteredPolicy(v, p...)
	return success, err
}
