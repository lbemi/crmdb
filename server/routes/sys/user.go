package sys

import (
	restfulspec "github.com/emicklei/go-restful-openapi/v2"
	"github.com/emicklei/go-restful/v3"
	"github.com/gin-gonic/gin"
	"github.com/lbemi/lbemi/api/v1/sys"
	"github.com/lbemi/lbemi/pkg/model/form"
	"github.com/lbemi/lbemi/pkg/rctx"
	"github.com/lbemi/lbemi/pkg/restfulx"
)

func NewUserRouter(router *gin.RouterGroup) {

	user := router.Group("/user")
	{
		// 用户退出登录
		//user.POST("/logout", sys.Logout)
		// 注册
		user.POST("/register", sys.Register)
		// 根据ID获取用户信息
		user.GET("/:id", sys.GetUserInfoById)
		// 获取用户列表
		user.GET("", func(c *gin.Context) {
			//rctx.NewReqCtx(c).Handle(sys.GetUserList)
		})
		// 删除
		user.DELETE("/:id", sys.DeleteUserByUserId)
		// 更新
		user.PUT("/:id", sys.UpdateUser)

		user.GET("/:id/roles", sys.GetUserRoles)  // 查询当前用户角色
		user.POST("/:id/roles", sys.SetUserRoles) // 根据用户id分配角色

		// 根据菜单ID获取当前用户的权限
		user.GET("/permissions", sys.GetButtonsByCurrentUser)
		// 根据用户ID获取用户的菜单
		user.GET("/menus", sys.GetLeftMenusByCurrentUser)
		//修改用户状态
		user.PUT("/:id/status/:status", sys.UpdateUserStatus)
	}

}
func UserRouter() *restful.WebService {
	ws := new(restful.WebService)
	ws.Path("/api/v1/users").Produces(restful.MIME_JSON)
	tags := []string{"users"}
	// 获取图片验证码
	ws.Route(ws.GET("/captcha").To(
		rctx.NewReqCtx().
			WithToken(false).
			WithCasbin(false).
			WithHandle(sys.GetCaptcha).Do()).
		Doc("captcha").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(sys.CaptchaInfo{}).
		Returns(200, "success", sys.CaptchaInfo{}).
		Returns(500, restfulx.ServerErr.Error(), restfulx.ServerErr))

	// 用户退出登录
	ws.Route(ws.POST("/logout").To(
		rctx.NewReqCtx().WithToken(true).WithCasbin(false).WithLog("logout").WithHandle(sys.Logout).Do()).
		Doc("logout").
		Metadata(restfulspec.KeyOpenAPITags, tags))

	ws.Route(ws.POST("/login").To(rctx.NewReqCtx().
		WithToken(false).
		WithCasbin(false).
		WithLog("login").
		WithHandle(sys.Login).
		Do()).
		Doc("login").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(form.UserLoginForm{}).Writes(form.LoginResp{}).
		Returns(200, "success", form.LoginResp{}).
		Returns(1001, restfulx.UserDeny.Error(), restfulx.UserDeny).
		Returns(1002, restfulx.PasswdWrong.Error(), restfulx.PasswdWrong).
		Returns(4001, restfulx.TokenExpire.Error(), restfulx.TokenExpire).
		Returns(1002, restfulx.TokenInvalid.Error(), restfulx.TokenInvalid))

	//// 注册
	//user.POST("/register", sys.Register)
	//// 根据ID获取用户信息
	//user.GET("/:id", sys.GetUserInfoById)
	//// 获取用户列表
	//user.GET("", func(c *gin.Context) {
	//	rctx.NewReqCtx(c).Handle(sys.GetUserList)
	//})
	//// 删除
	//user.DELETE("/:id", sys.DeleteUserByUserId)
	//// 更新
	//user.PUT("/:id", sys.UpdateUser)
	//
	//user.GET("/:id/roles", sys.GetUserRoles)  // 查询当前用户角色
	//user.POST("/:id/roles", sys.SetUserRoles) // 根据用户id分配角色
	//
	//// 根据菜单ID获取当前用户的权限
	//user.GET("/permissions", sys.GetButtonsByCurrentUser)
	//// 根据用户ID获取用户的菜单
	//user.GET("/menus", sys.GetLeftMenusByCurrentUser)
	////修改用户状态
	//user.PUT("/:id/status/:status", sys.UpdateUserStatus)

	return ws
}
