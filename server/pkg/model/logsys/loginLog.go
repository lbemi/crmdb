package logsys

import (
	"github.com/lbemi/lbemi/pkg/model"
	"time"

	"gorm.io/gorm"
)

type LogLogin struct {
	model.Model
	Username      string    `json:"username" gorm:"column:username;type:varchar(128);comment:用户名"`
	Status        string    `json:"status" gorm:"column:status;type:varchar(2);comment:状态"`
	Ipaddr        string    `json:"ipaddr" gorm:"column:ipaddr;type:varchar(255);comment:ip地址"`
	LoginLocation string    `json:"loginLocation" gorm:"column:loginLocation;type:varchar(255);comment:归属地"`
	Browser       string    `json:"browser" gorm:"column:browser;type:varchar(255);comment:浏览器"`
	Os            string    `json:"os" gorm:"column:os;type:varchar(255);comment:系统"`
	Platform      string    `json:"platform" gorm:"column:platform;type:varchar(255);comment:固件"`
	LoginTime     time.Time `json:"loginTime" gorm:"column:loginTime;type:timestamp;comment:登录时间"`
	Remark        string    `json:"remark" gorm:"column:remark;type:text;"` //备注
	Msg           string    `json:"msg" gorm:"column:msg;type:varchar(255);"`
}

// TableName 自定义表名
func (u *LogLogin) TableName() string {
	return "log_login"
}

// BeforeCreate 添加前
func (u *LogLogin) BeforeCreate(*gorm.DB) error {
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
	return nil
}

// BeforeUpdate 更新前
func (u *LogLogin) BeforeUpdate(*gorm.DB) error {
	u.UpdatedAt = time.Now()
	return nil
}
