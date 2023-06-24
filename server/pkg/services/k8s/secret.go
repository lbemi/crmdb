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

type SecretImp interface {
	List(ctx context.Context) []*v1.Secret
	Get(ctx context.Context, name string) *v1.Secret
	Delete(ctx context.Context, name string)
	Create(ctx context.Context, node *v1.Secret) *v1.Secret
	Update(ctx context.Context, secret *v1.Secret) *v1.Secret
}

type secret struct {
	client *cache.ClientConfig
	ns     string
}

func (s *secret) List(ctx context.Context) []*v1.Secret {
	list, err := s.client.SharedInformerFactory.Core().V1().Secrets().Lister().Secrets(s.ns).List(labels.Everything())
	restfulx.ErrNotNilDebug(err, restfulx.GetResourceErr)
	sort.Slice(list, func(i, j int) bool {
		return list[j].ObjectMeta.GetCreationTimestamp().Time.After(list[i].ObjectMeta.GetCreationTimestamp().Time)
	})
	return list
}

func (s *secret) Get(ctx context.Context, name string) *v1.Secret {
	res, err := s.client.SharedInformerFactory.Core().V1().Secrets().Lister().Secrets(s.ns).Get(name)
	restfulx.ErrNotNilDebug(err, restfulx.GetResourceErr)
	return res
}

func (s *secret) Delete(ctx context.Context, name string) {
	err := s.client.ClientSet.CoreV1().Secrets(s.ns).Delete(ctx, name, metav1.DeleteOptions{})
	restfulx.ErrNotNilDebug(err, restfulx.OperatorErr)
}

func (s *secret) Create(ctx context.Context, secret *v1.Secret) *v1.Secret {
	res, err := s.client.ClientSet.CoreV1().Secrets(s.ns).Create(ctx, secret, metav1.CreateOptions{})
	restfulx.ErrNotNilDebug(err, restfulx.OperatorErr)
	return res
}

func (s *secret) Update(ctx context.Context, secret *v1.Secret) *v1.Secret {
	res, err := s.client.ClientSet.CoreV1().Secrets(s.ns).Update(ctx, secret, metav1.UpdateOptions{})
	restfulx.ErrNotNilDebug(err, restfulx.OperatorErr)
	return res
}

func newSecret(client *cache.ClientConfig, namespace string) *secret {
	return &secret{client: client, ns: namespace}
}

type SecretHandle struct {
}

func NewSecretHandle() *SecretHandle {
	return &SecretHandle{}
}

func (s *SecretHandle) OnAdd(obj interface{}) {
	//TODO implement me
}

func (s *SecretHandle) OnUpdate(oldObj, newObj interface{}) {
	//TODO implement me
}

func (s *SecretHandle) OnDelete(obj interface{}) {
	//TODO implement me
}
