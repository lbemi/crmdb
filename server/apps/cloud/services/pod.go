package services

import (
	"context"
	"fmt"
	entity2 "github.com/lbemi/lbemi/apps/cloud/entity"
	"github.com/lbemi/lbemi/pkg/bootstrap/log"
	"github.com/lbemi/lbemi/pkg/cache"
	"github.com/lbemi/lbemi/pkg/common/entity"
	"github.com/lbemi/lbemi/pkg/restfulx"
	"github.com/lbemi/lbemi/pkg/util"
	policy "k8s.io/api/policy/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/scheme"
	"sort"
	"strings"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/remotecommand"
)

type PodGetter interface {
	Pods(namespace string) IPod
}

type IPod interface {
	List(ctx context.Context, query *entity.PageParam, name string, label string) *entity.PageResult
	Get(ctx context.Context, name string) *corev1.Pod
	Create(ctx context.Context, obj *corev1.Pod) *corev1.Pod
	Update(ctx context.Context, obj *corev1.Pod) *corev1.Pod
	Delete(ctx context.Context, name string)

	PodExec(ctx context.Context, namespace, pod, container string, command []string) remotecommand.Executor
	GetPodLog(ctx context.Context, pod, container string) *rest.Request
	GetPodEvent(ctx context.Context, name string) []*corev1.Event
	Search(ctx context.Context, key string, searchType int) []*corev1.Pod
	EvictsPod(ctx context.Context, name, namespace string)
	GetPodByNode(ctx context.Context, nodeName string) *corev1.PodList
}

type Pod struct {
	cli    *cache.ClientConfig
	ns     string
	Events IEvent
}

func NewPod(cli *cache.ClientConfig, ns string, event IEvent) *Pod {
	return &Pod{cli: cli, ns: ns, Events: event}
}

func (p *Pod) List(ctx context.Context, query *entity.PageParam, name string, label string) *entity.PageResult {
	data, err := p.cli.SharedInformerFactory.Core().V1().Pods().Lister().Pods(p.ns).List(labels.Everything())
	restfulx.ErrNotNilDebug(err, restfulx.GetResourceErr)
	res := &entity.PageResult{}
	var podList = make([]*corev1.Pod, 0)

	for _, item := range data {
		if (name == "" || strings.Contains(item.Name, name)) && (label == "" || strings.Contains(labels.FormatLabels(item.Labels), label)) {
			podList = append(podList, item)
		}
	}

	total := len(podList)
	//按时间排序
	sort.SliceStable(podList, func(i, j int) bool {
		return podList[j].ObjectMeta.GetCreationTimestamp().Time.Before(podList[i].ObjectMeta.GetCreationTimestamp().Time)
	})
	// 未传递分页查询参数
	if query.Limit == 0 && query.Page == 0 {
		res.Data = podList
	} else {
		if total <= query.Limit {
			res.Data = podList
		} else if query.Page*query.Limit >= total {
			res.Data = podList[(query.Page-1)*query.Limit : total]
		} else {
			res.Data = podList[(query.Page-1)*query.Limit : query.Page*query.Limit]
		}
	}

	res.Total = int64(total)
	return res
}

func (p *Pod) Get(ctx context.Context, name string) *corev1.Pod {
	dep, err := p.cli.SharedInformerFactory.Core().V1().Pods().Lister().Pods(p.ns).Get(name)
	util.RestoreGVK(dep)
	restfulx.ErrNotNilDebug(err, restfulx.OperatorErr)
	return dep
}

func (p *Pod) Create(ctx context.Context, obj *corev1.Pod) *corev1.Pod {
	newPod, err := p.cli.ClientSet.CoreV1().Pods(p.ns).Create(ctx, obj, metav1.CreateOptions{})
	restfulx.ErrNotNilDebug(err, restfulx.OperatorErr)
	util.RestoreGVK(newPod)
	return newPod
}

func (p *Pod) Update(ctx context.Context, obj *corev1.Pod) *corev1.Pod {
	updatePod, err := p.cli.ClientSet.CoreV1().Pods(p.ns).Update(ctx, obj, metav1.UpdateOptions{})
	restfulx.ErrNotNilDebug(err, restfulx.OperatorErr)
	util.RestoreGVK(updatePod)
	return updatePod
}

func (p *Pod) Delete(ctx context.Context, name string) {
	err := p.cli.ClientSet.CoreV1().Pods(p.ns).Delete(ctx, name, metav1.DeleteOptions{})
	restfulx.ErrNotNilDebug(err, restfulx.OperatorErr)
}

func (p *Pod) PodExec(ctx context.Context, namespace, pod, container string, command []string) remotecommand.Executor {
	option := &corev1.PodExecOptions{
		Container: container,
		Command:   command,
		Stderr:    true,
		Stdin:     true,
		Stdout:    true,
		TTY:       true,
	}
	request := p.cli.ClientSet.CoreV1().RESTClient().Post().Resource("pods").Namespace(namespace).
		Name(pod).SubResource("exec").Param("color", "true").
		VersionedParams(option, scheme.ParameterCodec)
	executor, err := remotecommand.NewSPDYExecutor(p.cli.Config, "POST", request.URL())
	restfulx.ErrNotNilDebug(err, restfulx.OperatorErr)
	return executor
}

func (p *Pod) GetPodLog(ctx context.Context, pod, container string) *rest.Request {
	var tailLine int64 = 100

	option := &corev1.PodLogOptions{
		Follow:    true,
		Container: container,
		TailLines: &tailLine,
	}
	return p.cli.ClientSet.CoreV1().Pods(p.ns).GetLogs(pod, option)
}

func (p *Pod) GetPodEvent(ctx context.Context, name string) []*corev1.Event {
	res := p.Events.List(ctx, &entity.PageParam{
		Limit: 0,
		Page:  0,
	})
	eventList, ok := res.Data.([]*corev1.Event)
	if !ok {
		return nil
	}
	events := make([]*corev1.Event, 0, len(eventList))
	for _, item := range eventList {
		if item.InvolvedObject.Kind == "Pod" && item.InvolvedObject.Name == name {
			events = append(events, item)
		}
	}
	return events
}

func (p *Pod) Search(ctx context.Context, key string, searchType int) []*corev1.Pod {
	var podList = make([]*corev1.Pod, 0)
	podResult := p.List(ctx, &entity.PageParam{
		Limit: 0,
		Page:  0,
	}, "", "")

	pods, ok := podResult.Data.([]*corev1.Pod)
	if !ok {
		return nil
	}
	switch searchType {
	case entity2.SearchByName:
		for _, item := range pods {
			if strings.Contains(item.Name, key) {
				podList = append(podList, item)
			}
		}
	case entity2.SearchByLabel:
		for _, item := range pods {
			for k, label := range item.Labels {
				if strings.Contains(label, key) || strings.Contains(k, key) {
					podList = append(podList, item)
					break
				}
			}
		}
	default:
		restfulx.ErrNotNilDebug(fmt.Errorf("参数错误"), restfulx.ParamErr)
	}

	sort.SliceStable(podList, func(i, j int) bool {
		return podList[j].ObjectMeta.GetCreationTimestamp().Time.Before(podList[i].ObjectMeta.GetCreationTimestamp().Time)
	})
	restoreGVKForList(podList)
	return podList
}

// EvictsPod 驱逐pod
func (p *Pod) EvictsPod(ctx context.Context, name, namespace string) {
	// Pod优雅退出时间, 默认退出时间30s, 如果未指定, 则默认为每个对象的值。0表示立即删除。
	var gracePeriodSeconds int64 = 0
	propagationPolicy := metav1.DeletePropagationForeground
	deleteOptions := &metav1.DeleteOptions{
		GracePeriodSeconds: &gracePeriodSeconds,
		PropagationPolicy:  &propagationPolicy,
	}
	err := p.cli.ClientSet.PolicyV1beta1().Evictions(namespace).Evict(ctx, &policy.Eviction{
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: namespace,
		},
		DeleteOptions: deleteOptions,
	})
	restfulx.ErrNotNilDebug(err, restfulx.OperatorErr)
}

func (p *Pod) GetPodByNode(ctx context.Context, nodeName string) *corev1.PodList {
	podList, err := p.cli.ClientSet.CoreV1().Pods("").List(ctx, metav1.ListOptions{
		FieldSelector: fmt.Sprintf("spec.nodeName=" + nodeName),
	})

	restfulx.ErrNotNilDebug(err, restfulx.GetResourceErr)
	return podList
}

type PodHandler struct {
	client      *cache.ClientConfig
	clusterName string
}

func NewPodHandler(client *cache.ClientConfig, clusterName string) *PodHandler {
	return &PodHandler{client: client, clusterName: clusterName}
}

func (p *PodHandler) OnAdd(obj interface{}, isInInitialList bool) {
	p.notifyPods(obj)
}

func (p *PodHandler) OnUpdate(oldObj, newObj interface{}) {
	p.notifyPods(newObj)
	//fmt.Println("Pod: OnUpdate: ", oldObj.(*corev1.Pod).Name, " --> ", newObj.(*corev1.Pod).Status.Phase)
}

func (p *PodHandler) OnDelete(obj interface{}) {

	p.notifyPods(obj)
}

func (p *PodHandler) notifyPods(obj interface{}) {
	namespace := obj.(*corev1.Pod).Namespace
	pods, err := p.client.SharedInformerFactory.Core().V1().Pods().Lister().Pods(namespace).List(labels.Everything())
	if err != nil {
		log.Logger.Error(err)
	}

	restoreGVKForList(pods)
	//按时间排序
	sort.Slice(pods, func(i, j int) bool {
		return pods[j].ObjectMeta.GetCreationTimestamp().Time.Before(pods[i].ObjectMeta.GetCreationTimestamp().Time)
	})

	go cache.WsClientMap.SendClusterResource(p.clusterName, "Pod", map[string]interface{}{
		"cluster": p.clusterName,
		"type":    "Pod",
		"result": map[string]interface{}{
			"namespace": namespace,
			"data":      pods,
		},
	})
}

//func restoreGVKForList(podList []*corev1.Pod) {
//	objects := make([]runtime.Object, len(podList))
//	for i, p := range podList {
//		objects[i] = p
//	}
//	util.RestoreGVKForList(objects)
//}
