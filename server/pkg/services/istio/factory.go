package istio

import "github.com/lbemi/lbemi/pkg/common/store"

type Interface interface {
	VirtualService() VirtualServiceImp
}

type Factory struct {
	client    *store.ClientConfig
	namespace string
}

func (f *Factory) VirtualService() VirtualServiceImp {
	return newVirtualService(f.client, f.namespace)
}

func NewIstioFactory(client *store.ClientConfig, namespace string) *Factory {
	return &Factory{client: client, namespace: namespace}
}
