package asset

import (
	"github.com/lbemi/lbemi/pkg/model"
	"gorm.io/gorm"
	"time"
)

type HostAccount struct {
	model.Model
	AccountId  uint64 `json:"account_id" gorm:"column:account_id;unique_index:uk_host_account_id;not null;comment:账号Id"`
	ResourceId uint64 `json:"resource_id" gorm:"column:resource_id;unique_index:uk_account_resource_id;not null;comment:资产ID"`
}

// TableName 自定义表名
func (h *HostAccount) TableName() string {
	return "host_accounts"
}

// BeforeCreate 添加前
func (h *HostAccount) BeforeCreate(*gorm.DB) error {
	h.CreatedAt = time.Now()
	h.UpdatedAt = time.Now()
	return nil
}

// BeforeUpdate 更新前
func (h *HostAccount) BeforeUpdate(*gorm.DB) error {
	h.UpdatedAt = time.Now()
	return nil
}
