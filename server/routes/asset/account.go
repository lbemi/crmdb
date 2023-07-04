package asset

import (
	restfulspec "github.com/emicklei/go-restful-openapi/v2"
	"github.com/emicklei/go-restful/v3"
	"github.com/lbemi/lbemi/api/asset"
	model "github.com/lbemi/lbemi/pkg/model/asset"
	"github.com/lbemi/lbemi/pkg/model/form"
	"github.com/lbemi/lbemi/pkg/rctx"
)

func AccountRoutes() *restful.WebService {
	ws := new(restful.WebService)
	ws.Path("/api/v1/accounts").Produces(restful.MIME_JSON)
	tags := []string{"accounts"}

	ws.Route(ws.GET("/{id}").To(func(request *restful.Request, response *restful.Response) {
		rctx.NewReqCtx(request, response).WithLog("accounts").WithHandle(asset.GetAccount).Do()
	}).
		Doc("根据账户ID获取账户信息").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Param(ws.PathParameter("id", "账户id").DataType("string")).
		Reads(model.Account{}).
		Returns(200, "success", model.Account{}))

	ws.Route(ws.GET("/").To(func(request *restful.Request, response *restful.Response) {
		rctx.NewReqCtx(request, response).WithLog("accounts").WithHandle(asset.ListAccount).Do()
	}).
		Doc("查看账户列表信息").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(form.PageResult{}).
		Returns(200, "success", form.PageResult{}))

	ws.Route(ws.POST("").To(func(request *restful.Request, response *restful.Response) {
		rctx.NewReqCtx(request, response).WithLog("accounts").WithHandle(asset.AddAccount).Do()
	}).
		Doc("创建账户").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(model.Account{}).
		Returns(200, "success", nil))

	ws.Route(ws.PUT("").To(func(request *restful.Request, response *restful.Response) {
		rctx.NewReqCtx(request, response).WithLog("accounts").WithHandle(asset.UpdateAccount).Do()
	}).
		Doc("更新账户").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(model.Account{}).
		Returns(200, "success", nil))

	ws.Route(ws.DELETE("/{id}").To(func(request *restful.Request, response *restful.Response) {
		rctx.NewReqCtx(request, response).WithLog("accounts").WithHandle(asset.DeleteAccount).Do()
	}).
		Doc("根据账户ID删除账户").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Param(ws.PathParameter("id", "账户id").DataType("string")).
		Returns(200, "success", nil))

	return ws
}
