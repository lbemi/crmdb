package k8s

import (
	"context"
	"github.com/lbemi/lbemi/pkg/bootstrap/log"
	"github.com/lbemi/lbemi/pkg/common/store"
	"github.com/lbemi/lbemi/pkg/common/store/wsstore"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/remotecommand"
	"sort"
)

type PodImp interface {
	List(ctx context.Context) ([]*corev1.Pod, error)
	Get(ctx context.Context, name string) (*corev1.Pod, error)
	Create(ctx context.Context, obj *corev1.Pod) (*corev1.Pod, error)
	Update(ctx context.Context, obj *corev1.Pod) (*corev1.Pod, error)
	Delete(ctx context.Context, name string) error
	GetPodByLabels(ctx context.Context, namespace string, label []map[string]string) ([]*corev1.Pod, error)
	PodExec(ctx context.Context, namespace, pod, container string, command []string) (remotecommand.Executor, error)
	GetPodLog(ctx context.Context, pod, container string) *rest.Request
}

type pod struct {
	cli *store.Clients
	ns  string
}

func newPod(cli *store.Clients, ns string) *pod {
	return &pod{cli: cli, ns: ns}
}

func (p *pod) List(ctx context.Context) ([]*corev1.Pod, error) {
	list, err := p.cli.SharedInformerFactory.Core().V1().Pods().Lister().Pods(p.ns).List(labels.Everything())
	if err != nil {
		log.Logger.Error(err)
	}

	sort.Slice(list, func(i, j int) bool {
		return list[j].ObjectMeta.CreationTimestamp.Time.Before(list[i].ObjectMeta.CreationTimestamp.Time)
	})

	return list, err
}

func (p *pod) Get(ctx context.Context, name string) (*corev1.Pod, error) {
	dep, err := p.cli.SharedInformerFactory.Core().V1().Pods().Lister().Pods(p.ns).Get(name)
	if err != nil {
		log.Logger.Error(err)
	}
	return dep, err
}

func (p *pod) Create(ctx context.Context, obj *corev1.Pod) (*corev1.Pod, error) {
	newPod, err := p.cli.ClientSet.CoreV1().Pods(p.ns).Create(ctx, obj, metav1.CreateOptions{})
	if err != nil {
		log.Logger.Error(err)
	}
	return newPod, err
}

func (p *pod) Update(ctx context.Context, obj *corev1.Pod) (*corev1.Pod, error) {
	updatePod, err := p.cli.ClientSet.CoreV1().Pods(p.ns).Update(ctx, obj, metav1.UpdateOptions{})
	if err != nil {
		log.Logger.Error(err)
	}
	return updatePod, err
}

func (p *pod) Delete(ctx context.Context, name string) error {
	err := p.cli.ClientSet.CoreV1().Pods(p.ns).Delete(ctx, name, metav1.DeleteOptions{})
	if err != nil {
		log.Logger.Error(err)
	}
	return err
}

func (p *pod) GetPodByLabels(ctx context.Context, namespace string, label []map[string]string) ([]*corev1.Pod, error) {

	res := make([]*corev1.Pod, 0)
	pods, err := p.cli.SharedInformerFactory.Core().V1().Pods().Lister().Pods(namespace).List(labels.Everything())

	if err != nil {
		log.Logger.Error(err)
		return nil, err
	}
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
	return res, nil
}

func (p *pod) PodExec(ctx context.Context, namespace, pod, container string, command []string) (remotecommand.Executor, error) {
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
	if err != nil {
		log.Logger.Error(err)
		return nil, err
	}
	return executor, nil
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

type PodHandler struct {
	client      *store.Clients
	clusterName string
}

func NewPodHandler(client *store.Clients, clusterName string) *PodHandler {
	return &PodHandler{client: client, clusterName: clusterName}
}

func (p *PodHandler) OnAdd(obj interface{}) {
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
