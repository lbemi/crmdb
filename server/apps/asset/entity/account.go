package entity

import (
	"github.com/lbemi/lbemi/pkg/common/entity"
	"gorm.io/gorm"
	"time"
)

type Account struct {
	entity.Model
	Name       string `json:"name" gorm:"column:name;comment:账号名称"`
	UserName   string `json:"user_name" gorm:"column:user_name;comment:登录名'"`
	Password   string `json:"password" gorm:"column:password;comment:登录密码"`
	AuthMethod string `json:"auth_method" gorm:"column:auth_method;type:varchar(2);not null;default:01;comment:登录方式,01:账号密码,02:密钥登录"`
	Secret     string `json:"secret" gorm:"column:secret;comment:登录密钥"`
	Status     int8   `json:"status" gorm:"column:status;type:tinyint(1);not null;default:1;comment:是否启用,1,2禁用"`
}

// TableName 表名
func (a *Account) TableName() string {
	return "accounts"
}

// BeforeCreate 添加前
func (a *Account) BeforeCreate(*gorm.DB) error {
	a.CreatedAt = time.Now()
	a.UpdatedAt = time.Now()
	return nil
}

// BeforeUpdate 更新前
func (a *Account) BeforeUpdate(*gorm.DB) error {
	a.UpdatedAt = time.Now()
	return nil
}
