package util

import (
	"errors"
	"github.com/lbemi/lbemi/apps/system/entity"
	"time"

	"github.com/lbemi/lbemi/pkg/restfulx"

	"github.com/golang-jwt/jwt"
)

type JwtUser interface {
	GetSnowID() string
}

type CustomClaims struct {
	User *entity.User
	jwt.StandardClaims
}

const (
	AppGuardName = "server"
	Key          = "3Bde3BGEbYqtqyEUzW3ry8jKFcaPH17fRmTmqE7MDr05Lwj95uruRKrrkb44TJ4s"
	TTL          = 30
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
func CreateToken(guardName string, user *entity.User) (tokenOut TokenOutPut) {
	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		CustomClaims{
			User: user,
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: time.Now().Unix() + 60*60*24*TTL,
				Id:        user.GetSnowID(),
				Issuer:    guardName,
				NotBefore: time.Now().Unix(),
			},
		},
	)
	tokenStr, err := token.SignedString([]byte(Key))
	if err != nil {
		restfulx.ErrNotNilDebug(err, restfulx.OperatorErr)
	}

	tokenOut.Token = tokenStr
	return
}

// ParseToken 解析token
func ParseToken(tokenStr string) (token *jwt.Token, claims *CustomClaims, err error) {
	token, err = jwt.ParseWithClaims(tokenStr, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(Key), nil
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
func RefreshToken(tokenStr string, user *entity.User) (tokenOut TokenOutPut, err error) {
	jwt.TimeFunc = func() time.Time {
		return time.Unix(0, 0)
	}
	token, err := jwt.ParseWithClaims(tokenStr, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(Key), nil
	})
	if err != nil {
		return
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		jwt.TimeFunc = time.Now
		claims.StandardClaims.ExpiresAt = time.Now().Add(10 * time.Hour).Unix()
		CreateToken(AppGuardName, user)
	}
	return tokenOut, TokenInvalid
}

//func getBlackListKey(tokenStr string) string {
//	return "jwt_black_list:" + MD5([]byte(tokenStr))
//}
//
//func JoinBlackList(token *jwt.Token) (err error) {
//	//nowUnix := time.Now().Unix()
//	//timer := time.Duration(token.Claims.(*CustomClaims).ExpiresAt-nowUnix) * time.Second
//	//return core.Core.Redis().SetNX(getBlackListKey(token.Raw), nowUnix, timer)
//	return nil
//}
//
//func IsInBlacklist(tokenStr string) bool {
//	//joinUnixStr, err := core.Core.Redis().Get(getBlackListKey(tokenStr)).Result()
//	//if err != nil || joinUnixStr == "" {
//	//
//	//	return false
//	//}
//
//	return true
//}
