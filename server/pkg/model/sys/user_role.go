package sys

import (
	"time"

	"gorm.io/gorm"

	"github.com/lbemi/lbemi/pkg/model/basemodel"
)

type UserRole struct {
	basemodel.Model

	UserID uint64 `gorm:"column:user_id;unique_index:uk_user_role_user_id;not null;comment:管理员ID" json:"user_id"` // 管理员ID
	RoleID uint64 `gorm:"column:role_id;unique_index:uk_user_role_user_id;not null;comment:角色ID" json:"role_id"`   // 角色ID
}

// TableName 自定义表名
func (u *UserRole) TableName() string {
	return "user_roles"
}

// BeforeCreate 添加前
func (u *UserRole) BeforeCreate(*gorm.DB) error {
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
	return nil
}

// BeforeUpdate 更新前
func (u *UserRole) BeforeUpdate(*gorm.DB) error {
	u.UpdatedAt = time.Now()
	return nil
}
