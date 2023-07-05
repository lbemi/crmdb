package asset

import (
	restfulspec "github.com/emicklei/go-restful-openapi/v2"
	"github.com/emicklei/go-restful/v3"
	"github.com/lbemi/lbemi/api/asset"
	form2 "github.com/lbemi/lbemi/api/form"
	model "github.com/lbemi/lbemi/pkg/model/asset"
	"github.com/lbemi/lbemi/pkg/model/form"
	"github.com/lbemi/lbemi/pkg/rctx"
)

func HostRotes() *restful.WebService {
	ws := new(restful.WebService)
	ws.Path("/api/v1/hosts").Produces(restful.MIME_JSON)
	tags := []string{"hosts"}
	ws.Route(ws.GET("").To(func(request *restful.Request, response *restful.Response) {
		rctx.NewReqCtx(request, response).WithLog("hosts").WithHandle(asset.ListHosts).
			Do()
	}).Doc("获取主机列表").Metadata(restfulspec.KeyOpenAPITags, tags).
		Param(ws.QueryParameter("page", "page").DataType("int")).
		Param(ws.QueryParameter("limit", "limit").DataType("int")).
		Param(ws.QueryParameter("groups", "分组id").DataType("string")).
		Writes(form.PageHost{}).
		Returns(200, "success", form.PageHost{}))

	ws.Route(ws.GET("/{id}").To(func(request *restful.Request, response *restful.Response) {
		rctx.NewReqCtx(request, response).WithLog("hosts").WithHandle(asset.GetHostById).
			Do()
	}).Doc("根据id获取主机").Metadata(restfulspec.KeyOpenAPITags, tags).
		Param(ws.PathParameter("id", "主机id").DataType("int")).
		Writes(model.Host{}).
		Returns(200, "success", model.Host{}))

	ws.Route(ws.GET("/groups").To(func(request *restful.Request, response *restful.Response) {
		rctx.NewReqCtx(request, response).WithLog("hosts").WithHandle(asset.GetHostByGroups).
			Do()
	}).Doc("根据id获取主机").Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(form2.GroupVo{}).
		Writes(model.Host{}).
		Returns(200, "success", model.Host{}))

	ws.Route(ws.POST("").To(func(request *restful.Request, response *restful.Response) {
		rctx.NewReqCtx(request, response).WithLog("hosts").WithHandle(asset.AddHost).Do()
	}).Doc("添加主机").Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(model.HostReq{}).
		Returns(200, "success", nil))

	ws.Route(ws.PUT("").To(func(request *restful.Request, response *restful.Response) {
		rctx.NewReqCtx(request, response).WithLog("hosts").WithHandle(asset.UpdateHost).Do()
	}).Doc("根据id修改主机").Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(model.Host{}).
		Returns(200, "success", nil))

	ws.Route(ws.DELETE("/{id}").To(func(request *restful.Request, response *restful.Response) {
		rctx.NewReqCtx(request, response).WithLog("hosts").WithHandle(asset.DeleteHost).Do()
	}).Doc("根据id删除主机").Metadata(restfulspec.KeyOpenAPITags, tags).
		Param(ws.PathParameter("id", "主机id").DataType("int")).
		Returns(200, "success", nil))

	ws.Route(ws.GET("/{id}/accounts").To(func(request *restful.Request, response *restful.Response) {
		rctx.NewReqCtx(request, response).WithLog("hosts").WithHandle(asset.GetHostAccounts).Do()
	}).Doc("根据id获取主机账户").Metadata(restfulspec.KeyOpenAPITags, tags).
		Param(ws.PathParameter("id", "主机id").DataType("int")).
		Reads([]model.Account{}).
		Returns(200, "success", []model.Account{}))

	ws.Route(ws.GET("/{id}/shell/{account}").To(func(request *restful.Request, response *restful.Response) {
		rctx.NewReqCtx(request, response).WithNoRes().WithLog("websocket").WithHandle(asset.WsShell).Do()
	}).Doc("webshell").Metadata(restfulspec.KeyOpenAPITags, tags).
		Param(ws.PathParameter("id", "主机id").DataType("int")).
		Param(ws.PathParameter("account", "用户id").DataType("int")).
		Returns(200, "success", nil))

	return ws
}
