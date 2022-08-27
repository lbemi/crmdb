package sys

import (
	"github.com/lbemi/lbemi/pkg/model/basemodel"
	"gorm.io/gorm"
	"time"
)

// RoleMenu 角色-菜单
type RoleMenu struct {
	basemodel.Model
	RoleID uint64 `gorm:"column:role_id;unique_index:uk_role_menu_role_id;not null;"` // 角色ID
	MenuID uint64 `gorm:"column:menu_id;unique_index:uk_role_menu_role_id;not null;"` // 菜单ID
}

// TableName 表名
func (RoleMenu) TableName() string {
	return "role_menu"
}

// BeforeCreate 添加前
func (m *RoleMenu) BeforeCreate(scope *gorm.DB) error {
	m.CreatedAt = time.Now()
	m.UpdatedAt = time.Now()
	return nil
}

// BeforeUpdate 更新前
func (m *RoleMenu) BeforeUpdate(scope *gorm.DB) error {
	m.UpdatedAt = time.Now()
	return nil
}
