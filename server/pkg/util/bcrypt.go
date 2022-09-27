package util

import (
	"github.com/lbemi/lbemi/pkg/global"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

// BcryptMake 加密
func BcryptMake(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.DefaultCost)
	if err != nil {
		global.App.Log.Error("err", zap.Error(err))
	}
	return string(hash)
}

// BcryptMakeCheck 加密验证
func BcryptMakeCheck(pwd []byte, hashedPwd string) bool {
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, pwd)
	if err != nil {
		return false
	}
	return true
}
