package middleware

import (
	"github.com/gin-gonic/gin"
)

func CasbinMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		//// 用户ID
		//uid, isExit := c.Get("id")
		//if !isExit {
		//	response.Fail(c, response.InvalidToken)
		//	return
		//}
		//role := sys.Role{}
		//err := role.GetRoleByUId(uid.(string))
		//if err != nil {
		//	global.App.Log.Error(err.Error())
		//	response.Fail(c, response.StatusInternalServerError)
		//	c.Abort()
		//	return
		//}
		////if uid == "1" {
		////	c.Next()
		////	return
		////}
		//p := c.Request.URL.Path
		//m := c.Request.Method
		//fmt.Printf("------------url: %v, method: %v, UID: %v,  Role: %v \n", p, m, uid, role.Name)
		////ok, err := global.App.Enforcer.Enforce(role.Name, p, m)
		//ok, err := global.App.Enforcer.Enforce(uid, p, m)
		//if err != nil {
		//	global.App.Log.Fatal(err.Error())
		//	response.Fail(c, response.StatusInternalServerError)
		//	c.Abort()
		//	return
		//}
		//if !ok {
		//	response.Fail(c, response.NoPermission)
		//	c.Abort()
		//	return
		//}
		//c.Next()
	}
}
