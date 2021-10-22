package main

import (
	"embed"
	"log"
	"os"
	"strings"

	"go.uber.org/zap"

	"eas-go-service/global"
	"eas-go-service/initialize"
	"eas-go-service/main/eas-source-zkill/core"
	"eas-go-service/utils"
)

//go:embed config.yaml
var staticFile embed.FS

const (
	configFileName = "config.yaml"

	clientName = utils.SystemName

	// Client Type

	redisQ    = "RedisQ"
	webSocket = "WebSocket"
)

func init() {
	global.EASStaticFile = staticFile
}

var client core.Source

// os.Args
func main() {
	global.EASViper = initialize.Viper(configFileName)
	global.EASLog = initialize.Zap()
	// global.EASMySql = initialize.Mysql()
	// global.EASRedis = initialize.Redis()

	if len(os.Args) == 1 {
		log.Print("Default Select RedisQ Client")
		client = core.NewRedisQClient(clientName)
	} else {
		clientType := os.Args[1]
		if strings.EqualFold(clientType, redisQ) {
			log.Print("Select RedisQ Client")
			client = core.NewRedisQClient(clientName)
		} else if strings.EqualFold(clientType, webSocket) {
			client = core.NewWebSocketClient(clientName)
			log.Print("Select WebSocket Client")
		}
	}

	run(client)
}

func run(client core.Source) {
	defer func() {
		if err := recover(); err != nil {
			global.EASLog.Error("Client error", zap.Any("err", err))
		}
	}()

	client.Listening(listeningFunc)
}

var listeningFunc = func(msg string) {
	// TODO 将msg发送到Kafka
	go func(msg string) {
		global.EASLog.Info("Kafka receive msg", zap.String("msg", msg))
	}(msg)
}
