package cache

import (
	"errors"
	"sync"

	"github.com/lbemi/lbemi/pkg/restfulx"

	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/metrics/pkg/client/clientset/versioned"
)

type ClientConfig struct {
	ClientSet             *kubernetes.Clientset
	MetricSet             *versioned.Clientset
	SharedInformerFactory informers.SharedInformerFactory
	IsInit                bool
	Config                *rest.Config
	StopChan              chan struct{}
}

type ClientMap struct {
	data map[string]*ClientConfig
	lock sync.Mutex
}

type ClientStoreImp interface {
	Add(key string, resource *ClientConfig)
	Get(key string) *ClientConfig
	Delete(key string)
}

func (c *ClientMap) Add(key string, resource *ClientConfig) {
	c.lock.Lock()
	defer c.lock.Unlock()

	if key == "" || resource == nil {
		restfulx.ErrNotNilDebug(errors.New("key or value is null"), restfulx.OperatorErr)
	}
	// 如果key已存在则覆盖
	c.data[key] = resource
}

func (c *ClientMap) Get(key string) *ClientConfig {
	c.lock.Lock()
	defer c.lock.Unlock()
	return c.data[key]
}

func (c *ClientMap) Delete(key string) {
	c.lock.Lock()
	defer c.lock.Unlock()
	delete(c.data, key)
}

func NewClientStore() *ClientMap {
	return &ClientMap{
		data: map[string]*ClientConfig{},
	}
}
