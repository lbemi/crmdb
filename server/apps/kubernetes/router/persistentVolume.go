package router

import (
	restfulspec "github.com/emicklei/go-restful-openapi/v2"
	"github.com/emicklei/go-restful/v3"
	"github.com/lbemi/lbemi/apps/kubernetes/api"
	"github.com/lbemi/lbemi/pkg/common/entity"
	corev1 "k8s.io/api/core/v1"

	"github.com/lbemi/lbemi/pkg/rctx"
)

func KubernetesPersistentVolumeRoutes() *restful.WebService {
	ws := new(restful.WebService)
	ws.Path("/api/v1/persistentvolumes").Produces(restful.MIME_JSON)
	tags := []string{"kubernetes-pv"}

	ws.Route(ws.GET("").To(func(request *restful.Request, response *restful.Response) {
		rctx.NewReqCtx(request, response).WithLog("pv").
			WithHandle(api.ListPersistentVolume).Do()
	}).Doc("获取pv列表").Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(&entity.PageResult{}).
		Param(ws.QueryParameter("cloud", "集群名称").Required(true).DataType("string")).
		Param(ws.QueryParameter("page", "页码").Required(false).DataType("string")).
		Param(ws.QueryParameter("limit", "每页条数").Required(false).DataType("string")).
		Param(ws.QueryParameter("name", "过滤名称，支持模糊查询").DataType("string")).
		Param(ws.QueryParameter("label", "过滤标签，支持模糊查询").DataType("string")).
		Returns(200, "success", &entity.PageResult{}))

	ws.Route(ws.GET("/{name}").To(func(request *restful.Request, response *restful.Response) {
		rctx.NewReqCtx(request, response).WithLog("pv").
			WithHandle(api.GetPersistentVolume).Do()
	}).Doc("获取pv信息").Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(corev1.PersistentVolume{}).
		Param(ws.QueryParameter("cloud", "集群名称").Required(true).DataType("string")).
		Param(ws.PathParameter("name", "pv名称").Required(true).DataType("string")).
		Returns(200, "success", corev1.PersistentVolume{}))

	ws.Route(ws.POST("").To(func(request *restful.Request, response *restful.Response) {
		rctx.NewReqCtx(request, response).WithLog("pv").
			WithHandle(api.CreatePersistentVolume).Do()
	}).Doc("创建pv").Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(entity.PageResult{}).
		Reads(corev1.PersistentVolume{}).
		Param(ws.QueryParameter("cloud", "集群名称").Required(true).DataType("string")).
		Returns(200, "success", corev1.PersistentVolume{}))

	ws.Route(ws.DELETE("/{name}").To(func(request *restful.Request, response *restful.Response) {
		rctx.NewReqCtx(request, response).WithLog("pv").
			WithHandle(api.DeletePersistentVolume).Do()
	}).Doc("删除pv").Metadata(restfulspec.KeyOpenAPITags, tags).
		Param(ws.PathParameter("name", "pv名称").Required(true).DataType("string")).
		Param(ws.QueryParameter("cloud", "集群名称").Required(true).DataType("string")).
		Returns(200, "success", nil))

	return ws
}
