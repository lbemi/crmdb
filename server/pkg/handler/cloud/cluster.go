package cloud

import (
	"context"
	"errors"
	"github.com/lbemi/lbemi/pkg/bootstrap/log"
	"github.com/lbemi/lbemi/pkg/handler/kuberntetes"
	"github.com/lbemi/lbemi/pkg/model/cloud"
	"github.com/lbemi/lbemi/pkg/model/form"
	"github.com/lbemi/lbemi/pkg/services"
	cloud2 "github.com/lbemi/lbemi/pkg/services/cloud"
	"github.com/lbemi/lbemi/pkg/util"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
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

	GenerateClient(name, config string) (*cloud2.Clients, error)

	kuberntetes.DeploymentGetter
	kuberntetes.NodeGetter
	kuberntetes.ServiceGetter
}

type cluster struct {
	factory     services.IDbFactory
	clusterName string
}

func (c *cluster) Service(namespace string) kuberntetes.IService {
	return kuberntetes.NewService(c.getClient(c.clusterName).ClientSet, namespace)
}

func (c *cluster) Nodes() kuberntetes.INode {
	return kuberntetes.NewNode(c.getClient(c.clusterName).ClientSet)
}

func (c *cluster) Deployments(namespace string) kuberntetes.IDeployment {
	return kuberntetes.NewDeployment(c.getClient(c.clusterName).ClientSet, namespace)
}

func (c *cluster) Create(ctx context.Context, config *form.ClusterReq) error {

	clients, err := c.factory.Cluster().GenerateClient(config.Name, config.KubeConfig)
	if err != nil {
		log.Logger.Error(err)
		return err
	}
	var conf cloud.Config
	list, err := clients.ClientSet.CoreV1().Nodes().List(ctx, metav1.ListOptions{})
	if err != nil {
		log.Logger.Error(err)
		return errors.New("server is not health")
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

	conf.Name = config.Name
	conf.KubeConfig = config.KubeConfig
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

func (c *cluster) GenerateClient(name, config string) (*cloud2.Clients, error) {
	clients, err := c.factory.Cluster().GenerateClient(name, config)
	if err != nil {
		log.Logger.Error(err)
		return nil, err
	}
	return clients, nil
}

func NewCluster(factory services.IDbFactory, clusterName string) *cluster {
	return &cluster{factory: factory, clusterName: clusterName}
}
