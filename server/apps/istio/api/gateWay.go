package api

import (
	"github.com/lbemi/lbemi/pkg/core"
	"github.com/lbemi/lbemi/pkg/rctx"
	"istio.io/client-go/pkg/apis/networking/v1beta1"
)

func ListGateways(rc *rctx.ReqCtx) {
	c := rc.Request.Request.Context()
	namespace := rc.PathParam("namespace")
	clusterName := rc.QueryCloud()
	pageParam := rc.GetPageQueryParam()
	name := rc.Query("name")
	label := rc.Query("label")
	rc.ResData = core.V1.Cluster(clusterName).Istio().Gateways(namespace).List(c, pageParam, name, label)
}

func GetGateway(rc *rctx.ReqCtx) {
	c := rc.Request.Request.Context()
	clusterName := rc.Query("cloud")
	namespace := rc.PathParam("namespace")
	GatewayName := rc.PathParam("name")
	rc.ResData = core.V1.Cluster(clusterName).Istio().Gateways(namespace).Get(c, GatewayName)
}

func CreateGateway(rc *rctx.ReqCtx) {
	c := rc.Request.Request.Context()
	clusterName := rc.Query("cloud")
	var Gateway *v1beta1.Gateway
	rc.ShouldBind(&Gateway)
	rc.ResData = core.V1.Cluster(clusterName).Istio().Gateways(Gateway.Namespace).Create(c, Gateway)
}

func UpdateGateway(rc *rctx.ReqCtx) {
	c := rc.Request.Request.Context()
	clusterName := rc.Query("cloud")
	var Gateway *v1beta1.Gateway
	rc.ShouldBind(&Gateway)
	rc.ResData = core.V1.Cluster(clusterName).Istio().Gateways(Gateway.Namespace).Update(c, Gateway)
}

func DeleteGateway(rc *rctx.ReqCtx) {
	c := rc.Request.Request.Context()
	clusterName := rc.Query("cloud")
	namespace := rc.PathParam("namespace")
	GatewayName := rc.PathParam("name")
	core.V1.Cluster(clusterName).Istio().Gateways(namespace).Delete(c, GatewayName)
}
