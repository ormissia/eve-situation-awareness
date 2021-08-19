package main

import (
	"eas-source-zkill/global"
	"eas-source-zkill/initialize"
)

func main() {
	global.EASViper = initialize.Viper()
	global.EASLog = initialize.Zap()
	global.EASMySql = initialize.Mysql()
	global.EASRedis = initialize.Redis()
}
