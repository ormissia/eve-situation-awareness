package initialize

import (
	"strings"
	"sync"

	"github.com/Shopify/sarama"
	"go.uber.org/zap"

	"eas-go-service/global"
)

var once sync.Once

func init() {
	once.Do(func() {
		if global.EASKafka == nil {
			global.EASKafka = new(global.KafkaClient)
		}
	})
}

func KafkaConsumer() (consumer sarama.ConsumerGroup) {
	config := sarama.NewConfig()
	config.Version = sarama.V3_0_0_0

	consumer, err := sarama.NewConsumerGroup(strings.Split(global.EASConfig.Kafka.Path, ","), global.EASConfig.Kafka.Group, config)
	if err != nil {
		global.EASLog.Error("Kafka consumer init err", zap.String("err", err.Error()))
	}

	return
}

func KafkaProducer() (producer sarama.AsyncProducer) {
	defer func() {
		if err := recover(); err != nil {
			global.EASLog.Error("Kafka producer init err", zap.Any("err", err))
		}
	}()

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
	if global.EASConfig.Kafka.Username != "" && global.EASConfig.Kafka.Password != "" {
		config.Net.SASL.Enable = true
		config.Net.SASL.User = global.EASConfig.Kafka.Username
		config.Net.SASL.Password = global.EASConfig.Kafka.Password
	}

	bootstrapServers := strings.Split(global.EASConfig.Kafka.Path, ",")
	producer, err := sarama.NewAsyncProducer(bootstrapServers, config)
	if err != nil {
		global.EASLog.Error("Kafka producer init err", zap.String("err", err.Error()))
		return nil
	}

	go func(producer sarama.AsyncProducer) {
		global.EASLog.Info("start monitor kafka producer status...")
		for {
			select {
			case success, ok := <-producer.Successes():
				if ok {
					global.EASLog.Info("Kafka producer success", zap.Any("msg", success))
				}
			case errors, ok := <-producer.Errors():
				if ok {
					global.EASLog.Error("Kafka producer err", zap.Any("err", errors))
				}
			}
		}
	}(producer)

	return producer
}
