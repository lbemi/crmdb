package entity

import (
	"github.com/lbemi/lbemi/pkg/common/entity"
	"time"

	"gorm.io/gorm"
)

type LogOperator struct {
	entity.Model
	Title        string `json:"title" gorm:"column:title;type:varchar(128);comment:操作的模块"`
	BusinessType string `json:"businessType" gorm:"column:businessType;type:varchar(2);comment:04其它 01新增 02修改 03删除"`
	Method       string `json:"method" gorm:"column:method;type:varchar(255);comment:请求方法"`
	Name         string `json:"name" gorm:"column:name;type:varchar(255);comment:操作人员"`
	Url          string `json:"url" gorm:"column:url;type:varchar(255);comment:操作url"`
	Ip           string `json:"ip" gorm:"column:ip;type:varchar(255);comment:操作IP"`
	Location     string `json:"location" gorm:"column:location;type:varchar(255);comment:操作地点"`
	Param        string `json:"param" gorm:"column:param;type:varchar(255);comment:请求参数"` //
	Status       int16  `json:"status" gorm:"column:status;comment:请求状态吗"`
	ErrMsg       string `json:"errMsg" gorm:"column:errMsg;type:text"`
}

// TableName 自定义表名
func (u *LogOperator) TableName() string {
	return "log_operator"
}

// BeforeCreate 添加前
func (u *LogOperator) BeforeCreate(*gorm.DB) error {
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
	return nil
}

// BeforeUpdate 更新前
func (u *LogOperator) BeforeUpdate(*gorm.DB) error {
	u.UpdatedAt = time.Now()
	return nil
}
