package kuberntetes

import (
	"context"
	"github.com/lbemi/lbemi/pkg/bootstrap/log"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type EventGetter interface {
	Events(namespace string) IEvent
}

type IEvent interface {
	List(ctx context.Context) (*v1.EventList, error)
	Get(ctx context.Context, name string) (*v1.Event, error)
}

type event struct {
	cli *kubernetes.Clientset
	ns  string
}

func (e *event) List(ctx context.Context) (*v1.EventList, error) {
	eventList, err := e.cli.CoreV1().Events(e.ns).List(ctx, metav1.ListOptions{})
	if err != nil {
		log.Logger.Error(err)
	}
	return eventList, err
}

func (e *event) Get(ctx context.Context, name string) (*v1.Event, error) {
	event, err := e.cli.CoreV1().Events(e.ns).Get(ctx, name, metav1.GetOptions{})
	if err != nil {
		log.Logger.Error(err)
	}
	return event, err
}

func NewEvent(client *kubernetes.Clientset, namespace string) *event {
	return &event{cli: client, ns: namespace}
}
