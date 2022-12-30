package kuberntetes

import (
	"context"
	"github.com/lbemi/lbemi/pkg/bootstrap/log"
	"github.com/lbemi/lbemi/pkg/common/store"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
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
	cli *store.Clients
	ns  string
}

func (n *namespace) List(ctx context.Context) ([]*v1.Namespace, error) {
	list, err := n.cli.SharedInformerFactory.Core().V1().Namespaces().Lister().List(labels.Everything())
	if err != nil {
		log.Logger.Error(err)
	}

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

func NewNamespace(cli *store.Clients) *namespace {
	return &namespace{cli: cli}
}
