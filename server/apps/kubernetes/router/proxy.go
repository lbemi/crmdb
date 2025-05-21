package router

import (
	restfulspec "github.com/emicklei/go-restful-openapi/v2"
	"github.com/emicklei/go-restful/v3"
	"github.com/lbemi/lbemi/apps/kubernetes/api"
	"github.com/lbemi/lbemi/pkg/rctx"
)

// KubernetesProxyRoutes kubernetes proxy
// TODO: optimize methods functions to one
func KubernetesProxyRoutes() *restful.WebService {
	ws := new(restful.WebService)
	ws.Path("/api/v1/proxy").Produces(restful.MIME_JSON)
	tags := []string{"kubernetes-proxy"}

	// GET 方法
	ws.Route(ws.GET("/{clusterName}/{*}").To(func(request *restful.Request, response *restful.Response) {
		rctx.NewReqCtx(request, response).WithLog("proxy").
			WithHandle(api.Proxy).WithNoRes().Do()
	}).Doc("代理请求").Metadata(restfulspec.KeyOpenAPITags, tags).
		Param(ws.PathParameter("clusterName", "集群名称").Required(true).DataType("string")).
		Param(ws.PathParameter("*", "代理路径").Required(true).DataType("string")).
		Returns(200, "success", nil))

	// PUT 方法
	ws.Route(ws.PUT("/{clusterName}/{*}").To(func(request *restful.Request, response *restful.Response) {
		rctx.NewReqCtx(request, response).WithLog("proxy").
			WithHandle(api.Proxy).WithNoRes().Do()
	}).Doc("代理请求").Metadata(restfulspec.KeyOpenAPITags, tags).
		Param(ws.PathParameter("clusterName", "集群名称").Required(true).DataType("string")).
		Param(ws.PathParameter("*", "代理路径").Required(true).DataType("string")).
		Returns(200, "success", nil))

	ws.Route(ws.DELETE("/{clusterName}/{*}").To(func(request *restful.Request, response *restful.Response) {
		rctx.NewReqCtx(request, response).WithLog("proxy").
			WithHandle(api.Proxy).WithNoRes().Do()
	}).Doc("代理请求").Metadata(restfulspec.KeyOpenAPITags, tags).
		Param(ws.PathParameter("clusterName", "集群名称").Required(true).DataType("string")).
		Param(ws.PathParameter("*", "代理路径").Required(true).DataType("string")).
		Returns(200, "success", nil))

	ws.Route(ws.POST("/{clusterName}/{*}").To(func(request *restful.Request, response *restful.Response) {
		rctx.NewReqCtx(request, response).WithLog("proxy").
			WithHandle(api.Proxy).WithNoRes().Do()
	}).Doc("代理请求").Metadata(restfulspec.KeyOpenAPITags, tags).
		Param(ws.PathParameter("clusterName", "集群名称").Required(true).DataType("string")).
		Param(ws.PathParameter("*", "代理路径").Required(true).DataType("string")).
		Returns(200, "success", nil))

	ws.Route(ws.PATCH("/{clusterName}/{*}").To(func(request *restful.Request, response *restful.Response) {
		rctx.NewReqCtx(request, response).WithLog("proxy").
			WithHandle(api.Proxy).WithNoRes().Do()
	}).Doc("代理请求").Metadata(restfulspec.KeyOpenAPITags, tags).
		Param(ws.PathParameter("clusterName", "集群名称").Required(true).DataType("string")).
		Param(ws.PathParameter("*", "代理路径").Required(true).DataType("string")).
		Returns(200, "success", nil))

	return ws
}
