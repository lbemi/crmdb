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
			response.Fail(c, 2003, "登录过期")
			c.Abort()
			return
		}
		token, claims, err := util.ParseToken(tokenStr)
		if err != nil || util.IsInBlacklist(tokenStr) {
			response.Fail(c, 2004, "token已失效")
			c.Abort()
			return
		}
		c.Set("id", claims.Id)
		c.Set("token", token)
	}
}
