package main

import (
	"admin/global"
	"admin/initialize"
)

func main() {
	global.EASViper = initialize.Viper()
	global.EASLog = initialize.Zap()
	global.EASMySql = initialize.Mysql()
	global.EASRedis = initialize.Redis()

	r := initialize.Routers()

	if global.EASMySql != nil {
		// 程序结束前关闭数据库链接
		db, _ := global.EASMySql.DB()
		defer db.Close()
	}
	err := r.Run(":8080")
	if err != nil {
		panic(err)
	}
}
