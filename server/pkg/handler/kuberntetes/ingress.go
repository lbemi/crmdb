package kuberntetes

import (
	"context"
	"github.com/lbemi/lbemi/pkg/bootstrap/log"
	"github.com/lbemi/lbemi/pkg/services/cloud"
	v1 "k8s.io/api/networking/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
)

type IngressesGetter interface {
	Ingresses(namespace string) IIngresses
}

type IIngresses interface {
	List(ctx context.Context) ([]*v1.Ingress, error)
	Get(ctx context.Context, name string) (*v1.Ingress, error)
	Delete(ctx context.Context, name string) error
	Create(ctx context.Context, node *v1.Ingress) (*v1.Ingress, error)
	Update(ctx context.Context, ingresses *v1.Ingress) (*v1.Ingress, error)
}

type ingresses struct {
	client *cloud.Clients
	ns     string
}

func (s *ingresses) List(ctx context.Context) ([]*v1.Ingress, error) {
	nodeList, err := s.client.Factory.Networking().V1().Ingresses().Lister().Ingresses(s.ns).List(labels.Everything())
	if err != nil {
		log.Logger.Error(err)
	}
	return nodeList, err
}

func (s *ingresses) Get(ctx context.Context, name string) (*v1.Ingress, error) {
	res, err := s.client.Factory.Networking().V1().Ingresses().Lister().Ingresses(s.ns).Get(name)
	if err != nil {
		log.Logger.Error(err)
	}
	return res, err
}

func (s *ingresses) Delete(ctx context.Context, name string) error {
	err := s.client.ClientSet.NetworkingV1().Ingresses(s.ns).Delete(ctx, name, metav1.DeleteOptions{})
	if err != nil {
		log.Logger.Error(err)
	}
	return err
}

func (s *ingresses) Create(ctx context.Context, ingresses *v1.Ingress) (*v1.Ingress, error) {
	res, err := s.client.ClientSet.NetworkingV1().Ingresses(s.ns).Create(ctx, ingresses, metav1.CreateOptions{})
	if err != nil {
		log.Logger.Error(err)
	}
	return res, err
}

func (s *ingresses) Update(ctx context.Context, ingresses *v1.Ingress) (*v1.Ingress, error) {
	res, err := s.client.ClientSet.NetworkingV1().Ingresses(s.ns).Update(ctx, ingresses, metav1.UpdateOptions{})
	if err != nil {
		log.Logger.Error(err)
	}
	return res, err
}

func NewIngress(client *cloud.Clients, namespace string) *ingresses {
	return &ingresses{client: client, ns: namespace}
}
