package v1beat1

import (
	"github.com/lbemi/lbemi/pkg/core"
	"github.com/lbemi/lbemi/pkg/model/form"
	"github.com/lbemi/lbemi/pkg/rctx"
)

func CreateCluster(rc *rctx.ReqCtx) {
	bytes := rc.FormFile("file")
	var req form.ClusterReq
	req.Name = rc.PostForm("name")
	req.KubeConfig = string(bytes)
	core.V1.Cluster(req.Name).Create(&req)
}

func ListCluster(rc *rctx.ReqCtx) {
	rc.ResData = core.V1.Cluster("").List()
}

func GetCluster(rc *rctx.ReqCtx) {
	clusterName := rc.PathParam("name")
	rc.ResData = core.V1.Cluster("").GetByName(clusterName)
}

func DeleteCluster(rc *rctx.ReqCtx) {
	id := rc.PathParamUint64("id")
	core.V1.Cluster("").Delete(id)
}
