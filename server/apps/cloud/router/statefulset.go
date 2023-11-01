package router

import (
	restfulspec "github.com/emicklei/go-restful-openapi/v2"
	"github.com/emicklei/go-restful/v3"
	"github.com/lbemi/lbemi/apps/cloud/api"
	"github.com/lbemi/lbemi/pkg/common/entity"
	v1 "k8s.io/api/apps/v1"

	"github.com/lbemi/lbemi/pkg/rctx"
)

func KubernetesStatefulSetRoutes() *restful.WebService {
	ws := new(restful.WebService)
	ws.Path("/api/v1/statefulsets").Produces(restful.MIME_JSON)
	tags := []string{"kubernetes-StatefulSet"}

	ws.Route(ws.GET("/namespaces/{namespace}").To(func(request *restful.Request, response *restful.Response) {
		rctx.NewReqCtx(request, response).WithLog("StatefulSet").
			WithHandle(api.ListStatefulSets).Do()
	}).Doc("获取StatefulSet列表").Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(&entity.PageResult{}).
		Param(ws.QueryParameter("cloud", "集群名称").Required(true).DataType("string")).
		Param(ws.QueryParameter("page", "页码").Required(false).DataType("string")).
		Param(ws.QueryParameter("limit", "每页条数").Required(false).DataType("string")).
		Param(ws.QueryParameter("name", "过滤名称，支持模糊查询").DataType("string")).
		Param(ws.QueryParameter("label", "过滤标签，支持模糊查询").DataType("string")).
		Param(ws.PathParameter("namespace", "命名空间").Required(true).DataType("string")).
		Returns(200, "success", &entity.PageResult{}))

	ws.Route(ws.GET("/namespaces/{namespace}/{name}").To(func(request *restful.Request, response *restful.Response) {
		rctx.NewReqCtx(request, response).WithLog("StatefulSet").
			WithHandle(api.GetStatefulSet).Do()
	}).Doc("获取StatefulSet信息").Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(v1.StatefulSet{}).
		Param(ws.QueryParameter("cloud", "集群名称").Required(true).DataType("string")).
		Param(ws.PathParameter("namespace", "命名空间").Required(true).DataType("string")).
		Param(ws.PathParameter("name", "StatefulSet名称").Required(true).DataType("string")).
		Returns(200, "success", v1.StatefulSet{}))

	ws.Route(ws.POST("").To(func(request *restful.Request, response *restful.Response) {
		rctx.NewReqCtx(request, response).WithLog("StatefulSet").
			WithHandle(api.CreateStatefulSet).Do()
	}).Doc("创建StatefulSet").Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(entity.PageResult{}).
		Reads(v1.StatefulSet{}).
		Param(ws.QueryParameter("cloud", "集群名称").Required(true).DataType("string")).
		Returns(200, "success", v1.StatefulSet{}))

	ws.Route(ws.PUT("").To(func(request *restful.Request, response *restful.Response) {
		rctx.NewReqCtx(request, response).WithLog("StatefulSet").
			WithHandle(api.UpdateStatefulSet).Do()
	}).Doc("修改StatefulSet").Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(entity.PageResult{}).
		Reads(v1.StatefulSet{}).
		Param(ws.QueryParameter("cloud", "集群名称").Required(true).DataType("string")).
		Returns(200, "success", v1.StatefulSet{}))

	ws.Route(ws.DELETE("/namespaces/{namespace}/{name}").To(func(request *restful.Request, response *restful.Response) {
		rctx.NewReqCtx(request, response).WithLog("StatefulSet").
			WithHandle(api.DeleteStatefulSet).Do()
	}).Doc("删除StatefulSet").Metadata(restfulspec.KeyOpenAPITags, tags).
		Param(ws.PathParameter("namespace", "命名空间").Required(true).DataType("string")).
		Param(ws.PathParameter("name", "StatefulSet名称").Required(true).DataType("string")).
		Param(ws.QueryParameter("cloud", "集群名称").Required(true).DataType("string")).
		Returns(200, "success", nil))

	return ws
}
