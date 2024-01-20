package router

import (
	restfulspec "github.com/emicklei/go-restful-openapi/v2"
	"github.com/emicklei/go-restful/v3"
	"github.com/lbemi/lbemi/apps/kubernetes/api"
	"github.com/lbemi/lbemi/pkg/common/entity"
	"github.com/lbemi/lbemi/pkg/rctx"
	"k8s.io/api/storage/v1"
)

func KubernetesStorageClassRoutes() *restful.WebService {
	ws := new(restful.WebService)
	ws.Path("/api/v1/storageclasses").Produces(restful.MIME_JSON)
	tags := []string{"kubernetes-storageClass"}

	ws.Route(ws.GET("").To(func(request *restful.Request, response *restful.Response) {
		rctx.NewReqCtx(request, response).WithLog("storageClass").
			WithHandle(api.ListStorageClass).Do()
	}).Doc("获取storageClass列表").Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(&entity.PageResult{}).
		Param(ws.QueryParameter("cloud", "集群名称").Required(true).DataType("string")).
		Param(ws.QueryParameter("page", "页码").Required(false).DataType("string")).
		Param(ws.QueryParameter("limit", "每页条数").Required(false).DataType("string")).
		Param(ws.QueryParameter("name", "过滤名称，支持模糊查询").DataType("string")).
		Param(ws.QueryParameter("label", "过滤标签，支持模糊查询").DataType("string")).
		Returns(200, "success", &entity.PageResult{}))

	ws.Route(ws.GET("/{name}").To(func(request *restful.Request, response *restful.Response) {
		rctx.NewReqCtx(request, response).WithLog("storageClass").
			WithHandle(api.GetStorageClass).Do()
	}).Doc("获取storageClass信息").Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(v1.StorageClass{}).
		Param(ws.QueryParameter("cloud", "集群名称").Required(true).DataType("string")).
		Param(ws.PathParameter("name", "storageClass名称").Required(true).DataType("string")).
		Returns(200, "success", v1.StorageClass{}))

	ws.Route(ws.POST("").To(func(request *restful.Request, response *restful.Response) {
		rctx.NewReqCtx(request, response).WithLog("storageClass").
			WithHandle(api.CreateStorageClass).Do()
	}).Doc("创建storageClass").Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(entity.PageResult{}).
		Reads(v1.StorageClass{}).
		Param(ws.QueryParameter("cloud", "集群名称").Required(true).DataType("string")).
		Returns(200, "success", v1.StorageClass{}))

	ws.Route(ws.DELETE("/{name}").To(func(request *restful.Request, response *restful.Response) {
		rctx.NewReqCtx(request, response).WithLog("storageClass").
			WithHandle(api.DeleteStorageClass).Do()
	}).Doc("删除storageClass").Metadata(restfulspec.KeyOpenAPITags, tags).
		Param(ws.PathParameter("name", "storageClass名称").Required(true).DataType("string")).
		Param(ws.QueryParameter("cloud", "集群名称").Required(true).DataType("string")).
		Returns(200, "success", nil))

	return ws
}
