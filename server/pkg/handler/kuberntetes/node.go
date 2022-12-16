package kuberntetes

import (
	"context"
	"encoding/json"
	"github.com/lbemi/lbemi/pkg/bootstrap/log"
	"github.com/lbemi/lbemi/pkg/services/cloud"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
)

type NodeGetter interface {
	Nodes() INode
}

type INode interface {
	List(ctx context.Context) (*v1.NodeList, error)
	Get(ctx context.Context, name string) (*v1.Node, error)
	Delete(ctx context.Context, name string) error
	Create(ctx context.Context, node *v1.Node) (*v1.Node, error)
	Update(ctx context.Context, node *v1.Node) (*v1.Node, error)
	Patch(ctx context.Context, name string, playLoad map[string]string) (*v1.Node, error)
}

type node struct {
	cli *cloud.Clients
}

func (n *node) List(ctx context.Context) (*v1.NodeList, error) {
	nodeList, err := n.cli.ClientSet.CoreV1().Nodes().List(ctx, metav1.ListOptions{})
	if err != nil {
		log.Logger.Error(err)
	}
	return nodeList, err
}

func (n *node) Get(ctx context.Context, name string) (*v1.Node, error) {
	node, err := n.cli.ClientSet.CoreV1().Nodes().Get(ctx, name, metav1.GetOptions{})
	if err != nil {
		log.Logger.Error(err)
	}
	return node, err
}

func (n *node) Delete(ctx context.Context, name string) error {
	err := n.cli.ClientSet.CoreV1().Nodes().Delete(ctx, name, metav1.DeleteOptions{})
	if err != nil {
		log.Logger.Error(err)
	}
	return err
}

func (n *node) Create(ctx context.Context, node *v1.Node) (*v1.Node, error) {
	res, err := n.cli.ClientSet.CoreV1().Nodes().Create(ctx, node, metav1.CreateOptions{})
	if err != nil {
		log.Logger.Error(err)
	}
	return res, err
}

func (n *node) Update(ctx context.Context, node *v1.Node) (*v1.Node, error) {

	res, err := n.cli.ClientSet.CoreV1().Nodes().Update(ctx, node, metav1.UpdateOptions{})
	if err != nil {
		log.Logger.Error(err)
	}
	return res, err
}

func (n *node) Patch(ctx context.Context, name string, labels map[string]string) (*v1.Node, error) {
	patchData := map[string]interface{}{"metadata": map[string]map[string]string{"labels": labels}}
	playLoadBytes, err := json.Marshal(patchData)
	if err != nil {
		log.Logger.Error(err)
		return nil, err
	}

	res, err := n.cli.ClientSet.CoreV1().Nodes().Patch(ctx, name, types.StrategicMergePatchType, playLoadBytes, metav1.PatchOptions{})
	if err != nil {
		log.Logger.Error(err)
	}
	return res, err
}

func NewNode(cli *cloud.Clients) *node {
	return &node{cli: cli}
}
