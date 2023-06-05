package logsys

import (
	restfulspec "github.com/emicklei/go-restful-openapi/v2"
	"github.com/emicklei/go-restful/v3"
	"github.com/lbemi/lbemi/api/v1/logsys"
	"github.com/lbemi/lbemi/pkg/model/form"
	logModel "github.com/lbemi/lbemi/pkg/model/logsys"
	"github.com/lbemi/lbemi/pkg/rctx"
)

func OperatorLogRoutes() *restful.WebService {
	ws := new(restful.WebService)
	ws.Path("/api/v1/logs/operator").Produces(restful.MIME_JSON)
	tags := []string{"logs"}
	// 获取操作日志列表
	ws.Route(ws.GET("").To(
		func(request *restful.Request, response *restful.Response) {
			rctx.NewReqCtx(request, response).WithLog("operatorLog").WithHandle(logsys.ListOperatorLog).Do()
		}).
		Doc("获取操作日志列表").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Param(ws.QueryParameter("page", "page").DataType("int")).
		Param(ws.QueryParameter("limit", "limit").DataType("int")).
		Param(ws.QueryParameter("type", "过滤状态").DataType("string")).
		Param(ws.QueryParameter("title", "过滤名称").DataType("string")).
		Param(ws.QueryParameter("name", "过滤名称").DataType("string")).
		Writes(form.PageResult{}).
		Returns(200, "success", form.PageResult{}))

	// 根据ID获取操作日志
	ws.Route(ws.GET("/{id}").To(func(request *restful.Request, response *restful.Response) {
		rctx.NewReqCtx(request, response).WithLog("operatorLog").WithHandle(logsys.GetOperatorLog).Do()
	}).Doc("根据id获取操作日志").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Param(ws.QueryParameter("id", "操作日志id").DataType("int")).
		Writes(logModel.LogOperator{}).
		Returns(200, "success", logModel.LogOperator{}))

	// 删除指定日志
	ws.Route(ws.DELETE("").To(
		func(request *restful.Request, response *restful.Response) {
			rctx.NewReqCtx(request, response).
				WithLog("operatorLog").
				WithHandle(logsys.DeleteOperatorLogs).
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
				WithLog("operatorLog").
				WithHandle(logsys.DeleteAllOperatorLog).
				Do()
		}).
		Doc("删除所有日志").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Returns(200, "success", nil))

	return ws
}
