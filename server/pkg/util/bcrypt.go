package util

import (
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"

	"github.com/lbemi/lbemi/pkg/bootstrap/log"
)

// BcryptMake 加密
func BcryptMake(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.DefaultCost)
	if err != nil {
		log.Logger.Error("err", zap.Error(err))
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
