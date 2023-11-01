package entity

import (
	"github.com/lbemi/lbemi/pkg/common/entity"
	"gorm.io/gorm"
	"time"
)

type Group struct {
	entity.Model
	Name     string  `json:"name" gorm:"column:name;type:varchar(100);not null;unique"`
	ParentId uint64  `json:"parent_id" gorm:"column:parent_id"`
	Sequence int     `gorm:"column:sequence;not null;comment:排序值" json:"sequence" form:"sequence"`                                   // 排序值
	Memo     string  `gorm:"column:memo;size:128;comment:备注" json:"memo" form:"memo"`                                                // 备注
	Status   string  `gorm:"column:status;type:varchar(2);not null;default:01;comment:状态(01:启用 02:不启用)" json:"status" form:"status"` // 状态(01:启用 02:不启用)
	Children []Group `json:"children,omitempty" gorm:"-"`
}

// TableName 自定义表名
func (g *Group) TableName() string {
	return "groups"
}

// BeforeCreate 添加前
func (g *Group) BeforeCreate(*gorm.DB) error {
	g.CreatedAt = time.Now()
	g.UpdatedAt = time.Now()
	return nil
}

// BeforeUpdate 更新前
func (g *Group) BeforeUpdate(*gorm.DB) error {
	g.UpdatedAt = time.Now()
	return nil
}
