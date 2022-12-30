package cloud

import (
	"context"
	"errors"
	"github.com/lbemi/lbemi/pkg/bootstrap/log"
	"github.com/lbemi/lbemi/pkg/common/store"
	"github.com/lbemi/lbemi/pkg/model/cloud"
	"github.com/lbemi/lbemi/pkg/util"
	"gorm.io/gorm"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"time"
)

type ICluster interface {
	GenerateClient(name, config string) (*store.Clients, *cloud.Config, error)

	Create(config *cloud.Config) error
	Delete(id uint64) error
	Update(id uint64, config *cloud.Config) error
	Get(id uint64) (*cloud.Config, error)
	GetByName(name string) (*cloud.Config, error)
	List() (*[]cloud.Config, error)
	GetClient(name string) *store.Clients
	ChangeStatus(id uint64, status bool) error

	RemoveFromStore(name string)
}

type cluster struct {
	db    *gorm.DB
	store *store.ClientStore
}

func NewCluster(db *gorm.DB, store *store.ClientStore) *cluster {
	return &cluster{
		db:    db,
		store: store,
	}
}

func (c *cluster) GenerateClient(name, config string) (*store.Clients, *cloud.Config, error) {

	//如果已经存在或者已经初始化client则退出
	if c.store.Get(name) != nil && c.store.Get(name).IsInit {
		return nil, nil, errors.New("client has already been initialized")
	}

	var client store.Clients
	clientConfig, err := clientcmd.RESTConfigFromKubeConfig([]byte(config))
	if err != nil {
		c.store.Delete(name)
		log.Logger.Error(err)
		return nil, nil, err
	}

	//生成clientSet
	clientSet, err := kubernetes.NewForConfig(clientConfig)
	if err != nil {
		c.store.Delete(name)
		log.Logger.Error(err)
		return nil, nil, err
	}

	var conf cloud.Config
	withTimeout, _ := context.WithTimeout(context.TODO(), time.Second*3)
	list, err := clientSet.CoreV1().Nodes().List(withTimeout, metav1.ListOptions{})
	if err != nil {
		c.store.Delete(name)
		log.Logger.Error(err)
		return nil, nil, errors.New("cluster is not health")
	}

	conf.PodCidr = list.Items[0].Spec.PodCIDR
	conf.RunTime = list.Items[0].Status.NodeInfo.ContainerRuntimeVersion
	conf.Version = list.Items[0].Status.NodeInfo.KubeletVersion
	conf.Status = true
	conf.Nodes = len(list.Items)
	conf.InternalIP = list.Items[0].Status.Addresses[0].Address
	conf.CPU = 0
	conf.Memory = 0

	for _, node := range list.Items {
		conf.CPU = conf.CPU + node.Status.Capacity.Cpu().AsApproximateFloat64()
		conf.Memory = conf.Memory + node.Status.Capacity.Memory().AsApproximateFloat64()
	}

	conf.Memory = conf.Memory / 1024

	conf.Name = name
	conf.KubeConfig = config

	client.ClientSet = clientSet
	//生成informer factory
	client.SharedInformerFactory = informers.NewSharedInformerFactory(clientSet, time.Second*60)
	client.IsInit = true
	c.store.Add(name, &client)

	// 设置需要启动的informer gvr资源
	gvrs := []schema.GroupVersionResource{
		{Group: "", Version: "v1", Resource: "pods"},
		{Group: "", Version: "v1", Resource: "nodes"},
		{Group: "", Version: "v1", Resource: "services"},
		{Group: "", Version: "v1", Resource: "namespaces"},
		{Group: "", Version: "v1", Resource: "events"},
		{Group: "", Version: "v1", Resource: "secrets"},
		{Group: "", Version: "v1", Resource: "configmaps"},
		{Group: "apps", Version: "v1", Resource: "deployments"},
		{Group: "apps", Version: "v1", Resource: "statefulsets"},
		{Group: "apps", Version: "v1", Resource: "daemonsets"},
		{Group: "apps", Version: "v1", Resource: "replicasets"},
		{Group: "networking.k8s.io", Version: "v1beta1", Resource: "ingresses"},

		{Group: "batch", Version: "v1beta1", Resource: "cronjobs"},
		{Group: "batch", Version: "v1", Resource: "jobs"},
	}

	for _, gvr := range gvrs {
		// 实例化informer
		_, err := client.SharedInformerFactory.ForResource(gvr)
		if err != nil {
			log.Logger.Error("informer init failed. err: ", err)
		}
	}
	stopChan := make(chan struct{})
	// 启动informer
	client.SharedInformerFactory.Start(stopChan)
	// 等待informer同步完成
	client.SharedInformerFactory.WaitForCacheSync(stopChan)

	return &client, &conf, nil
}

func (c *cluster) Create(config *cloud.Config) error {

	sec := util.Encrypt(config.KubeConfig)
	config.KubeConfig = sec

	err := c.db.Model(&cloud.Config{}).Create(&config).Error
	if err != nil {
		log.Logger.Error(err)
		return err
	}

	return nil
}

func (c *cluster) Delete(id uint64) error {

	err := c.db.Where("id = ?", id).Delete(&cloud.Config{}).Error
	if err != nil {
		return err
	}

	return nil
}

func (c *cluster) Update(id uint64, config *cloud.Config) error {
	return c.db.Where("id = ?", id).Updates(&config).Error
}

func (c *cluster) Get(id uint64) (*cloud.Config, error) {
	var clu cloud.Config
	err := c.db.Where("id = ?", id).First(&clu).Error
	if err != nil {
		return nil, err
	}
	return &clu, err
}

func (c *cluster) GetByName(name string) (*cloud.Config, error) {
	var clu cloud.Config
	err := c.db.Where("name = ?", name).First(&clu).Error
	if err != nil {
		return nil, err
	}
	return &clu, err
}
func (c *cluster) List() (*[]cloud.Config, error) {
	var clu []cloud.Config
	err := c.db.Find(&clu).Error
	if err != nil {
		return nil, err
	}
	return &clu, nil
}

func (c *cluster) GetClient(name string) *store.Clients {
	return c.store.Get(name)
}

func (c *cluster) ChangeStatus(id uint64, status bool) error {
	return c.db.Model(&cloud.Config{}).Where("id = ?", id).Update("status", status).Error
}

func (c *cluster) RemoveFromStore(name string) {
	c.store.Delete(name)
}
