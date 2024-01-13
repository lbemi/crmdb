package router

import (
	restfulspec "github.com/emicklei/go-restful-openapi/v2"
	"github.com/emicklei/go-restful/v3"
	"github.com/lbemi/lbemi/apps/kubernetes/api"
	"github.com/lbemi/lbemi/pkg/common/entity"
	appsv1 "k8s.io/api/apps/v1"
	v1 "k8s.io/api/core/v1"

	"github.com/lbemi/lbemi/pkg/rctx"
)

func KubernetesDaemonSetRoutes() *restful.WebService {
	ws := new(restful.WebService)
	ws.Path("/api/v1/daemonsets").Produces(restful.MIME_JSON)
	tags := []string{"kubernetes-daemonset"}

	ws.Route(ws.GET("/namespaces/{namespace}").To(func(request *restful.Request, response *restful.Response) {
		rctx.NewReqCtx(request, response).WithLog("daemonset").
			WithHandle(api.ListDaemonSets).Do()
	}).Doc("获取daemonset列表").Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(&entity.PageResult{}).
		Param(ws.QueryParameter("cloud", "集群名称").Required(true).DataType("string")).
		Param(ws.PathParameter("namespace", "命名空间").Required(true).DataType("string")).
		Param(ws.QueryParameter("page", "页码").Required(false).DataType("string")).
		Param(ws.QueryParameter("limit", "每页条数").Required(false).DataType("string")).
		Param(ws.QueryParameter("name", "过滤名称，支持模糊查询").DataType("string")).
		Param(ws.QueryParameter("label", "过滤标签，支持模糊查询").DataType("string")).
		Returns(200, "success", &entity.PageResult{}))

	ws.Route(ws.GET("/namespaces/{namespace}/{name}").To(func(request *restful.Request, response *restful.Response) {
		rctx.NewReqCtx(request, response).WithLog("daemonset").
			WithHandle(api.GetDaemonSet).Do()
	}).Doc("获取daemonset信息").Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(appsv1.DaemonSet{}).
		Param(ws.QueryParameter("cloud", "集群名称").Required(true).DataType("string")).
		Param(ws.PathParameter("namespace", "命名空间").Required(true).DataType("string")).
		Param(ws.PathParameter("name", "daemonset名称").Required(true).DataType("string")).
		Returns(200, "success", appsv1.DaemonSet{}))
	ws.Route(ws.GET("/namespaces/{namespace}/{name}/pods").To(func(request *restful.Request, response *restful.Response) {
		rctx.NewReqCtx(request, response).WithLog("deployment").
			WithHandle(api.GetDaemonSetPods).Do()
	}).Doc("获取daemonset所属pod列表").Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(map[string]interface{}{
			"pods":        []*v1.Pod{},
			"replicaSets": []*appsv1.ReplicaSet{},
		}).
		Param(ws.QueryParameter("cloud", "集群名称").Required(true).DataType("string")).
		Param(ws.PathParameter("namespace", "命名空间").Required(true).DataType("string")).
		Param(ws.PathParameter("name", "deployment名称").Required(true).DataType("string")).
		Returns(200, "success", map[string]interface{}{
			"pods":        []*v1.Pod{},
			"replicaSets": []*appsv1.ReplicaSet{},
		}))

	ws.Route(ws.GET("/namespaces/{namespace}/{name}/events").To(func(request *restful.Request, response *restful.Response) {
		rctx.NewReqCtx(request, response).WithLog("deployment").
			WithHandle(api.GetDaemonSetEvents).Do()
	}).Doc("获取daemonSet所属事件列表").Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes([]*v1.Event{}).
		Param(ws.QueryParameter("cloud", "集群名称").Required(true).DataType("string")).
		Param(ws.PathParameter("namespace", "命名空间").Required(true).DataType("string")).
		Param(ws.PathParameter("name", "daemonSet名称").Required(true).DataType("string")).
		Returns(200, "success", []*v1.Event{}))

	ws.Route(ws.POST("").To(func(request *restful.Request, response *restful.Response) {
		rctx.NewReqCtx(request, response).WithLog("daemonset").
			WithHandle(api.CreateDaemonSet).Do()
	}).Doc("创建daemonset").Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(entity.PageResult{}).
		Reads(appsv1.DaemonSet{}).
		Param(ws.QueryParameter("cloud", "集群名称").Required(true).DataType("string")).
		Returns(200, "success", appsv1.DaemonSet{}))

	ws.Route(ws.PUT("").To(func(request *restful.Request, response *restful.Response) {
		rctx.NewReqCtx(request, response).WithLog("daemonset").
			WithHandle(api.UpdateDaemonSet).Do()
	}).Doc("修改daemonset").Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(entity.PageResult{}).
		Reads(appsv1.DaemonSet{}).
		Param(ws.QueryParameter("cloud", "集群名称").Required(true).DataType("string")).
		Returns(200, "success", appsv1.DaemonSet{}))

	ws.Route(ws.DELETE("/namespaces/{namespace}/{name}").To(func(request *restful.Request, response *restful.Response) {
		rctx.NewReqCtx(request, response).WithLog("daemonset").
			WithHandle(api.DeleteDaemonSet).Do()
	}).Doc("删除daemonset").Metadata(restfulspec.KeyOpenAPITags, tags).
		Param(ws.QueryParameter("cloud", "集群名称").Required(true).DataType("string")).
		Param(ws.PathParameter("namespace", "命名空间").Required(true).DataType("string")).
		Param(ws.PathParameter("name", "daemonset名称").Required(true).DataType("string")).
		Returns(200, "success", nil))

	return ws
}
