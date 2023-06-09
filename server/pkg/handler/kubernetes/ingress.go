package kubernetes

import (
	"context"
	"github.com/lbemi/lbemi/pkg/services/k8s"
	v1 "k8s.io/api/networking/v1"
	"sort"
)

type IngressesGetter interface {
	Ingresses(namespace string) IIngresses
}

type IIngresses interface {
	List(ctx context.Context) []*v1.Ingress
	Get(ctx context.Context, name string) *v1.Ingress
	Delete(ctx context.Context, name string)
	Create(ctx context.Context, node *v1.Ingress) *v1.Ingress
	Update(ctx context.Context, ingresses *v1.Ingress) *v1.Ingress
}

type ingresses struct {
	k8s *k8s.Factory
}

func NewIngresses(k8s *k8s.Factory) *ingresses {
	return &ingresses{k8s: k8s}
}

func (s *ingresses) List(ctx context.Context) []*v1.Ingress {
	ingressList := s.k8s.Ingress().List(ctx)
	sort.Slice(ingressList, func(i, j int) bool {
		return ingressList[j].ObjectMeta.CreationTimestamp.Time.Before(ingressList[i].ObjectMeta.CreationTimestamp.Time)
	})
	return ingressList
}

func (s *ingresses) Get(ctx context.Context, name string) *v1.Ingress {
	return s.k8s.Ingress().Get(ctx, name)
}

func (s *ingresses) Delete(ctx context.Context, name string) {
	s.k8s.Ingress().Delete(ctx, name)
}

func (s *ingresses) Create(ctx context.Context, ingresses *v1.Ingress) *v1.Ingress {
	return s.k8s.Ingress().Create(ctx, ingresses)
}

func (s *ingresses) Update(ctx context.Context, ingresses *v1.Ingress) *v1.Ingress {
	return s.k8s.Ingress().Update(ctx, ingresses)
}
