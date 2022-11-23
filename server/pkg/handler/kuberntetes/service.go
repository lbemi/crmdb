package kuberntetes

import (
	"context"
	"github.com/lbemi/lbemi/pkg/bootstrap/log"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type ServiceGetter interface {
	Service(namespace string) IService
}

type IService interface {
	List(ctx context.Context) (*v1.ServiceList, error)
	Get(ctx context.Context, name string) (*v1.Service, error)
	Delete(ctx context.Context, name string) error
	Create(ctx context.Context, node *v1.Service) (*v1.Service, error)
}

type service struct {
	clientSet *kubernetes.Clientset
	ns        string
}

func (s *service) List(ctx context.Context) (*v1.ServiceList, error) {
	nodeList, err := s.clientSet.CoreV1().Services(s.ns).List(ctx, metav1.ListOptions{})
	if err != nil {
		log.Logger.Error(err)
	}
	return nodeList, err
}

func (s *service) Get(ctx context.Context, name string) (*v1.Service, error) {
	res, err := s.clientSet.CoreV1().Services(s.ns).Get(ctx, name, metav1.GetOptions{})
	if err != nil {
		log.Logger.Error(err)
	}
	return res, err
}

func (s *service) Delete(ctx context.Context, name string) error {
	err := s.clientSet.CoreV1().Services(s.ns).Delete(ctx, name, metav1.DeleteOptions{})
	if err != nil {
		log.Logger.Error(err)
	}
	return err
}

func (s *service) Create(ctx context.Context, service *v1.Service) (*v1.Service, error) {
	res, err := s.clientSet.CoreV1().Services(s.ns).Create(ctx, service, metav1.CreateOptions{})
	if err != nil {
		log.Logger.Error(err)
	}
	return res, err
}

func NewService(client *kubernetes.Clientset, namespace string) *service {
	return &service{clientSet: client, ns: namespace}
}
