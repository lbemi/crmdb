package router

import (
	restfulspec "github.com/emicklei/go-restful-openapi/v2"
	"github.com/emicklei/go-restful/v3"
	"github.com/lbemi/lbemi/apps/istio/api"
	"github.com/lbemi/lbemi/apps/istio/api/vo"
	"github.com/lbemi/lbemi/pkg/common/entity"
	"github.com/lbemi/lbemi/pkg/rctx"
	"istio.io/client-go/pkg/apis/networking/v1beta1"
)

func IstioGatewayRoutes() *restful.WebService {
	ws := new(restful.WebService)
	ws.Path("/api/v1/gateways").Produces(restful.MIME_JSON)
	tags := []string{"kubernetes-Gateway"}

	ws.Route(ws.GET("/namespaces/{namespace}").To(func(request *restful.Request, response *restful.Response) {
		rctx.NewReqCtx(request, response).WithLog("Gateway").
			WithHandle(api.ListGateways).Do()
	}).Doc("获取Gateway列表").Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(&vo.PageGateway{}).
		Param(ws.QueryParameter("istio", "集群名称").Required(true).DataType("string")).
		Param(ws.QueryParameter("page", "页码").Required(false).DataType("string")).
		Param(ws.QueryParameter("limit", "每页条数").Required(false).DataType("string")).
		Param(ws.QueryParameter("name", "过滤名称，支持模糊查询").DataType("string")).
		Param(ws.QueryParameter("label", "过滤标签，支持模糊查询").DataType("string")).
		Param(ws.PathParameter("namespace", "命名空间").Required(true).DataType("string")).
		Returns(200, "success", &vo.PageGateway{}))

	ws.Route(ws.GET("/namespaces/{namespace}/{name}").To(func(request *restful.Request, response *restful.Response) {
		rctx.NewReqCtx(request, response).WithLog("Gateway").
			WithHandle(api.GetGateway).Do()
	}).Doc("获取Gateway信息").Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(v1beta1.Gateway{}).
		Param(ws.QueryParameter("istio", "集群名称").Required(true).DataType("string")).
		Param(ws.PathParameter("namespace", "命名空间").Required(true).DataType("string")).
		Param(ws.PathParameter("name", "Gateway名称").Required(true).DataType("string")).
		Returns(200, "success", v1beta1.Gateway{}))

	ws.Route(ws.POST("").To(func(request *restful.Request, response *restful.Response) {
		rctx.NewReqCtx(request, response).WithLog("Gateway").
			WithHandle(api.CreateGateway).Do()
	}).Doc("创建Gateway").Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(entity.PageResult{}).
		Reads(v1beta1.Gateway{}).
		Param(ws.QueryParameter("istio", "集群名称").Required(true).DataType("string")).
		Returns(200, "success", v1beta1.Gateway{}))

	ws.Route(ws.PUT("").To(func(request *restful.Request, response *restful.Response) {
		rctx.NewReqCtx(request, response).WithLog("Gateway").
			WithHandle(api.UpdateGateway).Do()
	}).Doc("修改Gateway").Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(entity.PageResult{}).
		Reads(v1beta1.Gateway{}).
		Param(ws.QueryParameter("istio", "集群名称").Required(true).DataType("string")).
		Returns(200, "success", v1beta1.Gateway{}))

	ws.Route(ws.DELETE("/namespaces/{namespace}/{name}").To(func(request *restful.Request, response *restful.Response) {
		rctx.NewReqCtx(request, response).WithLog("Gateway").
			WithHandle(api.DeleteGateway).Do()
	}).Doc("删除Gateway").Metadata(restfulspec.KeyOpenAPITags, tags).
		Param(ws.PathParameter("namespace", "命名空间").Required(true).DataType("string")).
		Param(ws.PathParameter("name", "Gateway名称").Required(true).DataType("string")).
		Param(ws.QueryParameter("istio", "集群名称").Required(true).DataType("string")).
		Returns(200, "success", nil))

	return ws
}
