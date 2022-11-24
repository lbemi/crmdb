package kuberntetes

import (
	"context"
	"github.com/lbemi/lbemi/pkg/bootstrap/log"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type SecretGetter interface {
	Secrets(namespace string) ISecret
}

type ISecret interface {
	List(ctx context.Context) (*v1.SecretList, error)
	Get(ctx context.Context, name string) (*v1.Secret, error)
	Delete(ctx context.Context, name string) error
	Create(ctx context.Context, node *v1.Secret) (*v1.Secret, error)
	Update(ctx context.Context, secret *v1.Secret) (*v1.Secret, error)
}

type secret struct {
	clientSet *kubernetes.Clientset
	ns        string
}

func (s *secret) List(ctx context.Context) (*v1.SecretList, error) {
	nodeList, err := s.clientSet.CoreV1().Secrets(s.ns).List(ctx, metav1.ListOptions{})
	if err != nil {
		log.Logger.Error(err)
	}
	return nodeList, err
}

func (s *secret) Get(ctx context.Context, name string) (*v1.Secret, error) {
	res, err := s.clientSet.CoreV1().Secrets(s.ns).Get(ctx, name, metav1.GetOptions{})
	if err != nil {
		log.Logger.Error(err)
	}
	return res, err
}

func (s *secret) Delete(ctx context.Context, name string) error {
	err := s.clientSet.CoreV1().Secrets(s.ns).Delete(ctx, name, metav1.DeleteOptions{})
	if err != nil {
		log.Logger.Error(err)
	}
	return err
}

func (s *secret) Create(ctx context.Context, secret *v1.Secret) (*v1.Secret, error) {
	res, err := s.clientSet.CoreV1().Secrets(s.ns).Create(ctx, secret, metav1.CreateOptions{})
	if err != nil {
		log.Logger.Error(err)
	}
	return res, err
}

func (s *secret) Update(ctx context.Context, secret *v1.Secret) (*v1.Secret, error) {
	res, err := s.clientSet.CoreV1().Secrets(s.ns).Update(ctx, secret, metav1.UpdateOptions{})
	if err != nil {
		log.Logger.Error(err)
	}
	return res, err
}

func NewSecret(client *kubernetes.Clientset, namespace string) *secret {
	return &secret{clientSet: client, ns: namespace}
}
