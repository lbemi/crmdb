package handler

import (
	r "github.com/go-redis/redis"
	"github.com/lbemi/lbemi/pkg/handler/asset"
	"github.com/lbemi/lbemi/pkg/handler/cloud"
	"github.com/lbemi/lbemi/pkg/handler/policy"
	"github.com/lbemi/lbemi/pkg/handler/redis"
	"github.com/lbemi/lbemi/pkg/handler/sys"
	"github.com/lbemi/lbemi/pkg/model/config"
	"github.com/lbemi/lbemi/pkg/services"
)

type Getter interface {
	redis.RedisGeeter
	policy.PolicyGetter
	sys.UserGetter
	sys.RoleGetter
	sys.MenuGetter
	asset.HostGetter
	asset.TerminalGetter
	asset.WsGetter
	cloud.ClusterGetter
	//cloud.ResourceGetter
}

type Handler struct {
	Config    *config.Config
	DbFactory services.IDbFactory
	RedisCli  *r.Client
}

//func (c *Handler) Resource() cloud.IResource {
//	return cloud.NewResource(c.DbFactory)
//}

func NewHandler(factory services.IDbFactory, redisCli *r.Client) Getter {
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

func (c *Handler) Terminal() asset.ITerminal {
	return asset.NewTerminal(c.DbFactory)
}

func (c *Handler) Ws() asset.IWs {
	return asset.NewWs()
}

func (c *Handler) Cluster(clusterName string) cloud.ICluster {
	return cloud.NewCluster(c.DbFactory, clusterName)
}
