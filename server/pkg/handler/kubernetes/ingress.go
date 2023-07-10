package kubernetes

import (
	"context"
	"strings"

	"github.com/lbemi/lbemi/pkg/model"
	"github.com/lbemi/lbemi/pkg/model/form"
	"github.com/lbemi/lbemi/pkg/services/k8s"

	v1 "k8s.io/api/networking/v1"
	"k8s.io/apimachinery/pkg/labels"
)

type IngressesGetter interface {
	Ingresses(namespace string) IIngresses
}

type IIngresses interface {
	List(ctx context.Context, query *model.PageParam, name string, label string) *form.PageResult
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

func (s *ingresses) List(ctx context.Context, query *model.PageParam, name string, label string) *form.PageResult {
	data := s.k8s.Ingress().List(ctx)
	res := &form.PageResult{}
	var ingressList = make([]*v1.Ingress, 0)

	for _, item := range data {
		if (name == "" || strings.Contains(item.Name, name)) && (label == "" || strings.Contains(labels.FormatLabels(item.Labels), label)) {
			ingressList = append(ingressList, item)
		}
	}

	total := len(ingressList)
	if query.Limit == 0 && query.Page == 0 {
		res.Data = ingressList
	} else {
		if total <= query.Limit {
			res.Data = ingressList
		} else if query.Page*query.Limit >= total {
			res.Data = ingressList[(query.Page-1)*query.Limit : total]
		} else {
			res.Data = ingressList[(query.Page-1)*query.Limit : query.Page*query.Limit]
		}
	}

	res.Total = int64(total)
	return res
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
