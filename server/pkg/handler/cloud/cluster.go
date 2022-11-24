package cloud

import (
	"context"
	"github.com/lbemi/lbemi/pkg/bootstrap/log"
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
	Delete(ctx context.Context, id uint64) error
	Update(ctx context.Context, id uint64, config *cloud.Config) error
	Get(ctx context.Context, id uint64) (*cloud.Config, error)
	List(ctx context.Context) (*[]cloud.Config, error)
	ChangeStatus(id uint64, status bool) error
	CheckHealth(ctx context.Context) bool
	//GenerateClient(name, config string) (*cloud2.Clients, *cloud.Config, error)

	// 注册资源接口

	kuberntetes.DeploymentGetter
	kuberntetes.NodeGetter
	kuberntetes.ServiceGetter
	kuberntetes.NamespaceGetter
	kuberntetes.SecretGetter
}

type cluster struct {
	factory     services.IDbFactory
	clusterName string
}

// kubernetes 资源接口

func (c *cluster) Secrets(namespace string) kuberntetes.ISecret {
	if namespace == "" {
		namespace = "default"
	}
	return kuberntetes.NewSecret(c.getClient(c.clusterName).ClientSet, namespace)
}

func (c *cluster) Namespaces() kuberntetes.INamespace {
	return kuberntetes.NewNamespace(c.getClient(c.clusterName))
}

func (c *cluster) Service(namespace string) kuberntetes.IService {
	if namespace == "" {
		namespace = "default"
	}
	return kuberntetes.NewService(c.getClient(c.clusterName).ClientSet, namespace)
}

func (c *cluster) Nodes() kuberntetes.INode {
	return kuberntetes.NewNode(c.getClient(c.clusterName))
}

func (c *cluster) Deployments(namespace string) kuberntetes.IDeployment {
	if namespace == "" {
		namespace = "default"
	}
	return kuberntetes.NewDeployment(c.getClient(c.clusterName), namespace)
}

func (c *cluster) Create(ctx context.Context, config *form.ClusterReq) error {

	_, conf, err := c.factory.Cluster().GenerateClient(config.Name, config.KubeConfig)
	if err != nil || conf == nil {
		log.Logger.Error(err)
		return err
	}

	util.WithErrorLog(c.factory.Cluster().Create(conf))
	return nil
}

func (c *cluster) Delete(ctx context.Context, id uint64) error {
	util.WithErrorLog(c.factory.Cluster().Delete(id))
	return nil
}

func (c *cluster) Update(ctx context.Context, id uint64, config *cloud.Config) error {
	util.WithErrorLog(c.factory.Cluster().Update(id, config))
	return nil
}

func (c *cluster) Get(ctx context.Context, id uint64) (*cloud.Config, error) {
	config, err := c.factory.Cluster().Get(id)
	util.WithErrorLog(err)
	return config, nil
}

func (c *cluster) List(ctx context.Context) (*[]cloud.Config, error) {
	list, err := c.factory.Cluster().List()
	util.WithErrorLog(err)
	return list, nil
}

func (c *cluster) CheckHealth(ctx context.Context) bool {
	config, err := c.factory.Cluster().GetByName(c.clusterName)
	if err != nil || config == nil {
		return false
	}

	if !config.Status {
		return false
	}

	return true
}

func (c *cluster) ChangeStatus(id uint64, status bool) error {
	return c.factory.Cluster().ChangeStatus(id, status)
}

func (c *cluster) getClient(name string) *cloud2.Clients {
	return c.factory.Cluster().GetClient(name)
}

func NewCluster(factory services.IDbFactory, clusterName string) *cluster {
	return &cluster{factory: factory, clusterName: clusterName}
}
