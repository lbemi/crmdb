package util

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

type ValidatorMessages map[string]string

type Validator interface {
	GetMessages() ValidatorMessages
}

// GetErrorMsg 检查字段非法错误信息
func GetErrorMsg(request interface{}, err error) string {
	if _, isValidatorErrors := err.(validator.ValidationErrors); isValidatorErrors {
		_, isValidator := request.(Validator)
		for _, v := range err.(validator.ValidationErrors) {
			if isValidator {
				if message, exit := request.(Validator).GetMessages()[v.Field()+"."+v.Tag()]; exit {
					return message
				}
			}
			return v.Error()
		}
	}
	return "Parameter error"
}

// ValidateMobile 验证手机号是否合法
func ValidateMobile(fl validator.FieldLevel) bool {
	mobile := fl.Field().String()
	ok, _ := regexp.MatchString(`^(13[0-9]|14[01456879]|15[0-35-9]|16[2567]|17[0-8]|18[0-9]|19[0-35-9])\d{8}$`, mobile)
	if ok {
		return true
	}
	return false
}

// ValidateEmail 验证电子邮箱是否合法
func ValidateEmail(fl validator.FieldLevel) bool {
	email := fl.Field().String()
	ok, _ := regexp.MatchString(`^(\w+([-.][A-Za-z0-9]+)*){3,18}@\w+([-.][A-Za-z0-9]+)*\.\w+([-.][A-Za-z0-9]+)*$`, email)
	if ok {
		return true
	}
	return false
}
