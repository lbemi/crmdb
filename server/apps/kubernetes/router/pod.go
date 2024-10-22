package router

import (
	restfulspec "github.com/emicklei/go-restful-openapi/v2"
	"github.com/emicklei/go-restful/v3"
	"github.com/lbemi/lbemi/apps/kubernetes/api"
	"github.com/lbemi/lbemi/pkg/common/entity"
	v1 "k8s.io/api/core/v1"

	"github.com/lbemi/lbemi/pkg/rctx"
)

func KubernetesPodRoutes() *restful.WebService {
	ws := new(restful.WebService)
	ws.Path("/api/v1/pods").Produces(restful.MIME_JSON, restful.MIME_OCTET).Produces(restful.MIME_JSON, restful.MIME_OCTET)
	tags := []string{"kubernetes-Pod"}

	ws.Route(ws.GET("/namespaces/{namespace}").To(func(request *restful.Request, response *restful.Response) {
		rctx.NewReqCtx(request, response).WithLog("Pod").
			WithHandle(api.ListPods).Do()
	}).Doc("获取Pod列表").Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(&entity.PageResult{}).
		Param(ws.QueryParameter("cloud", "集群名称").Required(true).DataType("string")).
		Param(ws.PathParameter("namespace", "命名空间").Required(true).DataType("string")).
		Param(ws.QueryParameter("page", "页码").Required(false).DataType("string")).
		Param(ws.QueryParameter("limit", "每页条数").Required(false).DataType("string")).
		Param(ws.QueryParameter("name", "过滤名称，支持模糊查询").DataType("string")).
		Param(ws.QueryParameter("label", "过滤标签，支持模糊查询").DataType("string")).
		Returns(200, "success", &entity.PageResult{}))

	ws.Route(ws.GET("/namespaces/{namespace}/{name}").To(func(request *restful.Request, response *restful.Response) {
		rctx.NewReqCtx(request, response).WithLog("Pod").
			WithHandle(api.GetPod).Do()
	}).Doc("获取Pod信息").Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(v1.Pod{}).
		Param(ws.QueryParameter("cloud", "集群名称").Required(true).DataType("string")).
		Param(ws.PathParameter("namespace", "命名空间").Required(true).DataType("string")).
		Param(ws.PathParameter("name", "Pod名称").Required(true).DataType("string")).
		Returns(200, "success", v1.Pod{}))

	ws.Route(ws.GET("/namespaces/{namespace}/{name}/events").To(func(request *restful.Request, response *restful.Response) {
		rctx.NewReqCtx(request, response).WithLog("Pod").
			WithHandle(api.GetPodEvents).Do()
	}).Doc("获取Pod事件").Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes([]*v1.Event{}).
		Param(ws.QueryParameter("cloud", "集群名称").Required(true).DataType("string")).
		Param(ws.PathParameter("namespace", "命名空间").Required(true).DataType("string")).
		Param(ws.PathParameter("name", "Pod名称").Required(true).DataType("string")).
		Returns(200, "success", []*v1.Event{}))

	ws.Route(ws.GET("/namespaces/{namespace}/exec/{name}/{container}").To(func(request *restful.Request, response *restful.Response) {
		rctx.NewReqCtx(request, response).WithLog("websocket").WithNoRes().
			WithHandle(api.PodExec).Do()
	}).Doc("pod exec命令行 websocket").Metadata(restfulspec.KeyOpenAPITags, tags).
		Param(ws.QueryParameter("cloud", "集群名称").Required(true).DataType("string")).
		Param(ws.PathParameter("namespace", "命名空间").Required(true).DataType("string")).
		Param(ws.PathParameter("name", "Pod名称").Required(true).DataType("string")).
		Param(ws.PathParameter("container", "container名称").Required(true).DataType("string")).
		Returns(200, "success", nil))

	ws.Route(ws.GET("/namespaces/{namespace}/logs/{name}/{container}").To(func(request *restful.Request, response *restful.Response) {
		rctx.NewReqCtx(request, response).WithLog("websocket").WithNoRes().
			WithHandle(api.GetPodLog).Do()
	}).Doc("获取pod日志 websocket").Metadata(restfulspec.KeyOpenAPITags, tags).
		Param(ws.QueryParameter("cloud", "集群名称").Required(true).DataType("string")).
		Param(ws.PathParameter("namespace", "命名空间").Required(true).DataType("string")).
		Param(ws.PathParameter("name", "Pod名称").Required(true).DataType("string")).
		Param(ws.PathParameter("container", "container名称").Required(true).DataType("string")).
		Returns(200, "success", nil))
	ws.Route(ws.POST("").To(func(request *restful.Request, response *restful.Response) {
		rctx.NewReqCtx(request, response).WithLog("Pod").
			WithHandle(api.CreatePod).Do()
	}).Doc("创建Pod").Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(entity.PageResult{}).
		Reads(v1.Pod{}).
		Param(ws.QueryParameter("cloud", "集群名称").Required(true).DataType("string")).
		Returns(200, "success", v1.Pod{}))

	ws.Route(ws.PUT("").To(func(request *restful.Request, response *restful.Response) {
		rctx.NewReqCtx(request, response).WithLog("Pod").
			WithHandle(api.UpdatePod).Do()
	}).Doc("修改Pod").Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(entity.PageResult{}).
		Reads(v1.Pod{}).
		Param(ws.QueryParameter("cloud", "集群名称").Required(true).DataType("string")).
		Returns(200, "success", v1.Pod{}))

	ws.Route(ws.DELETE("/namespaces/{namespace}/{name}").To(func(request *restful.Request, response *restful.Response) {
		rctx.NewReqCtx(request, response).WithLog("Pod").
			WithHandle(api.DeletePod).Do()
	}).Doc("删除Pod").Metadata(restfulspec.KeyOpenAPITags, tags).
		Param(ws.PathParameter("namespace", "命名空间").Required(true).DataType("string")).
		Param(ws.PathParameter("name", "Pod名称").Required(true).DataType("string")).
		Param(ws.QueryParameter("cloud", "集群名称").Required(true).DataType("string")).
		Returns(200, "success", nil))

	ws.Route(ws.GET("/namespaces/{namespace}/files/{name}/{container}").To(func(request *restful.Request, response *restful.Response) {
		rctx.NewReqCtx(request, response).WithLog("files").
			WithHandle(api.GetPodFileList).Do()
	}).Doc("获取Pod文件列表").Metadata(restfulspec.KeyOpenAPITags, tags).
		Param(ws.QueryParameter("cloud", "集群名称").Required(true).DataType("string")).
		Param(ws.PathParameter("namespace", "命名空间").Required(true).DataType("string")).
		Param(ws.PathParameter("name", "Pod名称").Required(true).DataType("string")).
		Param(ws.PathParameter("container", "container名称").Required(true).DataType("string")).
		Param(ws.QueryParameter("path", "路径").Required(true).DataType("string")).
		Returns(200, "success", nil))

	ws.Route(ws.GET("/namespaces/{namespace}/files/{name}/{container}/read").To(func(request *restful.Request, response *restful.Response) {
		rctx.NewReqCtx(request, response).WithLog("files").
			WithHandle(api.ReadFileInfo).Do()
	}).Doc("获取Pod文件内容").Metadata(restfulspec.KeyOpenAPITags, tags).
		Param(ws.QueryParameter("cloud", "集群名称").Required(true).DataType("string")).
		Param(ws.PathParameter("namespace", "命名空间").Required(true).DataType("string")).
		Param(ws.PathParameter("name", "Pod名称").Required(true).DataType("string")).
		Param(ws.PathParameter("container", "container名称").Required(true).DataType("string")).
		Param(ws.QueryParameter("file", "文件路径").Required(true).DataType("string")).
		Returns(200, "success", []byte{}))

	ws.Route(ws.PUT("/namespaces/{namespace}/files/{name}/{container}/mv").To(func(request *restful.Request, response *restful.Response) {
		rctx.NewReqCtx(request, response).WithLog("exec").
			WithHandle(api.UpdateFileName).Do()
	}).Doc("修改Pod文件名").Metadata(restfulspec.KeyOpenAPITags, tags).
		Param(ws.QueryParameter("cloud", "集群名称").Required(true).DataType("string")).
		Param(ws.PathParameter("namespace", "命名空间").Required(true).DataType("string")).
		Param(ws.PathParameter("name", "Pod名称").Required(true).DataType("string")).
		Param(ws.PathParameter("container", "container名称").Required(true).DataType("string")).
		Param(ws.QueryParameter("src", "源文件路径").Required(true).DataType("string")).
		Param(ws.QueryParameter("dst", "目标文件路径").Required(true).DataType("string")).
		Returns(200, "success", nil))

	ws.Route(ws.POST("/namespaces/{namespace}/files/{name}/{container}/create").To(func(request *restful.Request, response *restful.Response) {
		rctx.NewReqCtx(request, response).WithLog("exec").
			WithHandle(api.CreateDir).Do()
	}).Doc("Pod创建目录").Metadata(restfulspec.KeyOpenAPITags, tags).
		Param(ws.QueryParameter("cloud", "集群名称").Required(true).DataType("string")).
		Param(ws.PathParameter("namespace", "命名空间").Required(true).DataType("string")).
		Param(ws.PathParameter("name", "Pod名称").Required(true).DataType("string")).
		Param(ws.QueryParameter("path", "路径").Required(true).DataType("string")).
		Param(ws.PathParameter("container", "container名称").Required(true).DataType("string")).
		Returns(200, "success", nil))

	ws.Route(ws.DELETE("/namespaces/{namespace}/files/{name}/{container}/remove").To(func(request *restful.Request, response *restful.Response) {
		rctx.NewReqCtx(request, response).WithLog("exec").
			WithHandle(api.RemoveFileOrDir).Do()
	}).Doc("删除Pod文件或目录").Metadata(restfulspec.KeyOpenAPITags, tags).
		Param(ws.QueryParameter("cloud", "集群名称").Required(true).DataType("string")).
		Param(ws.PathParameter("namespace", "命名空间").Required(true).DataType("string")).
		Param(ws.PathParameter("name", "Pod名称").Required(true).DataType("string")).
		Param(ws.QueryParameter("dst", "目标文件").Required(true).DataType("string")).
		Param(ws.PathParameter("container", "container名称").Required(true).DataType("string")).
		Returns(200, "success", nil))

	ws.Route(ws.GET("/namespaces/{namespace}/files/{name}/{container}/download").To(func(request *restful.Request, response *restful.Response) {
		rctx.NewReqCtx(request, response).WithLog("exec").WithNoRes().
			WithHandle(api.DownloadFile).Do()
	}).Doc("下载pod中的文件").Metadata(restfulspec.KeyOpenAPITags, tags).
		Param(ws.QueryParameter("cloud", "集群名称").Required(true).DataType("string")).
		Param(ws.PathParameter("namespace", "命名空间").Required(true).DataType("string")).
		Param(ws.PathParameter("name", "Pod名称").Required(true).DataType("string")).
		Param(ws.QueryParameter("file", "下载的文件").Required(true).DataType("string")).
		Param(ws.PathParameter("container", "container名称").Required(true).DataType("string")).
		Returns(200, "success", nil))

	return ws
}
