package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lbemi/lbemi/pkg/common/response"
	"github.com/lbemi/lbemi/pkg/util"
)

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenStr := c.Request.Header.Get("Authorization")
		fmt.Println("Authorization", tokenStr)
		if tokenStr == "" {
			response.Fail(c, response.ErrCodeNotLogin)
			c.Abort()
			return
		}
		token, claims, err := util.ParseToken(tokenStr)
		if err != nil || util.IsInBlacklist(tokenStr) {
			response.Fail(c, response.InvalidToken)
			c.Abort()
			return
		}
		c.Set("id", claims.Id)
		c.Set("token", token)
	}
}
