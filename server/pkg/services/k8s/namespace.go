package k8s

import (
	"context"
	"sort"

	"github.com/lbemi/lbemi/pkg/common/cache"
	"github.com/lbemi/lbemi/pkg/restfulx"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
)

type NamespaceImp interface {
	List(ctx context.Context) []*v1.Namespace
	Get(ctx context.Context, name string) *v1.Namespace
	Create(ctx context.Context, obj *v1.Namespace) *v1.Namespace
	Update(ctx context.Context, obj *v1.Namespace) *v1.Namespace
	Delete(ctx context.Context, name string)
}

type namespace struct {
	cli *cache.ClientConfig
}

func (n *namespace) List(ctx context.Context) []*v1.Namespace {
	list, err := n.cli.SharedInformerFactory.Core().V1().Namespaces().Lister().List(labels.Everything())
	restfulx.ErrNotNilDebug(err, restfulx.GetResourceErr)
	sort.Slice(list, func(i, j int) bool {
		return list[j].ObjectMeta.CreationTimestamp.Time.Before(list[i].ObjectMeta.CreationTimestamp.Time)
	})
	return list
}

func (n *namespace) Get(ctx context.Context, name string) *v1.Namespace {
	dep, err := n.cli.SharedInformerFactory.Core().V1().Namespaces().Lister().Get(name)
	restfulx.ErrNotNilDebug(err, restfulx.GetResourceErr)
	return dep
}

func (n *namespace) Create(ctx context.Context, obj *v1.Namespace) *v1.Namespace {
	newNamespace, err := n.cli.ClientSet.CoreV1().Namespaces().Create(ctx, obj, metav1.CreateOptions{})
	restfulx.ErrNotNilDebug(err, restfulx.OperatorErr)
	return newNamespace
}

func (n *namespace) Update(ctx context.Context, obj *v1.Namespace) *v1.Namespace {
	newNamespace, err := n.cli.ClientSet.CoreV1().Namespaces().Update(ctx, obj, metav1.UpdateOptions{})
	restfulx.ErrNotNilDebug(err, restfulx.OperatorErr)
	return newNamespace
}

func (n *namespace) Delete(ctx context.Context, name string) {
	restfulx.ErrNotNilDebug(n.cli.ClientSet.CoreV1().Namespaces().Delete(ctx, name, metav1.DeleteOptions{}), restfulx.OperatorErr)
}

func newNamespace(cli *cache.ClientConfig) *namespace {
	return &namespace{cli: cli}
}

type NameSpaceHandler struct {
	cli *cache.ClientConfig
	ns  string
}

func NewNameSpaceHandler(cli *cache.ClientConfig, ns string) *NameSpaceHandler {
	return &NameSpaceHandler{cli: cli, ns: ns}
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
