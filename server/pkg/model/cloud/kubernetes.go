package cloud

import (
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
)

type KubernetesClient struct {
	ID              uint64
	ClientSet       *kubernetes.Clientset
	InformerFactory informers.SharedInformerFactory
	SopChan         chan struct{}
	IsInit          bool
}
