package router

import (
	restfulspec "github.com/emicklei/go-restful-openapi/v2"
	"github.com/emicklei/go-restful/v3"
	"github.com/gin-gonic/gin"
	"github.com/lbemi/lbemi/apps/system/api"
	form2 "github.com/lbemi/lbemi/apps/system/api/form"
	model "github.com/lbemi/lbemi/apps/system/entity"
	"github.com/lbemi/lbemi/pkg/common/entity"

	"github.com/lbemi/lbemi/pkg/rctx"
)

func NewRoleRouter(router *gin.RouterGroup) {
}

func RoleRoutes() *restful.WebService {
	role := new(restful.WebService)
	role.Path("/api/v1/roles").Produces(restful.MIME_JSON)
	tags := []string{"roles"}

	// 根据角色ID获取角色信息
	role.Route(role.GET("/{id}").To(func(request *restful.Request, response *restful.Response) {
		rctx.NewReqCtx(request, response).WithLog("roles").WithHandle(api.GetRole).Do()
	}).
		Doc("根据角色ID获取角色信息").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Param(role.PathParameter("id", "角色id").DataType("string")).
		Reads(model.Role{}).
		Returns(200, "success", model.Role{}))

	// 获取所有角色
	role.Route(role.GET("").To(func(request *restful.Request, response *restful.Response) {
		rctx.NewReqCtx(request, response).WithLog("roles").WithHandle(api.ListRoles).Do()
	}).
		Doc("获取角色列表").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Param(role.QueryParameter("page", "页码").Required(false).DataType("string")).
		Param(role.QueryParameter("limit", "每页条数").Required(false).DataType("string")).
		Param(role.QueryParameter("name", "角色名称-模糊查询").Required(false).DataType("string")).
		Param(role.QueryParameter("status", "过滤状态").DataType("int")).
		Reads(entity.PageResult{}).
		Returns(200, "success", entity.PageResult{}))

	// 根据角色ID获取角色权限
	role.Route(role.GET("{id}/menus").To(func(request *restful.Request, response *restful.Response) {
		rctx.NewReqCtx(request, response).WithLog("roles").WithHandle(api.GetMenusByRole).Do()
	}).
		Doc("根据角色ID获取角色权限").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Param(role.PathParameter("id", "角色id").DataType("string")).
		Param(role.QueryParameter("limit", "每页条数").Required(false).DataType("string")).
		Reads(entity.PageResult{}).
		Returns(200, "success", entity.PageResult{}))

	// 添加角色
	role.Route(role.POST("").To(func(request *restful.Request, response *restful.Response) {
		rctx.NewReqCtx(request, response).WithLog("roles").WithHandle(api.AddRole).Do()
	}).
		Doc("添加角色").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(form2.RoleReq{}).
		Returns(200, "success", nil))

	// 根据角色ID设置角色权限
	role.Route(role.POST("/{id}/menus").To(func(request *restful.Request, response *restful.Response) {
		rctx.NewReqCtx(request, response).WithLog("roles").WithHandle(api.SetRoleMenus).Do()
	}).
		Doc("根据角色ID设置角色权限").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Param(role.PathParameter("id", "角色id").DataType("string")).
		Writes(form2.Menus{}).
		Returns(200, "success", nil))

	// 根据角色ID更新角色信息
	role.Route(role.PUT("/{id}").To(func(request *restful.Request, response *restful.Response) {
		rctx.NewReqCtx(request, response).WithLog("roles").WithHandle(api.UpdateRole).Do()
	}).
		Doc("根据角色ID更新角色信息").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Param(role.PathParameter("id", "角色id").DataType("string")).
		Writes(form2.UpdateRoleReq{}).
		Returns(200, "success", nil))

	// 修改角色状态
	role.Route(role.PUT("/{id}/status/{status}").To(func(request *restful.Request, response *restful.Response) {
		rctx.NewReqCtx(request, response).WithLog("roles").WithHandle(api.UpdateRoleStatus).Do()
	}).
		Doc("根据角色ID更新角色信息").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Param(role.PathParameter("id", "角色id").DataType("string")).
		Param(role.PathParameter("status", "状态： 0 为禁用，1为启用").DataType("string")).
		Returns(200, "success", nil))

	// 删除角色
	role.Route(role.DELETE("/{id}").To(func(request *restful.Request, response *restful.Response) {
		rctx.NewReqCtx(request, response).WithLog("roles").WithHandle(api.DeleteRole).Do()
	}).
		Doc("删除角色").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Param(role.PathParameter("id", "角色id").DataType("string")).
		Returns(200, "success", nil))

	return role
}
