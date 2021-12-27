package main

import (
	"embed"

	"aeon/global"
	"aeon/initialize"
	"aeon/main/esa-kafka-transfer/transfer"
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
	global.ESAKafka.Producer = initialize.KafkaProducer()
	// global.ESAMySql = initialize.Mysql()
	// global.ESARedis = initialize.Redis()
	consumer := new(transfer.ZkillConsumer)
	bc := make(chan bool)
	consumer.SetReady(&bc)
	initialize.KafkaConsume(consumer)
}
