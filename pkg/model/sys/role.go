package sys

import (
	"github.com/lbemi/lbemi/pkg/model/basemodel"
	"time"

	"gorm.io/gorm"
)

// Role 角色
type Role struct {
	basemodel.Model
	Memo     string `gorm:"column:memo;size:64;" json:"memo" form:"memo"`                 // 备注
	Name     string `gorm:"column:name;size:32;not null;" json:"name" form:"name"`        // 名称
	Sequence int    `gorm:"column:sequence;not null;" json:"sequence" form:"sequence"`    // 排序值
	ParentID uint64 `gorm:"column:parent_id;not null;" json:"parent_id" form:"parent_id"` // 父级ID
}

// TableName 自定义表名
func (Role) TableName() string {
	return "role"
}

// BeforeCreate 添加前
func (m *Role) BeforeCreate(scope *gorm.DB) error {
	m.CreatedAt = time.Now()
	m.UpdatedAt = time.Now()
	return nil
}

// BeforeUpdate 更新前
func (m *Role) BeforeUpdate(scope *gorm.DB) error {
	m.UpdatedAt = time.Now()
	return nil
}
