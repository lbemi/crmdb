package router

import (
	restfulspec "github.com/emicklei/go-restful-openapi/v2"
	"github.com/emicklei/go-restful/v3"
	"github.com/lbemi/lbemi/apps/cloud/api"
	"github.com/lbemi/lbemi/pkg/common/entity"
	v1 "k8s.io/api/core/v1"

	"github.com/lbemi/lbemi/pkg/rctx"
)

func KubernetesConfigMapRoutes() *restful.WebService {
	ws := new(restful.WebService)
	ws.Path("/api/v1/configmaps").Produces(restful.MIME_JSON)
	tags := []string{"kubernetes-configmap"}

	ws.Route(ws.GET("/namespaces/{namespace}").To(func(request *restful.Request, response *restful.Response) {
		rctx.NewReqCtx(request, response).WithLog("configmap").
			WithHandle(api.ListConfigMaps).Do()
	}).Doc("获取configmap列表").Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(&entity.PageConfigMap{}).
		Param(ws.QueryParameter("cloud", "集群名称").Required(true).DataType("string")).
		Param(ws.QueryParameter("page", "页码").Required(false).DataType("string")).
		Param(ws.QueryParameter("limit", "每页条数").Required(false).DataType("string")).
		Param(ws.QueryParameter("name", "过滤名称，支持模糊查询").DataType("string")).
		Param(ws.QueryParameter("label", "过滤标签，支持模糊查询").DataType("string")).
		Param(ws.PathParameter("namespace", "命名空间").Required(true).DataType("string")).
		Returns(200, "success", &entity.PageConfigMap{}))

	ws.Route(ws.GET("/namespaces/{namespace}/{name}").To(func(request *restful.Request, response *restful.Response) {
		rctx.NewReqCtx(request, response).WithLog("configmap").
			WithHandle(api.GetConfigMap).Do()
	}).Doc("获取configmap信息").Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(v1.ConfigMap{}).
		Param(ws.QueryParameter("cloud", "集群名称").Required(true).DataType("string")).
		Param(ws.PathParameter("namespace", "命名空间").Required(true).DataType("string")).
		Param(ws.PathParameter("name", "configmap名称").Required(true).DataType("string")).
		Returns(200, "success", v1.ConfigMap{}))

	ws.Route(ws.POST("").To(func(request *restful.Request, response *restful.Response) {
		rctx.NewReqCtx(request, response).WithLog("configmap").
			WithHandle(api.CreateConfigMap).Do()
	}).Doc("创建configmap").Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(entity.PageResult{}).
		Reads(v1.ConfigMap{}).
		Param(ws.QueryParameter("cloud", "集群名称").Required(true).DataType("string")).
		Returns(200, "success", v1.ConfigMap{}))

	ws.Route(ws.PUT("").To(func(request *restful.Request, response *restful.Response) {
		rctx.NewReqCtx(request, response).WithLog("configmap").
			WithHandle(api.UpdateConfigMap).Do()
	}).Doc("修改configmap").Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(entity.PageResult{}).
		Reads(v1.ConfigMap{}).
		Param(ws.QueryParameter("cloud", "集群名称").Required(true).DataType("string")).
		Returns(200, "success", v1.ConfigMap{}))

	ws.Route(ws.DELETE("/namespaces/{namespace}/{name}").To(func(request *restful.Request, response *restful.Response) {
		rctx.NewReqCtx(request, response).WithLog("configmap").
			WithHandle(api.DeleteConfigMap).Do()
	}).Doc("删除configmap").Metadata(restfulspec.KeyOpenAPITags, tags).
		Param(ws.PathParameter("namespace", "命名空间").Required(true).DataType("string")).
		Param(ws.PathParameter("name", "configmap名称").Required(true).DataType("string")).
		Param(ws.QueryParameter("cloud", "集群名称").Required(true).DataType("string")).
		Returns(200, "success", nil))

	return ws
}
