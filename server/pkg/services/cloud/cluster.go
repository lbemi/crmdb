package cloud

import (
	"errors"
	"github.com/lbemi/lbemi/pkg/bootstrap/log"
	"github.com/lbemi/lbemi/pkg/model/cloud"
	"github.com/lbemi/lbemi/pkg/util"
	"gorm.io/gorm"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"time"
)

type ICluster interface {
	GenerateClient(name, config string) (*Clients, error)

	Create(config *cloud.Config) error
	Delete(id *uint64) error
	Update(id *uint64, config *cloud.Config) error
	Get(id *uint64) (*cloud.Config, error)
	List() (*[]cloud.Config, error)
	GetClient(name string) *Clients
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

func (c *cluster) GenerateClient(name, config string) (*Clients, error) {

	//如果已经存在或者已经初始化client则退出
	if c.store.Get(name) != nil && c.store.Get(name).IsInit {

		return nil, errors.New("client has invited")
	}
	var client Clients
	clientConfig, err := clientcmd.RESTConfigFromKubeConfig([]byte(config))
	if err != nil {
		log.Logger.Error(err)
		return nil, err
	}
	//生成clientSet
	clientSet, err := kubernetes.NewForConfig(clientConfig)
	if err != nil {
		log.Logger.Error(err)
		return nil, err
	}

	client.ClientSet = clientSet
	//生成informer factory
	client.Factory = informers.NewSharedInformerFactory(clientSet, time.Second*30)
	client.IsInit = true
	c.store.Add(name, &client)

	return &client, nil
}

func (c *cluster) Create(config *cloud.Config) error {

	//_, err := c.GenerateClient(config.Name, config.KubeConfig)
	//if err != nil {
	//	log.Logger.Error(err)
	//	return err
	//}

	sec := util.Encrypt(config.KubeConfig)
	config.KubeConfig = sec

	err := c.db.Model(&cloud.Config{}).Create(&config).Error
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
	return c.store.Get(name)
}
