package redis

import (
	"github.com/go-redis/redis"
	"github.com/lbemi/lbemi/pkg/ginx"
	"time"
)

type RedisGeeter interface {
	Redis() IRedis
}

type Redis struct {
	cli *redis.Client
}

func NewRedis(cli *redis.Client) IRedis {
	return &Redis{
		cli: cli,
	}
}

type IRedis interface {
	Get(key string) *redis.StringCmd
	Set(key string, value interface{}, expiration time.Duration)
	SetNX(key string, value interface{}, expiration time.Duration)
}

func (r *Redis) Get(key string) *redis.StringCmd {
	return r.cli.Get(key)
}

func (r *Redis) Set(key string, value interface{}, expiration time.Duration) {
	ginx.ErrIsNil(r.cli.Set(key, value, expiration).Err(), "set redis failed")
}

func (r *Redis) SetNX(key string, value interface{}, expiration time.Duration) {
	ginx.ErrIsNil(r.cli.SetNX(key, value, expiration).Err(), "set expiration key failed")
}
