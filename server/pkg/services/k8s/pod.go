package k8s

import (
	"context"
	"github.com/lbemi/lbemi/pkg/bootstrap/log"
	"github.com/lbemi/lbemi/pkg/common/store"
	"github.com/lbemi/lbemi/pkg/common/store/wsstore"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"sort"
)

type PodImp interface {
	List(ctx context.Context) ([]*corev1.Pod, error)
	Get(ctx context.Context, name string) (*corev1.Pod, error)
	Create(ctx context.Context, obj *corev1.Pod) (*corev1.Pod, error)
	Update(ctx context.Context, obj *corev1.Pod) (*corev1.Pod, error)
	Delete(ctx context.Context, name string) error
	GetPodByLabels(ctx context.Context, namespace string, label []map[string]string) ([]*corev1.Pod, error)
}

type pod struct {
	cli *store.Clients
	ns  string
}

func newPod(cli *store.Clients, ns string) *pod {
	return &pod{cli: cli, ns: ns}
}

func (d *pod) List(ctx context.Context) ([]*corev1.Pod, error) {
	list, err := d.cli.SharedInformerFactory.Core().V1().Pods().Lister().Pods(d.ns).List(labels.Everything())
	if err != nil {
		log.Logger.Error(err)
	}

	sort.Slice(list, func(i, j int) bool {
		return list[j].ObjectMeta.CreationTimestamp.Time.Before(list[i].ObjectMeta.CreationTimestamp.Time)
	})

	return list, err
}

func (d *pod) Get(ctx context.Context, name string) (*corev1.Pod, error) {
	dep, err := d.cli.SharedInformerFactory.Core().V1().Pods().Lister().Pods(d.ns).Get(name)
	if err != nil {
		log.Logger.Error(err)
	}
	return dep, err
}

func (d *pod) Create(ctx context.Context, obj *corev1.Pod) (*corev1.Pod, error) {
	newPod, err := d.cli.ClientSet.CoreV1().Pods(d.ns).Create(ctx, obj, metav1.CreateOptions{})
	if err != nil {
		log.Logger.Error(err)
	}
	return newPod, err
}

func (d *pod) Update(ctx context.Context, obj *corev1.Pod) (*corev1.Pod, error) {
	updatePod, err := d.cli.ClientSet.CoreV1().Pods(d.ns).Update(ctx, obj, metav1.UpdateOptions{})
	if err != nil {
		log.Logger.Error(err)
	}
	return updatePod, err
}

func (d *pod) Delete(ctx context.Context, name string) error {
	err := d.cli.ClientSet.CoreV1().Pods(d.ns).Delete(ctx, name, metav1.DeleteOptions{})
	if err != nil {
		log.Logger.Error(err)
	}
	return err
}
func (d *pod) GetPodByLabels(ctx context.Context, namespace string, label []map[string]string) ([]*corev1.Pod, error) {

	res := make([]*corev1.Pod, 0)
	pods, err := d.cli.SharedInformerFactory.Core().V1().Pods().Lister().Pods(namespace).List(labels.Everything())

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
