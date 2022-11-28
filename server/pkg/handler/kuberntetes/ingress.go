package kuberntetes

import (
	"context"
	"github.com/lbemi/lbemi/pkg/bootstrap/log"
	v1 "k8s.io/api/networking/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type IngressesGetter interface {
	Ingresses(namespace string) IIngresses
}

type IIngresses interface {
	List(ctx context.Context) (*v1.IngressList, error)
	Get(ctx context.Context, name string) (*v1.Ingress, error)
	Delete(ctx context.Context, name string) error
	Create(ctx context.Context, node *v1.Ingress) (*v1.Ingress, error)
	Update(ctx context.Context, ingresses *v1.Ingress) (*v1.Ingress, error)
}

type ingresses struct {
	clientSet *kubernetes.Clientset
	ns        string
}

func (s *ingresses) List(ctx context.Context) (*v1.IngressList, error) {
	nodeList, err := s.clientSet.NetworkingV1().Ingresses(s.ns).List(ctx, metav1.ListOptions{})
	if err != nil {
		log.Logger.Error(err)
	}
	return nodeList, err
}

func (s *ingresses) Get(ctx context.Context, name string) (*v1.Ingress, error) {
	res, err := s.clientSet.NetworkingV1().Ingresses(s.ns).Get(ctx, name, metav1.GetOptions{})
	if err != nil {
		log.Logger.Error(err)
	}
	return res, err
}

func (s *ingresses) Delete(ctx context.Context, name string) error {
	err := s.clientSet.NetworkingV1().Ingresses(s.ns).Delete(ctx, name, metav1.DeleteOptions{})
	if err != nil {
		log.Logger.Error(err)
	}
	return err
}

func (s *ingresses) Create(ctx context.Context, ingresses *v1.Ingress) (*v1.Ingress, error) {
	res, err := s.clientSet.NetworkingV1().Ingresses(s.ns).Create(ctx, ingresses, metav1.CreateOptions{})
	if err != nil {
		log.Logger.Error(err)
	}
	return res, err
}

func (s *ingresses) Update(ctx context.Context, ingresses *v1.Ingress) (*v1.Ingress, error) {
	res, err := s.clientSet.NetworkingV1().Ingresses(s.ns).Update(ctx, ingresses, metav1.UpdateOptions{})
	if err != nil {
		log.Logger.Error(err)
	}
	return res, err
}

func NewIngress(client *kubernetes.Clientset, namespace string) *ingresses {
	return &ingresses{clientSet: client, ns: namespace}
}
