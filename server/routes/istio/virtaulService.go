package istio

import (
	restfulspec "github.com/emicklei/go-restful-openapi/v2"
	"github.com/emicklei/go-restful/v3"
	"github.com/lbemi/lbemi/api/istio"
	"github.com/lbemi/lbemi/pkg/model/form"
	"github.com/lbemi/lbemi/pkg/rctx"
	"istio.io/client-go/pkg/apis/networking/v1alpha3"
)

func IstioVirtualServiceRoutes() *restful.WebService {
	ws := new(restful.WebService)
	ws.Path("/api/v1/virtualservices").Produces(restful.MIME_JSON)
	tags := []string{"kubernetes-virtualservice"}

	ws.Route(ws.GET("/namespaces/{namespace}").To(func(request *restful.Request, response *restful.Response) {
		rctx.NewReqCtx(request, response).WithLog("virtualservice").
			WithHandle(istio.ListVirtualServices).Do()
	}).Doc("获取virtualservice列表").Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(&form.PageVirtualService{}).
		Param(ws.QueryParameter("istio", "集群名称").Required(true).DataType("string")).
		Param(ws.QueryParameter("page", "页码").Required(false).DataType("string")).
		Param(ws.QueryParameter("limit", "每页条数").Required(false).DataType("string")).
		Param(ws.QueryParameter("name", "过滤名称，支持模糊查询").DataType("string")).
		Param(ws.QueryParameter("label", "过滤标签，支持模糊查询").DataType("string")).
		Param(ws.PathParameter("namespace", "命名空间").Required(true).DataType("string")).
		Returns(200, "success", &form.PageVirtualService{}))

	ws.Route(ws.GET("/namespaces/{namespace}/{name}").To(func(request *restful.Request, response *restful.Response) {
		rctx.NewReqCtx(request, response).WithLog("virtualservice").
			WithHandle(istio.GetVirtualService).Do()
	}).Doc("获取virtualservice信息").Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(v1alpha3.VirtualService{}).
		Param(ws.QueryParameter("istio", "集群名称").Required(true).DataType("string")).
		Param(ws.PathParameter("namespace", "命名空间").Required(true).DataType("string")).
		Param(ws.PathParameter("name", "virtualservice名称").Required(true).DataType("string")).
		Returns(200, "success", v1alpha3.VirtualService{}))

	ws.Route(ws.POST("").To(func(request *restful.Request, response *restful.Response) {
		rctx.NewReqCtx(request, response).WithLog("virtualservice").
			WithHandle(istio.CreateVirtualService).Do()
	}).Doc("创建virtualservice").Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(form.PageResult{}).
		Reads(v1alpha3.VirtualService{}).
		Param(ws.QueryParameter("istio", "集群名称").Required(true).DataType("string")).
		Returns(200, "success", v1alpha3.VirtualService{}))

	ws.Route(ws.PUT("").To(func(request *restful.Request, response *restful.Response) {
		rctx.NewReqCtx(request, response).WithLog("virtualservice").
			WithHandle(istio.UpdateVirtualService).Do()
	}).Doc("修改virtualservice").Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(form.PageResult{}).
		Reads(v1alpha3.VirtualService{}).
		Param(ws.QueryParameter("istio", "集群名称").Required(true).DataType("string")).
		Returns(200, "success", v1alpha3.VirtualService{}))

	ws.Route(ws.DELETE("/namespaces/{namespace}/{name}").To(func(request *restful.Request, response *restful.Response) {
		rctx.NewReqCtx(request, response).WithLog("virtualservice").
			WithHandle(istio.DeleteVirtualService).Do()
	}).Doc("删除virtualservice").Metadata(restfulspec.KeyOpenAPITags, tags).
		Param(ws.PathParameter("namespace", "命名空间").Required(true).DataType("string")).
		Param(ws.PathParameter("name", "virtualservice名称").Required(true).DataType("string")).
		Param(ws.QueryParameter("istio", "集群名称").Required(true).DataType("string")).
		Returns(200, "success", nil))

	return ws
}
