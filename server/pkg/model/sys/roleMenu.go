package sys

import (
	"github.com/lbemi/lbemi/pkg/model"
	"time"

	"gorm.io/gorm"
)

// RoleMenu 角色-菜单
type RoleMenu struct {
	model.Model

	RoleID uint64 `gorm:"column:roleID;unique_index:uk_role_menu_role_id;not null;comment:角色ID" json:"roleID"`  // 角色ID
	MenuID uint64 `gorm:"column:menuID;unique_index:uk_role_menu_role_id;not null;comment:菜单ID" json:"menuID'"` // 菜单ID
}

// TableName 表名
func (r *RoleMenu) TableName() string {
	return "role_menus"
}

// BeforeCreate 添加前
func (r *RoleMenu) BeforeCreate(*gorm.DB) error {
	r.CreatedAt = time.Now()
	r.UpdatedAt = time.Now()
	return nil
}

// BeforeUpdate 更新前
func (r *RoleMenu) BeforeUpdate(*gorm.DB) error {
	r.UpdatedAt = time.Now()
	return nil
}
