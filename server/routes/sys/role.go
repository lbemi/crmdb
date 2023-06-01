package sys

import (
	restfulspec "github.com/emicklei/go-restful-openapi/v2"
	"github.com/emicklei/go-restful/v3"
	"github.com/gin-gonic/gin"
	"github.com/lbemi/lbemi/api/v1/sys"
	"github.com/lbemi/lbemi/pkg/model/form"
	model "github.com/lbemi/lbemi/pkg/model/sys"
	"github.com/lbemi/lbemi/pkg/rctx"
)

func NewRoleRouter(router *gin.RouterGroup) {
	//role := router.Group("/role")
	//{
	//role.POST("", sys.AddRole)          // 添加角色
	//role.PUT("/:id", sys.UpdateRole)    // 根据角色ID更新角色信息
	//role.DELETE("/:id", sys.DeleteRole) // 删除角色
	//
	//role.POST("/:id/menus", sys.SetRoleMenus)             // 根据角色ID设置角色权限
	//role.PUT("/:id/status/:status", sys.UpdateRoleStatus) // 修改角色状态
	//}
}

func RoleRoutes() *restful.WebService {
	role := new(restful.WebService)
	role.Path("/api/v1/roles").Produces(restful.MIME_JSON)
	tags := []string{"roles"}

	// 根据角色ID获取角色信息
	role.Route(role.GET("/{id}").To(
		rctx.NewReqCtx().WithLog("roles").WithHandle(sys.GetRole).Do()).
		Doc("根据角色ID获取角色信息").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Param(role.PathParameter("id", "角色id").DataType("string")).
		Reads(model.Role{}).
		Returns(200, "success", model.Role{}))

	//role.GET("", sys.ListRoles)                // 获取所有角色
	role.Route(role.GET("").To(
		rctx.NewReqCtx().WithLog("roles").WithHandle(sys.ListRoles).Do()).
		Doc("获取角色列表").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Param(role.QueryParameter("page", "页码").Required(false).DataType("string")).
		Param(role.QueryParameter("limit", "每页条数").Required(false).DataType("string")).
		Param(role.QueryParameter("name", "角色名称-模糊查询").Required(false).DataType("string")).
		Reads(form.PageResult{}).
		Returns(200, "success", form.PageResult{}))

	//role.GET("/:id/menus", sys.GetMenusByRole) // 根据角色ID获取角色权限
	role.Route(role.GET("{id}/menus").To(
		rctx.NewReqCtx().WithLog("roles").WithHandle(sys.GetMenusByRole).Do()).
		Doc("根据角色ID获取角色权限").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Param(role.PathParameter("id", "角色id").DataType("string")).
		Param(role.QueryParameter("limit", "每页条数").Required(false).DataType("string")).
		Reads(form.PageResult{}).
		Returns(200, "success", form.PageResult{}))

	return role
}
