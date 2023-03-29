package k8s

import (
	"context"
	"github.com/lbemi/lbemi/pkg/bootstrap/log"
	"github.com/lbemi/lbemi/pkg/common/store"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"sort"
)

type NamespaceImp interface {
	List(ctx context.Context) ([]*v1.Namespace, error)
	Get(ctx context.Context, name string) (*v1.Namespace, error)
	Create(ctx context.Context, obj *v1.Namespace) (*v1.Namespace, error)
	Update(ctx context.Context, obj *v1.Namespace) (*v1.Namespace, error)
	Delete(ctx context.Context, name string) error
}

type namespace struct {
	cli *store.Clients
}

func (n *namespace) List(ctx context.Context) ([]*v1.Namespace, error) {
	list, err := n.cli.SharedInformerFactory.Core().V1().Namespaces().Lister().List(labels.Everything())

	if err != nil {
		log.Logger.Error(err)
	}

	sort.Slice(list, func(i, j int) bool {
		return list[j].ObjectMeta.CreationTimestamp.Time.Before(list[i].ObjectMeta.CreationTimestamp.Time)
	})

	return list, err
}

func (n *namespace) Get(ctx context.Context, name string) (*v1.Namespace, error) {
	dep, err := n.cli.SharedInformerFactory.Core().V1().Namespaces().Lister().Get(name)
	if err != nil {
		log.Logger.Error(err)
	}
	return dep, err
}

func (n *namespace) Create(ctx context.Context, obj *v1.Namespace) (*v1.Namespace, error) {
	newNamespace, err := n.cli.ClientSet.CoreV1().Namespaces().Create(ctx, obj, metav1.CreateOptions{})
	if err != nil {
		log.Logger.Error(err)
	}
	return newNamespace, err
}

func (n *namespace) Update(ctx context.Context, obj *v1.Namespace) (*v1.Namespace, error) {
	newNamespace, err := n.cli.ClientSet.CoreV1().Namespaces().Update(ctx, obj, metav1.UpdateOptions{})
	if err != nil {
		log.Logger.Error(err)
	}
	return newNamespace, err
}

func (n *namespace) Delete(ctx context.Context, name string) error {
	err := n.cli.ClientSet.CoreV1().Namespaces().Delete(ctx, name, metav1.DeleteOptions{})
	if err != nil {
		log.Logger.Error(err)
	}
	return err
}

func newNamespace(cli *store.Clients) *namespace {
	return &namespace{cli: cli}
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