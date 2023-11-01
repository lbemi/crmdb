package router

import (
	restfulspec "github.com/emicklei/go-restful-openapi/v2"
	"github.com/emicklei/go-restful/v3"
	"github.com/lbemi/lbemi/apps/asset/api"
	model "github.com/lbemi/lbemi/apps/asset/entity"
	"github.com/lbemi/lbemi/pkg/common/entity"
	"github.com/lbemi/lbemi/pkg/rctx"
)

func GroupRoutes() *restful.WebService {
	ws := new(restful.WebService)
	ws.Path("/api/v1/groups").Produces(restful.MIME_JSON)
	tags := []string{"groups"}

	ws.Route(ws.GET("/{id}").To(func(request *restful.Request, response *restful.Response) {
		rctx.NewReqCtx(request, response).WithLog("groups").WithHandle(api.GetGroup).Do()
	}).
		Doc("根据分组ID获取分组信息").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Param(ws.PathParameter("id", "分组id").DataType("string")).
		Reads(model.Group{}).
		Returns(200, "success", model.Group{}))

	ws.Route(ws.GET("/").To(func(request *restful.Request, response *restful.Response) {
		rctx.NewReqCtx(request, response).WithLog("groups").WithHandle(api.ListGroup).Do()
	}).
		Doc("查看分组列表信息").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(entity.PageResult{}).
		Returns(200, "success", entity.PageResult{}))

	ws.Route(ws.POST("").To(func(request *restful.Request, response *restful.Response) {
		rctx.NewReqCtx(request, response).WithLog("groups").WithHandle(api.AddGroup).Do()
	}).
		Doc("创建分组").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(model.Group{}).
		Returns(200, "success", nil))

	ws.Route(ws.PUT("").To(func(request *restful.Request, response *restful.Response) {
		rctx.NewReqCtx(request, response).WithLog("groups").WithHandle(api.UpdateGroup).Do()
	}).
		Doc("更新分组").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(model.Group{}).
		Returns(200, "success", nil))

	ws.Route(ws.DELETE("/{id}").To(func(request *restful.Request, response *restful.Response) {
		rctx.NewReqCtx(request, response).WithLog("groups").WithHandle(api.DeleteGroup).Do()
	}).
		Doc("根据分组ID删除分组").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Param(ws.PathParameter("id", "分组id").DataType("string")).
		Returns(200, "success", nil))

	return ws
}
