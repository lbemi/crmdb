package router

import (
	restfulspec "github.com/emicklei/go-restful-openapi/v2"
	"github.com/emicklei/go-restful/v3"
	"github.com/lbemi/lbemi/apps/kubernetes/api"
	entity2 "github.com/lbemi/lbemi/apps/kubernetes/entity"
	"github.com/lbemi/lbemi/pkg/common/entity"
	v1 "k8s.io/api/core/v1"

	"github.com/lbemi/lbemi/pkg/rctx"
)

func KubernetesServiceRoutes() *restful.WebService {
	ws := new(restful.WebService)
	ws.Path("/api/v1/services").Produces(restful.MIME_JSON)
	tags := []string{"kubernetes-Service"}

	ws.Route(ws.GET("/namespaces/{namespace}").To(func(request *restful.Request, response *restful.Response) {
		rctx.NewReqCtx(request, response).WithLog("Service").
			WithHandle(api.ListServices).Do()
	}).Doc("获取Service列表").Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(&entity.PageResult{}).
		Param(ws.QueryParameter("cloud", "集群名称").Required(true).DataType("string")).
		Param(ws.QueryParameter("page", "页码").Required(false).DataType("string")).
		Param(ws.QueryParameter("limit", "每页条数").Required(false).DataType("string")).
		Param(ws.QueryParameter("name", "过滤名称，支持模糊查询").DataType("string")).
		Param(ws.QueryParameter("label", "过滤标签，支持模糊查询").DataType("string")).
		Param(ws.PathParameter("namespace", "命名空间").Required(true).DataType("string")).
		Returns(200, "success", &entity.PageResult{}))

	ws.Route(ws.GET("/namespaces/{namespace}/{name}").To(func(request *restful.Request, response *restful.Response) {
		rctx.NewReqCtx(request, response).WithLog("Service").
			WithHandle(api.GetService).Do()
	}).Doc("获取Service信息").Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(v1.Service{}).
		Param(ws.QueryParameter("cloud", "集群名称").Required(true).DataType("string")).
		Param(ws.PathParameter("namespace", "命名空间").Required(true).DataType("string")).
		Param(ws.PathParameter("name", "Service名称").Required(true).DataType("string")).
		Returns(200, "success", v1.Service{}))

	ws.Route(ws.GET("/namespaces/{namespace}/workload/{name}").To(func(request *restful.Request, response *restful.Response) {
		rctx.NewReqCtx(request, response).WithLog("Service").
			WithHandle(api.GetServiceWorkLoad).Do()
	}).Doc("获取Service的所有负载信息").Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(entity2.ServiceWorkLoad{}).
		Param(ws.QueryParameter("cloud", "集群名称").Required(true).DataType("string")).
		Param(ws.PathParameter("namespace", "命名空间").Required(true).DataType("string")).
		Param(ws.PathParameter("name", "Service名称").Required(true).DataType("string")).
		Returns(200, "success", entity2.ServiceWorkLoad{}))

	ws.Route(ws.POST("").To(func(request *restful.Request, response *restful.Response) {
		rctx.NewReqCtx(request, response).WithLog("Service").
			WithHandle(api.CreateService).Do()
	}).Doc("创建Service").Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(entity.PageResult{}).
		Reads(v1.Service{}).
		Param(ws.QueryParameter("cloud", "集群名称").Required(true).DataType("string")).
		Returns(200, "success", v1.Service{}))

	ws.Route(ws.PUT("").To(func(request *restful.Request, response *restful.Response) {
		rctx.NewReqCtx(request, response).WithLog("Service").
			WithHandle(api.UpdateService).Do()
	}).Doc("修改Service").Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(entity.PageResult{}).
		Reads(v1.Service{}).
		Param(ws.QueryParameter("cloud", "集群名称").Required(true).DataType("string")).
		Returns(200, "success", v1.Service{}))

	ws.Route(ws.DELETE("/namespaces/{namespace}/{name}").To(func(request *restful.Request, response *restful.Response) {
		rctx.NewReqCtx(request, response).WithLog("Service").
			WithHandle(api.DeleteService).Do()
	}).Doc("删除Service").Metadata(restfulspec.KeyOpenAPITags, tags).
		Param(ws.PathParameter("namespace", "命名空间").Required(true).DataType("string")).
		Param(ws.PathParameter("name", "Service名称").Required(true).DataType("string")).
		Param(ws.QueryParameter("cloud", "集群名称").Required(true).DataType("string")).
		Returns(200, "success", nil))

	return ws
}
