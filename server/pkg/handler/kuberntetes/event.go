package kuberntetes

import (
	"context"
	"github.com/lbemi/lbemi/pkg/bootstrap/log"
	"github.com/lbemi/lbemi/pkg/common/store"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/labels"
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
	cli *store.Clients
	ns  string
}

func (e *event) List(ctx context.Context) ([]*v1.Event, error) {
	eventList, err := e.cli.SharedInformerFactory.Core().V1().Events().Lister().Events(e.ns).List(labels.Everything())
	if err != nil {
		log.Logger.Error(err)
	}
	sort.Slice(eventList, func(i, j int) bool {
		return eventList[j].ObjectMeta.CreationTimestamp.Time.Before(eventList[i].ObjectMeta.CreationTimestamp.Time)
	})
	return eventList, err
}

func (e *event) Get(ctx context.Context, name string) (*v1.Event, error) {
	event, err := e.cli.SharedInformerFactory.Core().V1().Events().Lister().Events(e.ns).Get(name)
	if err != nil {
		log.Logger.Error(err)
	}
	return event, err
}

func NewEvent(client *store.Clients, namespace string) *event {
	return &event{cli: client, ns: namespace}
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
