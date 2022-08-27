package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/lbemi/lbemi/pkg/common/response"
	"github.com/lbemi/lbemi/pkg/global"
	"github.com/lbemi/lbemi/pkg/services"
)

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenStr := c.Request.Header.Get("Authorization")
		if tokenStr == "" {
			response.Fail(c, 2003, "登录过期")
			c.Abort()
			return
		}
		fmt.Println(tokenStr)
		token, err := jwt.ParseWithClaims(tokenStr, &services.CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(global.App.Config.Jwt.Key), nil
		})
		if err != nil {
			global.App.Log.Error(err.Error())
			response.Fail(c, 2004, "解析token失败")
			c.Abort()
			return
		}
		claims := token.Claims.(*services.CustomClaims)
		c.Set("token", token)
		c.Set("id", claims.Id)
	}
}
