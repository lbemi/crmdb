package handler

import (
	"github.com/casbin/casbin/v2"
	r "github.com/go-redis/redis"
	"github.com/lbemi/lbemi/pkg/common/store"
	"github.com/lbemi/lbemi/pkg/handler/asset"
	"github.com/lbemi/lbemi/pkg/handler/cloud"
	"github.com/lbemi/lbemi/pkg/handler/logsys"
	"github.com/lbemi/lbemi/pkg/handler/policy"
	"github.com/lbemi/lbemi/pkg/handler/redis"
	"github.com/lbemi/lbemi/pkg/handler/sys"
	"github.com/lbemi/lbemi/pkg/model/config"
	"gorm.io/gorm"
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
	Config      *config.Config
	RedisCli    *r.Client
	DB          *gorm.DB
	Enforcer    *casbin.SyncedEnforcer
	ClientStore *store.ClientMap
}

func NewHandler(redisCli *r.Client, db *gorm.DB, enforcer *casbin.SyncedEnforcer, clientStore *store.ClientMap) Getter {
	return &Handler{
		RedisCli:    redisCli,
		DB:          db,
		Enforcer:    enforcer,
		ClientStore: clientStore,
	}
}

func (c *Handler) Redis() redis.IRedis {
	return redis.NewRedis(c.RedisCli)
}

func (c *Handler) Policy() policy.IPolicy {
	return policy.NewPolicy(c.DB, c.Enforcer)
}

func (c *Handler) User() sys.IUSer {
	return sys.NewUser(c.DB, c.Policy(), c.Menu(), c.Login())
}

func (c *Handler) Role() sys.IRole {
	return sys.NewRole(c.DB, c.Policy(), c.Menu())
}

func (c *Handler) Menu() sys.IMenu {
	return sys.NewMenu(c.DB, c.Policy())
}

func (c *Handler) Host() asset.IHost {
	return asset.NewHost(c.DB, c.ResourceBindAccount(), c.Account())
}

func (c *Handler) Account() asset.IAccount {
	return asset.NewAccount(c.DB)
}

func (c *Handler) Group() asset.IGroup {
	return asset.NewGroup(c.DB)
}

func (c *Handler) Terminal() asset.ITerminal {
	return asset.NewTerminal(c.Host(), c.Account())
}

func (c *Handler) Ws() asset.IWs {
	return asset.NewWs()
}

func (c *Handler) Cluster(clusterName string) cloud.ICluster {
	return cloud.NewCluster(c.DB, c.ClientStore, clusterName)
}

func (c *Handler) Login() logsys.ILoginLog {
	return logsys.NewLogin(c.DB)
}

func (c *Handler) Operator() logsys.IOperatorLog {
	return logsys.NewOperator(c.DB)
}

func (c *Handler) ResourceBindAccount() asset.IResourceBindAccount {
	return asset.NewResourceBindAccount(c.DB)
}
