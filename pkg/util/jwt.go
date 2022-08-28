package util

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
	AppGuardName = "app"
)

var (
	TokenExpired     = errors.New("Token已过期")
	TokenNotValidYet = errors.New("Token未生效")
	TokenMalformed   = errors.New("无效token")
	TokenInvalid     = errors.New("非法的Token")
)

type TokenOutPut struct {
	Token string `json:"token"`
}

//CreateToken 生成token
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

//ParseToken 解析token
func ParseToken(tokenStr string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(global.App.Config.Jwt.Key), nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}
		}
	}
	if token != nil {
		if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
			return claims, nil
		}
		return nil, TokenInvalid
	} else {
		return nil, TokenInvalid
	}
}

//RefreshToken 刷新token
func RefreshToken(tokenStr string, user JwtUser) (tokenOut TokenOutPut, err error) {
	jwt.TimeFunc = func() time.Time {
		return time.Unix(0, 0)
	}
	token, err := jwt.ParseWithClaims(tokenStr, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(global.App.Config.Jwt.Key), nil
	})
	if err != nil {
		return
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		jwt.TimeFunc = time.Now
		claims.StandardClaims.ExpiresAt = time.Now().Add(10 * time.Hour).Unix()
		return CreateToken(AppGuardName, user)
	}
	return tokenOut, TokenInvalid
}
