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

func NewMenuRouter(router *gin.RouterGroup) {

}

func MenuRoutes() *restful.WebService {
	menu := new(restful.WebService)
	menu.Path("/api/v1/menus").Produces(restful.MIME_JSON)
	tags := []string{"menus"}

	// 获取菜单/API/按钮列表
	menu.Route(menu.GET("").To(func(request *restful.Request, response *restful.Response) {
		rctx.NewReqCtx(request, response).WithLog("menus").WithHandle(sys.ListMenus).Do()
	}).
		Doc("获取菜单/API/按钮列表").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Param(menu.QueryParameter("page", "页码").Required(false).DataType("string")).
		Param(menu.QueryParameter("limit", "每页条数").Required(false).DataType("string")).
		Param(menu.QueryParameter("menuType", "菜单类型： 1:菜单，2:按钮，3:API, 默认为 1,2,3").Required(false).DataType("string")).
		Param(menu.QueryParameter("isTree", "是否返回树形结构，true / false，默认是 true").Required(false).DataType("string")).
		Param(menu.QueryParameter("status", "过滤状态").DataType("int")).
		Param(menu.QueryParameter("name", "过滤描述，支持模糊查询").DataType("string")).
		Reads(form.PageMenu{}).
		Returns(200, "success", form.PageMenu{}))

	// 根据菜单/API/按钮 ID获取详细信息
	menu.Route(menu.GET("/{id}").To(func(request *restful.Request, response *restful.Response) {
		rctx.NewReqCtx(request, response).WithLog("menus").WithHandle(sys.GetMenu).Do()
	}).
		Doc("根据菜单/API/按钮 ID获取详细信息").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Param(menu.PathParameter("id", "菜单/API/按钮 id").DataType("string")).
		Reads(model.Menu{}).
		Returns(200, "success", model.Menu{}))

	// 添加菜单/API/按钮
	menu.Route(menu.POST("").To(func(request *restful.Request, response *restful.Response) {
		rctx.NewReqCtx(request, response).WithLog("menus").WithHandle(sys.AddMenu).Do()
	}).
		Doc("添加菜单/API/按钮").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(form.MenusReq{}).
		Returns(200, "success", nil))
	// 根据角色ID更新角色信息
	menu.Route(menu.PUT("/{id}").To(func(request *restful.Request, response *restful.Response) {
		rctx.NewReqCtx(request, response).WithLog("menus").WithHandle(sys.UpdateMenu).Do()
	}).
		Doc("根据角色ID更新角色信息").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Param(menu.PathParameter("id", "菜单/API/按钮id").DataType("string")).
		Writes(form.UpdateRoleReq{}).
		Returns(200, "success", nil))

	// 修改菜单/API/按钮状态
	menu.Route(menu.PUT("/{id}/status/{status}").To(func(request *restful.Request, response *restful.Response) {
		rctx.NewReqCtx(request, response).WithLog("menus").WithHandle(sys.UpdateMenuStatus).Do()
	}).
		Doc("根据菜单/API/按钮ID更新信息").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Param(menu.PathParameter("id", "菜单/API/按钮id").DataType("string")).
		Param(menu.PathParameter("status", "状态： 0 为禁用，1为启用").DataType("string")).
		Returns(200, "success", nil))

	// 删除菜单/API/按钮
	menu.Route(menu.DELETE("/{id}").To(func(request *restful.Request, response *restful.Response) {
		rctx.NewReqCtx(request, response).WithLog("menus").WithHandle(sys.DeleteMenu).Do()
	}).
		Doc("删除菜单/API/按钮").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Param(menu.PathParameter("id", "菜单/API/按钮id").DataType("string")).
		Returns(200, "success", nil))

	return menu
}
