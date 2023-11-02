package services

import (
	"github.com/lbemi/lbemi/pkg/cache"
)

type K8SGetter interface {
	K8S() K8SInterface
}

type K8SInterface interface {
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
}

type K8S struct {
	clusterName string
	store       *cache.ClientStore
}

func NewK8S(clusterName string, store *cache.ClientStore) *K8S {
	return &K8S{clusterName: clusterName, store: store}
}

func (k *K8S) Events(namespace string) IEvent {
	if namespace == "all" {
		namespace = ""
	}
	return NewEvent(k.store.Get(k.clusterName), namespace)
}

func (k *K8S) Ingresses(namespace string) IIngresses {
	if namespace == "all" {
		namespace = ""
	}
	return NewIngresses(k.store.Get(k.clusterName), namespace)
}

func (k *K8S) ConfigMaps(namespace string) IConfigMap {
	if namespace == "all" {
		namespace = ""
	}
	return NewConfigMap(k.store.Get(k.clusterName), namespace)
}

func (k *K8S) CronJobs(namespace string) ICronJob {
	if namespace == "all" {
		namespace = ""
	}
	return NewCronJob(k.store.Get(k.clusterName), namespace)
}

func (k *K8S) Jobs(namespace string) IJob {
	if namespace == "all" {
		namespace = ""
	}
	return NewJob(k.store.Get(k.clusterName), namespace)
}

func (k *K8S) DaemonSets(namespace string) IDaemonSet {
	if namespace == "all" {
		namespace = ""
	}
	return NewDaemonSet(k.store.Get(k.clusterName), namespace)
}

func (k *K8S) StatefulSets(namespace string) IStatefulSet {
	if namespace == "all" {
		namespace = ""
	}
	return NewStatefulSet(k.store.Get(k.clusterName), namespace)
}

func (k *K8S) Pods(namespace string) IPod {
	if namespace == "all" {
		namespace = ""
	}
	return NewPod(k.store.Get(k.clusterName), namespace, k.Events(namespace))
}

// k8s 资源接口

func (k *K8S) Secrets(namespace string) ISecret {
	if namespace == "all" {
		namespace = ""
	}
	return NewSecret(k.store.Get(k.clusterName), namespace)
}

func (k *K8S) Namespaces() INamespace {
	return NewNamespace(k.store.Get(k.clusterName))
}

func (k *K8S) Service(namespace string) IService {
	if namespace == "all" {
		namespace = ""
	}
	return NewService(k.store.Get(k.clusterName), namespace)
}

func (k *K8S) Nodes() INode {
	return NewNode(k.store.Get(k.clusterName), k.Events(k.clusterName), k.Pods(k.clusterName))
}

func (k *K8S) Deployments(namespace string) IDeployment {
	if namespace == "all" {
		namespace = ""
	}
	return NewDeployment(k.store.Get(k.clusterName), namespace)
}
func (k *K8S) Replicaset(namespace string) ReplicasetImp {
	if namespace == "all" {
		namespace = ""
	}
	return NewReplicaset(k.store.Get(k.clusterName), namespace)
}

func (k *K8S) PersistentVolumeClaim(namespace string) PersistentVolumeClaimImp {
	if namespace == "all" {
		namespace = ""
	}
	return NewPersistentVolumeClaim(k.store.Get(k.clusterName), namespace)
}
