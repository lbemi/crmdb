package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lbemi/lbemi/pkg/common/response"
	"github.com/lbemi/lbemi/pkg/global"
)

func CasbinMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		// 用户ID
		uid, isExit := c.Get("id")
		if !isExit {
			response.Fail(c, response.InvalidToken)
			return
		}
		fmt.Println("********", uid)
		if uid == 1 {
			c.Next()
			return
		}
		p := c.Request.URL.Path
		m := c.Request.Method
		ok, err := global.App.Enforcer.Enforce(uid, p, m)
		if err != nil {
			global.App.Log.Fatal(err.Error())
			response.Fail(c, response.StatusInternalServerError)
			c.Abort()
			return
		}
		if !ok {
			response.Fail(c, response.NoPermission)
			c.Abort()
			return
		}
		c.Next()
	}
}
