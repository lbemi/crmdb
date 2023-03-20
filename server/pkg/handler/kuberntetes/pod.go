package kuberntetes

import (
	"context"
	"fmt"
	"github.com/lbemi/lbemi/pkg/bootstrap/log"
	"github.com/lbemi/lbemi/pkg/common/store"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
)

type PodGetter interface {
	Pods(namespace string) IPod
}

type IPod interface {
	List(ctx context.Context) ([]*corev1.Pod, error)
	Get(ctx context.Context, name string) (*corev1.Pod, error)
	Create(ctx context.Context, obj *corev1.Pod) (*corev1.Pod, error)
	Update(ctx context.Context, obj *corev1.Pod) (*corev1.Pod, error)
	Delete(ctx context.Context, name string) error
}

type pod struct {
	cli *store.Clients
	ns  string
}

func (d *pod) List(ctx context.Context) ([]*corev1.Pod, error) {
	list, err := d.cli.SharedInformerFactory.Core().V1().Pods().Lister().Pods(d.ns).List(labels.Everything())
	if err != nil {
		log.Logger.Error(err)
	}

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

func NewPod(cli *store.Clients, namespace string) *pod {
	return &pod{cli: cli, ns: namespace}
}

type PodHandler struct{}

func NewPodHandler() *PodHandler {
	return &PodHandler{}
}

func (p *PodHandler) OnAdd(obj interface{}) {
	fmt.Println("Pod: OnAdd :", obj.(*corev1.Pod).Name)
}

func (p *PodHandler) OnUpdate(oldObj, newObj interface{}) {

	fmt.Println("Pod: OnUpdate: ", oldObj.(*corev1.Pod).Name, " --> ", newObj.(*corev1.Pod).Status.Phase)
}

func (p *PodHandler) OnDelete(obj interface{}) {

	fmt.Println("Pod: OnDelete: ", obj.(*corev1.Pod).Name)
}
