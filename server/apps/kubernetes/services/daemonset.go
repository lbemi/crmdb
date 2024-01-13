package services

import (
	"context"
	"github.com/lbemi/lbemi/pkg/cache"
	"github.com/lbemi/lbemi/pkg/common/entity"
	"github.com/lbemi/lbemi/pkg/global"
	"github.com/lbemi/lbemi/pkg/restfulx"
	"github.com/lbemi/lbemi/pkg/util"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sort"
	"strings"

	appsv1 "k8s.io/api/apps/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
)

type DaemonSetGetter interface {
	DaemonSets(namespace string) IDaemonSet
}

type IDaemonSet interface {
	List(ctx context.Context, query *entity.PageParam, name string, label string) *entity.PageResult
	Get(ctx context.Context, name string) *appsv1.DaemonSet
	GetDaemonSetPods(ctx context.Context, name string) ([]*corev1.Pod, []*appsv1.ReplicaSet)
	GetDaemonSetEvent(ctx context.Context, name string) []*corev1.Event
	Create(ctx context.Context, obj *appsv1.DaemonSet) *appsv1.DaemonSet
	Update(ctx context.Context, obj *appsv1.DaemonSet) *appsv1.DaemonSet
	Delete(ctx context.Context, name string)
}

type daemonSet struct {
	client    *cache.ClientConfig
	namespace string
}

func NewDaemonSet(client *cache.ClientConfig, namespace string) IDaemonSet {
	return &daemonSet{client: client, namespace: namespace}
}

func (d *daemonSet) List(ctx context.Context, query *entity.PageParam, name string, label string) *entity.PageResult {
	data, err := d.client.SharedInformerFactory.Apps().V1().DaemonSets().Lister().DaemonSets(d.namespace).List(labels.Everything())
	restfulx.ErrNotNilDebug(err, restfulx.GetResourceErr)

	res := &entity.PageResult{}
	var daemonSetList = make([]*appsv1.DaemonSet, 0)

	if name != "" {
		for _, item := range data {
			if strings.Contains(item.Name, name) {
				daemonSetList = append(daemonSetList, item)
			}
		}
		data = daemonSetList
	}

	if label != "" {
		for _, item := range data {
			if strings.Contains(labels.FormatLabels(item.Labels), label) {
				daemonSetList = append(daemonSetList, item)
			}
		}
		data = daemonSetList
	}
	//按时间排序
	sort.SliceStable(data, func(i, j int) bool {
		return data[j].ObjectMeta.GetCreationTimestamp().Time.Before(data[i].ObjectMeta.GetCreationTimestamp().Time)
	})

	for _, item := range data {
		util.RestoreGVK(item)
	}
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

func (d *daemonSet) Get(ctx context.Context, name string) *appsv1.DaemonSet {
	dep, err := d.client.SharedInformerFactory.Apps().V1().DaemonSets().Lister().DaemonSets(d.namespace).Get(name)
	restfulx.ErrNotNilDebug(err, restfulx.GetResourceErr)

	return dep
}
func (d *daemonSet) GetDaemonSetPods(ctx context.Context, name string) ([]*corev1.Pod, []*appsv1.ReplicaSet) {
	dae := d.Get(ctx, name)

	res := make([]*corev1.Pod, 0)
	pods, err := d.client.SharedInformerFactory.Core().V1().Pods().Lister().Pods(d.namespace).List(labels.Everything())
	restfulx.ErrNotNilDebug(err, restfulx.OperatorErr)
	for _, item := range pods {
		if v1.IsControlledBy(item, dae) {
			res = append(res, item)
		}
	}

	restoreGVKForList(res)

	sort.Slice(res, func(i, j int) bool {
		// sort by creation timestamp in descending order
		if res[j].ObjectMeta.GetCreationTimestamp().Time.Before(res[i].ObjectMeta.GetCreationTimestamp().Time) {
			return true
		} else if res[i].ObjectMeta.GetCreationTimestamp().Time.Before(res[j].ObjectMeta.GetCreationTimestamp().Time) {
			return false
		}

		// if the creation timestamps are equal, sort by name in ascending order
		return res[i].ObjectMeta.GetName() < res[j].ObjectMeta.GetName()
	})

	return res, nil

}

func (d *daemonSet) GetDaemonSetEvent(ctx context.Context, name string) []*corev1.Event {
	eventList, err := d.client.SharedInformerFactory.Core().V1().Events().Lister().Events(d.namespace).List(labels.Everything())
	restfulx.ErrNotNilDebug(err, restfulx.GetResourceErr)
	events := make([]*corev1.Event, 0)
	for _, item := range eventList {
		if (item.InvolvedObject.Kind == "ReplicaSet" || item.InvolvedObject.Kind == "DaemonSet") && item.InvolvedObject.Name == name {
			events = append(events, item)
		}
	}
	return events

}
func (d *daemonSet) Create(ctx context.Context, obj *appsv1.DaemonSet) *appsv1.DaemonSet {
	newDaemonSet, err := d.client.ClientSet.AppsV1().DaemonSets(d.namespace).Create(ctx, obj, metav1.CreateOptions{})
	restfulx.ErrNotNilDebug(err, restfulx.OperatorErr)
	return newDaemonSet
}

func (d *daemonSet) Update(ctx context.Context, obj *appsv1.DaemonSet) *appsv1.DaemonSet {
	updateDaemonSet, err := d.client.ClientSet.AppsV1().DaemonSets(d.namespace).Update(ctx, obj, metav1.UpdateOptions{})
	restfulx.ErrNotNilDebug(err, restfulx.OperatorErr)
	return updateDaemonSet
}

func (d *daemonSet) Delete(ctx context.Context, name string) {
	err := d.client.ClientSet.AppsV1().DaemonSets(d.namespace).Delete(ctx, name, metav1.DeleteOptions{})
	restfulx.ErrNotNilDebug(err, restfulx.OperatorErr)
}

func (d *daemonSet) isRsFromDep(dep *appsv1.DaemonSet, rs *appsv1.ReplicaSet) bool {
	return v1.IsControlledBy(rs, dep)
}

type DaemonSetHandler struct {
	client      *cache.ClientConfig
	clusterName string
}

func NewDaemonSetHandler(client *cache.ClientConfig, clusterName string) *DaemonSetHandler {
	return &DaemonSetHandler{client: client, clusterName: clusterName}
}

func (d *DaemonSetHandler) OnAdd(obj interface{}, isInInitialList bool) {
	d.notifyDeployments(obj)
}

func (d *DaemonSetHandler) OnUpdate(oldObj, newObj interface{}) {
	d.notifyDeployments(newObj)
}

func (d *DaemonSetHandler) OnDelete(obj interface{}) {
	d.notifyDeployments(obj)
}

func (d *DaemonSetHandler) notifyDeployments(obj interface{}) {
	namespace := obj.(*appsv1.DaemonSet).Namespace
	daemonSets, err := d.client.SharedInformerFactory.Apps().V1().DaemonSets().Lister().DaemonSets(namespace).List(labels.Everything())
	if err != nil {
		global.Logger.Error(err)
	}

	//按时间排序
	sort.SliceStable(daemonSets, func(i, j int) bool {
		return daemonSets[j].ObjectMeta.GetCreationTimestamp().Time.Before(daemonSets[i].ObjectMeta.GetCreationTimestamp().Time)
	})

	go cache.WebsocketStore.SendClusterResource(d.clusterName, "daemonset", map[string]interface{}{
		"cluster": d.clusterName,
		"type":    "daemonset",
		"result": map[string]interface{}{
			"namespace": namespace,
			"data":      daemonSets,
		},
	})
}
