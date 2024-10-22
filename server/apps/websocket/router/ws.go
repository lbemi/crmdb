package router

import (
	restfulspec "github.com/emicklei/go-restful-openapi/v2"
	"github.com/emicklei/go-restful/v3"
	"github.com/lbemi/lbemi/apps/websocket/api"
	"github.com/lbemi/lbemi/pkg/rctx"
)

func WebSocketRoutes() *restful.WebService {
	ws := new(restful.WebService)
	ws.Path("/api/v1/ws").Produces(restful.MIME_JSON)
	tags := []string{"kubernetes-websocket"}

	ws.Route(ws.GET("/{cluster}/{type}").To(func(request *restful.Request, response *restful.Response) {
		rctx.NewReqCtx(request, response).WithLog("websocket").WithNoRes().
			WithHandle(api.WebSocket).Do()
	}).Doc("连接websocket").Metadata(restfulspec.KeyOpenAPITags, tags).
		Param(ws.PathParameter("cluster", "集群名称").Required(true).DataType("string")).
		Param(ws.PathParameter("type", "websocket类型").Required(true).DataType("string")).
		Returns(200, "success", nil))

	ws.Route(ws.GET("/send").To(func(request *restful.Request, response *restful.Response) {
		rctx.NewReqCtx(request, response).WithLog("websocket").WithNoRes().
			WithHandle(api.WsSendAll).Do()
	}).Doc("发送websocket消息").Metadata(restfulspec.KeyOpenAPITags, tags).
		Param(ws.QueryParameter("msg", "消息").Required(true).DataType("string")).
		Returns(200, "success", nil))

	return ws
}
