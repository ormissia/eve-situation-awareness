package main

import (
	"embed"
	"fmt"

	"eas-go-service/eas-source-zkill/core"
	"eas-go-service/global"
	"eas-go-service/initialize"
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
	// global.EASMySql = initialize.Mysql()
	// global.EASRedis = initialize.Redis()

	zKillClient := core.NewClient(1000)

	if zKillClient != nil {
		go zKillClient.Connect()

		zKillClient.Write <- `{"action":"sub","channel":"killstream"}`

		for msgByte := range zKillClient.Read {
			// TODO 收到数据了，写到Kafka里
			fmt.Printf("收到了：%s", string(msgByte))
		}
	}
}
