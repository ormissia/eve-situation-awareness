package initialize

import (
	"strings"

	"github.com/Shopify/sarama"
	"go.uber.org/zap"

	"eas-go-service/global"
)

func KafkaConsumer() (consumer sarama.ConsumerGroup) {
	config := sarama.NewConfig()
	config.Version = sarama.V3_0_0_0

	consumer, err := sarama.NewConsumerGroup(strings.Split(global.EASConfig.Kafka.Path, ","), global.EASConfig.Kafka.Group, config)
	if err != nil {
		global.EASLog.Error("Kafka consumer init err", zap.String("err", err.Error()))
	}

	return
}

func KafkaProducer() (producer sarama.SyncProducer) {
	config := sarama.NewConfig()
	config.Version = sarama.V3_0_0_0
	// 等待服务器所有副本都保存成功后的响应
	config.Producer.RequiredAcks = sarama.WaitForAll
	// 随机向partition发送消息
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	// 是否等待成功和失败后的响应,只有上面的RequireAcks设置不是NoReponse这里才有用.
	config.Producer.Return.Successes = true
	config.Producer.Return.Errors = true
	// 设置用户名密码
	config.Net.SASL.Enable = true
	config.Net.SASL.User = global.EASConfig.Kafka.Username
	config.Net.SASL.Password = global.EASConfig.Kafka.Password

	producer, err := sarama.NewSyncProducer(strings.Split(global.EASConfig.Kafka.Path, ","), config)
	if err != nil {
		global.EASLog.Error("Kafka producer init err", zap.String("err", err.Error()))
	}
	return producer
}
