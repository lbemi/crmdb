package bootstrap

import (
	"fmt"
	redis "github.com/go-redis/redis"
	"github.com/lbemi/lbemi/pkg/bootstrap/log"
	"github.com/lbemi/lbemi/pkg/model/configs"
	"go.uber.org/zap"
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
		log.Logger.Error("连接Redis失败", zap.Error(err))
	}
	return redisCli
}
