package main

import (
	"fmt"

	"eas-source-zkill/core"
	"eas-source-zkill/global"
	"eas-source-zkill/initialize"
)

func main() {
	global.EASViper = initialize.Viper()
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
