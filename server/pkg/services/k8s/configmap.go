package k8s

import (
	"context"
	"github.com/lbemi/lbemi/pkg/bootstrap/log"
	"github.com/lbemi/lbemi/pkg/common/store"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
)

type ConfigMapImp interface {
	List(ctx context.Context) ([]*v1.ConfigMap, error)
	Get(ctx context.Context, name string) (*v1.ConfigMap, error)
	Delete(ctx context.Context, name string) error
	Create(ctx context.Context, node *v1.ConfigMap) (*v1.ConfigMap, error)
	Update(ctx context.Context, configMap *v1.ConfigMap) (*v1.ConfigMap, error)
}

type configMap struct {
	client *store.ClientConfig
	ns     string
}

func (s *configMap) List(ctx context.Context) ([]*v1.ConfigMap, error) {
	nodeList, err := s.client.SharedInformerFactory.Core().V1().ConfigMaps().Lister().ConfigMaps(s.ns).List(labels.Everything())
	if err != nil {
		log.Logger.Error(err)
	}
	return nodeList, err
}

func (s *configMap) Get(ctx context.Context, name string) (*v1.ConfigMap, error) {
	res, err := s.client.SharedInformerFactory.Core().V1().ConfigMaps().Lister().ConfigMaps(s.ns).Get(name)
	if err != nil {
		log.Logger.Error(err)
	}
	return res, err
}

func (s *configMap) Delete(ctx context.Context, name string) error {
	err := s.client.ClientSet.CoreV1().ConfigMaps(s.ns).Delete(ctx, name, metav1.DeleteOptions{})
	if err != nil {
		log.Logger.Error(err)
	}
	return err
}

func (s *configMap) Create(ctx context.Context, configMap *v1.ConfigMap) (*v1.ConfigMap, error) {
	res, err := s.client.ClientSet.CoreV1().ConfigMaps(s.ns).Create(ctx, configMap, metav1.CreateOptions{})
	if err != nil {
		log.Logger.Error(err)
	}
	return res, err
}

func (s *configMap) Update(ctx context.Context, configMap *v1.ConfigMap) (*v1.ConfigMap, error) {
	res, err := s.client.ClientSet.CoreV1().ConfigMaps(s.ns).Update(ctx, configMap, metav1.UpdateOptions{})
	if err != nil {
		log.Logger.Error(err)
	}
	return res, err
}

func newConfigMap(client *store.ClientConfig, namespace string) *configMap {
	return &configMap{client: client, ns: namespace}
}

type ConfigMapHandler struct {
	client      *store.ClientConfig
	clusterName string
}

func NewConfigMapHandler(client *store.ClientConfig, clusterName string) *ConfigMapHandler {
	return &ConfigMapHandler{client: client, clusterName: clusterName}
}

func (c *ConfigMapHandler) OnAdd(obj interface{}) {
	//TODO implement me
}

func (c *ConfigMapHandler) OnUpdate(oldObj, newObj interface{}) {
	//TODO implement me
}

func (c *ConfigMapHandler) OnDelete(obj interface{}) {
	//TODO implement me
}
