package auth

import (
	"github.com/casbin/casbin/v2"
	"github.com/lbemi/lbemi/pkg/model/sys"

	"gorm.io/gorm"

	"strconv"
)

// TODO: 整体优化

type AuthenticationInterface interface {
	GetEnforce() *casbin.Enforcer
	AddRoleForUser(userID uint64, roleIDs []uint64) (err error)
	SetRolePermission(roleId uint64, menus *[]sys.Menu) (bool, error)
	DeleteRole(roleId uint64) error
	DeleteRolePermission(resource ...string) error
	DeleteRoleWithUser(uid, roleId uint64) error
	DeleteRolePermissionWithRole(roleId uint64, resource ...string) error
}

type authentication struct {
	db       *gorm.DB
	enforcer *casbin.Enforcer
}

func NewAuthentication(db *gorm.DB, e *casbin.Enforcer) *authentication {
	return &authentication{db, e}
}

func (c *authentication) GetEnforce() *casbin.Enforcer {
	return c.enforcer
}

// AddRoleForUser 分配用户角色
func (c *authentication) AddRoleForUser(userID uint64, roleIDs []uint64) (err error) {
	uidStr := strconv.FormatUint(userID, 10)
	_, err = c.enforcer.DeleteRolesForUser(uidStr)
	if err != nil {
		return
	}
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
	_, err := c.enforcer.DeletePermissionsForUser(strconv.FormatUint(roleId, 10))
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
	for _, menu := range *menus {
		if menu.MenuType == 2 || menu.MenuType == 3 {
			ok, err := c.enforcer.AddPermissionForUser(strconv.FormatUint(roleId, 10), menu.URL, menu.Method)
			if !ok || err != nil {
				return ok, err
			}
		}
	}
	return false, nil
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
