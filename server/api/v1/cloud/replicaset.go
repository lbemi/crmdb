package cloud

import (
	"github.com/lbemi/lbemi/pkg/core"
	"github.com/lbemi/lbemi/pkg/rctx"
)

func ListReplicaSets(rc *rctx.ReqCtx) {
	c := rc.Request.Request.Context()
	namespace := rc.PathParam("namespace")
	clusterName := rc.QueryCloud()
	pageParam := rc.GetPageQueryParam()
	name := rc.Query("name")
	label := rc.Query("label")
	rc.ResData = core.V1.Cluster(clusterName).Replicaset(namespace).List(c, pageParam, name, label)

}

func GetReplicaSet(rc *rctx.ReqCtx) {
	c := rc.Request.Request.Context()
	namespace := rc.PathParam("namespace")
	clusterName := rc.QueryCloud()
	replicasetName := rc.PathParam("name")
	rc.ResData = core.V1.Cluster(clusterName).Replicaset(namespace).Get(c, replicasetName)

}
