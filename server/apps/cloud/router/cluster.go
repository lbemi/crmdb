package router

import (
	restfulspec "github.com/emicklei/go-restful-openapi/v2"
	"github.com/emicklei/go-restful/v3"
	"github.com/lbemi/lbemi/apps/cloud/api"
	clusterModel "github.com/lbemi/lbemi/apps/cloud/entity"

	"github.com/lbemi/lbemi/pkg/rctx"
)

func ClusterRoutes() *restful.WebService {
	ws := new(restful.WebService)
	ws.Path("/api/v1/clusters").Produces(restful.MIME_JSON)
	tags := []string{"cluster"}
	ws.Route(ws.GET("").To(func(request *restful.Request, response *restful.Response) {
		rctx.NewReqCtx(request, response).WithLog("cluster").WithHandle(api.ListCluster).Do()
	}).Doc("获取集群列表").Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes([]clusterModel.Cluster{}).
		Returns(200, "success", []clusterModel.Cluster{}))

	ws.Route(ws.GET("/{name}").To(func(request *restful.Request, response *restful.Response) {
		rctx.NewReqCtx(request, response).WithLog("cluster").WithHandle(api.GetCluster).Do()
	}).Doc("根据id获取集群信息").Metadata(restfulspec.KeyOpenAPITags, tags).
		Param(ws.PathParameter("name", "集群id").DataType("string")).
		Writes(clusterModel.Cluster{}).
		Returns(200, "success", clusterModel.Cluster{}))

	ws.Route(ws.POST("").To(func(request *restful.Request, response *restful.Response) {
		rctx.NewReqCtx(request, response).WithLog("cluster").WithHandle(api.CreateCluster).Do()
	}).Doc("导入据群").Metadata(restfulspec.KeyOpenAPITags, tags).
		Param(ws.FormParameter("name", "集群名称").Required(true).DataType("string")).
		Param(ws.FormParameter("file", "kube config配置文件").Required(true).DataType("file")).
		//Param(ws.MultiPartFormParameter("file", "kube config配置文件").Required(true).DataType("file")).
		Returns(200, "success", nil))

	ws.Route(ws.DELETE("/{id}").To(func(request *restful.Request, response *restful.Response) {
		rctx.NewReqCtx(request, response).WithLog("cluster").WithHandle(api.DeleteCluster).Do()
	}).Doc("根据id删除集群").Metadata(restfulspec.KeyOpenAPITags, tags).
		Param(ws.PathParameter("id", "集群id").DataType("string")).
		Returns(200, "success", nil))

	return ws
}
