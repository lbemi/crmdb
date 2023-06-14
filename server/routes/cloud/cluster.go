package cloud

import (
	restfulspec "github.com/emicklei/go-restful-openapi/v2"
	"github.com/emicklei/go-restful/v3"
	"github.com/gin-gonic/gin"
	"github.com/lbemi/lbemi/api/v1/cloud"
	clusterModel "github.com/lbemi/lbemi/pkg/model/cloud"
	"github.com/lbemi/lbemi/pkg/rctx"
)

func NewClusterRoutes(router *gin.RouterGroup) {

	//cluster := router.Group("/cluster")
	//{
	//	cluster.POST("", cloud.CreateCluster)
	//	cluster.GET("", cloud.ListCluster)
	//	cluster.DELETE("/:id", cloud.DeleteCluster)
	//	cluster.GET("/:name", cloud.GetCluster)
	//}

}

func ClusterRoutes() *restful.WebService {
	ws := new(restful.WebService)
	ws.Path("/api/v1/clusters").Produces(restful.MIME_JSON)
	tags := []string{"cluster"}
	ws.Route(ws.GET("").To(func(request *restful.Request, response *restful.Response) {
		rctx.NewReqCtx(request, response).WithLog("cluster").WithHandle(cloud.ListCluster).Do()
	}).Doc("获取集群列表").Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes([]clusterModel.Cluster{}).
		Returns(200, "success", []clusterModel.Cluster{}))

	ws.Route(ws.GET("/{name}").To(func(request *restful.Request, response *restful.Response) {
		rctx.NewReqCtx(request, response).WithLog("cluster").WithHandle(cloud.GetCluster).Do()
	}).Doc("根据id获取集群信息").Metadata(restfulspec.KeyOpenAPITags, tags).
		Param(ws.PathParameter("name", "集群id").DataType("string")).
		Writes(clusterModel.Cluster{}).
		Returns(200, "success", clusterModel.Cluster{}))

	ws.Route(ws.POST("").To(func(request *restful.Request, response *restful.Response) {
		rctx.NewReqCtx(request, response).WithLog("cluster").WithHandle(cloud.CreateCluster).Do()
	}).Doc("导入据群").Metadata(restfulspec.KeyOpenAPITags, tags).
		Param(ws.FormParameter("name", "集群名称").Required(true).DataType("string")).
		Param(ws.FormParameter("file", "kube config配置文件").Required(true).DataType("file")).
		//Param(ws.MultiPartFormParameter("file", "kube config配置文件").Required(true).DataType("file")).
		Returns(200, "success", nil))

	ws.Route(ws.DELETE("/{id}").To(func(request *restful.Request, response *restful.Response) {
		rctx.NewReqCtx(request, response).WithLog("cluster").WithHandle(cloud.DeleteCluster).Do()
	}).Doc("根据id删除集群").Metadata(restfulspec.KeyOpenAPITags, tags).
		Param(ws.PathParameter("id", "集群id").DataType("string")).
		Returns(200, "success", nil))

	return ws
}
