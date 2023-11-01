package k8s

//
//import (
//	"context"
//	"sort"
//
//	"github.com/lbemi/lbemi/pkg/common/store"
//	"github.com/lbemi/lbemi/pkg/restfulx"
//
//	corev1 "k8s.io/api/core/v1"
//	"k8s.io/apimachinery/pkg/labels"
//)
//
//type EventImp interface {
//	List(ctx context.Context) []*corev1.Events
//	ListByLabels(ctx context.Context, labelsData labels.Set) []*corev1.Events
//	Get(ctx context.Context, name string) *corev1.Events
//}
//
//type event struct {
//	cli *store.ClientConfig
//	ns  string
//}
//
//func (e *event) ListByLabels(ctx context.Context, labelsData labels.Set) []*corev1.Events {
//	selector := labels.SelectorFromSet(labelsData)
//	eventList, err := e.cli.SharedInformerFactory.Core().V1().Events().Lister().Events(e.ns).List(selector)
//	restfulx.ErrNotNilDebug(err, restfulx.GetResourceErr)
//	return eventList
//}
//
//func (e *event) List(ctx context.Context) []*corev1.Events {
//	eventList, err := e.cli.SharedInformerFactory.Core().V1().Events().Lister().Events(e.ns).List(labels.Everything())
//	restfulx.ErrNotNilDebug(err, restfulx.GetResourceErr)
//	sort.Slice(eventList, func(i, j int) bool {
//		return eventList[j].ObjectMeta.CreationTimestamp.Time.Before(eventList[i].ObjectMeta.CreationTimestamp.Time)
//	})
//	return eventList
//}
//
//func (e *event) Get(ctx context.Context, name string) *corev1.Events {
//	res, err := e.cli.SharedInformerFactory.Core().V1().Events().Lister().Events(e.ns).Get(name)
//	restfulx.ErrNotNilDebug(err, restfulx.GetResourceErr)
//	return res
//}
//
//func newEvent(client *store.ClientConfig, namespace string) *event {
//	return &event{cli: client, ns: namespace}
//}
//
//type EventHandler struct {
//}
//
//func (e *EventHandler) OnAdd(obj interface{}) {
//	//TODO implement me
//}
//
//func (e *EventHandler) OnUpdate(oldObj, newObj interface{}) {
//	//TODO implement me
//}
//
//func (e *EventHandler) OnDelete(obj interface{}) {
//	//TODO implement me
//}
//
//func NewEventHandler() *EventHandler {
//	return &EventHandler{}
//}
