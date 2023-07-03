package asset

import (
	restfulspec "github.com/emicklei/go-restful-openapi/v2"
	"github.com/emicklei/go-restful/v3"
	"github.com/lbemi/lbemi/api/asset"
	model "github.com/lbemi/lbemi/pkg/model/asset"
	"github.com/lbemi/lbemi/pkg/model/form"
	"github.com/lbemi/lbemi/pkg/rctx"
)

func ResourceAccountRoutes() *restful.WebService {
	ws := new(restful.WebService)
	ws.Path("/api/v1/rbas").Produces(restful.MIME_JSON)
	tags := []string{"resourceBindAccount"}

	ws.Route(ws.GET("/{id}").To(func(request *restful.Request, response *restful.Response) {
		rctx.NewReqCtx(request, response).WithLog("resourceBindAccount").WithHandle(asset.Get).Do()
	}).
		Doc("根据ID获取资源绑定信息").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Param(ws.PathParameter("id", "id").DataType("string")).
		Reads(model.HostAccount{}).
		Returns(200, "success", model.HostAccount{}))

	ws.Route(ws.GET("").To(func(request *restful.Request, response *restful.Response) {
		rctx.NewReqCtx(request, response).WithLog("resourceBindAccount").WithHandle(asset.List).Do()
	}).
		Doc("查看资源绑定列表").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(form.PageResult{}).
		Returns(200, "success", form.PageResult{}))

	ws.Route(ws.POST("").To(func(request *restful.Request, response *restful.Response) {
		rctx.NewReqCtx(request, response).WithLog("resourceBindAccount").WithHandle(asset.BindAccount).Do()
	}).
		Doc("创建资源绑定").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(model.HostAccount{}).
		Returns(200, "success", nil))

	ws.Route(ws.PUT("").To(func(request *restful.Request, response *restful.Response) {
		rctx.NewReqCtx(request, response).WithLog("resourceBindAccount").WithHandle(asset.UpdateHostAccount).Do()
	}).
		Doc("更新资源绑定").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(model.HostAccount{}).
		Returns(200, "success", nil))

	ws.Route(ws.DELETE("/{id}").To(func(request *restful.Request, response *restful.Response) {
		rctx.NewReqCtx(request, response).WithLog("groups").WithHandle(asset.UnbindAccount).Do()
	}).
		Doc("根据分组ID删除资源绑定").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Param(ws.PathParameter("id", "id").DataType("string")).
		Returns(200, "success", nil))

	return ws
}
