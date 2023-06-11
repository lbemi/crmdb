package cloud

import (
	restfulspec "github.com/emicklei/go-restful-openapi/v2"
	"github.com/emicklei/go-restful/v3"
	"github.com/lbemi/lbemi/api/v1/cloud"
	"github.com/lbemi/lbemi/pkg/model/form"
	"github.com/lbemi/lbemi/pkg/rctx"
	"k8s.io/api/core/v1"
)

func KubernetesPodRoutes() *restful.WebService {
	ws := new(restful.WebService)
	ws.Path("/api/v1/pods")
	tags := []string{"kubernetes-Pod"}

	ws.Route(ws.GET("/namespaces/{namespace}").To(func(request *restful.Request, response *restful.Response) {
		rctx.NewReqCtx(request, response).WithLog("Pod").
			WithHandle(cloud.ListPods).Do()
	}).Doc("获取Pod列表").Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(&form.PageResult{}).
		Param(ws.QueryParameter("cloud", "集群名称").Required(true).DataType("string")).
		Param(ws.PathParameter("namespace", "命名空间").Required(true).DataType("string")).
		Param(ws.QueryParameter("page", "页码").Required(false).DataType("string")).
		Param(ws.QueryParameter("limit", "每页条数").Required(false).DataType("string")).
		Param(ws.QueryParameter("name", "过滤名称，支持模糊查询").DataType("string")).
		Param(ws.QueryParameter("label", "过滤标签，支持模糊查询").DataType("string")).
		Returns(200, "success", &form.PageResult{}))

	ws.Route(ws.GET("/namespaces/{namespace}/{name}").To(func(request *restful.Request, response *restful.Response) {
		rctx.NewReqCtx(request, response).WithLog("Pod").
			WithHandle(cloud.GetPod).Do()
	}).Doc("获取Pod信息").Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(v1.Pod{}).
		Param(ws.QueryParameter("cloud", "集群名称").Required(true).DataType("string")).
		Param(ws.PathParameter("namespace", "命名空间").Required(true).DataType("string")).
		Param(ws.PathParameter("name", "Pod名称").Required(true).DataType("string")).
		Returns(200, "success", v1.Pod{}))

	ws.Route(ws.GET("/namespaces/{namespace}/{name}/events").To(func(request *restful.Request, response *restful.Response) {
		rctx.NewReqCtx(request, response).WithLog("Pod").
			WithHandle(cloud.GetPodEvents).Do()
	}).Doc("获取Pod事件").Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes([]*v1.Event{}).
		Param(ws.QueryParameter("cloud", "集群名称").Required(true).DataType("string")).
		Param(ws.PathParameter("namespace", "命名空间").Required(true).DataType("string")).
		Param(ws.PathParameter("name", "Pod名称").Required(true).DataType("string")).
		Returns(200, "success", []*v1.Event{}))

	ws.Route(ws.GET("/namespaces/{namespace}/exec/{name}/{container}").To(func(request *restful.Request, response *restful.Response) {
		rctx.NewReqCtx(request, response).WithLog("websocket").WithNoRes().
			WithHandle(cloud.PodExec).Do()
	}).Doc("pod exec命令行 websocket").Metadata(restfulspec.KeyOpenAPITags, tags).
		Param(ws.QueryParameter("cloud", "集群名称").Required(true).DataType("string")).
		Param(ws.PathParameter("namespace", "命名空间").Required(true).DataType("string")).
		Param(ws.PathParameter("name", "Pod名称").Required(true).DataType("string")).
		Param(ws.PathParameter("container", "container名称").Required(true).DataType("string")).
		Returns(200, "success", nil))

	ws.Route(ws.GET("/namespaces/{namespace}/logs/{name}/{container}").To(func(request *restful.Request, response *restful.Response) {
		rctx.NewReqCtx(request, response).WithLog("websocket").WithNoRes().
			WithHandle(cloud.GetPodLog).Do()
	}).Doc("获取pod日志 websocket").Metadata(restfulspec.KeyOpenAPITags, tags).
		Param(ws.QueryParameter("cloud", "集群名称").Required(true).DataType("string")).
		Param(ws.PathParameter("namespace", "命名空间").Required(true).DataType("string")).
		Param(ws.PathParameter("name", "Pod名称").Required(true).DataType("string")).
		Param(ws.PathParameter("container", "container名称").Required(true).DataType("string")).
		Returns(200, "success", nil))
	ws.Route(ws.POST("").To(func(request *restful.Request, response *restful.Response) {
		rctx.NewReqCtx(request, response).WithLog("Pod").
			WithHandle(cloud.CreatePod).Do()
	}).Doc("创建Pod").Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(form.PageResult{}).
		Reads(v1.Pod{}).
		Param(ws.QueryParameter("cloud", "集群名称").Required(true).DataType("string")).
		Returns(200, "success", v1.Pod{}))

	ws.Route(ws.PUT("").To(func(request *restful.Request, response *restful.Response) {
		rctx.NewReqCtx(request, response).WithLog("Pod").
			WithHandle(cloud.UpdatePod).Do()
	}).Doc("修改Pod").Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(form.PageResult{}).
		Reads(v1.Pod{}).
		Param(ws.QueryParameter("cloud", "集群名称").Required(true).DataType("string")).
		Returns(200, "success", v1.Pod{}))

	ws.Route(ws.DELETE("/namespaces/{namespace}/{name}").To(func(request *restful.Request, response *restful.Response) {
		rctx.NewReqCtx(request, response).WithLog("Pod").
			WithHandle(cloud.DeletePod).Do()
	}).Doc("删除Pod").Metadata(restfulspec.KeyOpenAPITags, tags).
		Param(ws.PathParameter("namespace", "命名空间").Required(true).DataType("string")).
		Param(ws.PathParameter("name", "Pod名称").Required(true).DataType("string")).
		Param(ws.QueryParameter("cloud", "集群名称").Required(true).DataType("string")).
		Returns(200, "success", nil))

	return ws
}
