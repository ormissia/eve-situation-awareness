package global

import (
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"golang.org/x/sync/singleflight"
	"gorm.io/gorm"

	"eas-source-zkill/config"
)

var (
	EASConfig config.System
	EASViper  *viper.Viper
	EASLog    *zap.Logger
	EASMySql  *gorm.DB
	EASRedis  *redis.Client

	EASConcurrencyControl = &singleflight.Group{}
)
