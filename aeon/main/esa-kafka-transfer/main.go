package main

import (
	"embed"

	"aeon/global"
	"aeon/initialize"
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
	consumer := &ZkillConsumer{
		ready: make(chan bool),
	}
	initialize.KafkaConsume(consumer)
}
