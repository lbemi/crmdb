package cloud

import (
	"time"

	"github.com/lbemi/lbemi/pkg/core"
	"github.com/lbemi/lbemi/pkg/rctx"

	v1 "k8s.io/api/apps/v1"
)

func ListDeployments(rc *rctx.ReqCtx) {
	c := rc.Request.Request.Context()
	namespace := rc.PathParam("namespace")
	clusterName := rc.QueryCloud()
	pageParam := rc.GetPageQueryParam()
	name := rc.Query("name")
	label := rc.Query("label")
	rc.ResData = core.V1.Cluster(clusterName).Deployments(namespace).List(c, pageParam, name, label)
}

func GetDeployment(rc *rctx.ReqCtx) {
	c := rc.Request.Request.Context()
	namespace := rc.PathParam("namespace")
	clusterName := rc.QueryCloud()
	deploymentName := rc.PathParam("name")
	rc.ResData = core.V1.Cluster(clusterName).Deployments(namespace).Get(c, deploymentName)
}

func CreateDeployment(rc *rctx.ReqCtx) {
	c := rc.Request.Request.Context()
	clusterName := rc.QueryCloud()
	var deployment *v1.Deployment
	rc.ShouldBind(&deployment)
	rc.ResData = core.V1.Cluster(clusterName).Deployments(deployment.Namespace).Create(c, deployment)
}

func UpdateDeployment(rc *rctx.ReqCtx) {
	c := rc.Request.Request.Context()
	clusterName := rc.QueryCloud()
	var deployment *v1.Deployment
	rc.ShouldBind(&deployment)
	rc.ResData = core.V1.Cluster(clusterName).Deployments(deployment.Namespace).Update(c, deployment)
}

func RollBackDeployment(rc *rctx.ReqCtx) {
	c := rc.Request.Request.Context()
	clusterName := rc.QueryCloud()
	namespace := rc.PathParam("namespace")
	deploymentName := rc.PathParam("name")
	reversion := rc.PathParamInt64("reversion")
	rc.ResData = core.V1.Cluster(clusterName).Deployments(namespace).RollBack(c, deploymentName, reversion)
}
func ReDeployDeployment(rc *rctx.ReqCtx) {
	c := rc.Request.Request.Context()
	clusterName := rc.QueryCloud()
	namespace := rc.PathParam("namespace")
	deploymentName := rc.PathParam("name")
	deployment := core.V1.Cluster(clusterName).Deployments(namespace).Get(c, deploymentName)
	deployment.Spec.Template.Annotations = map[string]string{
		"lbemi.io/restartAt": time.Now().String(),
	}
	rc.ResData = core.V1.Cluster(clusterName).Deployments(deployment.Namespace).Update(c, deployment)
}

func DeleteDeployment(rc *rctx.ReqCtx) {
	c := rc.Request.Request.Context()
	clusterName := rc.QueryCloud()
	namespace := rc.PathParam("namespace")
	deploymentName := rc.PathParam("name")
	core.V1.Cluster(clusterName).Deployments(namespace).Delete(c, deploymentName)
}

func ScaleDeployments(rc *rctx.ReqCtx) {
	c := rc.Request.Request.Context()
	clusterName := rc.QueryCloud()
	namespace := rc.PathParam("namespace")
	deploymentName := rc.PathParam("name")
	replica := rc.PathParamInt("replica")
	core.V1.Cluster(clusterName).Deployments(namespace).Scale(c, deploymentName, int32(replica))
}

func GetDeploymentPods(rc *rctx.ReqCtx) {
	c := rc.Request.Request.Context()
	clusterName := rc.QueryCloud()
	namespace := rc.PathParam("namespace")
	deploymentName := rc.PathParam("name")
	pods, replicaSets := core.V1.Cluster(clusterName).Deployments(namespace).GetDeploymentPods(c, deploymentName)
	rc.ResData = map[string]interface{}{
		"pods":        pods,
		"replicaSets": replicaSets,
	}
}

func GetDeploymentEvents(rc *rctx.ReqCtx) {
	c := rc.Request.Request.Context()
	clusterName := rc.QueryCloud()
	namespace := rc.PathParam("namespace")
	deploymentName := rc.PathParam("name")
	rc.ResData = core.V1.Cluster(clusterName).Deployments(namespace).GetDeploymentEvent(c, deploymentName)
}

//
//func SearchDeployments(rc *rctx.ReqCtx) {
//	c := rc.Request.Request.Context()
//	clusterName := rc.QueryCloud()
//	namespace := rc.PathParam("namespace")
//	pageParam := rc.GetPageQueryParam()
//	key := rc.Query("key")
//	searchType := rc.QueryDefaultInt("type", 0)
//
//	deploymentList := core.V1.Cluster(clusterName).Deployments(namespace).Search(c, key, searchType)
//	if err != nil {
//		response.Fail(c, response.ErrOperateFailed)
//		return
//	}
//	// 处理分页
//	var pageQuery types.PageQuery
//	pageQuery.Total = len(deploymentList)
//
//	if pageQuery.Total <= limit {
//		pageQuery.Data = deploymentList
//	} else if page*limit >= pageQuery.Total {
//		pageQuery.Data = deploymentList[(page-1)*limit : pageQuery.Total]
//	} else {
//		pageQuery.Data = deploymentList[(page-1)*limit : page*limit]
//	}
//
//	response.Success(c, response.StatusOK, pageQuery)
//}
