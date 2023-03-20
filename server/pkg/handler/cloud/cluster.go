package cloud

import (
	"context"
	"github.com/lbemi/lbemi/pkg/bootstrap/log"
	cloud2 "github.com/lbemi/lbemi/pkg/common/store"
	"github.com/lbemi/lbemi/pkg/handler/kuberntetes"
	"github.com/lbemi/lbemi/pkg/model/cloud"
	"github.com/lbemi/lbemi/pkg/model/form"
	"github.com/lbemi/lbemi/pkg/services"
	"github.com/lbemi/lbemi/pkg/util"
)

type ClusterGetter interface {
	Cluster(clusterName string) ICluster
}

type ICluster interface {
	Create(ctx context.Context, config *form.ClusterReq) error
	Delete(ctx context.Context, id uint64) error
	Update(ctx context.Context, id uint64, config *cloud.Config) error
	Get(ctx context.Context, id uint64) (*cloud.Config, error)
	List(ctx context.Context) (*[]cloud.Config, error)
	GetByName(ctx context.Context, name string) (*cloud.Config, error)
	ChangeStatus(id uint64, status bool) error
	CheckHealth(ctx context.Context) bool
	//GenerateClient(name, config string) (*cloud2.Clients, *cloud.Config, error)

	// 注册资源接口

	kuberntetes.DeploymentGetter
	kuberntetes.StatefulSetGetter
	kuberntetes.DaemonSetGetter
	kuberntetes.NodeGetter
	kuberntetes.ServiceGetter
	kuberntetes.NamespaceGetter
	kuberntetes.SecretGetter
	kuberntetes.PodGetter
	kuberntetes.JobGetter
	kuberntetes.CronJobGetter
	kuberntetes.ConfigMapGetter
	kuberntetes.IngressesGetter
	kuberntetes.EventGetter
}

type cluster struct {
	factory     services.IDbFactory
	clusterName string
}

func (c *cluster) Events(namespace string) kuberntetes.IEvent {
	if namespace == "all" {
		namespace = ""
	}
	return kuberntetes.NewEvent(c.getClient(c.clusterName), namespace)
}
func (c *cluster) Ingresses(namespace string) kuberntetes.IIngresses {
	if namespace == "" {
		namespace = "default"
	}
	return kuberntetes.NewIngress(c.getClient(c.clusterName), namespace)
}

func (c *cluster) ConfigMaps(namespace string) kuberntetes.IConfigMap {
	if namespace == "" {
		namespace = "default"
	}
	return kuberntetes.NewConfigMap(c.getClient(c.clusterName), namespace)
}

func (c *cluster) CronJobs(namespace string) kuberntetes.ICronJob {
	if namespace == "" {
		namespace = "default"
	}
	return kuberntetes.NewCronJob(c.getClient(c.clusterName), namespace)
}

func (c *cluster) Jobs(namespace string) kuberntetes.IJob {
	if namespace == "all" {
		namespace = ""
	}
	return kuberntetes.NewJob(c.getClient(c.clusterName), namespace)
}

func (c *cluster) DaemonSets(namespace string) kuberntetes.IDaemonSet {
	if namespace == "all" {
		namespace = ""
	}
	return kuberntetes.NewDaemonSet(c.getClient(c.clusterName), namespace)
}

func (c *cluster) StatefulSets(namespace string) kuberntetes.IStatefulSet {
	if namespace == "all" {
		namespace = ""
	}
	return kuberntetes.NewStatefulSet(c.getClient(c.clusterName), namespace)
}

func (c *cluster) Pods(namespace string) kuberntetes.IPod {
	if namespace == "all" {
		namespace = ""
	}
	return kuberntetes.NewPod(c.getClient(c.clusterName), namespace)
}

// kubernetes 资源接口

func (c *cluster) Secrets(namespace string) kuberntetes.ISecret {
	if namespace == "all" {
		namespace = ""
	}
	return kuberntetes.NewSecret(c.getClient(c.clusterName), namespace)
}

func (c *cluster) Namespaces() kuberntetes.INamespace {
	return kuberntetes.NewNamespace(c.getClient(c.clusterName))
}

func (c *cluster) Service(namespace string) kuberntetes.IService {
	if namespace == "all" {
		namespace = ""
	}
	return kuberntetes.NewService(c.getClient(c.clusterName), namespace)
}

func (c *cluster) Nodes() kuberntetes.INode {
	return kuberntetes.NewNode(c.getClient(c.clusterName))
}

func (c *cluster) Deployments(namespace string) kuberntetes.IDeployment {
	if namespace == "all" {
		namespace = ""
	}

	dep := kuberntetes.NewDeployment(c.getClient(c.clusterName), namespace)
	//dep.Start()
	return dep
}

func (c *cluster) Create(ctx context.Context, config *form.ClusterReq) error {

	_, conf, err := c.factory.Cluster().GenerateClient(config.Name, config.KubeConfig)
	if err != nil || conf == nil {
		log.Logger.Error(err)
		return err
	}

	util.WithErrorLog(c.factory.Cluster().Create(conf))
	return nil
}

func (c *cluster) Delete(ctx context.Context, id uint64) error {
	info, err := c.factory.Cluster().Get(id)
	if err != nil {
		return err
	}
	go c.factory.Cluster().RemoveFromStore(info.Name)
	util.WithErrorLog(c.factory.Cluster().Delete(id))

	return nil
}

func (c *cluster) Update(ctx context.Context, id uint64, config *cloud.Config) error {
	util.WithErrorLog(c.factory.Cluster().Update(id, config))
	return nil
}

func (c *cluster) Get(ctx context.Context, id uint64) (*cloud.Config, error) {
	config, err := c.factory.Cluster().Get(id)
	util.WithErrorLog(err)
	return config, nil
}

func (c *cluster) List(ctx context.Context) (*[]cloud.Config, error) {
	list, err := c.factory.Cluster().List()
	util.WithErrorLog(err)
	return list, nil
}

func (c *cluster) GetByName(ctx context.Context, name string) (*cloud.Config, error) {
	clusterInfo, err := c.factory.Cluster().GetByName(name)
	util.WithErrorLog(err)
	return clusterInfo, nil
}

func (c *cluster) CheckHealth(ctx context.Context) bool {

	// 获取集群信息
	config, err := c.factory.Cluster().GetByName(c.clusterName)
	if err != nil || config == nil {
		return false
	}

	health := c.factory.Cluster().CheckCusterHealth(c.clusterName)
	if health && !config.Status {
		err = c.ChangeStatus(config.ID, true)
		if err != nil {
			return false
		}
		return true
	}

	if !health && config.Status {
		err = c.ChangeStatus(config.ID, false)
		if err != nil {
			return false
		}
		return false
	}

	return true
}

func (c *cluster) ChangeStatus(id uint64, status bool) error {
	return c.factory.Cluster().ChangeStatus(id, status)
}

func (c *cluster) getClient(name string) *cloud2.Clients {
	return c.factory.Cluster().GetClient(name)
}

func NewCluster(factory services.IDbFactory, clusterName string) *cluster {
	return &cluster{factory: factory, clusterName: clusterName}
}
