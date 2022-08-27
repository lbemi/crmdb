package util

import (
	"github.com/go-playground/validator/v10"
	"regexp"
)

type ValidatorMessages map[string]string

type Validator interface {
	GetMessages() ValidatorMessages
}

func GetErrorMsg(request interface{}, err error) string {
	if _, isvalidatorErrors := err.(validator.ValidationErrors); isvalidatorErrors {
		_, isvalidator := request.(Validator)
		for _, v := range err.(validator.ValidationErrors) {
			if isvalidator {
				if message, exit := request.(Validator).GetMessages()[v.Field()+"."+v.Tag()]; exit {
					return message
				}
			}
			return v.Error()
		}
	}
	return "Parameter error"
}

func ValidateMobile(fl validator.FieldLevel) bool {
	mobile := fl.Field().String()
	ok, _ := regexp.MatchString(`^(13[0-9]|14[01456879]|15[0-35-9]|16[2567]|17[0-8]|18[0-9]|19[0-35-9])\d{8}$`, mobile)
	if ok {
		return true
	}
	return false
}
