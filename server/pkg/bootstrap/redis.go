package bootstrap

import (
	"fmt"
	redis "github.com/go-redis/redis"
	"github.com/lbemi/lbemi/pkg/global"
	"go.uber.org/zap"
)

func InitializeRedis() *redis.Client {
	addr := fmt.Sprintf("%s:%d", global.App.Config.Redis.Addr, global.App.Config.Redis.Port)
	redisCli := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: global.App.Config.Redis.Password,
		DB:       0,
	})
	_, err := redisCli.Ping().Result()
	if err != nil {
		global.App.Log.Error("连接Redis失败", zap.Error(err))
	}
	return redisCli
}
