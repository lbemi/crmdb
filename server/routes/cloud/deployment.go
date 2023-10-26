package cloud

import (
	restfulspec "github.com/emicklei/go-restful-openapi/v2"
	"github.com/emicklei/go-restful/v3"
	"github.com/lbemi/lbemi/api/cloud"
	appsv1 "k8s.io/api/apps/v1"
	v1 "k8s.io/api/core/v1"

	"github.com/lbemi/lbemi/pkg/model/form"
	"github.com/lbemi/lbemi/pkg/rctx"
)

func KubernetesDeploymentRoutes() *restful.WebService {
	ws := new(restful.WebService)
	ws.Path("/api/v1/deployments").Produces(restful.MIME_JSON)
	tags := []string{"kubernetes-deployment"}

	ws.Route(ws.GET("/namespaces/{namespace}").To(func(request *restful.Request, response *restful.Response) {
		rctx.NewReqCtx(request, response).WithLog("deployment").
			WithHandle(cloud.ListDeployments).Do()
	}).Doc("获取deployment列表").Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(&form.PageResult{}).
		Param(ws.QueryParameter("cloud", "集群名称").Required(true).DataType("string")).
		Param(ws.PathParameter("namespace", "命名空间").Required(true).DataType("string")).
		Param(ws.QueryParameter("page", "页码").Required(false).DataType("string")).
		Param(ws.QueryParameter("limit", "每页条数").Required(false).DataType("string")).
		Param(ws.QueryParameter("name", "过滤名称，支持模糊查询").DataType("string")).
		Param(ws.QueryParameter("label", "过滤标签，支持模糊查询").DataType("string")).
		Returns(200, "success", &form.PageResult{}))

	ws.Route(ws.GET("/namespaces/{namespace}/{name}").To(func(request *restful.Request, response *restful.Response) {
		rctx.NewReqCtx(request, response).WithLog("deployment").
			WithHandle(cloud.GetDeployment).Do()
	}).Doc("获取deployment信息").Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(appsv1.Deployment{}).
		Param(ws.QueryParameter("cloud", "集群名称").Required(true).DataType("string")).
		Param(ws.PathParameter("namespace", "命名空间").Required(true).DataType("string")).
		Param(ws.PathParameter("name", "deployment名称").Required(true).DataType("string")).
		Returns(200, "success", appsv1.Deployment{}))
	//	deployment.GET("/:namespace/:deploymentName/pod", cloud2.GetDeploymentPods)
	ws.Route(ws.GET("/namespaces/{namespace}/{name}/pods").To(func(request *restful.Request, response *restful.Response) {
		rctx.NewReqCtx(request, response).WithLog("deployment").
			WithHandle(cloud.GetDeploymentPods).Do()
	}).Doc("获取deployment所属pod列表").Metadata(restfulspec.KeyOpenAPITags, tags).
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

	//	deployment.GET("/:namespace/:deploymentName/event", cloud2.GetDeploymentEvents)
	ws.Route(ws.GET("/namespaces/{namespace}/{name}/events").To(func(request *restful.Request, response *restful.Response) {
		rctx.NewReqCtx(request, response).WithLog("deployment").
			WithHandle(cloud.GetDeploymentEvents).Do()
	}).Doc("获取deployment所属事件列表").Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes([]*v1.Event{}).
		Param(ws.QueryParameter("cloud", "集群名称").Required(true).DataType("string")).
		Param(ws.PathParameter("namespace", "命名空间").Required(true).DataType("string")).
		Param(ws.PathParameter("name", "deployment名称").Required(true).DataType("string")).
		Returns(200, "success", []*v1.Event{}))

	ws.Route(ws.POST("").To(func(request *restful.Request, response *restful.Response) {
		rctx.NewReqCtx(request, response).WithLog("deployment").
			WithHandle(cloud.CreateDeployment).Do()
	}).Doc("创建deployment").Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(form.PageResult{}).
		Reads(appsv1.Deployment{}).
		Param(ws.QueryParameter("cloud", "集群名称").Required(true).DataType("string")).
		Returns(200, "success", appsv1.Deployment{}))

	ws.Route(ws.PUT("").To(func(request *restful.Request, response *restful.Response) {
		rctx.NewReqCtx(request, response).WithLog("deployment").
			WithHandle(cloud.UpdateDeployment).Do()
	}).Doc("修改deployment").Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(form.PageResult{}).
		Reads(appsv1.Deployment{}).
		Param(ws.QueryParameter("cloud", "集群名称").Required(true).DataType("string")).
		Returns(200, "success", appsv1.Deployment{}))

	ws.Route(ws.DELETE("/namespaces/{namespace}/{name}").To(func(request *restful.Request, response *restful.Response) {
		rctx.NewReqCtx(request, response).WithLog("deployment").
			WithHandle(cloud.DeleteDeployment).Do()
	}).Doc("删除deployment").Metadata(restfulspec.KeyOpenAPITags, tags).
		Param(ws.QueryParameter("cloud", "集群名称").Required(true).DataType("string")).
		Param(ws.PathParameter("namespace", "命名空间").Required(true).DataType("string")).
		Param(ws.PathParameter("name", "deployment名称").Required(true).DataType("string")).
		Returns(200, "success", nil))

	//	deployment.PUT("/redeploy/:namespace/:name", cloud2.ReDeployDeployment)
	ws.Route(ws.PUT("/namespaces/{namespace}/redeploy/{name}").To(func(request *restful.Request, response *restful.Response) {
		rctx.NewReqCtx(request, response).WithLog("deployment").
			WithHandle(cloud.ReDeployDeployment).Do()
	}).Doc("重新部署deployment").Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(form.PageResult{}).
		Reads(appsv1.Deployment{}).
		Param(ws.QueryParameter("cloud", "集群名称").Required(true).DataType("string")).
		Param(ws.PathParameter("namespace", "命名空间").Required(true).DataType("string")).
		Param(ws.PathParameter("name", "deployment名称").Required(true).DataType("string")).
		Returns(200, "success", appsv1.Deployment{}))
	//	deployment.PUT("/rollback/:namespace/:name/:reversion", cloud2.RollBackDeployment)
	ws.Route(ws.PUT("/namespaces/{namespace}/rollback/{name}/{reversion}").To(func(request *restful.Request, response *restful.Response) {
		rctx.NewReqCtx(request, response).WithLog("deployment").
			WithHandle(cloud.RollBackDeployment).Do()
	}).Doc("回滚deployment到指定版本").Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(form.PageResult{}).
		Reads(appsv1.Deployment{}).
		Param(ws.QueryParameter("cloud", "集群名称").Required(true).DataType("string")).
		Param(ws.PathParameter("namespace", "命名空间").Required(true).DataType("string")).
		Param(ws.PathParameter("name", "deployment名称").Required(true).DataType("string")).
		Param(ws.PathParameter("reversion", "deployment版本号").Required(true).DataType("string")).
		Returns(200, "success", appsv1.Deployment{}))
	//	deployment.PUT("/:namespace/:deploymentName/:scale", cloud2.ScaleDeployments)
	ws.Route(ws.PUT("/namespaces/{namespace}/{name}/scale/{replica}").To(func(request *restful.Request, response *restful.Response) {
		rctx.NewReqCtx(request, response).WithLog("deployment").
			WithHandle(cloud.ScaleDeployments).Do()
	}).Doc("修改deployment副本数").Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(form.PageResult{}).
		Reads(appsv1.Deployment{}).
		Param(ws.QueryParameter("cloud", "集群名称").Required(true).DataType("string")).
		Param(ws.PathParameter("namespace", "命名空间").Required(true).DataType("string")).
		Param(ws.PathParameter("name", "deployment名称").Required(true).DataType("string")).
		Param(ws.PathParameter("replica", "副本数").Required(true).DataType("int")).
		Returns(200, "success", appsv1.Deployment{}))
	return ws
}
