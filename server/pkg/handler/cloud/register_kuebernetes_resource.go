package cloud

import (
	istioHandler "github.com/lbemi/lbemi/pkg/handler/istio"
	"github.com/lbemi/lbemi/pkg/handler/kubernetes"
	"github.com/lbemi/lbemi/pkg/model/cloud"
	"github.com/lbemi/lbemi/pkg/model/form"
	"github.com/lbemi/lbemi/pkg/services"
	"github.com/lbemi/lbemi/pkg/services/istio"
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
	//GenerateClient(name, config string) (*cloud2.ClientConfig, *cloud.Cluster, error)

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

type cluster struct {
	factory services.Interface
	//k8s         k8s.Interface
	istio       istio.Interface
	clusterName string
}

func (c *cluster) Events(namespace string) kubernetes.IEvent {
	if namespace == "all" {
		namespace = ""
	}
	return kubernetes.NewEvent(c.getClient(c.clusterName), namespace)
}

func (c *cluster) Ingresses(namespace string) kubernetes.IIngresses {
	if namespace == "all" {
		namespace = ""
	}
	return kubernetes.NewIngresses(c.getClient(c.clusterName), namespace)
}

func (c *cluster) ConfigMaps(namespace string) kubernetes.IConfigMap {
	if namespace == "all" {
		namespace = ""
	}
	return kubernetes.NewConfigMap(c.getClient(c.clusterName), namespace)
}

func (c *cluster) CronJobs(namespace string) kubernetes.ICronJob {
	if namespace == "all" {
		namespace = ""
	}
	return kubernetes.NewCronJob(c.getClient(c.clusterName), namespace)
}

func (c *cluster) Jobs(namespace string) kubernetes.IJob {
	if namespace == "all" {
		namespace = ""
	}
	return kubernetes.NewJob(c.getClient(c.clusterName), namespace)
}

func (c *cluster) DaemonSets(namespace string) kubernetes.IDaemonSet {
	if namespace == "all" {
		namespace = ""
	}
	return kubernetes.NewDaemonSet(c.getClient(c.clusterName), namespace)
}

func (c *cluster) StatefulSets(namespace string) kubernetes.IStatefulSet {
	if namespace == "all" {
		namespace = ""
	}
	return kubernetes.NewStatefulSet(c.getClient(c.clusterName), namespace)
}

func (c *cluster) Pods(namespace string) kubernetes.IPod {
	if namespace == "all" {
		namespace = ""
	}
	return kubernetes.NewPod(c.getClient(c.clusterName), namespace, c.Events(namespace))
}

// k8s 资源接口

func (c *cluster) Secrets(namespace string) kubernetes.ISecret {
	if namespace == "all" {
		namespace = ""
	}
	return kubernetes.NewSecret(c.getClient(c.clusterName), namespace)
}

func (c *cluster) Namespaces() kubernetes.INamespace {
	return kubernetes.NewNamespace(c.getClient(c.clusterName))
}

func (c *cluster) Service(namespace string) kubernetes.IService {
	if namespace == "all" {
		namespace = ""
	}
	return kubernetes.NewService(c.getClient(c.clusterName), namespace)
}

func (c *cluster) Nodes() kubernetes.INode {
	return kubernetes.NewNode(c.getClient(c.clusterName), c.Events(c.clusterName), c.Pods(c.clusterName))
}

func (c *cluster) Deployments(namespace string) kubernetes.IDeployment {
	if namespace == "all" {
		namespace = ""
	}
	return kubernetes.NewDeployment(c.getClient(c.clusterName), namespace)
}
func (c *cluster) Replicaset(namespace string) kubernetes.ReplicasetImp {
	if namespace == "all" {
		namespace = ""
	}
	return kubernetes.NewReplicaset(c.getClient(c.clusterName), namespace)
}

func (c *cluster) PersistentVolumeClaim(namespace string) kubernetes.PersistentVolumeClaimImp {
	if namespace == "all" {
		namespace = ""
	}
	return kubernetes.NewPersistentVolumeClaim(c.getClient(c.clusterName), namespace)
}

func (c *cluster) VirtualServices(namespace string) istioHandler.IVirtualService {
	if namespace == "all" {
		namespace = ""
	}
	return istioHandler.NewVirtualService(istio.NewIstioFactory(c.getClient(c.clusterName), namespace))
}
