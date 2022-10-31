package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/lbemi/lbemi/pkg/bootstrap/log"
	"github.com/lbemi/lbemi/pkg/common/response"
	"github.com/lbemi/lbemi/pkg/lbemi"
)

func CasbinMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		enforcer := lbemi.CoreV1.Policy().GetEnforce()
		// 用户ID
		uid, isExit := c.Get("id")
		if !isExit {
			response.Fail(c, response.InvalidToken)
			return
		}
		//if uid == "1" {
		//	c.Next()
		//	return
		//}
		p := c.Request.URL.Path
		m := c.Request.Method
		log.Logger.Infof("casbin-rule: %v - %v - %v", uid, p, m)
		ok, err := enforcer.Enforce(uid, p, m)
		if err != nil {
			log.Logger.Error(err)
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
