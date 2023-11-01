package kubernetes

import (
	"context"
	"encoding/json"
	"github.com/lbemi/lbemi/pkg/bootstrap/log"
	"github.com/lbemi/lbemi/pkg/common/store"
	"github.com/lbemi/lbemi/pkg/restfulx"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	apitype "k8s.io/apimachinery/pkg/types"
	"sort"
	"strings"
	"time"

	"github.com/lbemi/lbemi/pkg/handler/types"
	"github.com/lbemi/lbemi/pkg/model"
	"github.com/lbemi/lbemi/pkg/model/form"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/labels"
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
	Patch(ctx context.Context, name string, playLoad map[string]interface{}) *corev1.Node
	Drain(ctx context.Context, name string) bool

	Schedulable(ctx context.Context, name string, unschedulable bool) bool
	GetPodByNode(ctx context.Context, nodeName string, query *model.PageParam) *form.PageResult
	GetNodeEvent(ctx context.Context, name string) []*corev1.Event
}

type Node struct {
	cli    *store.ClientConfig
	Pods   IPod
	Events IEvent
}

func NewNode(cli *store.ClientConfig, event IEvent, pod IPod) *Node {
	return &Node{cli: cli, Events: event, Pods: pod}
}

func (n *Node) List(ctx context.Context, query *model.PageParam, name string, label string) *form.PageResult {
	var nodes = make([]*types.Node, 0)
	nodeResult, err := n.cli.SharedInformerFactory.Core().V1().Nodes().Lister().List(labels.Everything())
	for _, node := range nodeResult {
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

	//data := n.k8s.Node().List(ctx)
	res := &form.PageResult{}
	var nodeList = make([]*types.Node, 0)
	if name != "" {
		for _, item := range nodes {
			if strings.Contains(item.Name, name) {
				nodeList = append(nodeList, item)
			}
		}
		nodes = nodeList
	}

	if label != "" {
		for _, item := range nodes {
			if strings.Contains(labels.FormatLabels(item.Labels), label) {
				nodeList = append(nodeList, item)
			}
		}
		nodes = nodeList
	}
	// 按创建时间排序排序
	sort.Slice(nodes, func(i, j int) bool {
		return nodes[j].ObjectMeta.GetCreationTimestamp().Time.After(nodes[i].ObjectMeta.GetCreationTimestamp().Time)
	})

	total := len(nodes)
	// 未传递分页查询参数
	if query.Limit == 0 && query.Page == 0 {
		res.Data = nodes
	} else {
		if total <= query.Limit {
			res.Data = nodes
		} else if query.Page*query.Limit >= total {
			res.Data = nodes[(query.Page-1)*query.Limit : total]
		} else {
			res.Data = nodes[(query.Page-1)*query.Limit : query.Page*query.Limit]
		}
	}
	res.Total = int64(total)
	return res
}

func (n *Node) Get(ctx context.Context, name string) *corev1.Node {
	node, err := n.cli.SharedInformerFactory.Core().V1().Nodes().Lister().Get(name)
	restfulx.ErrNotNilDebug(err, restfulx.GetResourceErr)
	return node
}

func (n *Node) Delete(ctx context.Context, name string) {
	err := n.cli.ClientSet.CoreV1().Nodes().Delete(ctx, name, metav1.DeleteOptions{})
	restfulx.ErrNotNilDebug(err, restfulx.OperatorErr)
}

func (n *Node) Create(ctx context.Context, node *corev1.Node) *corev1.Node {
	res, err := n.cli.ClientSet.CoreV1().Nodes().Create(ctx, node, metav1.CreateOptions{})
	restfulx.ErrNotNilDebug(err, restfulx.OperatorErr)
	return res
}

func (n *Node) Update(ctx context.Context, node *corev1.Node) *corev1.Node {
	res, err := n.cli.ClientSet.CoreV1().Nodes().Update(ctx, node, metav1.UpdateOptions{})
	restfulx.ErrNotNilDebug(err, restfulx.OperatorErr)
	return res
}

func (n *Node) Patch(ctx context.Context, name string, labels map[string]interface{}) *corev1.Node {
	patchData := map[string]interface{}{"metadata": map[string]map[string]interface{}{"labels": labels}}
	playLoadBytes, err := json.Marshal(patchData)
	restfulx.ErrNotNilDebug(err, restfulx.OperatorErr)

	res, err := n.cli.ClientSet.CoreV1().Nodes().Patch(ctx, name, apitype.StrategicMergePatchType, playLoadBytes, metav1.PatchOptions{})
	restfulx.ErrNotNilDebug(err, restfulx.OperatorErr)
	return res
}

func (n *Node) Drain(ctx context.Context, name string) bool {
	// 设置node不可调度
	status := n.Schedulable(ctx, name, true)
	if !status {
		return false
	}
	// 获取节点的上的pod信息
	podList := n.Pods.GetPodByNode(ctx, name)
	// 清除节点上的pod
	for _, item := range podList.Items {
		// 排除kube-system 空间的pod
		if item.Namespace == "kube-system" {
			continue
		}
		n.Pods.EvictsPod(ctx, item.Name, item.Namespace)
	}

	return true
}

func (n *Node) Schedulable(ctx context.Context, name string, unschedulable bool) bool {
	// 先查询node信息
	res := n.Get(ctx, name)
	// 设置Unschedulable状态
	res.Spec.Unschedulable = unschedulable
	// 更新Unschedulable状态
	_ = n.Update(ctx, res)
	return true
}

func (n *Node) GetPodByNode(ctx context.Context, nodeName string, query *model.PageParam) *form.PageResult {
	podList := n.Pods.GetPodByNode(ctx, nodeName)
	res := &form.PageResult{}
	data := podList.Items
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
	// 按创建时间排序
	sort.Slice(podList.Items, func(i, j int) bool {
		return podList.Items[j].ObjectMeta.CreationTimestamp.Time.Before(podList.Items[i].ObjectMeta.CreationTimestamp.Time)
	})
	return res

}

func (n *Node) GetNodeUsage(ctx context.Context, node *corev1.Node) (cpuUsage, memoryUsage float64) {

	// 如果1秒超时，则返回空
	withTimeout, cancelFunc := context.WithTimeout(ctx, 1*time.Second)
	defer cancelFunc()
	nodeMetric, err := n.cli.MetricSet.MetricsV1beta1().NodeMetricses().Get(withTimeout, node.Name, metav1.GetOptions{})
	restfulx.ErrNotNil(err, restfulx.GetResourceErr)

	cpuUsage = float64(nodeMetric.Usage.Cpu().MilliValue()) / float64(node.Status.Capacity.Cpu().MilliValue())
	memoryUsage = float64(nodeMetric.Usage.Memory().MilliValue()) / float64(node.Status.Capacity.Memory().MilliValue())
	return
}

func (n *Node) getPodNumByNode(ctx context.Context, nodeName string) int {
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

// TODO 是否需要这个，而且这个能否获取到相关node事件？
func (n *Node) GetNodeEvent(ctx context.Context, name string) []*corev1.Event {
	events := make([]*corev1.Event, 0)
	result := n.Events.List(ctx, &model.PageParam{
		Page:  0,
		Limit: 0,
	})
	eventList, ok := result.Data.([]*corev1.Event)
	if !ok {
		return nil
	}

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
