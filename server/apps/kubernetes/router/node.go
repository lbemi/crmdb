package router

import (
	restfulspec "github.com/emicklei/go-restful-openapi/v2"
	"github.com/emicklei/go-restful/v3"
	api2 "github.com/lbemi/lbemi/apps/kubernetes/api"
	"github.com/lbemi/lbemi/pkg/common/entity"
	v1 "k8s.io/api/core/v1"

	"github.com/lbemi/lbemi/pkg/rctx"
)

func KubernetesNodeRoutes() *restful.WebService {
	ws := new(restful.WebService)
	ws.Path("/api/v1/nodes").Produces(restful.MIME_JSON)
	tags := []string{"kubernetes-node"}

	ws.Route(ws.GET("").To(func(request *restful.Request, response *restful.Response) {
		rctx.NewReqCtx(request, response).WithLog("node").
			WithHandle(api2.ListNodes).Do()
	}).Doc("获取Node列表").Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(&entity.PageResult{}).
		Param(ws.QueryParameter("cloud", "集群名称").Required(true).DataType("string")).
		Param(ws.QueryParameter("page", "页码").Required(false).DataType("string")).
		Param(ws.QueryParameter("limit", "每页条数").Required(false).DataType("string")).
		Param(ws.QueryParameter("name", "过滤名称，支持模糊查询").DataType("string")).
		Param(ws.QueryParameter("label", "过滤标签，支持模糊查询").DataType("string")).
		Returns(200, "success", &entity.PageResult{}))

	ws.Route(ws.GET("/{name}/pods").To(func(request *restful.Request, response *restful.Response) {
		rctx.NewReqCtx(request, response).WithLog("node").
			WithHandle(api2.GetPodByNode).Do()
	}).Doc("获取Node的pod列表").Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(&entity.PageResult{}).
		Param(ws.QueryParameter("cloud", "集群名称").Required(true).DataType("string")).
		Param(ws.QueryParameter("page", "页码").Required(false).DataType("string")).
		Param(ws.QueryParameter("limit", "每页条数").Required(false).DataType("string")).
		Param(ws.PathParameter("name", "Node名称").DataType("string")).
		Returns(200, "success", &entity.PageResult{}))
	ws.Route(ws.GET("/{name}").To(func(request *restful.Request, response *restful.Response) {
		rctx.NewReqCtx(request, response).WithLog("node").
			WithHandle(api2.GetNode).Do()
	}).Doc("获取Node信息").Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(v1.Node{}).
		Param(ws.QueryParameter("cloud", "集群名称").Required(true).DataType("string")).
		Param(ws.PathParameter("name", "Node名称").Required(true).DataType("string")).
		Returns(200, "success", v1.Node{}))

	ws.Route(ws.GET("/{name}/events").To(func(request *restful.Request, response *restful.Response) {
		rctx.NewReqCtx(request, response).WithLog("node").
			WithHandle(api2.GetPodEvents).Do()
	}).Doc("获取Node节点的事件").Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes([]*v1.Event{}).
		Param(ws.QueryParameter("cloud", "集群名称").Required(true).DataType("string")).
		Param(ws.PathParameter("name", "Node名称").Required(true).DataType("string")).
		Returns(200, "success", []*v1.Node{}))

	ws.Route(ws.PUT("").To(func(request *restful.Request, response *restful.Response) {
		rctx.NewReqCtx(request, response).WithLog("node").
			WithHandle(api2.UpdateNode).Do()
	}).Doc("修改Node").Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(entity.PageResult{}).
		Reads(v1.Node{}).
		Param(ws.QueryParameter("cloud", "集群名称").Required(true).DataType("string")).
		Returns(200, "success", v1.Node{}))

	ws.Route(ws.PUT("/{name}/schedule/{unschedulable}").To(func(request *restful.Request, response *restful.Response) {
		rctx.NewReqCtx(request, response).WithLog("node").
			WithHandle(api2.Schedulable).Do()
	}).Doc("设定node节点是否可以调度").Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(entity.PageResult{}).
		Param(ws.QueryParameter("cloud", "集群名称").Required(true).DataType("string")).
		Param(ws.PathParameter("name", "Node名称").Required(true).DataType("string")).
		Param(ws.QueryParameter("unschedulable", "是否可以调度").Required(true).DataType("bool")).
		Returns(200, "success", v1.Node{}))

	ws.Route(ws.PATCH("/label").To(func(request *restful.Request, response *restful.Response) {
		rctx.NewReqCtx(request, response).WithLog("node").
			WithHandle(api2.PatchNode).Do()
	}).Doc("修改node节点标签").Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(v1.Node{}).
		Reads(entity.PatchNode{}).
		Param(ws.QueryParameter("cloud", "集群名称").Required(true).DataType("string")).
		Returns(200, "success", v1.Node{}))

	ws.Route(ws.POST("/{name}/drain").To(func(request *restful.Request, response *restful.Response) {
		rctx.NewReqCtx(request, response).WithLog("node").
			WithHandle(api2.Drain).Do()
	}).Doc("node节点排水").Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(v1.Node{}).
		Param(ws.QueryParameter("cloud", "集群名称").Required(true).DataType("string")).
		Param(ws.PathParameter("name", "Node名称").Required(true).DataType("string")).
		Returns(200, "success", v1.Node{}))

	return ws
}
