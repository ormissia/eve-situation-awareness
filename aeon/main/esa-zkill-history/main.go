package main

import (
	"embed"

	"aeon/global"
	"aeon/initialize"
	"aeon/main/esa-zkill-history/core"
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
	global.ESAMySqlESA = initialize.Mysql(global.ESAConfig.MysqlESA)
	core.HistoryDump()
}
