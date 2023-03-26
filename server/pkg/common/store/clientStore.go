package store

import (
	"errors"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/metrics/pkg/client/clientset/versioned"
	"sync"
)

type Clients struct {
	ClientSet             *kubernetes.Clientset
	MetricSet             *versioned.Clientset
	SharedInformerFactory informers.SharedInformerFactory
	IsInit                bool
	Config                *rest.Config
}

type ClientStore struct {
	data map[string]*Clients
	lock sync.Mutex
}

type ClientStoreImp interface {
	Add(key string, resource *Clients) error
	Get(key string) *Clients
	Delete(key string)
}

func (c *ClientStore) Add(key string, resource *Clients) error {
	c.lock.Lock()
	defer c.lock.Unlock()

	if key == "" || resource == nil {
		return errors.New("key or value is null")
	}
	// 如果key已存在则覆盖
	c.data[key] = resource
	return nil
}

func (c *ClientStore) Get(key string) *Clients {
	c.lock.Lock()
	defer c.lock.Unlock()
	return c.data[key]
}

func (c *ClientStore) Delete(key string) {
	c.lock.Lock()
	defer c.lock.Unlock()
	delete(c.data, key)
}

func NewClientStore() *ClientStore {
	return &ClientStore{
		data: map[string]*Clients{},
	}
}
