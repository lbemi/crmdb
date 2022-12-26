package kuberntetes

import (
	"context"
	"github.com/lbemi/lbemi/pkg/bootstrap/log"
	"github.com/lbemi/lbemi/pkg/services/cloud"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
)

type ServiceGetter interface {
	Service(namespace string) IService
}

type IService interface {
	List(ctx context.Context) ([]*v1.Service, error)
	Get(ctx context.Context, name string) (*v1.Service, error)
	Delete(ctx context.Context, name string) error
	Create(ctx context.Context, node *v1.Service) (*v1.Service, error)
	Update(ctx context.Context, service *v1.Service) (*v1.Service, error)
}

type service struct {
	client *cloud.Clients
	ns     string
}

func (s *service) List(ctx context.Context) ([]*v1.Service, error) {
	nodeList, err := s.client.Factory.Core().V1().Services().Lister().Services(s.ns).List(labels.Everything())

	if err != nil {
		log.Logger.Error(err)
	}
	return nodeList, err
}

func (s *service) Get(ctx context.Context, name string) (*v1.Service, error) {
	res, err := s.client.Factory.Core().V1().Services().Lister().Services(s.ns).Get(name)
	if err != nil {
		log.Logger.Error(err)
	}
	return res, err
}

func (s *service) Delete(ctx context.Context, name string) error {
	err := s.client.ClientSet.CoreV1().Services(s.ns).Delete(ctx, name, metav1.DeleteOptions{})
	if err != nil {
		log.Logger.Error(err)
	}
	return err
}

func (s *service) Create(ctx context.Context, service *v1.Service) (*v1.Service, error) {
	res, err := s.client.ClientSet.CoreV1().Services(s.ns).Create(ctx, service, metav1.CreateOptions{})
	if err != nil {
		log.Logger.Error(err)
	}
	return res, err
}

func (s *service) Update(ctx context.Context, service *v1.Service) (*v1.Service, error) {
	res, err := s.client.ClientSet.CoreV1().Services(s.ns).Update(ctx, service, metav1.UpdateOptions{})
	if err != nil {
		log.Logger.Error(err)
	}
	return res, err
}

func NewService(client *cloud.Clients, namespace string) *service {
	return &service{client: client, ns: namespace}
}
