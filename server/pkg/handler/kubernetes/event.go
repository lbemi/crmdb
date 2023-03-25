package kubernetes

import (
	"context"
	"github.com/lbemi/lbemi/pkg/bootstrap/log"
	"github.com/lbemi/lbemi/pkg/services/k8s"
	v1 "k8s.io/api/core/v1"
	"sort"
)

type EventGetter interface {
	Events(namespace string) IEvent
}

type IEvent interface {
	List(ctx context.Context) ([]*v1.Event, error)
	Get(ctx context.Context, name string) (*v1.Event, error)
}

type event struct {
	k8s *k8s.Factory
}

func NewEvent(k8s *k8s.Factory) *event {
	return &event{k8s: k8s}
}

func (e *event) List(ctx context.Context) ([]*v1.Event, error) {
	eventList, err := e.k8s.Event().List(ctx)
	if err != nil {
		log.Logger.Error(err)
	}
	sort.Slice(eventList, func(i, j int) bool {
		return eventList[j].ObjectMeta.CreationTimestamp.Time.Before(eventList[i].ObjectMeta.CreationTimestamp.Time)
	})
	return eventList, err
}

func (e *event) Get(ctx context.Context, name string) (*v1.Event, error) {
	event, err := e.k8s.Event().Get(ctx, name)
	if err != nil {
		log.Logger.Error(err)
	}
	return event, err
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
