package cloud

import (
	"context"
	"github.com/lbemi/lbemi/pkg/handler/kuberntetes"
	"github.com/lbemi/lbemi/pkg/model/cloud"
	"github.com/lbemi/lbemi/pkg/model/form"
	"github.com/lbemi/lbemi/pkg/services"
	cloud2 "github.com/lbemi/lbemi/pkg/services/cloud"
	"github.com/lbemi/lbemi/pkg/util"
)

type ClusterGetter interface {
	Cluster(clusterName string) ICluster
}

type ICluster interface {
	Create(ctx context.Context, config *form.ClusterReq) error
	Delete(ctx context.Context, id *uint64) error
	Update(ctx context.Context, id *uint64, config *cloud.Config) error
	Get(ctx context.Context, id *uint64) (*cloud.Config, error)
	List(ctx context.Context) (*[]cloud.Config, error)

	//GetClient(name string) *cloud2.Clients

	kuberntetes.DeploymentGetter
}

type cluster struct {
	factory     services.IDbFactory
	clusterName string
}

func (c *cluster) Deployments(namespace string) kuberntetes.IDeployment {
	return kuberntetes.NewDeployment(c.getClient(c.clusterName).ClientSet, namespace)
}

func (c *cluster) Create(ctx context.Context, config *form.ClusterReq) error {
	var conf cloud.Config
	conf.Name = config.Name
	conf.Config = config.Config

	util.WithErrorLog(c.factory.Cluster().Create(&conf))
	return nil
}

func (c *cluster) Delete(ctx context.Context, id *uint64) error {
	util.WithErrorLog(c.factory.Cluster().Delete(id))
	return nil
}

func (c *cluster) Update(ctx context.Context, id *uint64, config *cloud.Config) error {
	util.WithErrorLog(c.factory.Cluster().Update(id, config))
	return nil
}

func (c *cluster) Get(ctx context.Context, id *uint64) (*cloud.Config, error) {
	config, err := c.factory.Cluster().Get(id)
	util.WithErrorLog(err)
	return config, nil
}

func (c *cluster) List(ctx context.Context) (*[]cloud.Config, error) {
	list, err := c.factory.Cluster().List()
	util.WithErrorLog(err)
	return list, nil
}

func (c *cluster) getClient(name string) *cloud2.Clients {
	return c.factory.Cluster().GetClient(name)
}
func NewCluster(factory services.IDbFactory, clusterName string) *cluster {
	return &cluster{factory: factory, clusterName: clusterName}
}
