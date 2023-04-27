package k8s

import (
	"context"
	"github.com/lbemi/lbemi/pkg/bootstrap/log"
	"github.com/lbemi/lbemi/pkg/common/store"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/labels"
)

type EventImp interface {
	List(ctx context.Context) ([]*corev1.Event, error)
	ListByLabels(ctx context.Context, labelsData labels.Set) ([]*corev1.Event, error)
	Get(ctx context.Context, name string) (*corev1.Event, error)
}

type event struct {
	cli *store.Clients
	ns  string
}

func (e *event) ListByLabels(ctx context.Context, labelsData labels.Set) ([]*corev1.Event, error) {
	selector := labels.SelectorFromSet(labelsData)
	fileldS
	eventList, err := e.cli.SharedInformerFactory.Core().V1().Events().Lister().Events(e.ns).List(selector)
	if err != nil {
		log.Logger.Error(err)
	}

	return eventList, err
}
func (e *event) List(ctx context.Context) ([]*corev1.Event, error) {
	eventList, err := e.cli.SharedInformerFactory.Core().V1().Events().Lister().Events(e.ns).List(labels.Everything())
	if err != nil {
		log.Logger.Error(err)
	}

	return eventList, err
}

func (e *event) Get(ctx context.Context, name string) (*corev1.Event, error) {
	event, err := e.cli.SharedInformerFactory.Core().V1().Events().Lister().Events(e.ns).Get(name)
	if err != nil {
		log.Logger.Error(err)
	}
	return event, err
}

func newEvent(client *store.Clients, namespace string) *event {
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
