package cloud

import (
	"github.com/lbemi/lbemi/pkg/common/store"
	istioHandler "github.com/lbemi/lbemi/pkg/handler/istio"
	"github.com/lbemi/lbemi/pkg/handler/kubernetes"
	"github.com/lbemi/lbemi/pkg/model/cloud"
	"github.com/lbemi/lbemi/pkg/model/form"
	"gorm.io/gorm"
)

type ClusterGetter interface {
	Cluster(clusterName string) ICluster
}

type ICluster interface {
	Create(config *form.ClusterReq)
	Delete(id uint64)
	Update(id uint64, config *cloud.Cluster)
	Get(id uint64) *cloud.Cluster
	List() *[]cloud.Cluster
	GetByName(name string) *cloud.Cluster
	ChangeStatus(id uint64, status bool)
	CheckHealth() bool
	GenerateClient(name, config string) (*store.ClientConfig, *cloud.Cluster, error)

	// 注册资源接口

	kubernetes.DeploymentGetter
	kubernetes.StatefulSetGetter
	kubernetes.DaemonSetGetter
	kubernetes.NodeGetter
	kubernetes.ServiceGetter
	kubernetes.NamespaceGetter
	kubernetes.SecretGetter
	kubernetes.PodGetter
	kubernetes.JobGetter
	kubernetes.CronJobGetter
	kubernetes.ConfigMapGetter
	kubernetes.IngressesGetter
	kubernetes.EventGetter
	kubernetes.ReplicasetGetter
	kubernetes.PersistentVolumeClaimGetter

	istioHandler.VirtualServiceGetter
}

type Cluster struct {
	clusterName string
	db          *gorm.DB
	store       *store.ClientMap
}

func (c *Cluster) Events(namespace string) kubernetes.IEvent {
	if namespace == "all" {
		namespace = ""
	}
	return kubernetes.NewEvent(c.getClient(c.clusterName), namespace)
}

func (c *Cluster) Ingresses(namespace string) kubernetes.IIngresses {
	if namespace == "all" {
		namespace = ""
	}
	return kubernetes.NewIngresses(c.getClient(c.clusterName), namespace)
}

func (c *Cluster) ConfigMaps(namespace string) kubernetes.IConfigMap {
	if namespace == "all" {
		namespace = ""
	}
	return kubernetes.NewConfigMap(c.getClient(c.clusterName), namespace)
}

func (c *Cluster) CronJobs(namespace string) kubernetes.ICronJob {
	if namespace == "all" {
		namespace = ""
	}
	return kubernetes.NewCronJob(c.getClient(c.clusterName), namespace)
}

func (c *Cluster) Jobs(namespace string) kubernetes.IJob {
	if namespace == "all" {
		namespace = ""
	}
	return kubernetes.NewJob(c.getClient(c.clusterName), namespace)
}

func (c *Cluster) DaemonSets(namespace string) kubernetes.IDaemonSet {
	if namespace == "all" {
		namespace = ""
	}
	return kubernetes.NewDaemonSet(c.getClient(c.clusterName), namespace)
}

func (c *Cluster) StatefulSets(namespace string) kubernetes.IStatefulSet {
	if namespace == "all" {
		namespace = ""
	}
	return kubernetes.NewStatefulSet(c.getClient(c.clusterName), namespace)
}

func (c *Cluster) Pods(namespace string) kubernetes.IPod {
	if namespace == "all" {
		namespace = ""
	}
	return kubernetes.NewPod(c.getClient(c.clusterName), namespace, c.Events(namespace))
}

// k8s 资源接口

func (c *Cluster) Secrets(namespace string) kubernetes.ISecret {
	if namespace == "all" {
		namespace = ""
	}
	return kubernetes.NewSecret(c.getClient(c.clusterName), namespace)
}

func (c *Cluster) Namespaces() kubernetes.INamespace {
	return kubernetes.NewNamespace(c.getClient(c.clusterName))
}

func (c *Cluster) Service(namespace string) kubernetes.IService {
	if namespace == "all" {
		namespace = ""
	}
	return kubernetes.NewService(c.getClient(c.clusterName), namespace)
}

func (c *Cluster) Nodes() kubernetes.INode {
	return kubernetes.NewNode(c.getClient(c.clusterName), c.Events(c.clusterName), c.Pods(c.clusterName))
}

func (c *Cluster) Deployments(namespace string) kubernetes.IDeployment {
	if namespace == "all" {
		namespace = ""
	}
	return kubernetes.NewDeployment(c.getClient(c.clusterName), namespace)
}
func (c *Cluster) Replicaset(namespace string) kubernetes.ReplicasetImp {
	if namespace == "all" {
		namespace = ""
	}
	return kubernetes.NewReplicaset(c.getClient(c.clusterName), namespace)
}

func (c *Cluster) PersistentVolumeClaim(namespace string) kubernetes.PersistentVolumeClaimImp {
	if namespace == "all" {
		namespace = ""
	}
	return kubernetes.NewPersistentVolumeClaim(c.getClient(c.clusterName), namespace)
}

func (c *Cluster) VirtualServices(namespace string) istioHandler.IVirtualService {
	if namespace == "all" {
		namespace = ""
	}
	return istioHandler.NewVirtualService(c.getClient(c.clusterName), namespace)
}
