package entity

import (
	"github.com/lbemi/lbemi/pkg/common/entity"
	"gorm.io/gorm"
	"time"
)

type HostGroup struct {
	entity.Model
	GroupId    uint64 `json:"group_id" gorm:"column:group_id;unique_index:uk_host_group_id;not null;comment:分组ID"`
	ResourceId uint64 `json:"resource_id" gorm:"column:resource_id;unique_index:uk_group_resource_id;not null;comment:资产ID"`
}

// TableName 自定义表名
func (g *HostGroup) TableName() string {
	return "host_groups"
}

// BeforeCreate 添加前
func (g *HostGroup) BeforeCreate(*gorm.DB) error {
	g.CreatedAt = time.Now()
	g.UpdatedAt = time.Now()
	return nil
}

// BeforeUpdate 更新前
func (g *HostGroup) BeforeUpdate(*gorm.DB) error {
	g.UpdatedAt = time.Now()
	return nil
}
