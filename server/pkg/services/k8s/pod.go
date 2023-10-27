package k8s

import (
	"context"
	"fmt"
	"sort"
	"strings"

	"github.com/lbemi/lbemi/pkg/util"
	"k8s.io/apimachinery/pkg/runtime"

	"github.com/lbemi/lbemi/pkg/bootstrap/log"
	"github.com/lbemi/lbemi/pkg/common/store"
	"github.com/lbemi/lbemi/pkg/common/store/wsstore"
	"github.com/lbemi/lbemi/pkg/handler/types"
	"github.com/lbemi/lbemi/pkg/restfulx"

	corev1 "k8s.io/api/core/v1"
	policy "k8s.io/api/policy/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/remotecommand"
)

type PodImp interface {
	List(ctx context.Context) []*corev1.Pod
	Get(ctx context.Context, name string) *corev1.Pod
	Create(ctx context.Context, obj *corev1.Pod) *corev1.Pod
	Update(ctx context.Context, obj *corev1.Pod) *corev1.Pod
	Delete(ctx context.Context, name string)
	GetPodByLabels(ctx context.Context, namespace string, label []map[string]string) []*corev1.Pod
	PodExec(ctx context.Context, namespace, pod, container string, command []string) remotecommand.Executor
	GetPodLog(ctx context.Context, pod, container string) *rest.Request
	EvictsPod(ctx context.Context, name, namespace string)
	Search(ctx context.Context, key string, searchType int) []*corev1.Pod
}

type pod struct {
	cli *store.ClientConfig
	ns  string
}

func newPod(cli *store.ClientConfig, ns string) *pod {
	return &pod{cli: cli, ns: ns}
}

func (p *pod) List(ctx context.Context) []*corev1.Pod {
	list, err := p.cli.SharedInformerFactory.Core().V1().Pods().Lister().Pods(p.ns).List(labels.Everything())
	restfulx.ErrNotNilDebug(err, restfulx.GetResourceErr)
	restoreGVKForList(list)
	sort.Slice(list, func(i, j int) bool {
		return list[j].ObjectMeta.CreationTimestamp.Time.Before(list[i].ObjectMeta.CreationTimestamp.Time)
	})

	return list
}

func (p *pod) Get(ctx context.Context, name string) *corev1.Pod {
	dep, err := p.cli.SharedInformerFactory.Core().V1().Pods().Lister().Pods(p.ns).Get(name)
	util.RestoreGVK(dep)
	restfulx.ErrNotNilDebug(err, restfulx.OperatorErr)
	return dep
}

func (p *pod) Create(ctx context.Context, obj *corev1.Pod) *corev1.Pod {
	newPod, err := p.cli.ClientSet.CoreV1().Pods(p.ns).Create(ctx, obj, metav1.CreateOptions{})
	restfulx.ErrNotNilDebug(err, restfulx.OperatorErr)
	util.RestoreGVK(newPod)
	return newPod
}

func (p *pod) Update(ctx context.Context, obj *corev1.Pod) *corev1.Pod {
	updatePod, err := p.cli.ClientSet.CoreV1().Pods(p.ns).Update(ctx, obj, metav1.UpdateOptions{})
	restfulx.ErrNotNilDebug(err, restfulx.OperatorErr)
	util.RestoreGVK(updatePod)
	return updatePod
}

func (p *pod) Delete(ctx context.Context, name string) {
	err := p.cli.ClientSet.CoreV1().Pods(p.ns).Delete(ctx, name, metav1.DeleteOptions{})
	restfulx.ErrNotNilDebug(err, restfulx.OperatorErr)
}

func (p *pod) GetPodByLabels(ctx context.Context, namespace string, label []map[string]string) []*corev1.Pod {

	res := make([]*corev1.Pod, 0)
	pods, err := p.cli.SharedInformerFactory.Core().V1().Pods().Lister().Pods(namespace).List(labels.Everything())
	restfulx.ErrNotNilDebug(err, restfulx.OperatorErr)
	for _, item := range pods {
		for _, l := range label {
			i := 0
			for k1, v1 := range l {
				for k2, v2 := range item.Labels {
					if k1 == k2 && v1 == v2 {
						i++
					}
				}
			}
			if i == len(l) {
				res = append(res, item)
			}
		}
	}
	restoreGVKForList(res)
	return res
}

func (p *pod) PodExec(ctx context.Context, namespace, pod, container string, command []string) remotecommand.Executor {
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

func (p *pod) GetPodLog(ctx context.Context, pod, container string) *rest.Request {
	var tailLine int64 = 100

	option := &corev1.PodLogOptions{
		Follow:    true,
		Container: container,
		TailLines: &tailLine,
	}
	return p.cli.ClientSet.CoreV1().Pods(p.ns).GetLogs(pod, option)
}

// EvictsPod 驱逐pod
func (p *pod) EvictsPod(ctx context.Context, name, namespace string) {
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

func (p *pod) Search(ctx context.Context, key string, searchType int) []*corev1.Pod {
	var podList = make([]*corev1.Pod, 0)
	pods := p.List(ctx)

	switch searchType {
	case types.SearchByName:
		for _, item := range pods {
			if strings.Contains(item.Name, key) {
				podList = append(podList, item)
			}
		}
	case types.SearchByLabel:
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

type PodHandler struct {
	client      *store.ClientConfig
	clusterName string
}

func NewPodHandler(client *store.ClientConfig, clusterName string) *PodHandler {
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

	go wsstore.WsClientMap.SendClusterResource(p.clusterName, "pod", map[string]interface{}{
		"cluster": p.clusterName,
		"type":    "pod",
		"result": map[string]interface{}{
			"namespace": namespace,
			"data":      pods,
		},
	})
}

func restoreGVKForList(podList []*corev1.Pod) {
	objects := make([]runtime.Object, len(podList))
	for i, p := range podList {
		objects[i] = p
	}
	util.RestoreGVKForList(objects)
}
