package cloud

import (
	"context"
	"errors"
	"github.com/lbemi/lbemi/pkg/bootstrap/log"
	"github.com/lbemi/lbemi/pkg/model/cloud"
	"github.com/lbemi/lbemi/pkg/util"
	"gorm.io/gorm"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"time"
)

type ICluster interface {
	GenerateClient(name, config string) (*Clients, *cloud.Config, error)

	Create(config *cloud.Config) error
	Delete(id uint64) error
	Update(id uint64, config *cloud.Config) error
	Get(id uint64) (*cloud.Config, error)
	GetByName(name string) (*cloud.Config, error)
	List() (*[]cloud.Config, error)
	GetClient(name string) *Clients
	ChangeStatus(id uint64, status bool) error
}

type cluster struct {
	db    *gorm.DB
	store *ClientStore
}

func NewCluster(db *gorm.DB, store *ClientStore) *cluster {
	return &cluster{
		db:    db,
		store: store,
	}
}

func (c *cluster) GenerateClient(name, config string) (*Clients, *cloud.Config, error) {

	//如果已经存在或者已经初始化client则退出
	if c.store.Get(name) != nil && c.store.Get(name).IsInit {
		return nil, nil, errors.New("client has already been initialized")
	}

	var client Clients
	clientConfig, err := clientcmd.RESTConfigFromKubeConfig([]byte(config))
	if err != nil {
		log.Logger.Error(err)
		return nil, nil, err
	}

	//生成clientSet
	clientSet, err := kubernetes.NewForConfig(clientConfig)
	if err != nil {
		log.Logger.Error(err)
		return nil, nil, err
	}

	var conf cloud.Config
	list, err := clientSet.CoreV1().Nodes().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
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
	client.Factory = informers.NewSharedInformerFactory(clientSet, time.Second*30)
	client.IsInit = true
	// TODO Informers
	//client.Informers.
	//serviceInformer := client.Factory.Core().V1().Services()
	//podInformer := client.Factory.Core().V1().Pods()
	//nodeInformer := client.Factory.Core().V1().Nodes()
	//client.Factory.Core().V1().

	//deploymentInformer := client.Factory.Apps().V1().Deployments()
	//daemonSetInformer := client.Factory.Apps().V1().DaemonSets()
	//statefulSetInformer := client.Factory.Apps().V1().StatefulSets()
	c.store.Add(name, &client)

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

func (c *cluster) GetClient(name string) *Clients {
	return c.store.Get(name)
}

func (c *cluster) ChangeStatus(id uint64, status bool) error {
	return c.db.Model(&cloud.Config{}).Where("id = ?", id).Update("status", status).Error
}
