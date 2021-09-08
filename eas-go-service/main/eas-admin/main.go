package main

import (
	"embed"

	"eas-go-service/global"
	"eas-go-service/initialize"
	"eas-go-service/main/eas-admin/api"
)

//go:embed config.yaml
var staticFile embed.FS

const (
	configFileName = "config.yaml"
)

func init() {
	global.EASStaticFile = staticFile
}

func main() {
	global.EASViper = initialize.Viper(configFileName)
	global.EASLog = initialize.Zap()
	global.EASMySql = initialize.Mysql()
	global.EASRedis = initialize.Redis()

	r := api.Routers()

	if global.EASMySql != nil {
		// 程序结束前关闭数据库链接
		db, _ := global.EASMySql.DB()
		defer func() {
			_ = db.Close()
		}()
	}
	err := r.Run(":8080")
	if err != nil {
		panic(err)
	}
}
