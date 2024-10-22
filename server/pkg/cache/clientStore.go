package cache

import (
	"errors"
	tektonVersiond "github.com/tektoncd/pipeline/pkg/client/clientset/versioned"
	tektonExternalversions "github.com/tektoncd/pipeline/pkg/client/informers/externalversions"
	"istio.io/client-go/pkg/informers/externalversions"
	"k8s.io/client-go/discovery"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/dynamic/dynamicinformer"
	"sync"

	"github.com/lbemi/lbemi/pkg/restfulx"

	istio "istio.io/client-go/pkg/clientset/versioned"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/metrics/pkg/client/clientset/versioned"
)

type ClientConfig struct {
	ClientSet                    *kubernetes.Clientset
	MetricSet                    *versioned.Clientset
	DynamicSet                   *dynamic.DynamicClient
	DiscoveryClient              *discovery.DiscoveryClient
	SharedInformerFactory        informers.SharedInformerFactory
	DynamicSharedInformerFactory dynamicinformer.DynamicSharedInformerFactory
	IsInit                       bool
	Config                       *rest.Config
	StopChan                     chan struct{}
	IstioClient                  *istio.Clientset
	IstioSharedInformerFactory   externalversions.SharedInformerFactory
	TektonClient                 *tektonVersiond.Clientset
	TektonSharedInformerFactory  tektonExternalversions.SharedInformerFactory
}

type ClientStore struct {
	data map[string]*ClientConfig
	lock sync.Mutex
}

type ClientStoreImp interface {
	Add(key string, resource *ClientConfig)
	Get(key string) *ClientConfig
	Delete(key string)
}

func (c *ClientStore) Add(key string, resource *ClientConfig) {
	c.lock.Lock()
	defer c.lock.Unlock()

	if key == "" || resource == nil {
		restfulx.ErrNotNilDebug(errors.New("key or value is null"), restfulx.OperatorErr)
	}
	// 如果key已存在则覆盖
	c.data[key] = resource
}

func (c *ClientStore) Get(key string) *ClientConfig {
	c.lock.Lock()
	defer c.lock.Unlock()
	return c.data[key]
}

func (c *ClientStore) Delete(key string) {
	c.lock.Lock()
	defer c.lock.Unlock()
	// 关闭informer
	close(c.data[key].StopChan)
	delete(c.data, key)
}

func NewClientStore() *ClientStore {
	return &ClientStore{
		data: map[string]*ClientConfig{},
	}
}
