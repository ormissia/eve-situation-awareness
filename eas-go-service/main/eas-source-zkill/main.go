package main

import (
	"embed"
	"time"

	"go.uber.org/zap"

	"eas-go-service/global"
	"eas-go-service/initialize"
	"eas-go-service/main/eas-source-zkill/core"
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

	run()
}

func run() {
	zKillClient := core.NewClient(1000)

	if zKillClient != nil {
		go func() {
			defer func() {
				if err := recover(); err != nil {
					// 处理panic
					global.EASLog.Error("WebSocket panic", zap.Any("message", err))
					// 出现错误后5S进行重连
					time.Sleep(time.Second * 5)
					run()
				}
			}()
			zKillClient.Connect()
		}()

		zKillClient.Write <- `{"action":"sub","channel":"killstream"}`

		for msgByte := range zKillClient.Read {
			// TODO 收到数据了，写到Kafka里
			global.EASLog.Info("收到了：", zap.String("message", string(msgByte)))
		}
	}
}
