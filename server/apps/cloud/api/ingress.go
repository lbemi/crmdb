package api

import (
	"github.com/lbemi/lbemi/pkg/core"
	"github.com/lbemi/lbemi/pkg/rctx"

	v1 "k8s.io/api/networking/v1"
)

func ListIngresses(rc *rctx.ReqCtx) {
	c := rc.Request.Request.Context()
	namespace := rc.PathParam("namespace")
	clusterName := rc.QueryCloud()
	pageParam := rc.GetPageQueryParam()
	name := rc.Query("name")
	label := rc.Query("label")
	rc.ResData = core.V1.Cluster(clusterName).Ingresses(namespace).List(c, pageParam, name, label)
}

func GetIngress(rc *rctx.ReqCtx) {
	c := rc.Request.Request.Context()
	namespace := rc.PathParam("namespace")
	clusterName := rc.QueryCloud()
	ingressName := rc.PathParam("name")
	rc.ResData = core.V1.Cluster(clusterName).Ingresses(namespace).Get(c, ingressName)
}

func CreateIngress(rc *rctx.ReqCtx) {
	c := rc.Request.Request.Context()
	clusterName := rc.QueryCloud()
	var ingress v1.Ingress
	rc.ShouldBind(&ingress)
	rc.ResData = core.V1.Cluster(clusterName).Ingresses(ingress.Namespace).Create(c, &ingress)
}

func UpdateIngress(rc *rctx.ReqCtx) {
	c := rc.Request.Request.Context()
	clusterName := rc.QueryCloud()
	var ingress v1.Ingress
	rc.ShouldBind(&ingress)
	rc.ResData = core.V1.Cluster(clusterName).Ingresses(ingress.Namespace).Update(c, &ingress)
}

func DeleteIngress(rc *rctx.ReqCtx) {
	c := rc.Request.Request.Context()
	namespace := rc.PathParam("namespace")
	clusterName := rc.QueryCloud()
	ingressName := rc.PathParam("name")
	core.V1.Cluster(clusterName).Ingresses(namespace).Delete(c, ingressName)
}
