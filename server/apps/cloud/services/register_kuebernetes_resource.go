package services

import (
	"github.com/lbemi/lbemi/apps/cloud/api/form"
	"github.com/lbemi/lbemi/apps/cloud/entity"
	istioHandler "github.com/lbemi/lbemi/apps/istio/services"
	"github.com/lbemi/lbemi/pkg/cache"
	"gorm.io/gorm"
)

type ClusterGetter interface {
	Cluster(clusterName string) ICluster
}

type ICluster interface {
	Create(config *form.ClusterReq)
	Delete(id uint64)
	Update(id uint64, config *entity.Cluster)
	Get(id uint64) *entity.Cluster
	List() *[]entity.Cluster
	GetByName(name string) *entity.Cluster
	ChangeStatus(id uint64, status bool)
	CheckHealth() bool
	GenerateClient(name, config string) (*cache.ClientConfig, *entity.Cluster, error)

	// 注册资源接口

	DeploymentGetter
	StatefulSetGetter
	DaemonSetGetter
	NodeGetter
	ServiceGetter
	NamespaceGetter
	SecretGetter
	PodGetter
	JobGetter
	CronJobGetter
	ConfigMapGetter
	IngressesGetter
	EventGetter
	ReplicasetGetter
	PersistentVolumeClaimGetter

	istioHandler.VirtualServiceGetter
}

type Cluster struct {
	clusterName string
	db          *gorm.DB
	store       *cache.ClientMap
}

func (c *Cluster) Events(namespace string) IEvent {
	if namespace == "all" {
		namespace = ""
	}
	return NewEvent(c.getClient(c.clusterName), namespace)
}

func (c *Cluster) Ingresses(namespace string) IIngresses {
	if namespace == "all" {
		namespace = ""
	}
	return NewIngresses(c.getClient(c.clusterName), namespace)
}

func (c *Cluster) ConfigMaps(namespace string) IConfigMap {
	if namespace == "all" {
		namespace = ""
	}
	return NewConfigMap(c.getClient(c.clusterName), namespace)
}

func (c *Cluster) CronJobs(namespace string) ICronJob {
	if namespace == "all" {
		namespace = ""
	}
	return NewCronJob(c.getClient(c.clusterName), namespace)
}

func (c *Cluster) Jobs(namespace string) IJob {
	if namespace == "all" {
		namespace = ""
	}
	return NewJob(c.getClient(c.clusterName), namespace)
}

func (c *Cluster) DaemonSets(namespace string) IDaemonSet {
	if namespace == "all" {
		namespace = ""
	}
	return NewDaemonSet(c.getClient(c.clusterName), namespace)
}

func (c *Cluster) StatefulSets(namespace string) IStatefulSet {
	if namespace == "all" {
		namespace = ""
	}
	return NewStatefulSet(c.getClient(c.clusterName), namespace)
}

func (c *Cluster) Pods(namespace string) IPod {
	if namespace == "all" {
		namespace = ""
	}
	return NewPod(c.getClient(c.clusterName), namespace, c.Events(namespace))
}

// k8s 资源接口

func (c *Cluster) Secrets(namespace string) ISecret {
	if namespace == "all" {
		namespace = ""
	}
	return NewSecret(c.getClient(c.clusterName), namespace)
}

func (c *Cluster) Namespaces() INamespace {
	return NewNamespace(c.getClient(c.clusterName))
}

func (c *Cluster) Service(namespace string) IService {
	if namespace == "all" {
		namespace = ""
	}
	return NewService(c.getClient(c.clusterName), namespace)
}

func (c *Cluster) Nodes() INode {
	return NewNode(c.getClient(c.clusterName), c.Events(c.clusterName), c.Pods(c.clusterName))
}

func (c *Cluster) Deployments(namespace string) IDeployment {
	if namespace == "all" {
		namespace = ""
	}
	return NewDeployment(c.getClient(c.clusterName), namespace)
}
func (c *Cluster) Replicaset(namespace string) ReplicasetImp {
	if namespace == "all" {
		namespace = ""
	}
	return NewReplicaset(c.getClient(c.clusterName), namespace)
}

func (c *Cluster) PersistentVolumeClaim(namespace string) PersistentVolumeClaimImp {
	if namespace == "all" {
		namespace = ""
	}
	return NewPersistentVolumeClaim(c.getClient(c.clusterName), namespace)
}

func (c *Cluster) VirtualServices(namespace string) istioHandler.IVirtualService {
	if namespace == "all" {
		namespace = ""
	}
	return istioHandler.NewVirtualService(c.getClient(c.clusterName), namespace)
}
