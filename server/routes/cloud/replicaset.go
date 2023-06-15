package cloud

import (
	restfulspec "github.com/emicklei/go-restful-openapi/v2"
	"github.com/emicklei/go-restful/v3"
	v1 "k8s.io/api/apps/v1"

	"github.com/lbemi/lbemi/api/v1/cloud"
	"github.com/lbemi/lbemi/pkg/model/form"
	"github.com/lbemi/lbemi/pkg/rctx"
)

func KubernetesReplicaSetRoutes() *restful.WebService {
	ws := new(restful.WebService)
	ws.Path("/api/v1/replicasets")
	tags := []string{"kubernetes-ReplicaSet"}

	ws.Route(ws.GET("/namespaces/{namespace}").To(func(request *restful.Request, response *restful.Response) {
		rctx.NewReqCtx(request, response).WithLog("ReplicaSet").
			WithHandle(cloud.ListReplicaSets).Do()
	}).Doc("获取ReplicaSet列表").Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(&form.PageResult{}).
		Param(ws.QueryParameter("cloud", "集群名称").Required(true).DataType("string")).
		Param(ws.QueryParameter("page", "页码").Required(false).DataType("string")).
		Param(ws.QueryParameter("limit", "每页条数").Required(false).DataType("string")).
		Param(ws.QueryParameter("name", "过滤名称，支持模糊查询").DataType("string")).
		Param(ws.QueryParameter("label", "过滤标签，支持模糊查询").DataType("string")).
		Param(ws.PathParameter("namespace", "命名空间").Required(true).DataType("string")).
		Returns(200, "success", &form.PageResult{}))

	ws.Route(ws.GET("/namespaces/{namespace}/{name}").To(func(request *restful.Request, response *restful.Response) {
		rctx.NewReqCtx(request, response).WithLog("ReplicaSet").
			WithHandle(cloud.GetReplicaSet).Do()
	}).Doc("获取ReplicaSet信息").Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(v1.ReplicaSet{}).
		Param(ws.QueryParameter("cloud", "集群名称").Required(true).DataType("string")).
		Param(ws.PathParameter("namespace", "命名空间").Required(true).DataType("string")).
		Param(ws.PathParameter("name", "ReplicaSet名称").Required(true).DataType("string")).
		Returns(200, "success", v1.ReplicaSet{}))

	return ws
}
