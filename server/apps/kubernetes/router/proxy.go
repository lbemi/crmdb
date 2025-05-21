package router

import (
	restfulspec "github.com/emicklei/go-restful-openapi/v2"
	"github.com/emicklei/go-restful/v3"
	"github.com/lbemi/lbemi/apps/kubernetes/api"
	"github.com/lbemi/lbemi/pkg/rctx"
)

func KubernetesProxyRoutes() *restful.WebService {
	ws := new(restful.WebService)
	ws.Path("/api/v1/proxy").Produces(restful.MIME_JSON)
	tags := []string{"kubernetes-proxy"}

	// 使用 Method 方法处理所有 HTTP 方法
	ws.Route(ws.Method("*").Path("/{clusterName}/{*}").To(func(request *restful.Request, response *restful.Response) {
		rctx.NewReqCtx(request, response).WithLog("proxy").
			WithHandle(api.Proxy).Do()
	}).Doc("代理请求").Metadata(restfulspec.KeyOpenAPITags, tags).
		Param(ws.PathParameter("clusterName", "集群名称").Required(true).DataType("string")).
		Param(ws.PathParameter("*", "代理路径").Required(true).DataType("string")).
		Returns(200, "success", nil))

	return ws
}
