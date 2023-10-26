package cloud

import (
	restfulspec "github.com/emicklei/go-restful-openapi/v2"
	"github.com/emicklei/go-restful/v3"
	cloud2 "github.com/lbemi/lbemi/api/cloud"
	appsv1 "k8s.io/api/apps/v1"

	"github.com/lbemi/lbemi/pkg/model/form"
	"github.com/lbemi/lbemi/pkg/rctx"
)

func KubernetesDaemonSetRoutes() *restful.WebService {
	ws := new(restful.WebService)
	ws.Path("/api/v1/daemonsets").Produces(restful.MIME_JSON)
	tags := []string{"kubernetes-daemonset"}

	ws.Route(ws.GET("/namespaces/{namespace}").To(func(request *restful.Request, response *restful.Response) {
		rctx.NewReqCtx(request, response).WithLog("daemonset").
			WithHandle(cloud2.ListDaemonSets).Do()
	}).Doc("获取daemonset列表").Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(&form.PageResult{}).
		Param(ws.QueryParameter("cloud", "集群名称").Required(true).DataType("string")).
		Param(ws.PathParameter("namespace", "命名空间").Required(true).DataType("string")).
		Param(ws.QueryParameter("page", "页码").Required(false).DataType("string")).
		Param(ws.QueryParameter("limit", "每页条数").Required(false).DataType("string")).
		Param(ws.QueryParameter("name", "过滤名称，支持模糊查询").DataType("string")).
		Param(ws.QueryParameter("label", "过滤标签，支持模糊查询").DataType("string")).
		Returns(200, "success", &form.PageResult{}))

	ws.Route(ws.GET("/namespaces/{namespace}/{name}").To(func(request *restful.Request, response *restful.Response) {
		rctx.NewReqCtx(request, response).WithLog("daemonset").
			WithHandle(cloud2.GetCronJob).Do()
	}).Doc("获取daemonset信息").Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(appsv1.DaemonSet{}).
		Param(ws.QueryParameter("cloud", "集群名称").Required(true).DataType("string")).
		Param(ws.PathParameter("namespace", "命名空间").Required(true).DataType("string")).
		Param(ws.PathParameter("name", "daemonset名称").Required(true).DataType("string")).
		Returns(200, "success", appsv1.DaemonSet{}))

	ws.Route(ws.POST("").To(func(request *restful.Request, response *restful.Response) {
		rctx.NewReqCtx(request, response).WithLog("daemonset").
			WithHandle(cloud2.CreateCronJob).Do()
	}).Doc("创建daemonset").Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(form.PageResult{}).
		Reads(appsv1.DaemonSet{}).
		Param(ws.QueryParameter("cloud", "集群名称").Required(true).DataType("string")).
		Returns(200, "success", appsv1.DaemonSet{}))

	ws.Route(ws.PUT("").To(func(request *restful.Request, response *restful.Response) {
		rctx.NewReqCtx(request, response).WithLog("daemonset").
			WithHandle(cloud2.UpdateCronJob).Do()
	}).Doc("修改daemonset").Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(form.PageResult{}).
		Reads(appsv1.DaemonSet{}).
		Param(ws.QueryParameter("cloud", "集群名称").Required(true).DataType("string")).
		Returns(200, "success", appsv1.DaemonSet{}))

	ws.Route(ws.DELETE("/namespaces/{namespace}/{name}").To(func(request *restful.Request, response *restful.Response) {
		rctx.NewReqCtx(request, response).WithLog("daemonset").
			WithHandle(cloud2.DeleteCronJob).Do()
	}).Doc("删除daemonset").Metadata(restfulspec.KeyOpenAPITags, tags).
		Param(ws.QueryParameter("cloud", "集群名称").Required(true).DataType("string")).
		Param(ws.PathParameter("namespace", "命名空间").Required(true).DataType("string")).
		Param(ws.PathParameter("name", "daemonset名称").Required(true).DataType("string")).
		Returns(200, "success", nil))

	return ws
}
