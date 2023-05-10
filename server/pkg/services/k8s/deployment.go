package k8s

import (
	"context"
	"fmt"
	"github.com/lbemi/lbemi/pkg/bootstrap/log"
	"github.com/lbemi/lbemi/pkg/common/store"
	"github.com/lbemi/lbemi/pkg/common/store/wsstore"
	"github.com/lbemi/lbemi/pkg/handler/types"
	appsv1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"sort"
	"strings"
)

type DeploymentImp interface {
	List(ctx context.Context) ([]*appsv1.Deployment, error)
	Get(ctx context.Context, name string) (*appsv1.Deployment, error)
	Create(ctx context.Context, obj *appsv1.Deployment) (*appsv1.Deployment, error)
	Update(ctx context.Context, obj *appsv1.Deployment) (*appsv1.Deployment, error)
	Delete(ctx context.Context, name string) error
	Scale(ctx context.Context, name string, replicaNum int32) error
	Search(ctx context.Context, key string, searchType int) ([]*appsv1.Deployment, error)
}

type Deployment struct {
	cli       *store.Clients
	namespace string
}

func newDeployment(cli *store.Clients, namespace string) *Deployment {
	return &Deployment{cli: cli, namespace: namespace}
}

func (d *Deployment) List(ctx context.Context) ([]*appsv1.Deployment, error) {
	list, err := d.cli.SharedInformerFactory.Apps().V1().Deployments().Lister().Deployments(d.namespace).List(labels.Everything())
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
	dep, err := d.cli.SharedInformerFactory.Apps().V1().Deployments().Lister().Deployments(d.namespace).Get(name)
	if err != nil {
		log.Logger.Error(err)
	}
	return dep, err
}

func (d *Deployment) Create(ctx context.Context, obj *appsv1.Deployment) (*appsv1.Deployment, error) {
	newDeployment, err := d.cli.ClientSet.AppsV1().Deployments(d.namespace).Create(ctx, obj, metav1.CreateOptions{})
	if err != nil {
		log.Logger.Error(err)
	}
	return newDeployment, err
}

func (d *Deployment) Update(ctx context.Context, obj *appsv1.Deployment) (*appsv1.Deployment, error) {
	updateDeployment, err := d.cli.ClientSet.AppsV1().Deployments(d.namespace).Update(ctx, obj, metav1.UpdateOptions{})
	if err != nil {
		log.Logger.Error(err)
	}
	return updateDeployment, err
}

func (d *Deployment) Delete(ctx context.Context, name string) error {
	err := d.cli.ClientSet.AppsV1().Deployments(d.namespace).Delete(ctx, name, metav1.DeleteOptions{})
	if err != nil {
		log.Logger.Error(err)
	}
	return err
}

func (d *Deployment) Scale(ctx context.Context, name string, replicaNum int32) error {
	oldScale, err := d.cli.ClientSet.AppsV1().Deployments(d.namespace).GetScale(ctx, name, metav1.GetOptions{})
	if err != nil {
		log.Logger.Error(err)
		return err
	}
	oldScale.Spec.Replicas = replicaNum
	_, err = d.cli.ClientSet.AppsV1().Deployments(d.namespace).UpdateScale(ctx, name, oldScale, metav1.UpdateOptions{})
	if err != nil {
		log.Logger.Error(err)
	}
	return nil
}

func (d *Deployment) Search(ctx context.Context, key string, searchType int) ([]*appsv1.Deployment, error) {
	var deploymentList = make([]*appsv1.Deployment, 0)
	deployments, err := d.List(ctx)
	if err != nil {
		log.Logger.Error(err)
		return nil, err
	}
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
		return nil, fmt.Errorf("参数错误")
	}

	sort.Slice(deploymentList, func(i, j int) bool {
		return deploymentList[j].ObjectMeta.GetCreationTimestamp().Time.Before(deploymentList[i].ObjectMeta.GetCreationTimestamp().Time)
	})

	return deploymentList, nil
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
