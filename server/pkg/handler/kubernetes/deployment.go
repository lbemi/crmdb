package kubernetes

import (
	"context"
	"fmt"
	"github.com/lbemi/lbemi/pkg/bootstrap/log"
	"github.com/lbemi/lbemi/pkg/services/k8s"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sort"
	"strconv"
)

type DeploymentGetter interface {
	Deployments(namespace string) IDeployment
}

type IDeployment interface {
	List(ctx context.Context) ([]*appsv1.Deployment, error)
	Get(ctx context.Context, name string) (*appsv1.Deployment, error)
	Create(ctx context.Context, obj *appsv1.Deployment) (*appsv1.Deployment, error)
	Update(ctx context.Context, obj *appsv1.Deployment) (*appsv1.Deployment, error)
	RollBack(ctx context.Context, depName string, reversion string) (*appsv1.Deployment, error)
	Delete(ctx context.Context, name string) error
	Scale(ctx context.Context, name string, replicaNum int32) error

	GetDeploymentPods(ctx context.Context, name string) ([]*corev1.Pod, []*appsv1.ReplicaSet, error)
	GetDeploymentEvent(ctx context.Context, name string) ([]*corev1.Event, error)
	Search(ctx context.Context, key string, searchType int) ([]*appsv1.Deployment, error)
}

type Deployment struct {
	k8s *k8s.Factory
}

var deployment *Deployment

func NewDeployment(k8s *k8s.Factory) *Deployment {
	deployment = &Deployment{k8s: k8s}
	return deployment
}

func (d *Deployment) List(ctx context.Context) ([]*appsv1.Deployment, error) {
	list, err := d.k8s.Deployment().List(ctx)

	if err != nil {
		log.Logger.Error(err)
		return []*appsv1.Deployment{}, fmt.Errorf("record not found")
	}
	//按时间排序
	sort.Slice(list, func(i, j int) bool {
		return list[j].ObjectMeta.GetCreationTimestamp().Time.Before(list[i].ObjectMeta.GetCreationTimestamp().Time)
	})

	return list, nil
}

func (d *Deployment) Get(ctx context.Context, name string) (*appsv1.Deployment, error) {
	dep, err := d.k8s.Deployment().Get(ctx, name)
	if err != nil {
		log.Logger.Error(err)
	}
	return dep, err
}

func (d *Deployment) Create(ctx context.Context, obj *appsv1.Deployment) (*appsv1.Deployment, error) {
	newDeployment, err := d.k8s.Deployment().Create(ctx, obj)
	if err != nil {
		log.Logger.Error(err)
	}
	return newDeployment, err
}

func (d *Deployment) Update(ctx context.Context, obj *appsv1.Deployment) (*appsv1.Deployment, error) {
	updateDeployment, err := d.k8s.Deployment().Update(ctx, obj)
	if err != nil {
		log.Logger.Error(err)
	}
	return updateDeployment, err
}

func (d *Deployment) RollBack(ctx context.Context, depName string, reversion string) (*appsv1.Deployment, error) {
	parseInt, err := strconv.ParseInt(reversion, 10, 64)
	if err != nil {
		log.Logger.Error("reversion not fount")
		return nil, fmt.Errorf("reversion not fount")
	}
	if parseInt < 0 {
		log.Logger.Error("reversion not fount")
		return nil, fmt.Errorf("reversion not fount")
	}

	deployment, err := d.k8s.Deployment().Get(ctx, depName)
	if err != nil {
		log.Logger.Error(err)
	}
	replicaSets, err := d.k8s.Replicaset().List(ctx)
	if err != nil {
		return nil, err
	}

	var replicaSet *appsv1.ReplicaSet
	//找到对应的rs
	for _, item := range replicaSets {
		if d.isRsFromDep(deployment, item) {
			if item.ObjectMeta.Annotations["deployment.kubernetes.io/revision"] == reversion {
				replicaSet = item
			}
		}
	}

	if replicaSet == nil {
		return nil, fmt.Errorf("reversion not fount")
	}

	deployment.Spec.Template = replicaSet.Spec.Template

	updateDeployment, err := d.k8s.Deployment().Update(ctx, deployment)
	if err != nil {
		log.Logger.Error(err)
	}

	return updateDeployment, err
}

func (d *Deployment) Delete(ctx context.Context, name string) error {
	err := d.k8s.Deployment().Delete(ctx, name)
	if err != nil {
		log.Logger.Error(err)
	}
	return err
}

func (d *Deployment) Scale(ctx context.Context, name string, replicaNum int32) error {
	err := d.k8s.Deployment().Scale(ctx, name, replicaNum)
	if err != nil {
		log.Logger.Error(err)
		return err
	}
	return nil
}

func (d *Deployment) GetDeploymentPods(ctx context.Context, name string) ([]*corev1.Pod, []*appsv1.ReplicaSet, error) {
	dep, err := d.k8s.Deployment().Get(ctx, name)
	if err != nil {
		return nil, nil, err
	}

	replicaSets, err := d.k8s.Replicaset().List(ctx)
	if err != nil {
		return nil, nil, err
	}
	res := make([]map[string]string, 0)

	PodReplicaSets := make([]*appsv1.ReplicaSet, 0)

	for _, item := range replicaSets {
		if d.isRsFromDep(dep, item) {
			selectorAsMap, err := v1.LabelSelectorAsMap(item.Spec.Selector)
			if err != nil {
				return nil, nil, err
			}
			PodReplicaSets = append(PodReplicaSets, item)
			res = append(res, selectorAsMap)
		}
	}

	pods, err := d.k8s.Pod().GetPodByLabels(ctx, dep.Namespace, res)

	if err != nil {
		return nil, nil, err
	}

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
	return pods, PodReplicaSets, nil

}

func (d *Deployment) GetDeploymentEvent(ctx context.Context, name string) ([]*corev1.Event, error) {
	events := make([]*corev1.Event, 0)
	eventList, err := d.k8s.Event().List(ctx)
	if err != nil {
		return nil, err
	}

	for _, item := range eventList {
		if (item.InvolvedObject.Kind == "ReplicaSet" || item.InvolvedObject.Kind == "Deployment") && item.InvolvedObject.Name == name {
			events = append(events, item)
		}
	}

	return events, nil

}

func (d *Deployment) Search(ctx context.Context, key string, searchType int) ([]*appsv1.Deployment, error) {
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
