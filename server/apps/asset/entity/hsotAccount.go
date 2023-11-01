package entity

import (
	"database/sql/driver"
	"encoding/json"
	"github.com/lbemi/lbemi/pkg/common/entity"
	"gorm.io/gorm"
	"time"
)

type HostAccount struct {
	entity.Model
	RuleName   string `json:"rule_name" gorm:"column:rule_name;unique;comment:规则名称"`
	AccountId  Ids    `json:"account_id" gorm:"column:account_id;unique_index:uk_host_account_id;not null;comment:账号Id"`
	ResourceId Ids    `json:"resource_id" gorm:"column:resource_id;unique_index:uk_account_resource_id;not null;comment:资产ID"`
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

type Ids []uint64

func (t *Ids) Scan(value interface{}) error {
	bytesValue, ok := value.([]byte)
	if !ok || len(bytesValue) == 0 {
		return nil
	}
	result := Ids{}
	err := json.Unmarshal(bytesValue, &result)
	*t = Ids(result)
	return err
}

// Value 存入前转换为string
func (t Ids) Value() (driver.Value, error) {
	if len(t) == 0 {
		return "", nil
	}
	bytes, err := json.Marshal(t)
	if err != nil {
		return "", nil
	}
	return string(bytes), nil
}
