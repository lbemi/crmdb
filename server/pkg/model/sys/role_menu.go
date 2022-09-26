package sys

import (
	"time"

	"gorm.io/gorm"

	"github.com/lbemi/lbemi/pkg/model/basemodel"
)

// RoleMenu 角色-菜单
type RoleMenu struct {
	basemodel.Model

	RoleID int64 `gorm:"column:role_id;unique_index:uk_role_menu_role_id;not null;comment:角色ID" json:"role_id"`  // 角色ID
	MenuID int64 `gorm:"column:menu_id;unique_index:uk_role_menu_role_id;not null;comment:菜单ID" json:"menu_id'"` // 菜单ID
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
