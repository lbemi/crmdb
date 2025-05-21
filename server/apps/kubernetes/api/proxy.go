package api

import (
	"github.com/lbemi/lbemi/pkg/core"
	"github.com/lbemi/lbemi/pkg/rctx"
)

func Proxy(rc *rctx.ReqCtx) {
	clusterName := rc.QueryCloud()
	core.V1.Cluster(clusterName).K8S().Proxy().GET(*rc.Request.Request.URL, rc.Response.ResponseWriter, rc.Request.Request)
}
