package kuberntetes

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lbemi/lbemi/pkg/bootstrap/log"
	"github.com/lbemi/lbemi/pkg/common/store"
	"github.com/lbemi/lbemi/pkg/common/store/wsstore"
	appsv1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"sort"
	"sync"
)

type DeploymentGetter interface {
	Deployments(namespace string) IDeployment
}

type IDeployment interface {
	List(ctx context.Context) ([]*appsv1.Deployment, error)
	Get(ctx context.Context, name string) (*appsv1.Deployment, error)
	Create(ctx context.Context, obj *appsv1.Deployment) (*appsv1.Deployment, error)
	Update(ctx context.Context, obj *appsv1.Deployment) (*appsv1.Deployment, error)
	Delete(ctx context.Context, name string) error
	Scale(ctx context.Context, name string, replicaNum int32) error
}

type Deployment struct {
	cli *store.Clients
	ns  string
	//data []*appsv1.Deployment
	data  sync.Map
	IsDep bool
}

var deployment *Deployment

func NewDeployment(cli *store.Clients, namespace string) *Deployment {
	deployment = &Deployment{cli: cli, ns: namespace}
	return deployment
}

func (d *Deployment) List(ctx context.Context) ([]*appsv1.Deployment, error) {
	list, err := d.cli.SharedInformerFactory.Apps().V1().Deployments().Lister().Deployments(d.ns).List(labels.Everything())
	if err != nil {
		log.Logger.Error(err)
		return []*appsv1.Deployment{}, fmt.Errorf("record not found")
	}
	//按时间排序
	sort.Slice(list, func(i, j int) bool {
		return list[j].ObjectMeta.GetCreationTimestamp().Time.Before(list[i].ObjectMeta.GetCreationTimestamp().Time)
	})

	return list, nil
}

func (d *Deployment) Get(ctx context.Context, name string) (*appsv1.Deployment, error) {
	dep, err := d.cli.SharedInformerFactory.Apps().V1().Deployments().Lister().Deployments(d.ns).Get(name)
	if err != nil {
		log.Logger.Error(err)
	}
	return dep, err
}

func (d *Deployment) Create(ctx context.Context, obj *appsv1.Deployment) (*appsv1.Deployment, error) {
	newDeployment, err := d.cli.ClientSet.AppsV1().Deployments(d.ns).Create(ctx, obj, metav1.CreateOptions{})
	if err != nil {
		log.Logger.Error(err)
	}
	return newDeployment, err
}

func (d *Deployment) Update(ctx context.Context, obj *appsv1.Deployment) (*appsv1.Deployment, error) {
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

type DeploymentHandler struct {
	client      *store.Clients
	clusterName string
}

func NewDeploymentHandler(client *store.Clients, clusterName string) *DeploymentHandler {
	return &DeploymentHandler{client: client, clusterName: clusterName}
}

func (d *DeploymentHandler) OnAdd(obj interface{}) {
	d.notifyDeployments(obj)
	fmt.Println("OnAdd :", obj.(*appsv1.Deployment).Name)
}

func (d *DeploymentHandler) OnUpdate(oldObj, newObj interface{}) {
	d.notifyDeployments(newObj)
	fmt.Println("OnUpdate: ", oldObj.(*appsv1.Deployment).Name, " --> ", newObj.(*appsv1.Deployment).Status.AvailableReplicas)
}

func (d *DeploymentHandler) OnDelete(obj interface{}) {
	d.notifyDeployments(obj)
	fmt.Println("OnDelete: ", obj.(*appsv1.Deployment).Name)
}

func (d *DeploymentHandler) notifyDeployments(obj interface{}) {
	namespace := obj.(*appsv1.Deployment).Namespace
	deployments, err := d.client.SharedInformerFactory.Apps().V1().Deployments().Lister().Deployments(namespace).List(labels.Everything())
	if err != nil {
		log.Logger.Error(err)
	}

	//按时间排序
	sort.Slice(deployments, func(i, j int) bool {
		return deployments[j].ObjectMeta.GetCreationTimestamp().Time.Before(deployments[i].ObjectMeta.GetCreationTimestamp().Time)
	})

	wsstore.WsClientMap.SendClusterResource(d.clusterName, "deployment", gin.H{
		"cluster": d.clusterName,
		"type":    "deployment",
		"result": gin.H{
			"namespace": namespace,
			"data":      deployments,
		},
	})
}
