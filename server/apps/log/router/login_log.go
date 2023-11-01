package router

import (
	restfulspec "github.com/emicklei/go-restful-openapi/v2"
	"github.com/emicklei/go-restful/v3"
	"github.com/lbemi/lbemi/apps/log/api"
	logModel "github.com/lbemi/lbemi/apps/log/entity"
	"github.com/lbemi/lbemi/pkg/common/entity"

	"github.com/lbemi/lbemi/pkg/rctx"
)

func LoginLogRoutes() *restful.WebService {
	ws := new(restful.WebService)
	ws.Path("/api/v1/logs/login").Produces(restful.MIME_JSON)
	tags := []string{"logs"}
	// 获取登录日志列表
	ws.Route(ws.GET("").To(
		func(request *restful.Request, response *restful.Response) {
			rctx.NewReqCtx(request, response).WithLog("loginLog").WithHandle(api.ListLoginLog).Do()
		}).
		Doc("获取登录日志列表").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Param(ws.QueryParameter("page", "page").DataType("int")).
		Param(ws.QueryParameter("limit", "limit").DataType("int")).
		Param(ws.QueryParameter("status", "过滤状态").DataType("string")).
		Param(ws.QueryParameter("name", "过滤名称").DataType("string")).
		Writes(entity.PageResult{}).
		Returns(200, "success", entity.PageResult{}))

	// 根据ID获取登录日志
	ws.Route(ws.GET("/{id}").To(func(request *restful.Request, response *restful.Response) {
		rctx.NewReqCtx(request, response).WithLog("loginLog").WithHandle(api.GetLoginLog).Do()
	}).Doc("根据id获取登录日志").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Param(ws.QueryParameter("id", "登录日志id").DataType("int")).
		Writes(logModel.LogLogin{}).
		Returns(200, "success", logModel.LogLogin{}))

	// 删除指定日志
	ws.Route(ws.DELETE("").To(
		func(request *restful.Request, response *restful.Response) {
			rctx.NewReqCtx(request, response).
				WithLog("loginLog").
				WithHandle(api.DeleteLoginLogs).
				Do()
		}).
		Doc("删除指定日志").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Param(ws.QueryParameter("ids", "日志id列表，例如： '1,2,3' ").DataType("string")).
		Returns(200, "success", nil))
	// 删除所有日志
	ws.Route(ws.DELETE("/all").To(
		func(request *restful.Request, response *restful.Response) {
			rctx.NewReqCtx(request, response).
				WithLog("loginLog").
				WithHandle(api.DeleteAllLoginLog).
				Do()
		}).
		Doc("删除所有日志").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Returns(200, "success", nil))

	return ws
}
