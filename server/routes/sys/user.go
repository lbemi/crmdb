package sys

import (
	restfulspec "github.com/emicklei/go-restful-openapi/v2"
	"github.com/emicklei/go-restful/v3"
	"github.com/gin-gonic/gin"
	"github.com/lbemi/lbemi/api/v1/sys"
	"github.com/lbemi/lbemi/pkg/model/form"
	model "github.com/lbemi/lbemi/pkg/model/sys"
	"github.com/lbemi/lbemi/pkg/rctx"
	"github.com/lbemi/lbemi/pkg/restfulx"
)

func NewUserRouter(router *gin.RouterGroup) {

	user := router.Group("/user")
	{
		// 根据ID获取用户信息
		user.GET("/:id", sys.GetUserInfoById)

		// 更新
		user.PUT("/:id", sys.UpdateUser)

		user.GET("/:id/roles", sys.GetUserRoles)  // 查询当前用户角色
		user.POST("/:id/roles", sys.SetUserRoles) // 根据用户id分配角色

		// 根据菜单ID获取当前用户的权限
		user.GET("/permissions", sys.GetButtonsByCurrentUser)
		// 根据用户ID获取用户的菜单
		//user.GET("/menus", sys.GetLeftMenusByCurrentUser)
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
		rctx.NewReqCtx().WithToken(false).WithCasbin(false).WithHandle(sys.GetCaptcha).Do()).
		Doc("获取验证码").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(sys.CaptchaInfo{}).
		Returns(200, "success", sys.CaptchaInfo{}).
		Returns(500, restfulx.ServerErr.Error(), restfulx.ServerErr))

	// 用户退出登录
	ws.Route(ws.POST("/logout").To(
		rctx.NewReqCtx().WithToken(true).WithCasbin(false).WithLog("logout").WithHandle(sys.Logout).Do()).
		Doc("登出").
		Metadata(restfulspec.KeyOpenAPITags, tags))
	// 登录
	ws.Route(ws.POST("/login").To(rctx.NewReqCtx().
		WithToken(false).
		WithCasbin(false).
		WithLog("login").
		WithHandle(sys.Login).
		Do()).
		Doc("登录").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(form.UserLoginForm{}).Writes(form.LoginResp{}).
		Returns(200, "success", form.LoginResp{}).
		Returns(1001, restfulx.UserDeny.Error(), restfulx.UserDeny).
		Returns(1002, restfulx.PasswdWrong.Error(), restfulx.PasswdWrong).
		Returns(4001, restfulx.TokenExpire.Error(), restfulx.TokenExpire).
		Returns(1002, restfulx.TokenInvalid.Error(), restfulx.TokenInvalid))
	// 根据用户ID获取用户的菜单
	ws.Route(ws.GET("/menus").To(
		rctx.NewReqCtx().WithLog("users").WithHandle(sys.GetLeftMenusByCurrentUser).Do()).
		Doc("获取当前登录用户菜单").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Returns(200, "success", form.UserPermissionResp{}))

	// 获取用户列表
	ws.Route(ws.GET("").To(
		rctx.NewReqCtx().WithLog("users").WithHandle(sys.GetUserList).Do()).
		Doc("获取用户列表").
		Param(ws.QueryParameter("page", "page").DataType("int")).
		Param(ws.QueryParameter("limit", "limit").DataType("int")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Returns(200, "success", form.PageUser{}))

	// 注册
	ws.Route(ws.POST("/register").To(rctx.NewReqCtx().
		WithLog("users").
		WithHandle(sys.Register).
		Do()).
		Doc("登录").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(form.RegisterUserForm{}).
		Returns(200, "success", nil))

	// 删除
	ws.Route(ws.DELETE("/{id}").To(rctx.NewReqCtx().
		WithLog("users").
		WithHandle(sys.DeleteUserByUserId).
		Do()).
		Doc("删除用户").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Param(ws.PathParameter("id", "用户id").DataType("string")).
		Returns(200, "success", nil))

	// 根据ID获取用户信息
	ws.Route(ws.DELETE("/{id}").To(rctx.NewReqCtx().
		WithLog("users").
		WithHandle(sys.GetUserInfoById).
		Do()).
		Doc("根据用户ID获取用户信息").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Param(ws.PathParameter("id", "用户id").DataType("string")).
		Writes(model.User{}).
		Returns(200, "success", nil))

	// 更新
	user.PUT("/:id", sys.UpdateUser)

	user.GET("/:id/roles", sys.GetUserRoles)  // 查询当前用户角色
	user.POST("/:id/roles", sys.SetUserRoles) // 根据用户id分配角色

	// 根据菜单ID获取当前用户的权限
	user.GET("/permissions", sys.GetButtonsByCurrentUser)
	// 根据用户ID获取用户的菜单
	//user.GET("/menus", sys.GetLeftMenusByCurrentUser)
	//修改用户状态
	user.PUT("/:id/status/:status", sys.UpdateUserStatus)

	return ws
}
