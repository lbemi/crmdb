package services

import (
	"context"
	"fmt"
	entity2 "github.com/lbemi/lbemi/apps/kubernetes/entity"
	"github.com/lbemi/lbemi/pkg/cache"
	"github.com/lbemi/lbemi/pkg/common/entity"
	"github.com/lbemi/lbemi/pkg/global"
	"github.com/lbemi/lbemi/pkg/util"
	"k8s.io/apimachinery/pkg/runtime"
	"sort"
	"strconv"
	"strings"

	"github.com/lbemi/lbemi/pkg/restfulx"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
)

type DeploymentGetter interface {
	Deployments(namespace string) IDeployment
}

type IDeployment interface {
	List(ctx context.Context, query *entity.PageParam, name string, label string) *entity.PageResult
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
	client    *cache.ClientConfig
	namespace string
}

func NewDeployment(client *cache.ClientConfig, namespace string) *Deployment {
	return &Deployment{client: client, namespace: namespace}

}

func (d *Deployment) List(ctx context.Context, query *entity.PageParam, name string, label string) *entity.PageResult {
	data, err := d.client.SharedInformerFactory.Apps().V1().Deployments().Lister().Deployments(d.namespace).List(labels.Everything())
	restfulx.ErrNotNilDebug(err, restfulx.GetResourceErr)
	res := &entity.PageResult{}

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
			if strings.Contains(labels.FormatLabels(item.Labels), label) {
				deploymentList = append(deploymentList, item)
			}
		}
		data = deploymentList
	}
	//按时间排序
	sort.SliceStable(data, func(i, j int) bool {
		return data[j].ObjectMeta.GetCreationTimestamp().Time.Before(data[i].ObjectMeta.GetCreationTimestamp().Time)
	})
	total := len(data)
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
	dep, err := d.client.SharedInformerFactory.Apps().V1().Deployments().Lister().Deployments(d.namespace).Get(name)
	restfulx.ErrNotNilDebug(err, restfulx.GetResourceErr)
	return dep
}

func (d *Deployment) Create(ctx context.Context, obj *appsv1.Deployment) *appsv1.Deployment {
	newDeployment, err := d.client.ClientSet.AppsV1().Deployments(d.namespace).Create(ctx, obj, metav1.CreateOptions{})
	restfulx.ErrNotNilDebug(err, restfulx.OperatorErr)
	return newDeployment
}

func (d *Deployment) Update(ctx context.Context, obj *appsv1.Deployment) *appsv1.Deployment {
	result, err := d.client.ClientSet.AppsV1().Deployments(d.namespace).Update(ctx, obj, metav1.UpdateOptions{})
	restfulx.ErrNotNilDebug(err, restfulx.OperatorErr)
	return result
}

func (d *Deployment) RollBack(ctx context.Context, depName string, reversion int64) *appsv1.Deployment {
	if reversion < 0 {
		restfulx.ErrNotNilDebug(fmt.Errorf("reversion not fount"), restfulx.OperatorErr)
	}

	dep := d.Get(ctx, depName)
	replicaSets, err := d.client.SharedInformerFactory.Apps().V1().ReplicaSets().Lister().ReplicaSets(d.namespace).List(labels.Everything())
	restfulx.ErrNotNilDebug(err, restfulx.GetResourceErr)

	var replicaSet *appsv1.ReplicaSet
	//找到对应的rs
	for _, item := range replicaSets {
		if d.isRsFromDep(dep, item) {
			global.Logger.Infof("reversion:  ---> %s", item.ObjectMeta.Annotations["deployment.kubernetes.io/revision"])
			if item.ObjectMeta.Annotations["deployment.kubernetes.io/revision"] == strconv.FormatInt(reversion, 10) {
				replicaSet = item
			}
		}
	}
	if replicaSet == nil {
		restfulx.ErrNotNilDebug(fmt.Errorf("reversion not fount"), restfulx.OperatorErr)
	}
	dep.Spec.Template = replicaSet.Spec.Template
	updateDeployment := d.Update(ctx, dep)
	return updateDeployment
}

func (d *Deployment) Delete(ctx context.Context, name string) {
	err := d.client.ClientSet.AppsV1().Deployments(d.namespace).Delete(ctx, name, metav1.DeleteOptions{})
	restfulx.ErrNotNilDebug(err, restfulx.OperatorErr)
}

func (d *Deployment) Scale(ctx context.Context, name string, replicaNum int32) {
	oldScale, err := d.client.ClientSet.AppsV1().Deployments(d.namespace).GetScale(ctx, name, metav1.GetOptions{})
	restfulx.ErrNotNilDebug(err, restfulx.OperatorErr)
	oldScale.Spec.Replicas = replicaNum
	_, err = d.client.ClientSet.AppsV1().Deployments(d.namespace).UpdateScale(ctx, name, oldScale, metav1.UpdateOptions{})
	restfulx.ErrNotNilDebug(err, restfulx.OperatorErr)
}

func (d *Deployment) GetDeploymentPods(ctx context.Context, name string) ([]*corev1.Pod, []*appsv1.ReplicaSet) {
	dep := d.Get(ctx, name)
	replicaSets, err := d.client.SharedInformerFactory.Apps().V1().ReplicaSets().Lister().ReplicaSets(d.namespace).List(labels.Everything())
	restfulx.ErrNotNilDebug(err, restfulx.GetResourceErr)

	rsLabels := make([]map[string]string, 0)
	PodReplicaSets := make([]*appsv1.ReplicaSet, 0)

	// TODO 可以使用k8s内置方法v1.IsControlledBy()替换
	for _, item := range replicaSets {
		if d.isRsFromDep(dep, item) {
			selectorAsMap, err := v1.LabelSelectorAsMap(item.Spec.Selector)
			restfulx.ErrNotNilDebug(err, restfulx.GetResourceErr)
			PodReplicaSets = append(PodReplicaSets, item)
			rsLabels = append(rsLabels, selectorAsMap)
		}
	}

	res := make([]*corev1.Pod, 0)
	pods, err := d.client.SharedInformerFactory.Core().V1().Pods().Lister().Pods(d.namespace).List(labels.Everything())
	restfulx.ErrNotNilDebug(err, restfulx.OperatorErr)
	for _, item := range pods {
		for _, l := range rsLabels {
			i := 0
			for k1, v1 := range l {
				for k2, v2 := range item.Labels {
					if k1 == k2 && v1 == v2 {
						i++
					}
				}
			}
			if i == len(l) {
				res = append(res, item)
			}
		}
	}
	restoreGVKForList(res)

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
	eventList, err := d.client.SharedInformerFactory.Core().V1().Events().Lister().Events(d.namespace).List(labels.Everything())
	restfulx.ErrNotNilDebug(err, restfulx.GetResourceErr)
	events := make([]*corev1.Event, 0)
	for _, item := range eventList {
		if (item.InvolvedObject.Kind == "ReplicaSet" || item.InvolvedObject.Kind == "Deployment") && item.InvolvedObject.Name == name {
			events = append(events, item)
		}
	}
	return events

}

func (d *Deployment) Search(ctx context.Context, key string, searchType int) []*appsv1.Deployment {
	var deploymentList []*appsv1.Deployment
	deployments, err := d.client.SharedInformerFactory.Apps().V1().Deployments().Lister().Deployments(d.namespace).List(labels.Everything())
	restfulx.ErrNotNilDebug(err, restfulx.GetResourceErr)
	switch searchType {
	case entity2.SearchByName:
		// 遍历deployment，如果name包含key则保存返回
		for _, item := range deployments {
			if strings.Contains(item.Name, key) {
				deploymentList = append(deploymentList, item)
			}
		}
	case entity2.SearchByLabel:
		// 遍历deployment，如果name包含key则保存返回
		for _, item := range deployments {
			for k, label := range item.Labels {
				if strings.Contains(label, key) || strings.Contains(k, key) {
					deploymentList = append(deploymentList, item)
					break
				}
			}
		}
	default:
		restfulx.ErrNotNilDebug(fmt.Errorf("参数错误"), restfulx.ParamErr)
	}

	sort.Slice(deploymentList, func(i, j int) bool {
		return deploymentList[j].ObjectMeta.GetCreationTimestamp().Time.Before(deploymentList[i].ObjectMeta.GetCreationTimestamp().Time)
	})

	return deploymentList
}

func (d *Deployment) isRsFromDep(dep *appsv1.Deployment, rs *appsv1.ReplicaSet) bool {
	return v1.IsControlledBy(rs, dep)

	//for _, ref := range rs.OwnerReferences {
	//	if ref.Kind == "Deployment" && ref.Name == dep.Name {
	//		return true
	//	}
	//}
	//return false
}

type DeploymentHandler struct {
	client      *cache.ClientConfig
	clusterName string
}

func NewDeploymentHandler(client *cache.ClientConfig, clusterName string) *DeploymentHandler {
	return &DeploymentHandler{client: client, clusterName: clusterName}
}

func (d *DeploymentHandler) OnAdd(obj interface{}, isInInitialList bool) {
	d.notifyDeployments(obj)
}

func (d *DeploymentHandler) OnUpdate(oldObj, newObj interface{}) {
	d.notifyDeployments(newObj)
}

func (d *DeploymentHandler) OnDelete(obj interface{}) {
	d.notifyDeployments(obj)
}

func (d *DeploymentHandler) notifyDeployments(obj interface{}) {
	namespace := obj.(*appsv1.Deployment).Namespace
	deployments, err := d.client.SharedInformerFactory.Apps().V1().Deployments().Lister().Deployments(namespace).List(labels.Everything())
	if err != nil {
		global.Logger.Error(err)
	}

	//按时间排序
	sort.SliceStable(deployments, func(i, j int) bool {
		return deployments[j].ObjectMeta.GetCreationTimestamp().Time.Before(deployments[i].ObjectMeta.GetCreationTimestamp().Time)
	})

	go cache.WebsocketStore.SendClusterResource(d.clusterName, "deployment", map[string]interface{}{
		"cluster": d.clusterName,
		"type":    "deployment",
		"result": map[string]interface{}{
			"namespace": namespace,
			"data":      deployments,
		},
	})
}
func restoreGVKForList(podList []*corev1.Pod) {
	objects := make([]runtime.Object, len(podList))
	for i, p := range podList {
		objects[i] = p
	}
	util.RestoreGVKForList(objects)
}
