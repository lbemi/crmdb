package kubernetes

import (
	"context"
	"github.com/lbemi/lbemi/pkg/handler/types"
	"github.com/lbemi/lbemi/pkg/services/k8s"
	corev1 "k8s.io/api/core/v1"
	"sort"
)

type NodeGetter interface {
	Nodes() INode
}

type INode interface {
	List(ctx context.Context) []*types.Node
	Get(ctx context.Context, name string) *corev1.Node
	Delete(ctx context.Context, name string)
	Create(ctx context.Context, node *corev1.Node) *corev1.Node
	Update(ctx context.Context, node *corev1.Node) *corev1.Node
	Patch(ctx context.Context, name string, playLoad map[string]interface{}) *corev1.Node
	Drain(ctx context.Context, name string) bool

	Schedulable(ctx context.Context, name string, unschedulable bool) bool
	GetPodByNode(ctx context.Context, nodeName string) *corev1.PodList
}

type node struct {
	k8s *k8s.Factory
}

func NewNode(k8s *k8s.Factory) *node {
	return &node{k8s: k8s}
}

func (n *node) List(ctx context.Context) []*types.Node {
	nodeList := n.k8s.Node().List(ctx)
	// 按创建时间排序排序
	sort.Slice(nodeList, func(i, j int) bool {
		return nodeList[j].ObjectMeta.GetCreationTimestamp().Time.After(nodeList[i].ObjectMeta.GetCreationTimestamp().Time)
	})
	return nodeList
}

func (n *node) Get(ctx context.Context, name string) *corev1.Node {
	return n.k8s.Node().Get(ctx, name)
}

func (n *node) Delete(ctx context.Context, name string) {
	n.k8s.Node().Delete(ctx, name)
}

func (n *node) Create(ctx context.Context, node *corev1.Node) *corev1.Node {
	return n.k8s.Node().Create(ctx, node)
}

func (n *node) Update(ctx context.Context, node *corev1.Node) *corev1.Node {
	return n.k8s.Node().Update(ctx, node)
}

func (n *node) Patch(ctx context.Context, name string, labels map[string]interface{}) *corev1.Node {
	return n.k8s.Node().Patch(ctx, name, labels)
}

func (n *node) Drain(ctx context.Context, name string) bool {
	// 设置node不可调度
	status := n.Schedulable(ctx, name, true)
	if !status {
		return false
	}
	// 获取节点的上的pod信息
	podList := n.k8s.Node().GetPodByNode(ctx, name)
	// 清除节点上的pod
	for _, item := range podList.Items {
		// 排除kube-system 空间的pod
		if item.Namespace == "kube-system" {
			continue
		}
		n.k8s.Pod().EvictsPod(ctx, item.Name, item.Namespace)
	}
	return true
}

func (n *node) Schedulable(ctx context.Context, name string, unschedulable bool) bool {
	// 先查询node信息
	res := n.k8s.Node().Get(ctx, name)
	// 设置Unschedulable状态
	res.Spec.Unschedulable = unschedulable
	// 更新Unschedulable状态
	_ = n.k8s.Node().Update(ctx, res)
	return true
}

func (p *node) GetPodByNode(ctx context.Context, nodeName string) *corev1.PodList {
	return p.k8s.Node().GetPodByNode(ctx, nodeName)
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
