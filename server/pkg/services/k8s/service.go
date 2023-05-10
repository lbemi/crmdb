package k8s

import (
	"context"
	"github.com/lbemi/lbemi/pkg/bootstrap/log"
	"github.com/lbemi/lbemi/pkg/common/store"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
)

type ServiceImp interface {
	List(ctx context.Context) ([]*v1.Service, error)
	Get(ctx context.Context, name string) (*v1.Service, error)
	Delete(ctx context.Context, name string) error
	Create(ctx context.Context, node *v1.Service) (*v1.Service, error)
	Update(ctx context.Context, service *v1.Service) (*v1.Service, error)
}

type service struct {
	client *store.Clients
	ns     string
}

func (s *service) List(ctx context.Context) ([]*v1.Service, error) {
	nodeList, err := s.client.SharedInformerFactory.Core().V1().Services().Lister().Services(s.ns).List(labels.Everything())

	if err != nil {
		log.Logger.Error(err)
	}
	return nodeList, err
}

func (s *service) Get(ctx context.Context, name string) (*v1.Service, error) {
	res, err := s.client.SharedInformerFactory.Core().V1().Services().Lister().Services(s.ns).Get(name)
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

func newService(client *store.Clients, namespace string) *service {
	return &service{client: client, ns: namespace}
}

type ServiceHandle struct{}

func NewServiceHandle() *ServiceHandle {
	return &ServiceHandle{}
}

func (s *ServiceHandle) OnAdd(obj interface{}) {
	//TODO implement me
}

func (s *ServiceHandle) OnUpdate(oldObj, newObj interface{}) {
	//TODO implement me
}

func (s *ServiceHandle) OnDelete(obj interface{}) {
	//TODO implement me
}
