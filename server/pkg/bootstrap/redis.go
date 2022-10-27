package bootstrap

import (
	"fmt"
	redis "github.com/go-redis/redis"
	"github.com/lbemi/lbemi/pkg/model/configs"
)

func InitializeRedis(config configs.Redis) *redis.Client {
	addr := fmt.Sprintf("%s:%d", config.Addr, config.Port)
	redisCli := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: config.Password,
		DB:       0,
	})
	_, err := redisCli.Ping().Result()
	if err != nil {
		fmt.Println("连接Redis失败", err)
		//log.Logger.Error("连接Redis失败", err)
	}
	return redisCli
}
