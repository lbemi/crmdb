package kubernetes

import (
	"context"
	"github.com/lbemi/lbemi/pkg/common/store"
	"github.com/lbemi/lbemi/pkg/model"
	"github.com/lbemi/lbemi/pkg/model/form"
	"github.com/lbemi/lbemi/pkg/restfulx"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/labels"
	"sort"
)

type EventGetter interface {
	Events(namespace string) IEvent
}

type IEvent interface {
	List(ctx context.Context, query *model.PageParam) *form.PageResult
	Get(ctx context.Context, name string) *v1.Event
	ListByLabels(ctx context.Context, labelsData labels.Set) []*v1.Event
}

type Event struct {
	client    *store.ClientConfig
	namespace string
}

func NewEvent(client *store.ClientConfig, namespace string) *Event {
	return &Event{client: client, namespace: namespace}
}

func (e *Event) List(ctx context.Context, query *model.PageParam) *form.PageResult {
	data, err := e.client.SharedInformerFactory.Core().V1().Events().Lister().Events(e.namespace).List(labels.Everything())
	restfulx.ErrNotNilDebug(err, restfulx.GetResourceErr)
	//按时间排序
	sort.SliceStable(data, func(i, j int) bool {
		return data[j].ObjectMeta.GetCreationTimestamp().Time.Before(data[i].ObjectMeta.GetCreationTimestamp().Time)
	})
	res := &form.PageResult{}
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

func (e *Event) ListByLabels(ctx context.Context, labelsData labels.Set) []*v1.Event {
	selector := labels.SelectorFromSet(labelsData)
	eventList, err := e.client.SharedInformerFactory.Core().V1().Events().Lister().Events(e.namespace).List(selector)
	restfulx.ErrNotNilDebug(err, restfulx.GetResourceErr)
	return eventList
}

func (e *Event) Get(ctx context.Context, name string) *v1.Event {
	res, err := e.client.SharedInformerFactory.Core().V1().Events().Lister().Events(e.namespace).Get(name)
	restfulx.ErrNotNilDebug(err, restfulx.GetResourceErr)
	return res
}

type EventHandler struct {
}

func (e *EventHandler) OnAdd(obj interface{}) {
	//TODO implement me
}

func (e *EventHandler) OnUpdate(oldObj, newObj interface{}) {
	//TODO implement me
}

func (e *EventHandler) OnDelete(obj interface{}) {
	//TODO implement me
}

func NewEventHandler() *EventHandler {
	return &EventHandler{}
}
