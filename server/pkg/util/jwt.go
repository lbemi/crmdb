package util

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"

	"github.com/lbemi/lbemi/pkg/global"
)

type JwtUser interface {
	GetSnowID() string
}

type CustomClaims struct {
	jwt.StandardClaims
}

const (
	AppGuardName = "app"
)

var (
	TokenExpired     = errors.New("token已过期")
	TokenNotValidYet = errors.New("token未生效")
	TokenMalformed   = errors.New("无效token")
	TokenInvalid     = errors.New("非法的Token")
)

type TokenOutPut struct {
	Token string `json:"token"`
}

// CreateToken 生成token
func CreateToken(guardName string, user JwtUser) (tokenOut TokenOutPut, err error) {
	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		CustomClaims{
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: time.Now().Unix() + 60*60*24*global.App.Config.Jwt.TTL,
				Id:        user.GetSnowID(),
				Issuer:    guardName,
				NotBefore: time.Now().Unix(),
			},
		},
	)
	fmt.Println([]byte(global.App.Config.Jwt.Key), "---", time.Now().Unix())
	tokenStr, err := token.SignedString([]byte(global.App.Config.Jwt.Key))
	if err != nil {
		global.App.Log.Error(err.Error())

	}
	tokenOut.Token = tokenStr
	return
}

// ParseToken 解析token
func ParseToken(tokenStr string) (token *jwt.Token, claims *CustomClaims, err error) {
	token, err = jwt.ParseWithClaims(tokenStr, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(global.App.Config.Jwt.Key), nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return token, nil, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return token, nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return token, nil, TokenNotValidYet
			} else {
				return token, nil, TokenInvalid
			}
		}
	}
	if token != nil {
		if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
			return token, claims, nil
		}
		return token, nil, TokenInvalid
	} else {
		return token, nil, TokenInvalid
	}
}

// RefreshToken 刷新token
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

func getBlackListKey(tokenStr string) string {
	return "jwt_black_list:" + MD5([]byte(tokenStr))
}

func JoinBlackList(token *jwt.Token) (err error) {
	nowUnix := time.Now().Unix()
	timer := time.Duration(token.Claims.(*CustomClaims).ExpiresAt-nowUnix) * time.Second
	err = global.App.Redis.SetNX(getBlackListKey(token.Raw), nowUnix, timer).Err()
	return
}

func IsInBlacklist(tokenStr string) bool {
	joinUnixStr, err := global.App.Redis.Get(getBlackListKey(tokenStr)).Result()
	if err != nil || joinUnixStr == "" {

		return false
	}

	return true
}