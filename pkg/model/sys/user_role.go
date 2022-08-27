package sys

import (
	"github.com/lbemi/lbemi/pkg/model/basemodel"
	"gorm.io/gorm"
	"time"
)

type UserRole struct {
	basemodel.Model
	UserID uint64 `gorm:"column:user_id;unique_index:uk_user_role_user_id;not null;"` // 管理员ID
	RoleID uint64 `gorm:"column:role_id;unique_index:uk_user_role_user_id;not null;"` // 角色ID
}

func (UserRole) TableName() string {
	return "user_role"
}
func (u *UserRole) BeforeCreate(db *gorm.DB) error {
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
	return nil
}

func (u *UserRole) BeforeUpdate(db *gorm.DB) error {
	u.UpdatedAt = time.Now()
	return nil
}
