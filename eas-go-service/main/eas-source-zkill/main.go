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
	"eas-go-service/main/eas-source-zkill/storage"
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

var (
	sourceClient  core.Source
	storageClient []storage.Storage
)

// os.Args
func main() {
	global.EASViper = initialize.Viper(configFileName)
	global.EASLog = initialize.Zap()
	// global.EASMySql = initialize.Mysql()
	// global.EASRedis = initialize.Redis()

	storageClient = make([]storage.Storage, 0)
	if len(os.Args) == 1 {
		log.Print("Default Select RedisQ Client and Kafka Storage")
		sourceClient = core.NewRedisQClient(clientName)
		// TODO 默认使用Kafka
	} else {
		clientType := os.Args[1]
		if strings.EqualFold(clientType, redisQ) {
			log.Print("Select RedisQ Client")
			sourceClient = core.NewRedisQClient(clientName)
		} else if strings.EqualFold(clientType, webSocket) {
			sourceClient = core.NewWebSocketClient(clientName)
			log.Print("Select WebSocket Client")
		}
	}

	run(sourceClient)
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
	// TODO 将msg发送到所有客户端
	for _, sc := range storageClient {
		go func(s storage.Storage, msg string) {
			s.Save(msg)
			global.EASLog.Info("receive msg", zap.String("msg", msg))
		}(sc, msg)
	}
}
