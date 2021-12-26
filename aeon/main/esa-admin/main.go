package main

import (
	"embed"

	"aeon/global"
	"aeon/initialize"
)

//go:embed config.yaml
var staticFile embed.FS

const (
	configFileName = "config.yaml"
)

func init() {
	global.ESAStaticFile = staticFile
}

func main() {
	global.ESAViper = initialize.Viper(configFileName)
	global.ESALog = initialize.Zap()
	global.ESAMySqlESA = initialize.Mysql(global.ESAConfig.MysqlESA)
	global.ESARedis = initialize.Redis()

	if global.ESAMySqlESA != nil {
		// 程序结束前关闭数据库链接
		db, _ := global.ESAMySqlESA.DB()
		defer func() {
			_ = db.Close()
		}()
	}

	// r := api.Routers()
	// err := r.Run(":" + strconv.Itoa(global.ESAConfig.Server.Port))
	// if err != nil {
	// 	panic(err)
	// }
}
