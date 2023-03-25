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
}

type Factory struct {
	client    *store.Clients
	namespace string
}

func (k *Factory) Deployment() DeploymentImp {
	return newDeployment(k.client, k.namespace)
}

func (k *Factory) ConfigMap() ConfigMapImp {
	return newConfigMap(k.client, k.namespace)
}

func (k *Factory) CronJob() CronJobImp {
	return newCronJob(k.client, k.namespace)
}

func (k *Factory) DaemonSet() DaemonSetImp {
	return newDaemonSet(k.client, k.namespace)
}
func (k *Factory) Event() EventImp {
	return newEvent(k.client, k.namespace)
}
func (k *Factory) Ingress() IngressesImp {
	return newIngress(k.client, k.namespace)
}
func (k *Factory) Job() JobImp {
	return newJob(k.client, k.namespace)
}
func (k *Factory) Namespace() NamespaceImp {
	return newNamespace(k.client)
}
func (k *Factory) Node() NodeImp {
	return newNode(k.client)
}
func (k *Factory) Pod() PodImp {
	return newPod(k.client, k.namespace)
}

func (k *Factory) Replicaset() ReplicasetImp {
	return newReplicaset(k.client, k.namespace)
}
func (k *Factory) Secret() SecretImp {
	return newSecret(k.client, k.namespace)
}
func (k *Factory) Service() ServiceImp {
	return newService(k.client, k.namespace)
}

func (k *Factory) StatefulSet() StatefulSetImp {
	return newStatefulSet(k.client, k.namespace)
}
func NewK8sFactory(client *store.Clients, namespace string) *Factory {
	return &Factory{client: client, namespace: namespace}
}
