package cloud

import (
	"github.com/lbemi/lbemi/pkg/core"
	"github.com/lbemi/lbemi/pkg/rctx"
)

func ListEvents(rc *rctx.ReqCtx) {
	c := rc.Request.Request.Context()
	namespace := rc.PathParam("namespace")
	clusterName := rc.QueryCloud()
	pageParam := rc.GetPageQueryParam()
	rc.ResData = core.V1.Cluster(clusterName).Events(namespace).List(c, pageParam)
}

func GetEvent(rc *rctx.ReqCtx) {
	c := rc.Request.Request.Context()
	namespace := rc.PathParam("namespace")
	clusterName := rc.QueryCloud()
	eventName := rc.PathParam("name")
	rc.ResData = core.V1.Cluster(clusterName).Events(namespace).Get(c, eventName)
}
