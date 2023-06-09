package kubernetes

import (
	"context"
	"github.com/lbemi/lbemi/pkg/services/k8s"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/remotecommand"
	"sort"
)

type PodGetter interface {
	Pods(namespace string) IPod
}

type IPod interface {
	List(ctx context.Context) []*corev1.Pod
	Get(ctx context.Context, name string) *corev1.Pod
	Create(ctx context.Context, obj *corev1.Pod) *corev1.Pod
	Update(ctx context.Context, obj *corev1.Pod) *corev1.Pod
	Delete(ctx context.Context, name string)

	PodExec(ctx context.Context, namespace, pod, container string, command []string) remotecommand.Executor
	GetPodLog(ctx context.Context, pod, container string) *rest.Request
	GetPodEvent(ctx context.Context, name string) []*corev1.Event
	Search(ctx context.Context, key string, searchType int) []*corev1.Pod
}

type pod struct {
	k8s *k8s.Factory
}

func NewPod(k8s *k8s.Factory) *pod {
	return &pod{k8s: k8s}
}

func (p *pod) List(ctx context.Context) []*corev1.Pod {
	list := p.k8s.Pod().List(ctx)
	sort.Slice(list, func(i, j int) bool {
		return list[j].ObjectMeta.CreationTimestamp.Time.Before(list[i].ObjectMeta.CreationTimestamp.Time)
	})
	return list
}

func (p *pod) Get(ctx context.Context, name string) *corev1.Pod {
	return p.k8s.Pod().Get(ctx, name)
}

func (p *pod) Create(ctx context.Context, obj *corev1.Pod) *corev1.Pod {
	return p.k8s.Pod().Create(ctx, obj)
}

func (p *pod) Update(ctx context.Context, obj *corev1.Pod) *corev1.Pod {
	return p.k8s.Pod().Update(ctx, obj)
}

func (p *pod) Delete(ctx context.Context, name string) {
	p.k8s.Pod().Delete(ctx, name)
}

func (p *pod) PodExec(ctx context.Context, namespace, pod, container string, command []string) remotecommand.Executor {
	return p.k8s.Pod().PodExec(ctx, namespace, pod, container, command)
}

func (p *pod) GetPodLog(ctx context.Context, pod, container string) *rest.Request {
	return p.k8s.Pod().GetPodLog(ctx, pod, container)
}

func (p *pod) GetPodEvent(ctx context.Context, name string) []*corev1.Event {
	events := make([]*corev1.Event, 0)
	eventList := p.k8s.Event().List(ctx)
	for _, item := range eventList {
		if item.InvolvedObject.Kind == "Pod" && item.InvolvedObject.Name == name {
			events = append(events, item)
		}
	}
	return events
}

func (p *pod) Search(ctx context.Context, key string, searchType int) []*corev1.Pod {
	return p.k8s.Pod().Search(ctx, key, searchType)
}
