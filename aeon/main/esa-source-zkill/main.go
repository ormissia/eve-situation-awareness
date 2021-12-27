package main

import (
	"embed"
	"log"
	"os"
	"strings"

	"go.uber.org/zap"

	"aeon/global"
	"aeon/initialize"
	"aeon/main/esa-source-zkill/core"
	"aeon/main/esa-source-zkill/storage"
	"aeon/utils"
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
	global.ESAStaticFile = staticFile
}

var (
	sourceClient  core.Source
	storageClient []storage.Storage
)

// os.Args
func main() {
	global.ESAViper = initialize.Viper(configFileName)
	global.ESALog = initialize.Zap()
	global.ESAKafka.Producer = initialize.KafkaProducer()
	// global.ESAMySql = initialize.Mysql()
	// global.ESARedis = initialize.Redis()
	defer func() {
		if err := global.ESAKafka.Consumer.Close(); err != nil {
			global.ESALog.Error("Kafka consumer close err", zap.String("err", err.Error()))
		}
	}()

	storageClient = make([]storage.Storage, 0)
	if len(os.Args) == 1 {
		global.ESALog.Info("Default Select RedisQ Client and Kafka Storage")
		sourceClient = core.NewRedisQClient(clientName)
		factory, err := storage.Factory(storage.KafkaStorage)
		if err != nil {
			global.ESALog.Error("factory create kafka err", zap.Any("err", err))
		}
		storageClient = append(storageClient, factory)
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
			global.ESALog.Error("Client error", zap.Any("err", err))
		}
	}()

	client.Listening(listeningFunc)
}

var listeningFunc = func(msg []byte) {
	// 将msg发送到所有客户端
	for _, sc := range storageClient {
		go func(s storage.Storage, msg []byte) {
			s.Save(msg)
		}(sc, msg)
	}
}
