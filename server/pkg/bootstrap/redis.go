package bootstrap

import (
	"fmt"
	"github.com/lbemi/lbemi/pkg/config"
	"github.com/lbemi/lbemi/pkg/global"

	redis "github.com/go-redis/redis"
)

func InitializeRedis(config config.Redis) *redis.Client {
	addr := fmt.Sprintf("%s:%d", config.Addr, config.Port)
	redisCli := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: config.Password,
		DB:       0,
	})
	_, err := redisCli.Ping().Result()
	if err != nil {
		global.Logger.Error("连接Redis失败", err)
		//log.Logger.Error("连接Redis失败", err)
	}
	return redisCli
}
