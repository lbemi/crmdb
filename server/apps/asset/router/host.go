package router

import (
	restfulspec "github.com/emicklei/go-restful-openapi/v2"
	"github.com/emicklei/go-restful/v3"
	"github.com/lbemi/lbemi/apps/asset/api"
	"github.com/lbemi/lbemi/apps/asset/api/form"
	"github.com/lbemi/lbemi/apps/asset/entity"
	form2 "github.com/lbemi/lbemi/pkg/common/entity"
	"github.com/lbemi/lbemi/pkg/rctx"
)

func HostRotes() *restful.WebService {
	ws := new(restful.WebService)
	ws.Path("/api/v1/hosts").Produces(restful.MIME_JSON)
	tags := []string{"hosts"}
	ws.Route(ws.GET("").To(func(request *restful.Request, response *restful.Response) {
		rctx.NewReqCtx(request, response).WithLog("hosts").WithHandle(api.ListHosts).
			Do()
	}).Doc("获取主机列表").Metadata(restfulspec.KeyOpenAPITags, tags).
		Param(ws.QueryParameter("page", "page").DataType("int")).
		Param(ws.QueryParameter("limit", "limit").DataType("int")).
		Param(ws.QueryParameter("groups", "分组id").DataType("string")).
		Param(ws.QueryParameter("ip", "ip地址,模糊查询").DataType("string")).
		Param(ws.QueryParameter("label", "标签,模糊查询").DataType("string")).
		Param(ws.QueryParameter("description", "描述信息,模糊查询").DataType("string")).
		Writes(form.PageHost{}).
		Returns(200, "success", form.PageHost{}))

	ws.Route(ws.GET("/{id}").To(func(request *restful.Request, response *restful.Response) {
		rctx.NewReqCtx(request, response).WithLog("hosts").WithHandle(api.GetHostById).
			Do()
	}).Doc("根据id获取主机").Metadata(restfulspec.KeyOpenAPITags, tags).
		Param(ws.PathParameter("id", "主机id").DataType("int")).
		Writes(entity.Host{}).
		Returns(200, "success", entity.Host{}))

	ws.Route(ws.GET("/groups").To(func(request *restful.Request, response *restful.Response) {
		rctx.NewReqCtx(request, response).WithLog("hosts").WithHandle(api.GetHostByGroups).
			Do()
	}).Doc("根据id获取主机").Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(form2.GroupVo{}).
		Writes(entity.Host{}).
		Returns(200, "success", entity.Host{}))

	ws.Route(ws.POST("").To(func(request *restful.Request, response *restful.Response) {
		rctx.NewReqCtx(request, response).WithLog("hosts").WithHandle(api.AddHost).Do()
	}).Doc("添加主机").Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(entity.HostReq{}).
		Returns(200, "success", nil))

	ws.Route(ws.PUT("").To(func(request *restful.Request, response *restful.Response) {
		rctx.NewReqCtx(request, response).WithLog("hosts").WithHandle(api.UpdateHost).Do()
	}).Doc("根据id修改主机").Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(entity.Host{}).
		Returns(200, "success", nil))

	ws.Route(ws.DELETE("/{id}").To(func(request *restful.Request, response *restful.Response) {
		rctx.NewReqCtx(request, response).WithLog("hosts").WithHandle(api.DeleteHost).Do()
	}).Doc("根据id删除主机").Metadata(restfulspec.KeyOpenAPITags, tags).
		Param(ws.PathParameter("id", "主机id").DataType("int")).
		Returns(200, "success", nil))

	ws.Route(ws.GET("/{id}/accounts").To(func(request *restful.Request, response *restful.Response) {
		rctx.NewReqCtx(request, response).WithLog("hosts").WithHandle(api.GetHostAccounts).Do()
	}).Doc("根据id获取主机账户").Metadata(restfulspec.KeyOpenAPITags, tags).
		Param(ws.PathParameter("id", "主机id").DataType("int")).
		Reads([]entity.Account{}).
		Returns(200, "success", []entity.Account{}))

	ws.Route(ws.GET("/{id}/shell/{account}").To(func(request *restful.Request, response *restful.Response) {
		rctx.NewReqCtx(request, response).WithNoRes().WithLog("websocket").WithHandle(api.WsShell).Do()
	}).Doc("webshell").Metadata(restfulspec.KeyOpenAPITags, tags).
		Param(ws.PathParameter("id", "主机id").DataType("int")).
		Param(ws.PathParameter("account", "用户id").DataType("int")).
		Returns(200, "success", nil))

	return ws
}
