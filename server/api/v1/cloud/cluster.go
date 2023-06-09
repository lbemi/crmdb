package cloud

import (
	"github.com/lbemi/lbemi/pkg/core"
	"github.com/lbemi/lbemi/pkg/model/form"
	"github.com/lbemi/lbemi/pkg/rctx"
)

func CreateCluster(rc *rctx.ReqCtx) {

	bytes := rctx.FormFile(rc, "file")
	var req form.ClusterReq
	//rctx.ShouldBind(rc, &req)
	req.Name = rctx.PostForm(rc, "name")
	req.KubeConfig = string(bytes)

	core.V1.Cluster(req.Name).Create(&req)
}

func ListCluster(rc *rctx.ReqCtx) {
	rc.ResData = core.V1.Cluster("").List()
}

func GetCluster(rc *rctx.ReqCtx) {
	clusterName := rctx.PathParam(rc, "name")
	rc.ResData = core.V1.Cluster("").GetByName(clusterName)
}

func DeleteCluster(rc *rctx.ReqCtx) {
	id := rctx.PathParamUint64(rc, "id")
	core.V1.Cluster("").Delete(id)
}
