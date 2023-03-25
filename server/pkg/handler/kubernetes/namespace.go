package kubernetes

import (
	"context"
	"github.com/lbemi/lbemi/pkg/bootstrap/log"
	"github.com/lbemi/lbemi/pkg/services/k8s"
	v1 "k8s.io/api/core/v1"
	"sort"
)

type NamespaceGetter interface {
	Namespaces() INamespace
}

type INamespace interface {
	List(ctx context.Context) ([]*v1.Namespace, error)
	Get(ctx context.Context, name string) (*v1.Namespace, error)
	Create(ctx context.Context, obj *v1.Namespace) (*v1.Namespace, error)
	Update(ctx context.Context, obj *v1.Namespace) (*v1.Namespace, error)
	Delete(ctx context.Context, name string) error
}

type namespace struct {
	k8s *k8s.Factory
}

func NewNamespace(k8s *k8s.Factory) *namespace {
	return &namespace{k8s: k8s}
}

func (n *namespace) List(ctx context.Context) ([]*v1.Namespace, error) {
	list, err := n.k8s.Namespace().List(ctx)

	if err != nil {
		log.Logger.Error(err)
	}

	sort.Slice(list, func(i, j int) bool {
		return list[j].ObjectMeta.CreationTimestamp.Time.Before(list[i].ObjectMeta.CreationTimestamp.Time)
	})

	return list, err
}

func (n *namespace) Get(ctx context.Context, name string) (*v1.Namespace, error) {
	dep, err := n.k8s.Namespace().Get(ctx, name)
	if err != nil {
		log.Logger.Error(err)
	}
	return dep, err
}

func (n *namespace) Create(ctx context.Context, obj *v1.Namespace) (*v1.Namespace, error) {
	newNamespace, err := n.k8s.Namespace().Create(ctx, obj)
	if err != nil {
		log.Logger.Error(err)
	}
	return newNamespace, err
}

func (n *namespace) Update(ctx context.Context, obj *v1.Namespace) (*v1.Namespace, error) {
	newNamespace, err := n.k8s.Namespace().Update(ctx, obj)
	if err != nil {
		log.Logger.Error(err)
	}
	return newNamespace, err
}

func (n *namespace) Delete(ctx context.Context, name string) error {
	err := n.k8s.Namespace().Delete(ctx, name)
	if err != nil {
		log.Logger.Error(err)
	}
	return err
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
