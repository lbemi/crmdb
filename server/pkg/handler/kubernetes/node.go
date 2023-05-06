package kubernetes

import (
	"context"
	"github.com/lbemi/lbemi/pkg/bootstrap/log"
	"github.com/lbemi/lbemi/pkg/handler/types"
	"github.com/lbemi/lbemi/pkg/services/k8s"
	v1 "k8s.io/api/core/v1"
	"sort"
)

type NodeGetter interface {
	Nodes() INode
}

type INode interface {
	List(ctx context.Context) ([]*types.Node, error)
	Get(ctx context.Context, name string) (*v1.Node, error)
	Delete(ctx context.Context, name string) error
	Create(ctx context.Context, node *v1.Node) (*v1.Node, error)
	Update(ctx context.Context, node *v1.Node) (*v1.Node, error)
	Patch(ctx context.Context, name string, playLoad map[string]interface{}) (*v1.Node, error)
	Drain(ctx context.Context, name string) (bool, error)

	Schedulable(ctx context.Context, name string, unschedulable bool) (bool, error)
}

type node struct {
	k8s *k8s.Factory
}

func NewNode(k8s *k8s.Factory) *node {
	return &node{k8s: k8s}
}

func (n *node) List(ctx context.Context) ([]*types.Node, error) {
	nodeList, err := n.k8s.Node().List(ctx)

	// 按创建时间排序排序
	sort.Slice(nodeList, func(i, j int) bool {
		return nodeList[j].ObjectMeta.GetCreationTimestamp().Time.After(nodeList[i].ObjectMeta.GetCreationTimestamp().Time)
	})

	if err != nil {
		log.Logger.Error(err)
	}
	return nodeList, err
}

func (n *node) Get(ctx context.Context, name string) (*v1.Node, error) {
	node, err := n.k8s.Node().Get(ctx, name)
	if err != nil {
		log.Logger.Error(err)
	}
	return node, err
}

func (n *node) Delete(ctx context.Context, name string) error {
	err := n.k8s.Node().Delete(ctx, name)
	if err != nil {
		log.Logger.Error(err)
	}
	return err
}

func (n *node) Create(ctx context.Context, node *v1.Node) (*v1.Node, error) {
	res, err := n.k8s.Node().Create(ctx, node)
	if err != nil {
		log.Logger.Error(err)
	}
	return res, err
}

func (n *node) Update(ctx context.Context, node *v1.Node) (*v1.Node, error) {

	res, err := n.k8s.Node().Update(ctx, node)
	if err != nil {
		log.Logger.Error(err)
	}
	return res, err
}

func (n *node) Patch(ctx context.Context, name string, labels map[string]interface{}) (*v1.Node, error) {
	res, err := n.k8s.Node().Patch(ctx, name, labels)
	if err != nil {
		log.Logger.Error(err)
	}
	return res, err
}

func (n *node) Drain(ctx context.Context, name string) (bool, error) {
	// 设置node不可调度
	status, err := n.Schedulable(ctx, name, true)
	if err != nil || !status {
		log.Logger.Error(err)
		return false, err
	}
	// 获取节点的上的pod信息
	podList, err := n.k8s.Node().GetPodByNode(ctx, name)
	if err != nil {
		log.Logger.Error(err)
		return false, err
	}
	// 清除节点上的pod
	for _, item := range podList.Items {
		// 排除kube-system 空间的pod
		if item.Namespace == "kube-system" {
			continue
		}
		err := n.k8s.Pod().EvictsPod(ctx, item.Name, item.Namespace)
		if err != nil {
			log.Logger.Error("移出节点上的pod异常", err)
			return false, err
		}
	}

	return true, nil
}

func (n *node) Schedulable(ctx context.Context, name string, unschedulable bool) (bool, error) {
	// 先查询node信息
	node, err := n.k8s.Node().Get(ctx, name)
	if err != nil {
		log.Logger.Error(err)
		return false, err
	}
	// 设置Unschedulable状态
	node.Spec.Unschedulable = unschedulable
	// 更新Unschedulable状态
	_, err = n.k8s.Node().Update(ctx, node)
	if err != nil {
		log.Logger.Error(err)
		return false, err
	}
	return true, nil
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
