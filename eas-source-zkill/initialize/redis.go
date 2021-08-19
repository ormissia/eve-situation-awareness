package initialize

import (
	"context"

	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"

	"eas-source-zkill/global"
)

func Redis() (rdb *redis.Client) {
	redisConfig := global.EASConfig.Redis
	client := redis.NewClient(&redis.Options{
		Addr:     redisConfig.Path,
		Password: redisConfig.Password, //no password set
		DB:       redisConfig.DB,       //use default DB
	})

	pong, err := client.Ping(context.Background()).Result()
	if err != nil {
		global.EASLog.Error("redis connect ping failed, err:", zap.Any("err", err))
		rdb = nil
	} else {
		global.EASLog.Info("redis connect ping response:", zap.String("ping", pong))
		rdb = client
	}
	return
}
