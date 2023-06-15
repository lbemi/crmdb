package kubernetes

import (
	"context"

	"github.com/lbemi/lbemi/pkg/model"
	"github.com/lbemi/lbemi/pkg/model/form"
	"github.com/lbemi/lbemi/pkg/services/k8s"

	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/labels"
)

type EventGetter interface {
	Events(namespace string) IEvent
}

type IEvent interface {
	List(ctx context.Context, query *model.PageParam) *form.PageResult
	Get(ctx context.Context, name string) *v1.Event
	ListByLabels(ctx context.Context, labelsData labels.Set) []*v1.Event
}

type event struct {
	k8s *k8s.Factory
}

func NewEvent(k8s *k8s.Factory) *event {
	return &event{k8s: k8s}
}

func (e *event) List(ctx context.Context, query *model.PageParam) *form.PageResult {
	data := e.k8s.Event().List(ctx)
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

func (e *event) ListByLabels(ctx context.Context, labelsData labels.Set) []*v1.Event {
	return e.k8s.Event().ListByLabels(ctx, labelsData)
}

func (e *event) Get(ctx context.Context, name string) *v1.Event {
	return e.k8s.Event().Get(ctx, name)
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
