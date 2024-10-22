package services

import (
	"github.com/lbemi/lbemi/pkg/cache"
)

type IstioGetter interface {
	Istio() IstioInterface
}

type IstioInterface interface {
	VirtualServiceGetter
	GatewayGetter
}

type Istio struct {
	clusterName string
	store       *cache.ClientStore
}

func NewIstio(clusterName string, store *cache.ClientStore) *Istio {
	return &Istio{clusterName: clusterName, store: store}
}

func (c *Istio) VirtualServices(namespace string) IVirtualService {
	if namespace == "all" {
		namespace = ""
	}
	return NewVirtualService(c.store.Get(c.clusterName), namespace)
}

func (c *Istio) Gateways(namespace string) IGateway {
	if namespace == "all" {
		namespace = ""
	}
	return NewGateway(c.store.Get(c.clusterName), namespace)
}
