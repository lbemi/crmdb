package asset

import (
	"database/sql/driver"
	"encoding/json"
	"github.com/lbemi/lbemi/pkg/model"
	"time"

	"gorm.io/gorm"
)

// Host 机器信息
type Host struct {
	model.Model
	Labels    string `json:"labels" gorm:"labels;comment:标签"`
	Ip        string `json:"ip" gorm:"column:ip; not null;unique_index:uk_host_ip;"`        // IP地址
	Label     string `json:"label" gorm:"column:label;size:128;"`                           // 标签
	Remark    string `json:"remark" gorm:"column:remark;size:128;"`                         // 备注
	Port      int    `json:"port" gorm:"column:port; not null; default 22"`                 // 端口号
	Status    int8   `json:"status" gorm:"column:status;type:tinyint(1);not null;"`         // 状态 1:启用；2:停用
	EnableSSH int8   `json:"enable_ssh" gorm:"column:enable_ssh;type:tinyint(1);not null;"` // 是否允许SSH 1:启用；2:停用
}

// BeforeCreate 添加前
func (m *Host) BeforeCreate(*gorm.DB) error {
	m.CreatedAt = time.Now()
	m.UpdatedAt = time.Now()
	return nil
}

// BeforeUpdate 更新前
func (m *Host) BeforeUpdate(*gorm.DB) error {
	m.UpdatedAt = time.Now()
	return nil
}

type HostReq struct {
	Name      string `json:"name"`
	Label     string `json:"label"`      // 标签
	Remark    string `json:"remark"`     // 备注
	Ip        string `json:"ip"`         // IP地址
	Port      int    `json:"port"`       // 端口号
	Status    int8   `json:"status"`     // 状态 1:启用；2:停用
	EnableSSH int8   `json:"enable_ssh"` // 是否允许SSH 1:启用；2:停用
}

type Tag []string

func (t *Tag) Scan(value interface{}) error {
	bytesValue, _ := value.([]byte)
	return json.Unmarshal(bytesValue, t)
}

func (t Tag) Value() (driver.Value, error) {
	return json.Marshal(t)
}
