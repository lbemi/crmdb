package kubernetes

import (
	"context"
	"github.com/lbemi/lbemi/pkg/services/k8s"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/labels"
	"sort"
)

type EventGetter interface {
	Events(namespace string) IEvent
}

type IEvent interface {
	List(ctx context.Context) []*v1.Event
	Get(ctx context.Context, name string) *v1.Event
	ListByLabels(ctx context.Context, labelsData labels.Set) []*v1.Event
}

type event struct {
	k8s *k8s.Factory
}

func NewEvent(k8s *k8s.Factory) *event {
	return &event{k8s: k8s}
}

func (e *event) List(ctx context.Context) []*v1.Event {
	eventList := e.k8s.Event().List(ctx)
	sort.Slice(eventList, func(i, j int) bool {
		return eventList[j].ObjectMeta.CreationTimestamp.Time.Before(eventList[i].ObjectMeta.CreationTimestamp.Time)
	})
	return eventList
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
