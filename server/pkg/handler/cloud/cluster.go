package cloud

import (
	store "github.com/lbemi/lbemi/pkg/common/store"
	"github.com/lbemi/lbemi/pkg/model/cloud"
	"github.com/lbemi/lbemi/pkg/model/form"
	"github.com/lbemi/lbemi/pkg/restfulx"
	"github.com/lbemi/lbemi/pkg/services"
)

func (c *cluster) Create(config *form.ClusterReq) {
	_, conf, err := c.factory.Cluster().GenerateClient(config.Name, config.KubeConfig)
	restfulx.ErrNotNilDebug(err, restfulx.RegisterClusterErr)
	c.factory.Cluster().Create(conf)
}

func (c *cluster) Delete(id uint64) {
	info := c.factory.Cluster().Get(id)
	c.factory.Cluster().Delete(id)
	// 停止informer监听
	//c.factory.Cluster().ShutDownInformer(info.Name)
	c.factory.Cluster().RemoveFromStore(info.Name)
}

func (c *cluster) Update(id uint64, config *cloud.Cluster) {
	c.factory.Cluster().Update(id, config)
}

func (c *cluster) Get(id uint64) *cloud.Cluster {
	return c.factory.Cluster().Get(id)
}

func (c *cluster) List() *[]cloud.Cluster {
	return c.factory.Cluster().List()
}

func (c *cluster) GetByName(name string) *cloud.Cluster {
	return c.factory.Cluster().GetByName(name)
}

func (c *cluster) CheckHealth() bool {

	// 获取集群信息
	config := c.factory.Cluster().GetByName(c.clusterName)
	if config == nil {
		return false
	}

	health := c.factory.Cluster().CheckClusterHealth(c.clusterName)
	if health && !config.Status {
		c.ChangeStatus(config.ID, true)
	}

	if !health && config.Status {
		c.ChangeStatus(config.ID, false)
	}

	return true
}

func (c *cluster) ChangeStatus(id uint64, status bool) {
	c.factory.Cluster().ChangeStatus(id, status)
}

func (c *cluster) getClient(name string) *store.ClientConfig {
	return c.factory.Cluster().GetClient(name)
}

func NewCluster(factory services.FactoryImp, clusterName string) *cluster {
	return &cluster{factory: factory, clusterName: clusterName}
}
