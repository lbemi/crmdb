package kubernetes

import (
	"context"
	"github.com/lbemi/lbemi/pkg/bootstrap/log"
	"github.com/lbemi/lbemi/pkg/services/k8s"
	corev1 "k8s.io/api/core/v1"
	"sort"
)

type PodGetter interface {
	Pods(namespace string) IPod
}

type IPod interface {
	List(ctx context.Context) ([]*corev1.Pod, error)
	Get(ctx context.Context, name string) (*corev1.Pod, error)
	Create(ctx context.Context, obj *corev1.Pod) (*corev1.Pod, error)
	Update(ctx context.Context, obj *corev1.Pod) (*corev1.Pod, error)
	Delete(ctx context.Context, name string) error
}

type pod struct {
	k8s *k8s.Factory
}

func NewPod(k8s *k8s.Factory) *pod {
	return &pod{k8s: k8s}
}

func (d *pod) List(ctx context.Context) ([]*corev1.Pod, error) {
	list, err := d.k8s.Pod().List(ctx)
	if err != nil {
		log.Logger.Error(err)
	}

	sort.Slice(list, func(i, j int) bool {
		return list[j].ObjectMeta.CreationTimestamp.Time.Before(list[i].ObjectMeta.CreationTimestamp.Time)
	})

	return list, err
}

func (d *pod) Get(ctx context.Context, name string) (*corev1.Pod, error) {
	dep, err := d.k8s.Pod().Get(ctx, name)
	if err != nil {
		log.Logger.Error(err)
	}
	return dep, err
}

func (d *pod) Create(ctx context.Context, obj *corev1.Pod) (*corev1.Pod, error) {
	newPod, err := d.k8s.Pod().Create(ctx, obj)
	if err != nil {
		log.Logger.Error(err)
	}
	return newPod, err
}

func (d *pod) Update(ctx context.Context, obj *corev1.Pod) (*corev1.Pod, error) {
	updatePod, err := d.k8s.Pod().Update(ctx, obj)
	if err != nil {
		log.Logger.Error(err)
	}
	return updatePod, err
}

func (d *pod) Delete(ctx context.Context, name string) error {
	err := d.k8s.Pod().Delete(ctx, name)
	if err != nil {
		log.Logger.Error(err)
	}
	return err
}
