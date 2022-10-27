package controller

import (
	r "github.com/go-redis/redis"
	"github.com/lbemi/lbemi/pkg/controller/redis"
	"github.com/lbemi/lbemi/pkg/controller/sys/user"
	"github.com/lbemi/lbemi/pkg/factory"
	"github.com/lbemi/lbemi/pkg/model/configs"
)

type IController interface {
	user.UserGetter
	redis.RedisGeeter
}

type Controller struct {
	Config    *configs.Config
	DbFactory factory.IDbFactory
	RedisCli  *r.Client
}

func NewController(factory factory.IDbFactory, redisCli *r.Client) IController {
	return &Controller{
		DbFactory: factory,
		RedisCli:  redisCli,
	}
}

func (c *Controller) User() user.IUSer {
	return user.NewUser(c.DbFactory)
}

func (c *Controller) Redis() redis.IRedis {
	return redis.NewRedis(c.RedisCli)
}
