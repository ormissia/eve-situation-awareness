package main

import (
	"embed"
	"strconv"

	"aeon/global"
	"aeon/initialize"
	"aeon/main/esa-web/api"
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
	global.ESAMySql = initialize.Mysql()
	// global.ESARedis = initialize.Redis()
	r := api.Routers()

	// TODO 初始化缓存

	if global.ESAMySql != nil {
		// 程序结束前关闭数据库链接
		db, _ := global.ESAMySql.DB()
		defer func() {
			_ = db.Close()
		}()
	}
	err := r.Run(":" + strconv.Itoa(global.ESAConfig.Server.Port))
	if err != nil {
		panic(err)
	}
}
