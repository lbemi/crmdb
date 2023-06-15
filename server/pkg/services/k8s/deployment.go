package k8s

import (
	"context"
	"fmt"
	"sort"
	"strings"

	"github.com/lbemi/lbemi/pkg/bootstrap/log"
	"github.com/lbemi/lbemi/pkg/common/store"
	"github.com/lbemi/lbemi/pkg/common/store/wsstore"
	"github.com/lbemi/lbemi/pkg/handler/types"
	"github.com/lbemi/lbemi/pkg/restfulx"

	appsv1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
)

type DeploymentImp interface {
	List(ctx context.Context) []*appsv1.Deployment
	Get(ctx context.Context, name string) *appsv1.Deployment
	Create(ctx context.Context, obj *appsv1.Deployment) *appsv1.Deployment
	Update(ctx context.Context, obj *appsv1.Deployment) *appsv1.Deployment
	Delete(ctx context.Context, name string)
	Scale(ctx context.Context, name string, replicaNum int32)
	Search(ctx context.Context, key string, searchType int) []*appsv1.Deployment
}

type Deployment struct {
	cli       *store.ClientConfig
	namespace string
}

func newDeployment(cli *store.ClientConfig, namespace string) *Deployment {
	return &Deployment{cli: cli, namespace: namespace}
}

func (d *Deployment) List(ctx context.Context) []*appsv1.Deployment {
	list, err := d.cli.SharedInformerFactory.Apps().V1().Deployments().Lister().Deployments(d.namespace).List(labels.Everything())
	restfulx.ErrNotNilDebug(err, restfulx.GetResourceErr)

	//按时间排序
	sort.Slice(list, func(i, j int) bool {
		return list[j].ObjectMeta.GetCreationTimestamp().Time.Before(list[i].ObjectMeta.GetCreationTimestamp().Time)
	})
	return list
}

func (d *Deployment) Get(ctx context.Context, name string) *appsv1.Deployment {
	dep, err := d.cli.SharedInformerFactory.Apps().V1().Deployments().Lister().Deployments(d.namespace).Get(name)
	restfulx.ErrNotNilDebug(err, restfulx.GetResourceErr)
	return dep
}

func (d *Deployment) Create(ctx context.Context, obj *appsv1.Deployment) *appsv1.Deployment {
	newDeployment, err := d.cli.ClientSet.AppsV1().Deployments(d.namespace).Create(ctx, obj, metav1.CreateOptions{})
	restfulx.ErrNotNilDebug(err, restfulx.OperatorErr)
	return newDeployment
}

func (d *Deployment) Update(ctx context.Context, obj *appsv1.Deployment) *appsv1.Deployment {
	updateDeployment, err := d.cli.ClientSet.AppsV1().Deployments(d.namespace).Update(ctx, obj, metav1.UpdateOptions{})
	restfulx.ErrNotNilDebug(err, restfulx.OperatorErr)
	return updateDeployment
}

func (d *Deployment) Delete(ctx context.Context, name string) {
	err := d.cli.ClientSet.AppsV1().Deployments(d.namespace).Delete(ctx, name, metav1.DeleteOptions{})
	restfulx.ErrNotNilDebug(err, restfulx.OperatorErr)
}

func (d *Deployment) Scale(ctx context.Context, name string, replicaNum int32) {
	oldScale, err := d.cli.ClientSet.AppsV1().Deployments(d.namespace).GetScale(ctx, name, metav1.GetOptions{})
	restfulx.ErrNotNilDebug(err, restfulx.OperatorErr)
	oldScale.Spec.Replicas = replicaNum
	_, err = d.cli.ClientSet.AppsV1().Deployments(d.namespace).UpdateScale(ctx, name, oldScale, metav1.UpdateOptions{})
	restfulx.ErrNotNilDebug(err, restfulx.OperatorErr)
}

func (d *Deployment) Search(ctx context.Context, key string, searchType int) []*appsv1.Deployment {
	var deploymentList = make([]*appsv1.Deployment, 0)
	deployments := d.List(ctx)

	switch searchType {
	case types.SearchByName:
		// 遍历deployment，如果name包含key则保存返回
		for _, item := range deployments {
			if strings.Contains(item.Name, key) {
				deploymentList = append(deploymentList, item)
			}
		}
	case types.SearchByLabel:
		// 遍历deployment，如果name包含key则保存返回
		for _, item := range deployments {
			for k, label := range item.Labels {
				if strings.Contains(label, key) || strings.Contains(k, key) {
					deploymentList = append(deploymentList, item)
					break
				}
			}
		}
	default:
		restfulx.ErrNotNilDebug(fmt.Errorf("参数错误"), restfulx.ParamErr)
	}

	sort.Slice(deploymentList, func(i, j int) bool {
		return deploymentList[j].ObjectMeta.GetCreationTimestamp().Time.Before(deploymentList[i].ObjectMeta.GetCreationTimestamp().Time)
	})

	return deploymentList
}

type DeploymentHandler struct {
	client      *store.ClientConfig
	clusterName string
}

func NewDeploymentHandler(client *store.ClientConfig, clusterName string) *DeploymentHandler {
	return &DeploymentHandler{client: client, clusterName: clusterName}
}

func (d *DeploymentHandler) OnAdd(obj interface{}) {
	d.notifyDeployments(obj)
}

func (d *DeploymentHandler) OnUpdate(oldObj, newObj interface{}) {
	d.notifyDeployments(newObj)
}

func (d *DeploymentHandler) OnDelete(obj interface{}) {
	d.notifyDeployments(obj)
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

	go wsstore.WsClientMap.SendClusterResource(d.clusterName, "deployment", map[string]interface{}{
		"cluster": d.clusterName,
		"type":    "deployment",
		"result": map[string]interface{}{
			"namespace": namespace,
			"data":      deployments,
		},
	})
}
