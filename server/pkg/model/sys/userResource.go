package sys

import (
	"github.com/lbemi/lbemi/pkg/model/basemodel"
	"gorm.io/gorm"
	"time"
)

type UserResource struct {
	basemodel.Model
	UserID     uint64 `gorm:"column:user_id;not null;comment:用户ID" json:"user_id"` // 用户ID
	Type       string `json:"type" gorm:"column:type;comment:资源类型"`
	ResourceID uint64 `json:"resourceID" gorm:"resourceID;comment:资源ID"`
}

// TableName 表名
func (u *UserResource) TableName() string {
	return "user_resources"
}

// BeforeCreate 添加前
func (u *UserResource) BeforeCreate(*gorm.DB) error {
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
	return nil
}

// BeforeUpdate 更新前
func (u *UserResource) BeforeUpdate(tx *gorm.DB) error {
	u.UpdatedAt = time.Now()
	return nil
}
