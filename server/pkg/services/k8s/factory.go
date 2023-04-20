package k8s

import "github.com/lbemi/lbemi/pkg/common/store"

type FactoryImp interface {
	Deployment() DeploymentImp
	ConfigMap() ConfigMapImp
	CronJob() CronJobImp
	DaemonSet() DaemonSetImp
	Event() EventImp
	Ingress() IngressesImp
	Job() JobImp
	namespace() NamespaceImp
	Node() NodeImp
	Pod() PodImp
	Replicaset() ReplicasetImp
	Secret() SecretImp
	Service() ServiceImp
	StatefulSet() StatefulSetImp
	PersistentVolumeClaim() PersistentVolumeClaimImp
}

type Factory struct {
	client    *store.Clients
	namespace string
}

func (f *Factory) Deployment() DeploymentImp {
	return newDeployment(f.client, f.namespace)
}

func (f *Factory) ConfigMap() ConfigMapImp {
	return newConfigMap(f.client, f.namespace)
}

func (f *Factory) CronJob() CronJobImp {
	return newCronJob(f.client, f.namespace)
}

func (f *Factory) DaemonSet() DaemonSetImp {
	return newDaemonSet(f.client, f.namespace)
}
func (f *Factory) Event() EventImp {
	return newEvent(f.client, f.namespace)
}
func (f *Factory) Ingress() IngressesImp {
	return newIngress(f.client, f.namespace)
}
func (f *Factory) Job() JobImp {
	return newJob(f.client, f.namespace)
}
func (f *Factory) Namespace() NamespaceImp {
	return newNamespace(f.client)
}
func (f *Factory) Node() NodeImp {
	return newNode(f.client)
}
func (f *Factory) Pod() PodImp {
	return newPod(f.client, f.namespace)
}

func (f *Factory) Replicaset() ReplicasetImp {
	return newReplicaset(f.client, f.namespace)
}
func (f *Factory) Secret() SecretImp {
	return newSecret(f.client, f.namespace)
}
func (f *Factory) Service() ServiceImp {
	return newService(f.client, f.namespace)
}

func (f *Factory) StatefulSet() StatefulSetImp {
	return newStatefulSet(f.client, f.namespace)
}

func (f *Factory) PersistentVolumeClaim() PersistentVolumeClaimImp {
	return newPersistentVolumeClaim(f.client, f.namespace)
}
func NewK8sFactory(client *store.Clients, namespace string) *Factory {
	return &Factory{client: client, namespace: namespace}
}
