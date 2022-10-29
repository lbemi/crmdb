package controller

import (
	r "github.com/go-redis/redis"
	"github.com/lbemi/lbemi/pkg/controller/policy"
	"github.com/lbemi/lbemi/pkg/controller/redis"
	"github.com/lbemi/lbemi/pkg/controller/sys"
	"github.com/lbemi/lbemi/pkg/model/config"
	"github.com/lbemi/lbemi/pkg/services"
)

type IController interface {
	redis.RedisGeeter
	policy.PolicyGetter
	sys.UserGetter
	sys.RoleGetter
	sys.MenuGetter
}

type Controller struct {
	Config    *config.Config
	DbFactory services.IDbFactory
	RedisCli  *r.Client
}

func NewController(factory services.IDbFactory, redisCli *r.Client) IController {
	return &Controller{
		DbFactory: factory,
		RedisCli:  redisCli,
	}
}
func (c *Controller) Redis() redis.IRedis {
	return redis.NewRedis(c.RedisCli)
}
func (c *Controller) Policy() policy.PolicyInterface {
	return policy.NewPolicy(c.DbFactory)
}
func (c *Controller) User() sys.IUSer {
	return sys.NewUser(c.DbFactory)
}
func (c *Controller) Role() sys.IRole {
	return sys.NewRole(c.DbFactory)
}
func (c *Controller) Menu() sys.IMenu {
	return sys.NewMenu(c.DbFactory)
}
