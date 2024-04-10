package services

import (
	"context"
	"errors"
	"fmt"
	"github.com/lbemi/lbemi/apps/cloud/api/form"
	"github.com/lbemi/lbemi/apps/cloud/entity"
	istio2 "github.com/lbemi/lbemi/apps/istio/services"
	"github.com/lbemi/lbemi/apps/kubernetes/services"
	store "github.com/lbemi/lbemi/pkg/cache"
	"github.com/lbemi/lbemi/pkg/global"
	"github.com/lbemi/lbemi/pkg/restfulx"
	"github.com/lbemi/lbemi/pkg/util"
	tektonVersiond "github.com/tektoncd/pipeline/pkg/client/clientset/versioned"
	tektonExternalversions "github.com/tektoncd/pipeline/pkg/client/informers/externalversions"
	"gorm.io/gorm"
	istio "istio.io/client-go/pkg/clientset/versioned"
	"istio.io/client-go/pkg/informers/externalversions"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/discovery"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/dynamic/dynamicinformer"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/metrics/pkg/client/clientset/versioned"
	"strings"
	"time"
)

func NewCluster(db *gorm.DB, store *store.ClientStore, clusterName string) ICluster {
	return &Cluster{
		db:          db,
		store:       store,
		clusterName: clusterName,
	}
}
func (c *Cluster) Create(config *form.ClusterReq) {
	_, conf, err := c.GenerateClient(config.Name, config.KubeConfig)
	restfulx.ErrNotNilDebug(err, restfulx.RegisterClusterErr)
	conf.KubeConfig = util.Encrypt(conf.KubeConfig)
	err = c.db.Create(conf).Error
	restfulx.ErrNotNilDebug(err, restfulx.OperatorErr)
}

func (c *Cluster) Delete(id uint64) {
	info := c.Get(id)
	restfulx.ErrNotNilDebug(c.db.Where("id = ?", id).Delete(&entity.Cluster{}).Error, restfulx.OperatorErr)
	// 停止informer监听
	c.ShutDownInformer(info.Name)
	c.store.Delete(info.Name)
}

func (c *Cluster) Update(id uint64, config *entity.Cluster) {
	restfulx.ErrNotNilDebug(c.db.Where("id = ?", id).Updates(&config).Error, restfulx.OperatorErr)
}

func (c *Cluster) Get(id uint64) *entity.Cluster {
	var clu entity.Cluster
	restfulx.ErrNotNilDebug(c.db.Where("id = ?", id).First(&clu).Error, restfulx.OperatorErr)
	return &clu
}

func (c *Cluster) List() *[]entity.Cluster {
	var clu []entity.Cluster
	restfulx.ErrNotNilDebug(c.db.Find(&clu).Error, restfulx.OperatorErr)
	return &clu
}

func (c *Cluster) GetByName(name string) *entity.Cluster {
	var clu entity.Cluster
	restfulx.ErrNotNilDebug(c.db.Where("name = ?", name).First(&clu).Error, restfulx.OperatorErr)
	return &clu
}

func (c *Cluster) CheckHealth() bool {

	// 获取集群信息
	config := c.GetByName(c.clusterName)
	if config == nil {
		return false
	}

	health := c.CheckCusterHealth(c.clusterName)
	if health && !config.Status {
		c.ChangeStatus(config.ID, true)
	}

	if !health && config.Status {
		c.ChangeStatus(config.ID, false)
	}

	return true
}

func (c *Cluster) ChangeStatus(id uint64, status bool) {
	restfulx.ErrNotNilDebug(c.db.Model(&entity.Cluster{}).Where("id = ?", id).Update("status", status).Error, restfulx.OperatorErr)
}

func (c *Cluster) getClient(name string) *store.ClientConfig {
	health := c.CheckCusterHealth(name)
	restfulx.ErrNotTrue(health, restfulx.ClusterUnHealth)
	return c.store.Get(name)
}

func (c *Cluster) GenerateClient(name, config string) (*store.ClientConfig, *entity.Cluster, error) {

	//如果已经存在或者已经初始化client则退出
	clients := c.store.Get(name)
	if clients != nil && clients.IsInit {
		return nil, nil, errors.New("client has already been initialized")
	}

	clientConfig, err := clientcmd.RESTConfigFromKubeConfig([]byte(config))

	//TODO  开启的话，istio相关创建代码无法运行，istio client-go代码优化
	// 使用protobuf传输数据
	//clientConfig = metadata.ConfigFor(clientConfig)

	if err != nil {
		c.store.Delete(name)
		return nil, nil, err
	}
	discoveryClient, err := discovery.NewDiscoveryClientForConfig(clientConfig)
	if err != nil {
		c.store.Delete(name)
		return nil, nil, err
	}
	// 生产tekton client
	tektonClient, err := tektonVersiond.NewForConfig(clientConfig)
	if err != nil {
		c.store.Delete(name)
		return nil, nil, err
	}

	//生成clientSet
	clientSet, err := kubernetes.NewForConfig(clientConfig)
	if err != nil {
		c.store.Delete(name)
		return nil, nil, err
	}

	dynamicClient, err := dynamic.NewForConfig(clientConfig)
	if err != nil {
		c.store.Delete(name)
		return nil, nil, err
	}

	withTimeout, cancelFunc := context.WithTimeout(context.TODO(), time.Second*3)
	defer cancelFunc()

	response := clientSet.CoreV1().RESTClient().Get().AbsPath("/healthz").Do(withTimeout)
	//_, err = clientSet.Discovery().ServerVersion()
	if response.Error() != nil {
		return nil, nil, response.Error()
	}

	list, err := clientSet.CoreV1().Nodes().List(withTimeout, metav1.ListOptions{})
	if err != nil {
		c.store.Delete(name)
		return nil, nil, err
	}
	version, _ := clientSet.ServerVersion()

	clusterInfo := &entity.Cluster{
		Name:       name,
		KubeConfig: config,
		PodCidr:    list.Items[0].Spec.PodCIDR,
		RunTime:    list.Items[0].Status.NodeInfo.ContainerRuntimeVersion,
		Version:    version.String(),
		Status:     true,
		Nodes:      len(list.Items),
		InternalIP: list.Items[0].Status.Addresses[0].Address,
		CPU:        0,
		Memory:     0,
	}

	for _, node := range list.Items {
		clusterInfo.CPU = clusterInfo.CPU + node.Status.Capacity.Cpu().AsApproximateFloat64()
		clusterInfo.Memory = clusterInfo.Memory + node.Status.Capacity.Memory().AsApproximateFloat64()
	}

	clusterInfo.Memory = clusterInfo.Memory / 1000

	//初始化metricSet
	metricSet, err := versioned.NewForConfig(clientConfig)
	restfulx.ErrNotNilDebug(err, restfulx.RegisterClusterErr)

	dynamicSharedInformerFactory := dynamicinformer.NewDynamicSharedInformerFactory(dynamicClient, 0)
	istioClient := generateIstioClient(clientConfig)

	client := &store.ClientConfig{
		//生成SharedInformerFactory
		SharedInformerFactory:        informers.NewSharedInformerFactory(clientSet, 0),
		Config:                       clientConfig,
		DynamicSharedInformerFactory: dynamicSharedInformerFactory,
		IstioClient:                  istioClient,
		IstioSharedInformerFactory:   externalversions.NewSharedInformerFactoryWithOptions(istioClient, 0),
		IsInit:                       true,
		StopChan:                     make(chan struct{}),
		MetricSet:                    metricSet,
		ClientSet:                    clientSet,
		DynamicSet:                   dynamicClient,
		DiscoveryClient:              discoveryClient,
		TektonClient:                 tektonClient,
		TektonSharedInformerFactory:  tektonExternalversions.NewSharedInformerFactory(tektonClient, 0),
	}

	c.store.Add(name, client)
	//异步启动informer
	go c.StartInformer(name)
	return client, clusterInfo, nil
}
func (c *Cluster) CheckCusterHealth(name string) bool {

	clients := c.store.Get(name)
	if clients == nil {
		return false
	}

	withTimeout, cancelFunc := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancelFunc()

	// 健康检查
	response := clients.ClientSet.CoreV1().RESTClient().Get().AbsPath("/healthz").Do(withTimeout)
	if response.Error() != nil {
		return false
	}
	return true
}
func generateIstioClient(rc *rest.Config) *istio.Clientset {
	client, err := istio.NewForConfig(rc)
	if err != nil {
		global.Logger.Errorf("generate istio clientSet failed. err : %v", err)
	}
	return client
}
func (c *Cluster) StartInformer(clusterName string) {
	client := c.store.Get(clusterName)
	if client == nil {
		restfulx.ErrNotNilDebug(fmt.Errorf("初始化Informer失败"), restfulx.OperatorErr)
	}

	resources, err := client.DiscoveryClient.ServerPreferredResources()
	if err != nil {
		fmt.Println("初始化Informer失败")
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

			if strings.Contains(gvr.Group, "istio.io") {
				_, _ = client.IstioSharedInformerFactory.ForResource(gvr)
			}
			if strings.Contains(gvr.Group, "tekton.dev") {
				_, _ = client.TektonSharedInformerFactory.ForResource(gvr)
			}

			_, _ = client.SharedInformerFactory.ForResource(gvr)
			//if err != nil {
			//log.Logger.Error(err)
			//}

			//
			//informer := client.DynamicSharedInformerFactory.ForResource(gvr)
			//_ = informer
		}

	}

	client.SharedInformerFactory.Apps().V1().Deployments().Informer().AddEventHandler(services.NewDeploymentHandler(client, clusterName))
	client.SharedInformerFactory.Apps().V1().DaemonSets().Informer().AddEventHandler(services.NewDaemonSetHandler(client, clusterName))
	//client.SharedInformerFactory.Apps().V1().ReplicaSets().Informer().AddEventHandler(k8s.NewReplicasetHandler(client, clusterName))
	client.SharedInformerFactory.Core().V1().Pods().Informer().AddEventHandler(services.NewPodHandler(client, clusterName))
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

	// start k8s informer
	client.SharedInformerFactory.Start(client.StopChan)
	client.SharedInformerFactory.WaitForCacheSync(client.StopChan)

	_, err = client.IstioSharedInformerFactory.Networking().V1beta1().VirtualServices().Informer().AddEventHandler(istio2.NewVirtualServiceHandler(client, clusterName))
	if err != nil {
		global.Logger.Error("add informer handler failed. err:", err)
	}

	// start istio informer
	client.IstioSharedInformerFactory.Start(client.StopChan)
	client.IstioSharedInformerFactory.WaitForCacheSync(client.StopChan)

	client.TektonSharedInformerFactory.Start(client.StopChan)
	client.TektonSharedInformerFactory.WaitForCacheSync(client.StopChan)

}

func (c *Cluster) ShutDownInformer(clusterName string) {
	client := c.store.Get(clusterName)
	if client != nil {
		close(client.StopChan)
	} else {
		restfulx.ErrNotNilDebug(fmt.Errorf("获取client失败"), restfulx.OperatorErr)
	}
}
