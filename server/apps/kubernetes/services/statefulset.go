package services

import (
	"context"
	"github.com/lbemi/lbemi/pkg/cache"
	"github.com/lbemi/lbemi/pkg/common/entity"
	"github.com/lbemi/lbemi/pkg/restfulx"
	"github.com/lbemi/lbemi/pkg/util"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sort"
	"strings"

	v1 "k8s.io/api/apps/v1"
	"k8s.io/apimachinery/pkg/labels"
)

type StatefulSetGetter interface {
	StatefulSets(namespace string) IStatefulSet
}

type IStatefulSet interface {
	List(ctx context.Context, query *entity.PageParam, name string, label string) *entity.PageResult
	Get(ctx context.Context, name string) *v1.StatefulSet
	GetStatefulSetPods(ctx context.Context, name string) ([]*corev1.Pod, []*v1.ControllerRevision)
	GetStatefulSetEvent(ctx context.Context, name string) []*corev1.Event
	Create(ctx context.Context, obj *v1.StatefulSet) *v1.StatefulSet
	Update(ctx context.Context, obj *v1.StatefulSet) *v1.StatefulSet
	Delete(ctx context.Context, name string)
}

type StatefulSet struct {
	cli       *cache.ClientConfig
	namespace string
}

func NewStatefulSet(client *cache.ClientConfig, ns string) *StatefulSet {
	return &StatefulSet{cli: client, namespace: ns}
}

func (s *StatefulSet) List(ctx context.Context, query *entity.PageParam, name string, label string) *entity.PageResult {
	data, err := s.cli.SharedInformerFactory.Apps().V1().StatefulSets().Lister().StatefulSets(s.namespace).List(labels.Everything())
	restfulx.ErrNotNilDebug(err, restfulx.GetResourceErr)
	res := &entity.PageResult{}
	var statefulSetList = make([]*v1.StatefulSet, 0)
	if name != "" {
		for _, item := range data {
			if strings.Contains(item.Name, name) {
				statefulSetList = append(statefulSetList, item)
			}
		}
		data = statefulSetList
	}

	if label != "" {
		for _, item := range data {
			if strings.Contains(labels.FormatLabels(item.Labels), label) {
				statefulSetList = append(statefulSetList, item)
			}
		}
		data = statefulSetList
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

func (s *StatefulSet) Get(ctx context.Context, name string) *v1.StatefulSet {
	sts, err := s.cli.SharedInformerFactory.Apps().V1().StatefulSets().Lister().StatefulSets(s.namespace).Get(name)
	restfulx.ErrNotNilDebug(err, restfulx.GetResourceErr)
	return sts
}

func (s *StatefulSet) GetStatefulSetPods(ctx context.Context, name string) ([]*corev1.Pod, []*v1.ControllerRevision) {
	statefulSet := s.Get(ctx, name)
	pods := make([]*corev1.Pod, 0)
	podList, err := s.cli.SharedInformerFactory.Core().V1().Pods().Lister().Pods(s.namespace).List(labels.Everything())
	restfulx.ErrNotNilDebug(err, restfulx.OperatorErr)
	for _, item := range podList {
		if metav1.IsControlledBy(item, statefulSet) {
			pods = append(pods, item)
		}
	}
	restoreGVKForList(pods)

	sort.Slice(podList, func(i, j int) bool {
		// sort by creation timestamp in descending order
		if podList[j].ObjectMeta.GetCreationTimestamp().Time.Before(podList[i].ObjectMeta.GetCreationTimestamp().Time) {
			return true
		} else if podList[i].ObjectMeta.GetCreationTimestamp().Time.Before(podList[j].ObjectMeta.GetCreationTimestamp().Time) {
			return false
		}

		// if the creation timestamps are equal, sort by name in ascending order
		return podList[i].ObjectMeta.GetName() < podList[j].ObjectMeta.GetName()
	})

	controllerRevisionList, err := s.cli.SharedInformerFactory.Apps().V1().ControllerRevisions().Lister().List(labels.Everything())
	restfulx.ErrNotNilDebug(err, restfulx.OperatorErr)
	controllerRevisions := make([]*v1.ControllerRevision, 0)

	for _, item := range controllerRevisionList {
		if metav1.IsControlledBy(item, statefulSet) {
			controllerRevisions = append(controllerRevisions, item)
		}
	}

	return pods, controllerRevisions

}

func (s *StatefulSet) GetStatefulSetEvent(ctx context.Context, name string) []*corev1.Event {
	eventList, err := s.cli.SharedInformerFactory.Core().V1().Events().Lister().Events(s.namespace).List(labels.Everything())
	restfulx.ErrNotNilDebug(err, restfulx.GetResourceErr)
	events := make([]*corev1.Event, 0)
	for _, item := range eventList {
		if (item.InvolvedObject.Kind == "ReplicaSet" || item.InvolvedObject.Kind == "StatefulSet") && item.InvolvedObject.Name == name {
			events = append(events, item)
		}
	}
	return events

}

func (s *StatefulSet) Create(ctx context.Context, obj *v1.StatefulSet) *v1.StatefulSet {
	newStatefulSet, err := s.cli.ClientSet.AppsV1().StatefulSets(s.namespace).Create(ctx, obj, metav1.CreateOptions{})
	restfulx.ErrNotNilDebug(err, restfulx.OperatorErr)
	return newStatefulSet
}

func (s *StatefulSet) Update(ctx context.Context, obj *v1.StatefulSet) *v1.StatefulSet {
	updateStatefulSet, err := s.cli.ClientSet.AppsV1().StatefulSets(s.namespace).Update(ctx, obj, metav1.UpdateOptions{})
	restfulx.ErrNotNilDebug(err, restfulx.OperatorErr)
	return updateStatefulSet
}

func (s *StatefulSet) Delete(ctx context.Context, name string) {
	err := s.cli.ClientSet.AppsV1().StatefulSets(s.namespace).Delete(ctx, name, metav1.DeleteOptions{})
	restfulx.ErrNotNilDebug(err, restfulx.OperatorErr)
}

type StatefulSetHandler struct {
}

func NewStatefulSetHandle() *StatefulSetHandler {
	return &StatefulSetHandler{}
}

func (s *StatefulSetHandler) OnAdd(obj interface{}) {
	//TODO implement me
}

func (s *StatefulSetHandler) OnUpdate(oldObj, newObj interface{}) {
	//TODO implement me
}

func (s *StatefulSetHandler) OnDelete(obj interface{}) {
	//TODO implement me
}
