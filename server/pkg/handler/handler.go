package handler

import (
	r "github.com/go-redis/redis"
	"github.com/lbemi/lbemi/pkg/handler/asset"
	"github.com/lbemi/lbemi/pkg/handler/cloud"
	"github.com/lbemi/lbemi/pkg/handler/logsys"
	"github.com/lbemi/lbemi/pkg/handler/policy"
	"github.com/lbemi/lbemi/pkg/handler/redis"
	"github.com/lbemi/lbemi/pkg/handler/sys"
	"github.com/lbemi/lbemi/pkg/model/config"
	"github.com/lbemi/lbemi/pkg/services"
)

type Getter interface {
	redis.RedisGetter
	policy.PolicyGetter
	sys.UserGetter
	sys.RoleGetter
	sys.MenuGetter
	asset.HostGetter
	asset.AccountGetter
	asset.TerminalGetter
	asset.GroupGetter
	asset.WsGetter
	cloud.ClusterGetter
	logsys.LoginLogGetter
	logsys.OperatorLogGetter
	asset.ResourceAccountGetter
}

type Handler struct {
	Config    *config.Config
	DbFactory services.Interface
	RedisCli  *r.Client
}

func NewHandler(factory services.Interface, redisCli *r.Client) Getter {
	return &Handler{
		DbFactory: factory,
		RedisCli:  redisCli,
	}
}

func (c *Handler) Redis() redis.IRedis {
	return redis.NewRedis(c.RedisCli)
}

func (c *Handler) Policy() policy.PolicyInterface {
	return policy.NewPolicy(c.DbFactory)
}

func (c *Handler) User() sys.IUSer {
	return sys.NewUser(c.DbFactory)
}

func (c *Handler) Role() sys.IRole {
	return sys.NewRole(c.DbFactory)
}

func (c *Handler) Menu() sys.IMenu {
	return sys.NewMenu(c.DbFactory)
}

func (c *Handler) Host() asset.IHost {
	return asset.NewHost(c.DbFactory)
}

func (c *Handler) Account() asset.IAccount {
	return asset.NewAccount(c.DbFactory)
}

func (c *Handler) Group() asset.IGroup {
	return asset.NewGroup(c.DbFactory)
}

func (c *Handler) Terminal() asset.ITerminal {
	return asset.NewTerminal(c.DbFactory)
}

func (c *Handler) Ws() asset.IWs {
	return asset.NewWs()
}

func (c *Handler) Cluster(clusterName string) cloud.ICluster {
	return cloud.NewCluster(c.DbFactory, clusterName)
}

func (c *Handler) Login() logsys.LoginLogImp {
	return logsys.NewLogin(c.DbFactory)
}

func (c *Handler) Operator() logsys.OperatorLogImp {
	return logsys.NewOperator(c.DbFactory)
}

func (c *Handler) ResourceBindAccount() asset.IResourceBindAccount {
	return asset.NewResourceBindAccount(c.DbFactory)
}
