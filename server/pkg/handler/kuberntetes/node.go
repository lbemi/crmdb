package kuberntetes

import (
	"context"
	"github.com/lbemi/lbemi/pkg/bootstrap/log"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type NodeGetter interface {
	Nodes() INode
}

type INode interface {
	List(ctx context.Context) (*v1.NodeList, error)
	Get(ctx context.Context, name string) (*v1.Node, error)
	Delete(ctx context.Context, name string) error
	Create(ctx context.Context, node *v1.Node) (*v1.Node, error)
}

type node struct {
	clientSet *kubernetes.Clientset
}

func (n *node) List(ctx context.Context) (*v1.NodeList, error) {
	nodeList, err := n.clientSet.CoreV1().Nodes().List(ctx, metav1.ListOptions{})
	if err != nil {
		log.Logger.Error(err)
	}
	return nodeList, err
}

func (n *node) Get(ctx context.Context, name string) (*v1.Node, error) {
	node, err := n.clientSet.CoreV1().Nodes().Get(ctx, name, metav1.GetOptions{})
	if err != nil {
		log.Logger.Error(err)
	}
	return node, err
}

func (n *node) Delete(ctx context.Context, name string) error {
	err := n.clientSet.CoreV1().Nodes().Delete(ctx, name, metav1.DeleteOptions{})
	if err != nil {
		log.Logger.Error(err)
	}
	return err
}

func (n *node) Create(ctx context.Context, node *v1.Node) (*v1.Node, error) {
	res, err := n.clientSet.CoreV1().Nodes().Create(ctx, node, metav1.CreateOptions{})
	if err != nil {
		log.Logger.Error(err)
	}
	return res, err
}

func NewNode(clientSet *kubernetes.Clientset) *node {
	return &node{clientSet: clientSet}
}
