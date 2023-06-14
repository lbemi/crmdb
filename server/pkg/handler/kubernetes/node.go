package kubernetes

import (
	"context"
	"github.com/lbemi/lbemi/pkg/handler/types"
	"github.com/lbemi/lbemi/pkg/model"
	"github.com/lbemi/lbemi/pkg/model/form"
	"github.com/lbemi/lbemi/pkg/services/k8s"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/labels"
	"strings"
)

type NodeGetter interface {
	Nodes() INode
}

type INode interface {
	List(ctx context.Context, query *model.PageParam, name string, label string) *form.PageResult
	Get(ctx context.Context, name string) *corev1.Node
	Delete(ctx context.Context, name string)
	Create(ctx context.Context, node *corev1.Node) *corev1.Node
	Update(ctx context.Context, node *corev1.Node) *corev1.Node
	Patch(ctx context.Context, name string, playLoad map[string]string) *corev1.Node
	Drain(ctx context.Context, name string) bool

	Schedulable(ctx context.Context, name string, unschedulable bool) bool
	GetPodByNode(ctx context.Context, nodeName string, query *model.PageParam) *form.PageResult
	GetNodeEvent(ctx context.Context, name string) []*corev1.Event
}

type node struct {
	k8s *k8s.Factory
}

func NewNode(k8s *k8s.Factory) *node {
	return &node{k8s: k8s}
}

func (n *node) List(ctx context.Context, query *model.PageParam, name string, label string) *form.PageResult {
	data := n.k8s.Node().List(ctx)
	res := &form.PageResult{}
	var nodeList = make([]*types.Node, 0)
	if name != "" {
		for _, item := range data {
			if strings.Contains(item.Name, name) {
				nodeList = append(nodeList, item)
			}
		}
		data = nodeList
	}

	if label != "" {
		for _, item := range data {
			if strings.Contains(labels.FormatLabels(item.Labels), label) {
				nodeList = append(nodeList, item)
			}
		}
		data = nodeList
	}
	total := len(data)
	// 未传递分页查询参数
	if query.Limit == 0 && query.Page == 0 {
		res.Data = data
	} else {
		if total <= query.Limit {
			res.Data = data
		} else if query.Page*query.Limit >= total {
			res.Data = data[(query.Page-1)*query.Limit : total]
		} else {
			res.Data = data[(query.Page-1)*query.Limit : query.Page*query.Limit]
		}
	}
	res.Total = int64(total)
	return res
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

func (n *node) Patch(ctx context.Context, name string, labels map[string]string) *corev1.Node {
	nodeInfo := n.k8s.Node().Get(ctx, name)
	nodeInfo.Labels = labels
	return n.k8s.Node().Update(ctx, nodeInfo)
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

func (p *node) GetPodByNode(ctx context.Context, nodeName string, query *model.PageParam) *form.PageResult {
	list := p.k8s.Node().GetPodByNode(ctx, nodeName)
	res := &form.PageResult{}
	data := list.Items
	total := len(data)
	// 未传递分页查询参数
	if query.Limit == 0 && query.Page == 0 {
		res.Data = data
	} else {
		if total <= query.Limit {
			res.Data = data
		} else if query.Page*query.Limit >= total {
			res.Data = data[(query.Page-1)*query.Limit : total]
		} else {
			res.Data = data[(query.Page-1)*query.Limit : query.Page*query.Limit]
		}
	}
	res.Total = int64(total)
	return res
}

// TODO 是否需要这个，而且这个能否获取到相关node事件？
func (p *node) GetNodeEvent(ctx context.Context, name string) []*corev1.Event {
	events := make([]*corev1.Event, 0)
	eventList := p.k8s.Event().List(ctx)
	for _, item := range eventList {
		if item.InvolvedObject.Kind == "Node" && item.InvolvedObject.Name == name {
			events = append(events, item)
		}
	}
	return events
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
