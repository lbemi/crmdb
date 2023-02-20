package kuberntetes

import (
	"context"
	"fmt"
	"github.com/lbemi/lbemi/pkg/bootstrap/log"
	"github.com/lbemi/lbemi/pkg/common/store"
	v1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"sync"
)

type DeploymentGetter interface {
	Deployments(namespace string) IDeployment
}

type IDeployment interface {
	List(ctx context.Context) ([]*v1.Deployment, error)
	Get(ctx context.Context, name string) (*v1.Deployment, error)
	Create(ctx context.Context, obj *v1.Deployment) (*v1.Deployment, error)
	Update(ctx context.Context, obj *v1.Deployment) (*v1.Deployment, error)
	Delete(ctx context.Context, name string) error
	Scale(ctx context.Context, name string, replicaNum int32) error
}

type Deployment struct {
	cli  *store.Clients
	ns   string
	data []*v1.Deployment
	//data sync.Map
}

var once sync.Once

func (d *Deployment) OnAdd(obj interface{}) {
	fmt.Println("OnAdd :", obj.(*v1.Deployment).Name)
}

func (d *Deployment) OnUpdate(oldObj, newObj interface{}) {
	fmt.Println("OnUpdate: ", oldObj.(*v1.Deployment).Name, " --> ", newObj.(*v1.Deployment).Status.AvailableReplicas)
}

func (d *Deployment) OnDelete(obj interface{}) {
	fmt.Println("OnDelete: ", obj.(*v1.Deployment).Name)
}

func (d *Deployment) Start() {
	//once.Do(
	//	func() {
	log.Logger.Info("执行了初始化函数。。。。。。。。。。")
	d.cli.SharedInformerFactory.Apps().V1().Deployments().Informer().AddEventHandler(&Deployment{})
	//})
}

func NewDeployment(cli *store.Clients, namespace string) *Deployment {
	dep := &Deployment{cli: cli, ns: namespace, data: make([]*v1.Deployment, 0)}
	dep.Start()
	return dep
}

func (d *Deployment) List(ctx context.Context) ([]*v1.Deployment, error) {
	list, err := d.cli.SharedInformerFactory.Apps().V1().Deployments().Lister().Deployments(d.ns).List(labels.Everything())
	if err != nil {
		log.Logger.Error(err)
	}
	d.data = append(d.data, list...)
	return d.data, err
}

func (d *Deployment) Get(ctx context.Context, name string) (*v1.Deployment, error) {
	dep, err := d.cli.SharedInformerFactory.Apps().V1().Deployments().Lister().Deployments(d.ns).Get(name)
	//dep, err := d.cli.ClientSet.AppsV1().Deployments(d.ns).Get(ctx, name, metav1.GetOptions{})
	if err != nil {
		log.Logger.Error(err)
	}
	return dep, err
}

func (d *Deployment) Create(ctx context.Context, obj *v1.Deployment) (*v1.Deployment, error) {
	newDeployment, err := d.cli.ClientSet.AppsV1().Deployments(d.ns).Create(ctx, obj, metav1.CreateOptions{})
	if err != nil {
		log.Logger.Error(err)
	}
	return newDeployment, err
}

func (d *Deployment) Update(ctx context.Context, obj *v1.Deployment) (*v1.Deployment, error) {
	updateDeployment, err := d.cli.ClientSet.AppsV1().Deployments(d.ns).Update(ctx, obj, metav1.UpdateOptions{})
	if err != nil {
		log.Logger.Error(err)
	}
	return updateDeployment, err
}

func (d *Deployment) Delete(ctx context.Context, name string) error {
	err := d.cli.ClientSet.AppsV1().Deployments(d.ns).Delete(ctx, name, metav1.DeleteOptions{})
	if err != nil {
		log.Logger.Error(err)
	}
	return err
}

func (d *Deployment) Scale(ctx context.Context, name string, replicaNum int32) error {
	oldScale, err := d.cli.ClientSet.AppsV1().Deployments(d.ns).GetScale(ctx, name, metav1.GetOptions{})
	if err != nil {
		log.Logger.Error(err)
		return err
	}
	oldScale.Status.Replicas = replicaNum
	_, err = d.cli.ClientSet.AppsV1().Deployments(d.ns).UpdateScale(ctx, name, oldScale, metav1.UpdateOptions{})
	if err != nil {
		log.Logger.Error(err)
	}
	return nil
}
