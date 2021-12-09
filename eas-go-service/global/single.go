package global

import (
	"embed"

	"github.com/Shopify/sarama"
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"golang.org/x/sync/singleflight"
	"gorm.io/gorm"

	"eas-go-service/config"
)

var (
	EASConfig config.System
	EASViper  *viper.Viper
	EASLog    *zap.Logger
	EASMySql  *gorm.DB
	EASRedis  *redis.Client
	EASKafka  *KafkaClient
	// TODO
	EASElasticSearch string

	EASStaticFile embed.FS

	EASConcurrencyControl = &singleflight.Group{}
)

type KafkaClient struct {
	Producer sarama.SyncProducer
	Consumer sarama.ConsumerGroup
}
