package istio

import (
	"github.com/lbemi/lbemi/pkg/core"
	"github.com/lbemi/lbemi/pkg/rctx"
	"istio.io/client-go/pkg/apis/networking/v1beta1"
)

func ListVirtualServices(rc *rctx.ReqCtx) {
	c := rc.Request.Request.Context()
	namespace := rc.PathParam("namespace")
	clusterName := rc.QueryCloud()
	pageParam := rc.GetPageQueryParam()
	name := rc.Query("name")
	label := rc.Query("label")
	rc.ResData = core.V1.Cluster(clusterName).VirtualServices(namespace).List(c, pageParam, name, label)
}

func GetVirtualService(rc *rctx.ReqCtx) {
	c := rc.Request.Request.Context()
	clusterName := rc.Query("cloud")
	namespace := rc.PathParam("namespace")
	VirtualServiceName := rc.PathParam("VirtualServiceName")
	rc.ResData = core.V1.Cluster(clusterName).VirtualServices(namespace).Get(c, VirtualServiceName)
}

func CreateVirtualService(rc *rctx.ReqCtx) {
	c := rc.Request.Request.Context()
	clusterName := rc.Query("cloud")
	var VirtualService *v1beta1.VirtualService
	rc.ShouldBind(&VirtualService)
	rc.ResData = core.V1.Cluster(clusterName).VirtualServices(VirtualService.Namespace).Create(c, VirtualService)
}

func UpdateVirtualService(rc *rctx.ReqCtx) {
	c := rc.Request.Request.Context()
	clusterName := rc.Query("cloud")
	var VirtualService *v1beta1.VirtualService
	rc.ShouldBind(&VirtualService)
	rc.ResData = core.V1.Cluster(clusterName).VirtualServices(VirtualService.Namespace).Update(c, VirtualService)
}

func DeleteVirtualService(rc *rctx.ReqCtx) {
	c := rc.Request.Request.Context()
	clusterName := rc.Query("cloud")
	namespace := rc.PathParam("namespace")
	VirtualServiceName := rc.PathParam("name")
	core.V1.Cluster(clusterName).VirtualServices(namespace).Delete(c, VirtualServiceName)
}
