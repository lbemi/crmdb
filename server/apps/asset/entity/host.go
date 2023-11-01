package entity

import (
	"database/sql/driver"
	"encoding/json"
	"github.com/lbemi/lbemi/pkg/common/entity"
	"time"

	"gorm.io/gorm"
)

// Host 机器信息
type Host struct {
	entity.Model
	GroupId   uint64 `json:"group_id" gorm:"column:group_id;comment:组id"`
	Labels    Label  `json:"labels" gorm:"labels;comment:标签"`
	Ip        string `json:"ip" gorm:"column:ip; not null;unique_index:uk_host_ip;"`                  // IP地址
	Remark    string `json:"remark" gorm:"column:remark;size:128;"`                                   // 备注
	Port      int    `json:"port" gorm:"column:port; not null; default:22"`                           // 端口号
	Status    int8   `json:"status" gorm:"column:status;type:tinyint(1);default:1;not null;"`         // 状态 1:启用；2:停用
	EnableSSH int8   `json:"enable_ssh" gorm:"column:enable_ssh;type:tinyint(1);default:1;not null;"` // 是否允许SSH 1:启用；2:停用
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
	GroupId   uint64 `json:"groupId"`
	Name      string `json:"name"`
	Label     string `json:"label"`      // 标签
	Remark    string `json:"remark"`     // 备注
	Ip        string `json:"ip"`         // IP地址
	Port      int    `json:"port"`       // 端口号
	Status    int8   `json:"status"`     // 状态 1:启用；2:停用
	EnableSSH int8   `json:"enable_ssh"` // 是否允许SSH 1:启用；2:停用
}

type Label []string

func (l *Label) Scan(value interface{}) error {
	bytesValue, ok := value.([]byte)
	if !ok || len(bytesValue) == 0 {
		return nil
	}
	result := Label{}
	err := json.Unmarshal(bytesValue, &result)
	*l = Label(result)
	return err
}

// Value 存入前转换为string
func (l Label) Value() (driver.Value, error) {
	if len(l) == 0 {
		return "", nil
	}
	bytes, err := json.Marshal(l)
	if err != nil {
		return "", nil
	}
	return string(bytes), nil
}
