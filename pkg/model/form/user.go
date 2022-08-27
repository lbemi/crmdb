package form

import (
	"github.com/lbemi/lbemi/pkg/util"
)

type UserLoginForm struct {
	UserName string `json:"user_name" form:"user_name" binding:"required"` // 用户名
	Password string ` json:"password" form:"password" binding:"required"`
	Mobile   string `json:"mobile" form:"mobile" binding:"required,mobile"`
}

func (u UserLoginForm) GetMessages() util.ValidatorMessages {
	return util.ValidatorMessages{
		"user_name.required": "用户名不能为空",
		"password.required":  "密码不能为空",
		"mobile.required":    "手机号码不能为空",
		"mobile.mobile":      "手机号码格式错误",
	}
}
