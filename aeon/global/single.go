package global

import (
	"embed"

	"github.com/Shopify/sarama"
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"golang.org/x/sync/singleflight"
	"gorm.io/gorm"

	"aeon/config"
)

var (
	ESAConfig        config.System
	ESAViper         *viper.Viper
	ESALog           *zap.Logger
	ESAMySqlESA      *gorm.DB
	ESAMySqlBasic    *gorm.DB
	ESARedis         *redis.Client
	ESAKafka         *KafkaClient
	ESAElasticSearch string // TODO

	ESAStaticFile embed.FS

	ESAConcurrencyControl = &singleflight.Group{}
)

type KafkaClient struct {
	Producer sarama.AsyncProducer
}
