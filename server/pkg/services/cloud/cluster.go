package cloud

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"gorm.io/gorm"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/discovery"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/dynamic/dynamicinformer"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/metadata"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/metrics/pkg/client/clientset/versioned"

	"github.com/lbemi/lbemi/pkg/bootstrap/log"
	"github.com/lbemi/lbemi/pkg/common/store"
	"github.com/lbemi/lbemi/pkg/model/cloud"
	"github.com/lbemi/lbemi/pkg/restfulx"
	"github.com/lbemi/lbemi/pkg/services/k8s"
	"github.com/lbemi/lbemi/pkg/util"
)

type ICluster interface {
	GenerateClient(name, config string) (*store.ClientConfig, *cloud.Cluster, error)
	CheckClusterHealth(name string) bool

	Create(config *cloud.Cluster)
	Delete(id uint64)
	Update(id uint64, config *cloud.Cluster)
	Get(id uint64) *cloud.Cluster
	GetByName(name string) *cloud.Cluster
	List() *[]cloud.Cluster
	GetClient(name string) *store.ClientConfig
	ChangeStatus(id uint64, status bool)

	RemoveFromStore(name string)
	StartInformer(clusterName string)
	ShutDownInformer(clusterName string)
}

type Cluster struct {
	db    *gorm.DB
	store *store.ClientMap
}

func NewCluster(db *gorm.DB, store *store.ClientMap) *Cluster {
	return &Cluster{
		db:    db,
		store: store,
	}
}

// CheckClusterHealth checks the health of the cluster.
func (c *Cluster) CheckClusterHealth(name string) bool {
	client := c.store.Get(name)
	if client == nil {
		return false
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// Health check
	response := client.ClientSet.CoreV1().RESTClient().Get().AbsPath("/healthz").Do(ctx)
	return response.Error() == nil
}

func (c *Cluster) GenerateClient(name, config string) (*store.ClientConfig, *cloud.Cluster, error) {

	// Check if the client has already been initialized
	client := c.store.Get(name)
	if client != nil && client.IsInit {
		return nil, nil, errors.New("client has already been initialized")
	}

	// Parse the kubeconfig
	clientConfig, err := clientcmd.RESTConfigFromKubeConfig([]byte(config))
	if err != nil {
		c.store.Delete(name)
		return nil, nil, err
	}

	// Use protobuf for data transmission
	clientConfig = metadata.ConfigFor(clientConfig)

	discoveryClient, err := discovery.NewDiscoveryClientForConfig(clientConfig)
	if err != nil {
		c.store.Delete(name)
		return nil, nil, err
	}

	// Generate clientset
	clientSet, err := kubernetes.NewForConfig(clientConfig)
	if err != nil {
		c.store.Delete(name)
		return nil, nil, err
	}

	// Genrate dynamic client
	dynamicClient, err := dynamic.NewForConfig(clientConfig)
	if err != nil {
		c.store.Delete(name)
		return nil, nil, err
	}

	conf := cloud.Cluster{Name: name, KubeConfig: config}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	response := clientSet.CoreV1().RESTClient().Get().AbsPath("/healthz").Do(ctx)
	//_, err = clientSet.Discovery().ServerVersion()
	if response.Error() != nil {
		return nil, nil, response.Error()
	}

	list, err := clientSet.CoreV1().Nodes().List(ctx, metav1.ListOptions{})
	if err != nil {
		c.store.Delete(name)
		return nil, nil, err
	}

	conf.PodCidr = list.Items[0].Spec.PodCIDR
	conf.RunTime = list.Items[0].Status.NodeInfo.ContainerRuntimeVersion
	version, _ := clientSet.ServerVersion()
	conf.Version = version.String()
	conf.Status = true
	conf.Nodes = len(list.Items)
	conf.InternalIP = list.Items[0].Status.Addresses[0].Address
	conf.CPU = 0
	conf.Memory = 0

	for _, node := range list.Items {
		conf.CPU += node.Status.Capacity.Cpu().AsApproximateFloat64()
		conf.Memory += node.Status.Capacity.Memory().AsApproximateFloat64()
	}
	conf.Memory /= 1000

	//初始化metricSet
	metricSet, err := versioned.NewForConfig(clientConfig)
	restfulx.ErrNotNilDebug(err, restfulx.RegisterClusterErr)

	dynamicSharedInformerFactory := dynamicinformer.NewDynamicSharedInformerFactory(dynamicClient, 0)
	client = &store.ClientConfig{}

	//生成SharedInformerFactory
	client.SharedInformerFactory = informers.NewSharedInformerFactory(clientSet, 0)
	client.IsInit = true
	client.StopChan = make(chan struct{})
	c.store.Add(name, client)

	//异步启动informer
	go c.StartInformer(name)

	client.MetricSet = metricSet
	client.ClientSet = clientSet
	client.Config = clientConfig
	client.DynamicSet = dynamicClient
	client.DynamicSharedInformerFactory = dynamicSharedInformerFactory
	client.DiscoveryClient = discoveryClient
	return client, &conf, nil
}

func (c *Cluster) Create(config *cloud.Cluster) {

	sec := util.Encrypt(config.KubeConfig)
	config.KubeConfig = sec

	err := c.db.Model(&cloud.Cluster{}).Create(&config).Error
	restfulx.ErrNotNilDebug(err, restfulx.OperatorErr)
}

func (c *Cluster) Delete(id uint64) {
	restfulx.ErrNotNilDebug(c.db.Where("id = ?", id).Delete(&cloud.Cluster{}).Error, restfulx.OperatorErr)
}

func (c *Cluster) Update(id uint64, config *cloud.Cluster) {
	restfulx.ErrNotNilDebug(c.db.Where("id = ?", id).Updates(&config).Error, restfulx.OperatorErr)
}

func (c *Cluster) Get(id uint64) *cloud.Cluster {
	var clu cloud.Cluster
	restfulx.ErrNotNilDebug(c.db.Where("id = ?", id).First(&clu).Error, restfulx.OperatorErr)
	return &clu
}

func (c *Cluster) GetByName(name string) *cloud.Cluster {
	var clu cloud.Cluster
	restfulx.ErrNotNilDebug(c.db.Where("name = ?", name).First(&clu).Error, restfulx.OperatorErr)
	return &clu
}

func (c *Cluster) List() *[]cloud.Cluster {
	var clu []cloud.Cluster
	restfulx.ErrNotNilDebug(c.db.Find(&clu).Error, restfulx.OperatorErr)
	return &clu
}

func (c *Cluster) ChangeStatus(id uint64, status bool) {
	restfulx.ErrNotNilDebug(c.db.Model(&cloud.Cluster{}).Where("id = ?", id).Update("status", status).Error, restfulx.OperatorErr)
}

func (c *Cluster) RemoveFromStore(name string) {
	c.store.Delete(name)
}

func (c *Cluster) GetClient(name string) *store.ClientConfig {
	restfulx.ErrNotTrue(c.CheckClusterHealth(name), restfulx.ClusterUnHealth)
	return c.store.Get(name)
}

// StartInformer initializes and starts the informer for the given cluster.
func (c *Cluster) StartInformer(clusterName string) {
	client := c.store.Get(clusterName)
	if client != nil && client.DiscoveryClient != nil {
		resources, err := client.DiscoveryClient.ServerPreferredResources()
		if err != nil {
			restfulx.ErrNotNilDebug(fmt.Errorf("初始化 %v Informer失败", clusterName), restfulx.OperatorErr)
		}

		for _, apiResource := range resources {
			groupVersion := strings.Split(apiResource.GroupVersion, "/")
			gvr := schema.GroupVersionResource{}
			if len(groupVersion) == 1 {
				gvr.Group = ""
				gvr.Version = groupVersion[0]
			} else {
				gvr.Group = groupVersion[0]
				gvr.Version = groupVersion[1]
			}
			for _, v := range apiResource.APIResources {
				// v1 ComponentStatus is deprecated in v1.19+
				if v.Name == "componentstatuses" {
					continue
				}
				gvr.Resource = v.Name

				_, err := client.SharedInformerFactory.ForResource(gvr)
				if err != nil {
					log.Logger.Error(err)
				}
				//
				//informer := client.DynamicSharedInformerFactory.ForResource(gvr)
				//_ = informer
			}

		}

		client.SharedInformerFactory.Apps().V1().Deployments().Informer().AddEventHandler(k8s.NewDeploymentHandler(client, clusterName))
		//client.SharedInformerFactory.Apps().V1().ReplicaSets().Informer().AddEventHandler(k8s.NewReplicasetHandler(client, clusterName))
		client.SharedInformerFactory.Core().V1().Pods().Informer().AddEventHandler(k8s.NewPodHandler(client, clusterName))
		//client.SharedInformerFactory.Core().V1().Namespaces().Informer().AddEventHandler(k8s.NewNameSpaceHandler(client, clusterName))
		//client.SharedInformerFactory.Core().V1().Events().Informer().AddEventHandler(k8s.NewEventHandler())
		//client.SharedInformerFactory.Core().V1().Nodes().Informer().AddEventHandler(k8s.NewNodeHandler())
		//client.SharedInformerFactory.Core().V1().ConfigMaps().Informer().AddEventHandler(k8s.NewConfigMapHandler(client, clusterName))
		//client.SharedInformerFactory.Core().V1().Secrets().Informer().AddEventHandler(k8s.NewSecretHandle())
		//client.SharedInformerFactory.Core().V1().Services().Informer().AddEventHandler(k8s.NewServiceHandle())
		//client.SharedInformerFactory.Core().V1().PersistentVolumeClaims().Informer().AddEventHandler(k8s.NewPersistentVolumeClaimHandler(client, clusterName))
		//client.SharedInformerFactory.Networking().V1().Ingresses().Informer().AddEventHandler(k8s.NewIngressHandle())
		//client.SharedInformerFactory.Networking().V1beta1().Ingresses().Informer().AddEventHandler(k8s.NewIngressHandle())
		//client.SharedInformerFactory.Extensions().V1beta1().Ingresses().Informer().AddEventHandler(k8s.NewIngressHandle())
		//client.SharedInformerFactory.Apps().V1().StatefulSets().Informer().AddEventHandler(k8s.NewStatefulSetHandle())

		//client.DynamicSharedInformerFactory.Start(client.StopChan)
		//client.DynamicSharedInformerFactory.WaitForCacheSync(client.StopChan)
		// 启动informer
		client.SharedInformerFactory.Start(client.StopChan)
		// 等待informer同步完成
		client.SharedInformerFactory.WaitForCacheSync(client.StopChan)
	} else {
		restfulx.ErrNotNilDebug(fmt.Errorf("获取client失败"), restfulx.OperatorErr)
	}
}

// ShutDownInformer shuts down the informer for the given cluster.
func (c *Cluster) ShutDownInformer(clusterName string) {
	client := c.store.Get(clusterName)
	if client == nil {
		restfulx.ErrNotNilDebug(fmt.Errorf("获取client失败"), restfulx.OperatorErr)
		return
	}
	close(client.StopChan)
}
