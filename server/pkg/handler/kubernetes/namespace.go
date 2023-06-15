package kubernetes

import (
	"context"
	"strings"

	"github.com/lbemi/lbemi/pkg/model"
	"github.com/lbemi/lbemi/pkg/model/form"
	"github.com/lbemi/lbemi/pkg/services/k8s"

	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/labels"
)

type NamespaceGetter interface {
	Namespaces() INamespace
}

type INamespace interface {
	List(ctx context.Context, query *model.PageParam, name string, label string) *form.PageResult
	Get(ctx context.Context, name string) *v1.Namespace
	Create(ctx context.Context, obj *v1.Namespace) *v1.Namespace
	Update(ctx context.Context, obj *v1.Namespace) *v1.Namespace
	Delete(ctx context.Context, name string)
}

type namespace struct {
	k8s *k8s.Factory
}

func NewNamespace(k8s *k8s.Factory) *namespace {
	return &namespace{k8s: k8s}
}

func (n *namespace) List(ctx context.Context, query *model.PageParam, name string, label string) *form.PageResult {
	data := n.k8s.Namespace().List(ctx)
	res := &form.PageResult{}
	namespaceList := make([]*v1.Namespace, 0)
	if name != "" {
		for _, item := range data {
			if strings.Contains(item.Name, name) {
				namespaceList = append(namespaceList, item)
			}
		}
		data = namespaceList
	}

	if label != "" {
		for _, item := range data {
			if strings.Contains(labels.FormatLabels(item.Labels), label) {
				namespaceList = append(namespaceList, item)
			}
		}
		data = namespaceList
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

func (n *namespace) Get(ctx context.Context, name string) *v1.Namespace {
	return n.k8s.Namespace().Get(ctx, name)
}

func (n *namespace) Create(ctx context.Context, obj *v1.Namespace) *v1.Namespace {
	return n.k8s.Namespace().Create(ctx, obj)
}

func (n *namespace) Update(ctx context.Context, obj *v1.Namespace) *v1.Namespace {
	return n.k8s.Namespace().Update(ctx, obj)
}

func (n *namespace) Delete(ctx context.Context, name string) {
	n.k8s.Namespace().Delete(ctx, name)
}

type NameSpaceHandler struct {
}

func (n *NameSpaceHandler) OnAdd(obj interface{}) {
	//TODO implement me
}

func (n *NameSpaceHandler) OnUpdate(oldObj, newObj interface{}) {
	//TODO implement me
}

func (n *NameSpaceHandler) OnDelete(obj interface{}) {
	//TODO implement me
}

func NewNameSpaceHandler() *NameSpaceHandler {
	return &NameSpaceHandler{}
}
