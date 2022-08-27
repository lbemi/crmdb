package sys

import (
	"github.com/lbemi/lbemi/pkg/model/basemodel"
	"gorm.io/gorm"
	"strconv"
	"time"
)

// User 用户表
type User struct {
	basemodel.Model
	Memo     string `gorm:"column:memo;size:64;" json:"memo" form:"memo"`                                                   //备注
	UserName string `gorm:"column:user_name;size:32;unique_index:uk_user_name;not null;" json:"user_name" form:"user_name"` // 用户名
	RealName string `gorm:"column:real_name;size:32;" json:"real_name" form:"real_name"`                                    // 真实姓名
	Password string `gorm:"column:password;type:char(128);not null;" json:"password" form:"password"`                       // 密码(sha1(md5(明文))加密)
	Email    string `gorm:"column:email;size:64;" json:"email" form:"email"`                                                // 邮箱
	Mobile   string `gorm:"column:mobile;type:char(20);" json:"mobile" form:"mobile"`                                       // 手机号
	Status   uint8  `gorm:"column:status;type:tinyint(1);not null;" json:"status" form:"status"`                            // 状态(1:正常 2:未激活 3:暂停使用)
}

// TableName 自定义表名
func (User) TableName() string {
	return "user"
}

// BeforeCreate 添加前
func (m *User) BeforeCreate(*gorm.DB) error {
	m.CreatedAt = time.Now()
	m.UpdatedAt = time.Now()
	return nil
}

// BeforeUpdate 更新前
func (m *User) BeforeUpdate(*gorm.DB) error {
	m.UpdatedAt = time.Now()
	return nil
}
func (u User) GetUid() string {
	return strconv.Itoa(int(u.ID))
}
