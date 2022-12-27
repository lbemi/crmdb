package kuberntetes

import (
	"context"
	"github.com/lbemi/lbemi/pkg/bootstrap/log"
	"github.com/lbemi/lbemi/pkg/services/cloud"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/labels"
)

type EventGetter interface {
	Events(namespace string) IEvent
}

type IEvent interface {
	List(ctx context.Context) ([]*v1.Event, error)
	Get(ctx context.Context, name string) (*v1.Event, error)
}

type event struct {
	cli *cloud.Clients
	ns  string
}

func (e *event) List(ctx context.Context) ([]*v1.Event, error) {
	eventList, err := e.cli.SharedInformerFactory.Core().V1().Events().Lister().Events(e.ns).List(labels.Everything())
	if err != nil {
		log.Logger.Error(err)
	}
	return eventList, err
}

func (e *event) Get(ctx context.Context, name string) (*v1.Event, error) {
	event, err := e.cli.SharedInformerFactory.Core().V1().Events().Lister().Events(e.ns).Get(name)
	if err != nil {
		log.Logger.Error(err)
	}
	return event, err
}

func NewEvent(client *cloud.Clients, namespace string) *event {
	return &event{cli: client, ns: namespace}
}
