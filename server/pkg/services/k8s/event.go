package k8s

import (
	"context"
	"github.com/lbemi/lbemi/pkg/cache"
	"sort"

	"github.com/lbemi/lbemi/pkg/restfulx"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/labels"
)

type EventImp interface {
	List(ctx context.Context) []*corev1.Event
	ListByLabels(ctx context.Context, labelsData labels.Set) []*corev1.Event
	Get(ctx context.Context, name string) *corev1.Event
}

type event struct {
	cli *cache.ClientConfig
	ns  string
}

func (e *event) ListByLabels(ctx context.Context, labelsData labels.Set) []*corev1.Event {
	selector := labels.SelectorFromSet(labelsData)
	eventList, err := e.cli.SharedInformerFactory.Core().V1().Events().Lister().Events(e.ns).List(selector)
	restfulx.ErrNotNilDebug(err, restfulx.GetResourceErr)
	return eventList
}

func (e *event) List(ctx context.Context) []*corev1.Event {
	eventList, err := e.cli.SharedInformerFactory.Core().V1().Events().Lister().Events(e.ns).List(labels.Everything())
	restfulx.ErrNotNilDebug(err, restfulx.GetResourceErr)
	sort.Slice(eventList, func(i, j int) bool {
		return eventList[j].ObjectMeta.CreationTimestamp.Time.Before(eventList[i].ObjectMeta.CreationTimestamp.Time)
	})
	return eventList
}

func (e *event) Get(ctx context.Context, name string) *corev1.Event {
	res, err := e.cli.SharedInformerFactory.Core().V1().Events().Lister().Events(e.ns).Get(name)
	restfulx.ErrNotNilDebug(err, restfulx.GetResourceErr)
	return res
}

func newEvent(client *cache.ClientConfig, namespace string) *event {
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
