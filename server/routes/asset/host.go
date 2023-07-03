package asset

import (
	restfulspec "github.com/emicklei/go-restful-openapi/v2"
	"github.com/emicklei/go-restful/v3"
	"github.com/lbemi/lbemi/api/asset"
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
		Writes(form.PageHost{}).
		Returns(200, "success", form.PageHost{}))

	ws.Route(ws.GET("/{id}").To(func(request *restful.Request, response *restful.Response) {
		rctx.NewReqCtx(request, response).WithLog("hosts").WithHandle(asset.GetHostById).
			Do()
	}).Doc("根据id获取主机").Metadata(restfulspec.KeyOpenAPITags, tags).
		Param(ws.PathParameter("id", "主机id").DataType("int")).
		Writes(model.Host{}).
		Returns(200, "success", model.Host{}))

	ws.Route(ws.POST("").To(func(request *restful.Request, response *restful.Response) {
		rctx.NewReqCtx(request, response).WithLog("hosts").WithHandle(asset.AddHost).Do()
	}).Doc("添加主机").Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(model.HostReq{}).
		Returns(200, "success", nil))

	ws.Route(ws.PUT("/{id}").To(func(request *restful.Request, response *restful.Response) {
		rctx.NewReqCtx(request, response).WithLog("hosts").WithHandle(asset.UpdateHost).Do()
	}).Doc("根据id修改主机").Metadata(restfulspec.KeyOpenAPITags, tags).
		Param(ws.PathParameter("id", "主机id").DataType("int")).
		Reads(model.HostReq{}).
		Returns(200, "success", nil))

	ws.Route(ws.DELETE("/{id}").To(func(request *restful.Request, response *restful.Response) {
		rctx.NewReqCtx(request, response).WithLog("hosts").WithHandle(asset.DeleteHost).Do()
	}).Doc("根据id删除主机").Metadata(restfulspec.KeyOpenAPITags, tags).
		Param(ws.PathParameter("id", "主机id").DataType("int")).
		Returns(200, "success", nil))
	return ws
}
