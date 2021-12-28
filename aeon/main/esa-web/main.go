package main

import (
	"embed"
	"strconv"

	"go.uber.org/zap"

	"aeon/global"
	"aeon/initialize"
	"aeon/main/esa-web/api"
	"aeon/main/esa-web/service"
)

//go:embed config.yaml
var staticFile embed.FS

const (
	configFileName = "config.yaml"
)

func init() {
	global.ESAStaticFile = staticFile
}

// os.Args
func main() {
	global.ESAViper = initialize.Viper(configFileName)
	global.ESALog = initialize.Zap()
	// global.ESAKafka.Producer = initialize.KafkaProducer()
	global.ESAMySqlESA = initialize.Mysql(global.ESAConfig.MysqlESA)
	global.ESAMySqlBasic = initialize.Mysql(global.ESAConfig.MysqlBasic)
	global.ESARedis = initialize.Redis()

	initAllCache()

	r := api.Routers()

	if global.ESAMySqlESA != nil {
		// 程序结束前关闭数据库链接
		db, _ := global.ESAMySqlESA.DB()
		defer func() {
			_ = db.Close()
		}()
	}
	if global.ESAMySqlBasic != nil {
		// 程序结束前关闭数据库链接
		db, _ := global.ESAMySqlBasic.DB()
		defer func() {
			_ = db.Close()
		}()
	}
	err := r.Run(":" + strconv.Itoa(global.ESAConfig.Server.Port))
	if err != nil {
		panic(err)
	}
}

func initAllCache() {
	go func() {
		for {
			if err := service.InitSolarSystemRedisCache(); err != nil {
				global.ESALog.Error("InitSolarSystemRedisCache err", zap.Any("err", err))
			} else {
				global.ESALog.Info("InitSolarSystemRedisCache successful")
				break
			}
		}
	}()
}
