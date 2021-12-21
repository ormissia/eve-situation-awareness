package initialize

import (
	"context"

	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"

	"aeon/global"
)

func Redis() (rdb *redis.Client) {
	redisConfig := global.ESAConfig.Redis
	client := redis.NewClient(&redis.Options{
		Addr:     redisConfig.Path,
		Password: redisConfig.Password, // no password set
		DB:       redisConfig.DB,       // use default DB
	})

	pong, err := client.Ping(context.Background()).Result()
	if err != nil {
		global.ESALog.Error("redis connect ping failed, err:", zap.Any("err", err))
		rdb = nil
	} else {
		global.ESALog.Info("redis connect ping response:", zap.String("ping", pong))
		rdb = client
	}
	return
}
