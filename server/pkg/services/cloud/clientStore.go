package cloud

import (
	"errors"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"sync"
)

type Clients struct {
	ClientSet *kubernetes.Clientset
	Factory   informers.SharedInformerFactory
	IsInit    bool
}

type ClientStore struct {
	store map[string]*Clients
	lock  sync.Mutex
}

type IClientStore interface {
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
	c.store[key] = resource
	return nil
}

func (c *ClientStore) Get(key string) *Clients {
	c.lock.Lock()
	defer c.lock.Unlock()
	return c.store[key]
}

func (c *ClientStore) Delete(key string) {
	c.lock.Lock()
	defer c.lock.Unlock()
	delete(c.store, key)
}

func NewClientStore() *ClientStore {
	return &ClientStore{
		store: map[string]*Clients{},
	}
}
