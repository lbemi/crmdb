package k8s

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/lbemi/lbemi/pkg/bootstrap/log"
	"github.com/lbemi/lbemi/pkg/common/store"
	"github.com/lbemi/lbemi/pkg/handler/types"
	"github.com/lbemi/lbemi/pkg/restfulx"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	apitype "k8s.io/apimachinery/pkg/types"
	"sort"
	"time"
)

type NodeImp interface {
	List(ctx context.Context) []*types.Node
	Get(ctx context.Context, name string) *corev1.Node
	Delete(ctx context.Context, name string)
	Create(ctx context.Context, node *corev1.Node) *corev1.Node
	Update(ctx context.Context, node *corev1.Node) *corev1.Node
	Patch(ctx context.Context, name string, playLoad map[string]interface{}) *corev1.Node
	Drain(ctx context.Context, name string)
	GetPodByNode(ctx context.Context, nodeName string) *corev1.PodList
	GetNodeUsage(ctx context.Context, node *corev1.Node) (cpuUsage, memoryUsage float64)
}

type node struct {
	cli *store.ClientConfig
}

func (n *node) List(ctx context.Context) []*types.Node {
	var nodes = make([]*types.Node, 0)
	nodeList, err := n.cli.SharedInformerFactory.Core().V1().Nodes().Lister().List(labels.Everything())
	for _, node := range nodeList {
		cpuUsage, memoryUsage := n.GetNodeUsage(ctx, node)
		podNum := n.getPodNumByNode(ctx, node.Name)
		item := &types.Node{
			TypeMeta:   node.TypeMeta,
			ObjectMeta: node.ObjectMeta,
			Spec:       node.Spec,
			Status:     node.Status,
			Usage: types.Usage{
				Cpu:    cpuUsage,
				Memory: memoryUsage,
				Pod:    podNum,
			},
		}
		nodes = append(nodes, item)
	}
	restfulx.ErrNotNilDebug(err, restfulx.GetResourceErr)
	return nodes
}

func (n *node) Get(ctx context.Context, name string) *corev1.Node {
	node, err := n.cli.SharedInformerFactory.Core().V1().Nodes().Lister().Get(name)
	restfulx.ErrNotNilDebug(err, restfulx.GetResourceErr)
	return node
}

func (n *node) Delete(ctx context.Context, name string) {
	err := n.cli.ClientSet.CoreV1().Nodes().Delete(ctx, name, metav1.DeleteOptions{})
	restfulx.ErrNotNilDebug(err, restfulx.OperatorErr)
}

func (n *node) Create(ctx context.Context, node *corev1.Node) *corev1.Node {
	res, err := n.cli.ClientSet.CoreV1().Nodes().Create(ctx, node, metav1.CreateOptions{})
	restfulx.ErrNotNilDebug(err, restfulx.OperatorErr)
	return res
}

func (n *node) Update(ctx context.Context, node *corev1.Node) *corev1.Node {
	res, err := n.cli.ClientSet.CoreV1().Nodes().Update(ctx, node, metav1.UpdateOptions{})
	restfulx.ErrNotNilDebug(err, restfulx.OperatorErr)
	return res
}

func (n *node) Patch(ctx context.Context, name string, labels map[string]interface{}) *corev1.Node {
	patchData := map[string]interface{}{"metadata": map[string]map[string]interface{}{"labels": labels}}
	playLoadBytes, err := json.Marshal(patchData)
	restfulx.ErrNotNilDebug(err, restfulx.OperatorErr)

	res, err := n.cli.ClientSet.CoreV1().Nodes().Patch(ctx, name, apitype.StrategicMergePatchType, playLoadBytes, metav1.PatchOptions{})
	restfulx.ErrNotNilDebug(err, restfulx.OperatorErr)
	return res
}

func (n *node) GetNodeUsage(ctx context.Context, node *corev1.Node) (cpuUsage, memoryUsage float64) {

	// 如果两秒超时，则返回空
	withTimeout, cancelFunc := context.WithTimeout(ctx, 1*time.Second)
	defer cancelFunc()
	nodeMetric, err := n.cli.MetricSet.MetricsV1beta1().NodeMetricses().Get(withTimeout, node.Name, metav1.GetOptions{})
	restfulx.ErrNotNilDebug(err, restfulx.GetResourceErr)

	cpuUsage = float64(nodeMetric.Usage.Cpu().MilliValue()) / float64(node.Status.Capacity.Cpu().MilliValue())
	memoryUsage = float64(nodeMetric.Usage.Memory().MilliValue()) / float64(node.Status.Capacity.Memory().MilliValue())
	return
}

func (n *node) getPodNumByNode(ctx context.Context, nodeName string) int {
	list, err := n.cli.SharedInformerFactory.Core().V1().Pods().Lister().Pods("").List(labels.Everything())
	if err != nil {
		log.Logger.Error(err)
		return 0
	}

	var count = 0
	for _, item := range list {
		if item.Spec.NodeName == nodeName {
			count++
		}
	}

	return count
}

func (n *node) GetPodByNode(ctx context.Context, nodeName string) *corev1.PodList {
	podList, err := n.cli.ClientSet.CoreV1().Pods("").List(ctx, metav1.ListOptions{
		FieldSelector: fmt.Sprintf("spec.nodeName=" + nodeName),
	})

	restfulx.ErrNotNilDebug(err, restfulx.GetResourceErr)
	sort.Slice(podList.Items, func(i, j int) bool {
		return podList.Items[j].ObjectMeta.CreationTimestamp.Time.Before(podList.Items[i].ObjectMeta.CreationTimestamp.Time)
	})
	return podList
}

func (n *node) Drain(ctx context.Context, name string) {
	// 排水选项
	drainOptions := metav1.DeleteOptions{GracePeriodSeconds: int64Ptr(0)}

	// 获取该节点上的所有 Pod
	podList, err := n.cli.ClientSet.CoreV1().Pods("").List(ctx, metav1.ListOptions{
		FieldSelector: fmt.Sprintf("spec.nodeName=%s", name),
	})
	restfulx.ErrNotNilDebug(err, restfulx.OperatorErr)

	// 删除该节点上的所有 Pod
	for _, pod := range podList.Items {
		err = n.cli.ClientSet.CoreV1().Pods(pod.Namespace).Delete(ctx, pod.Name, drainOptions)
		if err != nil {
			if !errors.IsNotFound(err) {
				restfulx.ErrNotNilDebug(err, restfulx.OperatorErr)
			}
		}
	}
}

func newNode(cli *store.ClientConfig) *node {
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

func int64Ptr(i int64) *int64 {
	return &i
}
