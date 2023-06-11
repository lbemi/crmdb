package kubernetes

import (
	"context"
	"fmt"
	"github.com/lbemi/lbemi/pkg/bootstrap/log"
	"github.com/lbemi/lbemi/pkg/model"
	"github.com/lbemi/lbemi/pkg/model/form"
	"github.com/lbemi/lbemi/pkg/restfulx"
	"github.com/lbemi/lbemi/pkg/services/k8s"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sort"
	"strconv"
	"strings"
)

type DeploymentGetter interface {
	Deployments(namespace string) IDeployment
}

type IDeployment interface {
	List(ctx context.Context, query *model.PageParam, name string, label string) *form.PageResult
	Get(ctx context.Context, name string) *appsv1.Deployment
	Create(ctx context.Context, obj *appsv1.Deployment) *appsv1.Deployment
	Update(ctx context.Context, obj *appsv1.Deployment) *appsv1.Deployment
	RollBack(ctx context.Context, depName string, reversion int64) *appsv1.Deployment
	Delete(ctx context.Context, name string)
	Scale(ctx context.Context, name string, replicaNum int32)

	GetDeploymentPods(ctx context.Context, name string) ([]*corev1.Pod, []*appsv1.ReplicaSet)
	GetDeploymentEvent(ctx context.Context, name string) []*corev1.Event
	Search(ctx context.Context, key string, searchType int) []*appsv1.Deployment
}

type Deployment struct {
	k8s *k8s.Factory
}

var deployment *Deployment

func NewDeployment(k8s *k8s.Factory) *Deployment {
	deployment = &Deployment{k8s: k8s}
	return deployment
}

func (d *Deployment) List(ctx context.Context, query *model.PageParam, name string, label string) *form.PageResult {
	data := d.k8s.Deployment().List(ctx)
	res := &form.PageResult{}
	total := len(data)
	var deploymentList = make([]*appsv1.Deployment, 0)

	if name != "" {
		for _, item := range data {
			if strings.Contains(item.Name, name) {
				deploymentList = append(deploymentList, item)
			}
		}
		data = deploymentList
	}

	if label != "" {
		for _, item := range data {
			if strings.Contains(item.Name, label) {
				deploymentList = append(deploymentList, item)
			}
		}
		data = deploymentList
	}
	// 未传递分页查询参数
	if query.Limit == 0 && query.Page == 0 {
		res.Data = data
	} else {
		if total <= query.Limit {
			res.Data = data
		} else if query.Page*query.Limit >= total {
			res.Data = data[(query.Page-1)*query.Limit : total]
		} else {
			res.Data = data[(query.Page-1)*query.Limit : query.Page*query.Limit]
		}
	}
	res.Total = int64(total)

	return res
}

func (d *Deployment) Get(ctx context.Context, name string) *appsv1.Deployment {
	return d.k8s.Deployment().Get(ctx, name)
}

func (d *Deployment) Create(ctx context.Context, obj *appsv1.Deployment) *appsv1.Deployment {
	return d.k8s.Deployment().Create(ctx, obj)
}

func (d *Deployment) Update(ctx context.Context, obj *appsv1.Deployment) *appsv1.Deployment {
	return d.k8s.Deployment().Update(ctx, obj)
}

func (d *Deployment) RollBack(ctx context.Context, depName string, reversion int64) *appsv1.Deployment {
	if reversion < 0 {
		restfulx.ErrNotNilDebug(fmt.Errorf("reversion not fount"), restfulx.OperatorErr)
	}
	dep := d.k8s.Deployment().Get(ctx, depName)
	replicaSets := d.k8s.Replicaset().List(ctx)
	var replicaSet *appsv1.ReplicaSet
	//找到对应的rs
	for _, item := range replicaSets {
		if d.isRsFromDep(dep, item) {
			log.Logger.Infof("reversion:  ---> %s", item.ObjectMeta.Annotations["deployment.kubernetes.io/revision"])
			if item.ObjectMeta.Annotations["deployment.kubernetes.io/revision"] == strconv.FormatInt(reversion, 10) {
				replicaSet = item
			}
		}
	}
	if replicaSet == nil {
		restfulx.ErrNotNilDebug(fmt.Errorf("reversion not fount"), restfulx.OperatorErr)
	}
	dep.Spec.Template = replicaSet.Spec.Template
	updateDeployment := d.k8s.Deployment().Update(ctx, dep)
	return updateDeployment
}

func (d *Deployment) Delete(ctx context.Context, name string) {
	d.k8s.Deployment().Delete(ctx, name)
}

func (d *Deployment) Scale(ctx context.Context, name string, replicaNum int32) {
	d.k8s.Deployment().Scale(ctx, name, replicaNum)
}

func (d *Deployment) GetDeploymentPods(ctx context.Context, name string) ([]*corev1.Pod, []*appsv1.ReplicaSet) {
	dep := d.k8s.Deployment().Get(ctx, name)
	replicaSets := d.k8s.Replicaset().List(ctx)
	res := make([]map[string]string, 0)
	PodReplicaSets := make([]*appsv1.ReplicaSet, 0)
	for _, item := range replicaSets {
		if d.isRsFromDep(dep, item) {
			selectorAsMap, err := v1.LabelSelectorAsMap(item.Spec.Selector)
			restfulx.ErrNotNilDebug(err, restfulx.GetResourceErr)
			PodReplicaSets = append(PodReplicaSets, item)
			res = append(res, selectorAsMap)
		}
	}

	pods := d.k8s.Pod().GetPodByLabels(ctx, dep.Namespace, res)
	sort.Slice(pods, func(i, j int) bool {
		// sort by creation timestamp in descending order
		if pods[j].ObjectMeta.GetCreationTimestamp().Time.Before(pods[i].ObjectMeta.GetCreationTimestamp().Time) {
			return true
		} else if pods[i].ObjectMeta.GetCreationTimestamp().Time.Before(pods[j].ObjectMeta.GetCreationTimestamp().Time) {
			return false
		}

		// if the creation timestamps are equal, sort by name in ascending order
		return pods[i].ObjectMeta.GetName() < pods[j].ObjectMeta.GetName()
	})

	sort.Slice(PodReplicaSets, func(i, j int) bool {
		// sort by creation timestamp in descending order
		if PodReplicaSets[j].ObjectMeta.GetCreationTimestamp().Time.Before(PodReplicaSets[i].ObjectMeta.GetCreationTimestamp().Time) {
			return true
		} else if PodReplicaSets[i].ObjectMeta.GetCreationTimestamp().Time.Before(PodReplicaSets[j].ObjectMeta.GetCreationTimestamp().Time) {
			return false
		}

		// if the creation timestamps are equal, sort by name in ascending order
		return PodReplicaSets[i].ObjectMeta.GetName() < PodReplicaSets[j].ObjectMeta.GetName()
	})
	return pods, PodReplicaSets

}

func (d *Deployment) GetDeploymentEvent(ctx context.Context, name string) []*corev1.Event {
	events := make([]*corev1.Event, 0)
	eventList := d.k8s.Event().List(ctx)
	for _, item := range eventList {
		if (item.InvolvedObject.Kind == "ReplicaSet" || item.InvolvedObject.Kind == "Deployment") && item.InvolvedObject.Name == name {
			events = append(events, item)
		}
	}
	return events

}

func (d *Deployment) Search(ctx context.Context, key string, searchType int) []*appsv1.Deployment {
	return d.k8s.Deployment().Search(ctx, key, searchType)
}

func (d *Deployment) isRsFromDep(dep *appsv1.Deployment, rs *appsv1.ReplicaSet) bool {
	for _, ref := range rs.OwnerReferences {
		if ref.Kind == "Deployment" && ref.Name == dep.Name {
			return true
		}
	}
	return false
}
