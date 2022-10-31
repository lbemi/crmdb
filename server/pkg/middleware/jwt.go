package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/lbemi/lbemi/pkg/lbemi"
	"github.com/lbemi/lbemi/pkg/util"
	"time"

	"github.com/lbemi/lbemi/pkg/common/response"
)

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenStr := c.Request.Header.Get("Authorization")
		if tokenStr == "" {
			response.Fail(c, response.ErrCodeNotLogin)
			c.Abort()
			return
		}

		token, claims, err := util.ParseToken(tokenStr)
		if err != nil || isInBlacklist(tokenStr) {
			response.Fail(c, response.InvalidToken)
			c.Abort()
			return
		}

		c.Set("id", claims.Id)
		c.Set("token", token)
	}
}
func getBlackListKey(tokenStr string) string {
	return "jwt_black_list:" + util.MD5([]byte(tokenStr))
}

func JoinBlackList(token *jwt.Token) (err error) {
	nowUnix := time.Now().Unix()
	timer := time.Duration(token.Claims.(*util.CustomClaims).ExpiresAt-nowUnix) * time.Second
	return lbemi.CoreV1.Redis().SetNX(getBlackListKey(token.Raw), nowUnix, timer)
	return nil
}

func isInBlacklist(tokenStr string) bool {
	joinUnixStr, err := lbemi.CoreV1.Redis().Get(getBlackListKey(tokenStr)).Result()
	if err != nil || joinUnixStr == "" {

		return false
	}

	return true
}
