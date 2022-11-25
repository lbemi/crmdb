package kuberntetes

import (
	"context"
	"github.com/lbemi/lbemi/pkg/bootstrap/log"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type ConfigMapGetter interface {
	ConfigMaps(namespace string) IConfigMap
}

type IConfigMap interface {
	List(ctx context.Context) (*v1.ConfigMapList, error)
	Get(ctx context.Context, name string) (*v1.ConfigMap, error)
	Delete(ctx context.Context, name string) error
	Create(ctx context.Context, node *v1.ConfigMap) (*v1.ConfigMap, error)
	Update(ctx context.Context, configMap *v1.ConfigMap) (*v1.ConfigMap, error)
}

type configMap struct {
	clientSet *kubernetes.Clientset
	ns        string
}

func (s *configMap) List(ctx context.Context) (*v1.ConfigMapList, error) {
	nodeList, err := s.clientSet.CoreV1().ConfigMaps(s.ns).List(ctx, metav1.ListOptions{})
	if err != nil {
		log.Logger.Error(err)
	}
	return nodeList, err
}

func (s *configMap) Get(ctx context.Context, name string) (*v1.ConfigMap, error) {
	res, err := s.clientSet.CoreV1().ConfigMaps(s.ns).Get(ctx, name, metav1.GetOptions{})
	if err != nil {
		log.Logger.Error(err)
	}
	return res, err
}

func (s *configMap) Delete(ctx context.Context, name string) error {
	err := s.clientSet.CoreV1().ConfigMaps(s.ns).Delete(ctx, name, metav1.DeleteOptions{})
	if err != nil {
		log.Logger.Error(err)
	}
	return err
}

func (s *configMap) Create(ctx context.Context, configMap *v1.ConfigMap) (*v1.ConfigMap, error) {
	res, err := s.clientSet.CoreV1().ConfigMaps(s.ns).Create(ctx, configMap, metav1.CreateOptions{})
	if err != nil {
		log.Logger.Error(err)
	}
	return res, err
}

func (s *configMap) Update(ctx context.Context, configMap *v1.ConfigMap) (*v1.ConfigMap, error) {
	res, err := s.clientSet.CoreV1().ConfigMaps(s.ns).Update(ctx, configMap, metav1.UpdateOptions{})
	if err != nil {
		log.Logger.Error(err)
	}
	return res, err
}

func NewConfigMap(client *kubernetes.Clientset, namespace string) *configMap {
	return &configMap{clientSet: client, ns: namespace}
}
