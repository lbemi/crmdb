package v1beat1

import (
	"context"
	"k8s.io/api/extensions/v1beta1"
	"sort"

	"github.com/lbemi/lbemi/pkg/common/store"
	"github.com/lbemi/lbemi/pkg/restfulx"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
)

type IngressesImp interface {
	List(ctx context.Context) []*v1beta1.Ingress
	Get(ctx context.Context, name string) *v1beta1.Ingress
	Delete(ctx context.Context, name string)
	Create(ctx context.Context, node *v1beta1.Ingress) *v1beta1.Ingress
	Update(ctx context.Context, ingresses *v1beta1.Ingress) *v1beta1.Ingress
}

type ingresses struct {
	client *store.ClientConfig
	ns     string
}

func (s *ingresses) List(ctx context.Context) []*v1beta1.Ingress {
	ingressList, err := s.client.SharedInformerFactory.Extensions().V1beta1().Ingresses().Lister().Ingresses(s.ns).List(labels.Everything())
	restfulx.ErrNotNilDebug(err, restfulx.GetResourceErr)
	sort.Slice(ingressList, func(i, j int) bool {
		return ingressList[j].ObjectMeta.CreationTimestamp.Time.Before(ingressList[i].ObjectMeta.CreationTimestamp.Time)
	})
	return ingressList
}

func (s *ingresses) Get(ctx context.Context, name string) *v1beta1.Ingress {
	res, err := s.client.SharedInformerFactory.Extensions().V1beta1().Ingresses().Lister().Ingresses(s.ns).Get(name)
	restfulx.ErrNotNilDebug(err, restfulx.GetResourceErr)
	return res
}

func (s *ingresses) Delete(ctx context.Context, name string) {
	err := s.client.ClientSet.ExtensionsV1beta1().Ingresses(s.ns).Delete(ctx, name, metav1.DeleteOptions{})
	restfulx.ErrNotNilDebug(err, restfulx.OperatorErr)
}

func (s *ingresses) Create(ctx context.Context, ingresses *v1beta1.Ingress) *v1beta1.Ingress {
	res, err := s.client.ClientSet.ExtensionsV1beta1().Ingresses(s.ns).Create(ctx, ingresses, metav1.CreateOptions{})
	restfulx.ErrNotNilDebug(err, restfulx.OperatorErr)
	return res
}

func (s *ingresses) Update(ctx context.Context, ingresses *v1beta1.Ingress) *v1beta1.Ingress {
	res, err := s.client.ClientSet.ExtensionsV1beta1().Ingresses(s.ns).Update(ctx, ingresses, metav1.UpdateOptions{})
	restfulx.ErrNotNilDebug(err, restfulx.OperatorErr)
	return res
}

func newIngress(client *store.ClientConfig, namespace string) *ingresses {
	return &ingresses{client: client, ns: namespace}
}

type IngressHandle struct{}

func NewIngressHandle() *IngressHandle {
	return &IngressHandle{}
}

func (i *IngressHandle) OnAdd(obj interface{}) {
	//TODO implement me
}

func (i *IngressHandle) OnUpdate(oldObj, newObj interface{}) {
	//TODO implement me
}

func (i *IngressHandle) OnDelete(obj interface{}) {
	//TODO implement me
}
