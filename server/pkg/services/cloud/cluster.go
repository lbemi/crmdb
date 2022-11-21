package cloud

import (
	"github.com/lbemi/lbemi/pkg/bootstrap/log"
	"github.com/lbemi/lbemi/pkg/model/cloud"
	"gorm.io/gorm"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"sync"
	"time"
)

type ICluster interface {
	GenerateClient(name, config string) error

	Create(config *cloud.Config) error
	Delete(id *uint64) error
	Update(id *uint64, config *cloud.Config) error
	Get(id *uint64) (*cloud.Config, error)
	List() (*[]cloud.Config, error)
	GetClient(name string) *Clients
}

type cluster struct {
	db      *gorm.DB
	clients map[string]*Clients
	once    sync.Once
}

func NewCluster(db *gorm.DB) *cluster {
	return &cluster{
		db: db,
	}
}

func (c *cluster) GenerateClient(name, config string) error {
	c.once.Do(func() {
		c.clients = map[string]*Clients{}
	})
	//如果已经存在或者已经初始化client则退出
	if c.clients[name] != nil && c.clients[name].IsInit {
		return nil
	}
	var client Clients
	clientConfig, err := clientcmd.RESTConfigFromKubeConfig([]byte(config))
	if err != nil {
		log.Logger.Error(err)
		return err
	}
	//生成clientSet
	clientSet, err := kubernetes.NewForConfig(clientConfig)
	if err != nil {
		log.Logger.Error(err)
		return err
	}

	client.ClientSet = clientSet
	//生成informer factory
	client.Factory = informers.NewSharedInformerFactory(clientSet, time.Second*30)
	client.IsInit = true
	c.clients[name] = &client
	log.Logger.Info("@@@@@@@@@", c.clients)

	return nil
}

func (c *cluster) Create(config *cloud.Config) error {

	err := c.GenerateClient(config.Name, config.Config)
	if err != nil {
		log.Logger.Error(err)
		return err
	}

	err = c.db.Model(&cloud.Config{}).Create(&config).Error
	if err != nil {
		log.Logger.Error(err)
		return err
	}

	return nil
}

func (c *cluster) Delete(id *uint64) error {
	return c.db.Where("id = ?", id).Delete(&cloud.Config{}).Error
}

func (c *cluster) Update(id *uint64, config *cloud.Config) error {
	return c.db.Where("id = ?", id).Updates(&config).Error
}

func (c *cluster) Get(id *uint64) (*cloud.Config, error) {
	var clu cloud.Config
	err := c.db.Where("id = ?", id).First(&clu).Error
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
	log.Logger.Info("-------", c.clients["test"])
	return c.clients[name]
}
