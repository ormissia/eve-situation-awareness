package main

import (
	"embed"

	"aeon/global"
	"aeon/initialize"
	"aeon/main/esa-admin/api"
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
	global.ESAMySql = initialize.Mysql()
	global.ESARedis = initialize.Redis()

	r := api.Routers()

	if global.ESAMySql != nil {
		// 程序结束前关闭数据库链接
		db, _ := global.ESAMySql.DB()
		defer func() {
			_ = db.Close()
		}()
	}
	err := r.Run(":8080")
	if err != nil {
		panic(err)
	}
}
