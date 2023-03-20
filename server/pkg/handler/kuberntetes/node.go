package kuberntetes

import (
	"context"
	"encoding/json"
	"github.com/lbemi/lbemi/pkg/bootstrap/log"
	"github.com/lbemi/lbemi/pkg/common/store"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/types"
	"sort"
)

type NodeGetter interface {
	Nodes() INode
}

type INode interface {
	List(ctx context.Context) ([]*v1.Node, error)
	Get(ctx context.Context, name string) (*v1.Node, error)
	Delete(ctx context.Context, name string) error
	Create(ctx context.Context, node *v1.Node) (*v1.Node, error)
	Update(ctx context.Context, node *v1.Node) (*v1.Node, error)
	Patch(ctx context.Context, name string, playLoad map[string]string) (*v1.Node, error)

	GetNodeUsage(ctx context.Context, node *v1.Node) (cpuUsage, memoryUsage float64, err error)
}

type node struct {
	cli *store.Clients
}

func (n *node) List(ctx context.Context) ([]*v1.Node, error) {
	nodeList, err := n.cli.SharedInformerFactory.Core().V1().Nodes().Lister().List(labels.Everything())
	// 按创建时间排序排序
	sort.Slice(nodeList, func(i, j int) bool {
		return nodeList[j].ObjectMeta.GetCreationTimestamp().Time.Before(nodeList[i].ObjectMeta.GetCreationTimestamp().Time)
	})
	if err != nil {
		log.Logger.Error(err)
	}
	return nodeList, err
}

func (n *node) Get(ctx context.Context, name string) (*v1.Node, error) {
	node, err := n.cli.SharedInformerFactory.Core().V1().Nodes().Lister().Get(name)
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

func (n *node) GetNodeUsage(ctx context.Context, node *v1.Node) (cpuUsage, memoryUsage float64, err error) {
	nodeMetric, err := n.cli.MetricSet.MetricsV1beta1().NodeMetricses().Get(ctx, node.Name, metav1.GetOptions{})
	if err != nil {
		return
	}
	cpuUsage = float64(nodeMetric.Usage.Cpu().MilliValue()) / float64(node.Status.Capacity.Cpu().MilliValue())
	memoryUsage = float64(nodeMetric.Usage.Memory().MilliValue()) / float64(node.Status.Capacity.Memory().MilliValue())
	return
}

func NewNode(cli *store.Clients) *node {
	return &node{cli: cli}
}

type NodeHandler struct {
}

func (n *NodeHandler) OnAdd(obj interface{}) {
	//TODO implement me
}

func (n *NodeHandler) OnUpdate(oldObj, newObj interface{}) {
	//TODO implement me
}

func (n *NodeHandler) OnDelete(obj interface{}) {
	//TODO implement me
}

func NewNodeHandler() *NodeHandler {
	return &NodeHandler{}
}
