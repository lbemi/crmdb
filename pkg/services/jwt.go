package services

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/lbemi/lbemi/pkg/global"
	"time"
)

type JwtUser interface {
	GetUid() string
}

type CustomClaims struct {
	jwt.StandardClaims
}

const (
	tokenType    = "bearer"
	AppGuardName = "app"
)

type TokenOutPut struct {
	Token string `json:"token"`
}

func CreateToken(guardName string, user JwtUser) (tokenOut TokenOutPut, err error) {
	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		CustomClaims{
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: time.Now().Unix() + 60*60*24*global.App.Config.Jwt.TTL,
				Id:        user.GetUid(),
				Issuer:    guardName,
				NotBefore: time.Now().Unix(),
			},
		},
	)
	fmt.Println([]byte(global.App.Config.Jwt.Key), "---", time.Now().Unix())
	tokenStr, err := token.SignedString([]byte(global.App.Config.Jwt.Key))
	if err != nil {
		global.App.Log.Error(err.Error())
		err = errors.New("生成token失败")
	}
	tokenOut.Token = tokenStr
	return
}
