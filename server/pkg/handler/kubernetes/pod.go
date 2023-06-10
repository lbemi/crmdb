package kubernetes

import (
	"context"
	"github.com/lbemi/lbemi/pkg/model"
	"github.com/lbemi/lbemi/pkg/model/form"
	"github.com/lbemi/lbemi/pkg/services/k8s"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/remotecommand"
	"strings"
)

type PodGetter interface {
	Pods(namespace string) IPod
}

type IPod interface {
	List(ctx context.Context, query *model.PageParam, name string, label string) *form.PageResult
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

func (p *pod) List(ctx context.Context, query *model.PageParam, name string, label string) *form.PageResult {
	data := p.k8s.Pod().List(ctx)
	res := &form.PageResult{}
	var podList = make([]*corev1.Pod, 0)
	if name != "" {
		for _, item := range data {
			if strings.Contains(item.Name, name) {
				podList = append(podList, item)
			}
		}
		data = podList
	}

	if label != "" {
		for _, item := range data {
			if strings.Contains(item.Name, label) {
				podList = append(podList, item)
			}
		}
		data = podList
	}

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
