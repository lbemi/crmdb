package router

import (
	restfulspec "github.com/emicklei/go-restful-openapi/v2"
	"github.com/emicklei/go-restful/v3"
	"github.com/lbemi/lbemi/apps/asset/api"
	model "github.com/lbemi/lbemi/apps/asset/entity"
	"github.com/lbemi/lbemi/pkg/common/entity"
	"github.com/lbemi/lbemi/pkg/rctx"
)

// AccountRoutes returns a *restful.WebService.
//
// This function does not take any parameters.
// It returns a *restful.WebService.
func AccountRoutes() *restful.WebService {
	ws := new(restful.WebService)
	ws.Path("/api/v1/accounts").Produces(restful.MIME_JSON)
	tags := []string{"accounts"}

	ws.Route(ws.GET("/{id}").To(func(request *restful.Request, response *restful.Response) {
		rctx.NewReqCtx(request, response).WithLog("accounts").WithHandle(api.GetAccount).Do()
	}).
		Doc("根据账户ID获取账户信息").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Param(ws.PathParameter("id", "账户id").DataType("string")).
		Reads(model.Account{}).
		Returns(200, "success", model.Account{}))

	ws.Route(ws.GET("/").To(func(request *restful.Request, response *restful.Response) {
		rctx.NewReqCtx(request, response).WithLog("accounts").WithHandle(api.ListAccount).Do()
	}).
		Doc("查看账户列表信息").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Param(ws.QueryParameter("name", "名称,模糊查询").DataType("string")).
		Param(ws.QueryParameter("user_name", "登录名,模糊查询").DataType("string")).
		Reads(entity.PageResult{}).
		Returns(200, "success", entity.PageResult{}))

	ws.Route(ws.POST("").To(func(request *restful.Request, response *restful.Response) {
		rctx.NewReqCtx(request, response).WithLog("accounts").WithHandle(api.AddAccount).Do()
	}).
		Doc("创建账户").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(model.Account{}).
		Returns(200, "success", nil))

	ws.Route(ws.PUT("").To(func(request *restful.Request, response *restful.Response) {
		rctx.NewReqCtx(request, response).WithLog("accounts").WithHandle(api.UpdateAccount).Do()
	}).
		Doc("更新账户").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(model.Account{}).
		Returns(200, "success", nil))

	ws.Route(ws.DELETE("/{id}").To(func(request *restful.Request, response *restful.Response) {
		rctx.NewReqCtx(request, response).WithLog("accounts").WithHandle(api.DeleteAccount).Do()
	}).
		Doc("根据账户ID删除账户").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Param(ws.PathParameter("id", "账户id").DataType("string")).
		Returns(200, "success", nil))

	return ws
}
