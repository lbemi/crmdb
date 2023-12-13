package router

import (
	restfulspec "github.com/emicklei/go-restful-openapi/v2"
	"github.com/emicklei/go-restful/v3"
	"github.com/lbemi/lbemi/apps/tekton/api"
	"github.com/lbemi/lbemi/pkg/common/entity"
	v1 "k8s.io/api/core/v1"

	"github.com/lbemi/lbemi/pkg/rctx"
)

func TektonPipelinesRoutes() *restful.WebService {
	ws := new(restful.WebService)
	ws.Path("/api/v1/pipelines").Produces(restful.MIME_JSON)
	tags := []string{"tekton-pipeline"}

	ws.Route(ws.GET("/namespaces/{namespace}").To(func(request *restful.Request, response *restful.Response) {
		rctx.NewReqCtx(request, response).WithLog("pipeline").
			WithHandle(api.ListPipelines).Do()
	}).Doc("获取pipeline列表").Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(&entity.PageResult{}).
		Param(ws.QueryParameter("cloud", "集群名称").Required(true).DataType("string")).
		Param(ws.PathParameter("namespace", "命名空间").Required(true).DataType("string")).
		Param(ws.QueryParameter("page", "页码").Required(false).DataType("string")).
		Param(ws.QueryParameter("limit", "每页条数").Required(false).DataType("string")).
		Param(ws.QueryParameter("name", "过滤名称，支持模糊查询").DataType("string")).
		Param(ws.QueryParameter("label", "过滤标签，支持模糊查询").DataType("string")).
		Returns(200, "success", &entity.PageResult{}))

	ws.Route(ws.GET("/namespaces/{namespace}/{name}").To(func(request *restful.Request, response *restful.Response) {
		rctx.NewReqCtx(request, response).WithLog("pipeline").
			WithHandle(api.GetPipeline).Do()
	}).Doc("获取pipeline信息").Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(v1.ConfigMap{}).
		Param(ws.QueryParameter("cloud", "集群名称").Required(true).DataType("string")).
		Param(ws.PathParameter("namespace", "命名空间").Required(true).DataType("string")).
		Param(ws.PathParameter("name", "pipeline名称").Required(true).DataType("string")).
		Returns(200, "success", v1.ConfigMap{}))

	ws.Route(ws.POST("").To(func(request *restful.Request, response *restful.Response) {
		rctx.NewReqCtx(request, response).WithLog("pipeline").
			WithHandle(api.CreatePipeline).Do()
	}).Doc("创建pipeline").Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(entity.PageResult{}).
		Reads(v1.ConfigMap{}).
		Param(ws.QueryParameter("cloud", "集群名称").Required(true).DataType("string")).
		Returns(200, "success", v1.ConfigMap{}))

	ws.Route(ws.PUT("").To(func(request *restful.Request, response *restful.Response) {
		rctx.NewReqCtx(request, response).WithLog("pipeline").
			WithHandle(api.UpdatePipeline).Do()
	}).Doc("修改pipeline").Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(entity.PageResult{}).
		Reads(v1.ConfigMap{}).
		Param(ws.QueryParameter("cloud", "集群名称").Required(true).DataType("string")).
		Returns(200, "success", v1.ConfigMap{}))

	ws.Route(ws.DELETE("/namespaces/{namespace}/{name}").To(func(request *restful.Request, response *restful.Response) {
		rctx.NewReqCtx(request, response).WithLog("pipeline").
			WithHandle(api.DeletePipeline).Do()
	}).Doc("删除pipeline").Metadata(restfulspec.KeyOpenAPITags, tags).
		Param(ws.QueryParameter("cloud", "集群名称").Required(true).DataType("string")).
		Param(ws.PathParameter("namespace", "命名空间").Required(true).DataType("string")).
		Param(ws.PathParameter("name", "pipeline名称").Required(true).DataType("string")).
		Returns(200, "success", nil))

	return ws
}
