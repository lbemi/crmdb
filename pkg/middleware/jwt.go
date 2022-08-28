package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/lbemi/lbemi/pkg/common/response"
	"github.com/lbemi/lbemi/pkg/global"
	"github.com/lbemi/lbemi/pkg/util"
)

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenStr := c.Request.Header.Get("Authorization")
		if tokenStr == "" {
			response.Fail(c, 2003, "登录过期")
			c.Abort()
			return
		}
		claims, err := util.ParseToken(tokenStr)
		if err != nil {
			global.App.Log.Error(err.Error())
			response.Fail(c, 2004, err.Error())
			c.Abort()
			return
		}
		c.Set("id", claims.Id)
	}
}
