package global

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"

	"admin/config"
)

var (
	EASConfig config.System
	EASViper  *viper.Viper
	EASLog    *zap.Logger
	EASMySql  *gorm.DB
)
