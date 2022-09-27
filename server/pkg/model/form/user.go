package form

import (
	"github.com/lbemi/lbemi/pkg/util"
)

type UserLoginForm struct {
	UserName  string `json:"user_name" form:"user_name" binding:"required"` // 用户名
	Password  string ` json:"password" form:"password" binding:"required,min=5,max=20"`
	Mobile    string `json:"mobile" form:"mobile" binding:"mobile"`
	Captcha   string `json:"captcha" form:"captcha"  binding:"required,min=5,max=5"`
	CaptchaId string `json:"captcha_id" form:"captcha_id"  binding:"required"`
}

func (u UserLoginForm) GetMessages() util.ValidatorMessages {
	return util.ValidatorMessages{
		"user_name.required":  "用户名不能为空",
		"password.required":   "密码不能为空,最少5位",
		"mobile.required":     "手机号码不能为空",
		"mobile.mobile":       "手机号码格式错误",
		"captcha.required":    "验证码不能为空",
		"captcha_id.required": "captcha_id不存在",
	}
}

type RegisterUserForm struct {
	UserName string `json:"user_name" form:"user_name" binding:"required"` // 用户名
	Password string ` json:"password" form:"password" binding:"required,min=5,max=20"`
	Mobile   string `json:"mobile" form:"mobile" binding:"mobile"`
	Email    string `json:"email" form:"email" binding:"email"" `
}

func (u RegisterUserForm) GetMessages() util.ValidatorMessages {
	return util.ValidatorMessages{
		"user_name.required": "用户名不能为空",
		"password.required":  "密码不能为空,最少5位",
		"email.email":        "email格式错误",
		"mobile.mobile":      "手机号码格式错误",
	}
}
