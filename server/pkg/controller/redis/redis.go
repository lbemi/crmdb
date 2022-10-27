package redis

import (
	"github.com/go-redis/redis"
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
	Set(key string, value interface{}, expiration time.Duration) error
	SetNX(key string, value interface{}, expiration time.Duration) error
}

func (r *Redis) Get(key string) *redis.StringCmd {
	return r.cli.Get(key)
}

func (r *Redis) Set(key string, value interface{}, expiration time.Duration) error {
	return r.cli.Set(key, value, expiration).Err()
}

func (r *Redis) SetNX(key string, value interface{}, expiration time.Duration) error {
	return r.cli.SetNX(key, value, expiration).Err()
}
