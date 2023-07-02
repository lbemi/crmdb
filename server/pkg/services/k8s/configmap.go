package k8s

import (
	"context"
	"github.com/lbemi/lbemi/pkg/cache"

	"github.com/lbemi/lbemi/pkg/restfulx"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
)

type ConfigMapImp interface {
	List(ctx context.Context) []*v1.ConfigMap
	Get(ctx context.Context, name string) *v1.ConfigMap
	Delete(ctx context.Context, name string)
	Create(ctx context.Context, node *v1.ConfigMap) *v1.ConfigMap
	Update(ctx context.Context, configMap *v1.ConfigMap) *v1.ConfigMap
}

type configMap struct {
	client *cache.ClientConfig
	ns     string
}

func (s *configMap) List(ctx context.Context) []*v1.ConfigMap {
	nodeList, err := s.client.SharedInformerFactory.Core().V1().ConfigMaps().Lister().ConfigMaps(s.ns).List(labels.Everything())
	restfulx.ErrNotNilDebug(err, restfulx.GetResourceErr)
	return nodeList
}

func (s *configMap) Get(ctx context.Context, name string) *v1.ConfigMap {
	res, err := s.client.SharedInformerFactory.Core().V1().ConfigMaps().Lister().ConfigMaps(s.ns).Get(name)
	restfulx.ErrNotNilDebug(err, restfulx.GetResourceErr)
	return res
}

func (s *configMap) Delete(ctx context.Context, name string) {
	restfulx.ErrNotNilDebug(
		s.client.ClientSet.CoreV1().ConfigMaps(s.ns).Delete(ctx, name, metav1.DeleteOptions{}), restfulx.OperatorErr)
}

func (s *configMap) Create(ctx context.Context, configMap *v1.ConfigMap) *v1.ConfigMap {
	res, err := s.client.ClientSet.CoreV1().ConfigMaps(s.ns).Create(ctx, configMap, metav1.CreateOptions{})
	restfulx.ErrNotNilDebug(err, restfulx.GetResourceErr)
	return res
}

func (s *configMap) Update(ctx context.Context, configMap *v1.ConfigMap) *v1.ConfigMap {
	res, err := s.client.ClientSet.CoreV1().ConfigMaps(s.ns).Update(ctx, configMap, metav1.UpdateOptions{})
	restfulx.ErrNotNilDebug(err, restfulx.GetResourceErr)
	return res
}

func newConfigMap(client *cache.ClientConfig, namespace string) *configMap {
	return &configMap{client: client, ns: namespace}
}

type ConfigMapHandler struct {
	client      *cache.ClientConfig
	clusterName string
}

func NewConfigMapHandler(client *cache.ClientConfig, clusterName string) *ConfigMapHandler {
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
