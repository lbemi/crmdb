package api

import (
	"github.com/lbemi/lbemi/pkg/core"
	"github.com/lbemi/lbemi/pkg/rctx"

	v1 "k8s.io/api/core/v1"
)

func ListNamespace(rc *rctx.ReqCtx) {
	c := rc.Request.Request.Context()
	clusterName := rc.QueryCloud()
	pageParam := rc.GetPageQueryParam()
	name := rc.Query("name")
	label := rc.Query("label")
	rc.ResData = core.V1.Cluster(clusterName).Namespaces().List(c, pageParam, name, label)
}

func GetNamespace(rc *rctx.ReqCtx) {
	c := rc.Request.Request.Context()
	clusterName := rc.QueryCloud()
	name := rc.PathParam("name")
	rc.ResData = core.V1.Cluster(clusterName).Namespaces().Get(c, name)
}

func DeleteNamespace(rc *rctx.ReqCtx) {
	c := rc.Request.Request.Context()
	clusterName := rc.QueryCloud()
	name := rc.PathParam("name")
	core.V1.Cluster(clusterName).Namespaces().Delete(c, name)
}

func CreateNamespace(rc *rctx.ReqCtx) {
	c := rc.Request.Request.Context()
	clusterName := rc.QueryCloud()
	var namespace v1.Namespace
	rc.ShouldBind(&namespace)
	rc.ResData = core.V1.Cluster(clusterName).Namespaces().Create(c, &namespace)

}

func UpdateNamespace(rc *rctx.ReqCtx) {
	c := rc.Request.Request.Context()
	clusterName := rc.QueryCloud()
	var namespace v1.Namespace
	rc.ShouldBind(&namespace)
	rc.ResData = core.V1.Cluster(clusterName).Namespaces().Update(c, &namespace)
}
