package cloud

import (
	restfulspec "github.com/emicklei/go-restful-openapi/v2"
	"github.com/emicklei/go-restful/v3"
	v1 "k8s.io/api/core/v1"

	"github.com/lbemi/lbemi/api/v1/cloud"
	"github.com/lbemi/lbemi/pkg/model/form"
	"github.com/lbemi/lbemi/pkg/rctx"
)

func KubernetesNamespaceRoutes() *restful.WebService {
	ws := new(restful.WebService)
	ws.Path("/api/v1/namespaces")
	tags := []string{"kubernetes-namespace"}

	ws.Route(ws.GET("").To(func(request *restful.Request, response *restful.Response) {
		rctx.NewReqCtx(request, response).WithLog("namespace").
			WithHandle(cloud.ListNamespace).Do()
	}).Doc("获取Namespace列表").Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(&form.PageResult{}).
		Param(ws.QueryParameter("cloud", "集群名称").Required(true).DataType("string")).
		Param(ws.QueryParameter("page", "页码").Required(false).DataType("string")).
		Param(ws.QueryParameter("limit", "每页条数").Required(false).DataType("string")).
		Param(ws.QueryParameter("name", "过滤名称，支持模糊查询").DataType("string")).
		Param(ws.QueryParameter("label", "过滤标签，支持模糊查询").DataType("string")).
		Returns(200, "success", &form.PageResult{}))

	ws.Route(ws.GET("/{name}").To(func(request *restful.Request, response *restful.Response) {
		rctx.NewReqCtx(request, response).WithLog("namespace").
			WithHandle(cloud.GetNamespace).Do()
	}).Doc("获取Namespace信息").Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(v1.Namespace{}).
		Param(ws.QueryParameter("cloud", "集群名称").Required(true).DataType("string")).
		Param(ws.PathParameter("name", "Namespace名称").Required(true).DataType("string")).
		Returns(200, "success", v1.Namespace{}))

	ws.Route(ws.POST("").To(func(request *restful.Request, response *restful.Response) {
		rctx.NewReqCtx(request, response).WithLog("namespace").
			WithHandle(cloud.CreateNamespace).Do()
	}).Doc("创建Namespace").Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(form.PageResult{}).
		Reads(v1.Namespace{}).
		Param(ws.QueryParameter("cloud", "集群名称").Required(true).DataType("string")).
		Returns(200, "success", v1.Namespace{}))

	ws.Route(ws.PUT("").To(func(request *restful.Request, response *restful.Response) {
		rctx.NewReqCtx(request, response).WithLog("namespace").
			WithHandle(cloud.UpdateNamespace).Do()
	}).Doc("修改Namespace").Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(form.PageResult{}).
		Reads(v1.Namespace{}).
		Param(ws.QueryParameter("cloud", "集群名称").Required(true).DataType("string")).
		Returns(200, "success", v1.Namespace{}))

	ws.Route(ws.DELETE("/{name}").To(func(request *restful.Request, response *restful.Response) {
		rctx.NewReqCtx(request, response).WithLog("namespace").
			WithHandle(cloud.DeleteNamespace).Do()
	}).Doc("删除Namespace").Metadata(restfulspec.KeyOpenAPITags, tags).
		Param(ws.QueryParameter("cloud", "集群名称").Required(true).DataType("string")).
		Param(ws.PathParameter("name", "namespace名称").Required(true).DataType("string")).
		Returns(200, "success", nil))

	return ws
}
