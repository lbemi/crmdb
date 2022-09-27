package sys

import (
	"strconv"
	"time"

	"gorm.io/gorm"

	"github.com/lbemi/lbemi/pkg/model/basemodel"
	"github.com/lbemi/lbemi/pkg/util"
)

// User 用户表
type User struct {
	basemodel.Model
	UserName string `gorm:"column:user_name;size:32;unique_index:uk_user_name;not null;" json:"user_name" form:"user_name"` // 用户名
	RealName string `gorm:"column:real_name;size:32;" json:"real_name" form:"real_name"`                                    // 真实姓名
	Password string `gorm:"column:password;type:char(128);not null;" json:"-" form:"password"`                              // 密码(sha1(md5(明文))加密)
	Email    string `gorm:"column:email;size:64;" json:"email" form:"email"`                                                // 邮箱
	Mobile   string `gorm:"column:mobile;type:char(20);" json:"mobile" form:"mobile"`                                       // 手机号
	Status   uint8  `gorm:"column:status;default:1;type:tinyint(1);not null;" json:"status" form:"status"`                  // 状态(1:正常 2:未激活 3:暂停使用)
}

// TableName 自定义表名
func (u *User) TableName() string {
	return "users"
}

// BeforeCreate 添加前
func (u *User) BeforeCreate(*gorm.DB) error {
	u.ID = util.GetSnowID()
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
	return nil
}

// BeforeUpdate 更新前
func (u *User) BeforeUpdate(*gorm.DB) error {
	u.UpdatedAt = time.Now()
	return nil
}

func (u User) GetSnowID() string {
	return strconv.Itoa(int(u.ID))
}
