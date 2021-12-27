package initialize

import (
	"strings"
	"sync"

	"github.com/Shopify/sarama"
	"go.uber.org/zap"

	"aeon/global"
)

var once sync.Once

func init() {
	once.Do(func() {
		if global.ESAKafkaIn == nil {
			global.ESAKafkaIn = new(global.KafkaClient)
		}
		if global.ESAKafkaOut == nil {
			global.ESAKafkaOut = new(global.KafkaClient)
		}
	})
}

func KafkaConsumer() (consumer sarama.ConsumerGroup) {
	config := sarama.NewConfig()
	config.Version = sarama.V3_0_0_0

	consumer, err := sarama.NewConsumerGroup(strings.Split(global.ESAConfig.KafkaOut.Path, ","), global.ESAConfig.KafkaOut.Group, config)
	if err != nil {
		global.ESALog.Error("Kafka consumer init err", zap.String("err", err.Error()))
	}

	return
}

func KafkaProducer() (producer sarama.AsyncProducer) {
	defer func() {
		if err := recover(); err != nil {
			global.ESALog.Error("Kafka producer init err", zap.Any("err", err))
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
	if global.ESAConfig.KafkaIn.Username != "" && global.ESAConfig.KafkaIn.Password != "" {
		config.Net.SASL.Enable = true
		config.Net.SASL.User = global.ESAConfig.KafkaIn.Username
		config.Net.SASL.Password = global.ESAConfig.KafkaIn.Password
	}

	bootstrapServers := strings.Split(global.ESAConfig.KafkaIn.Path, ",")
	producer, err := sarama.NewAsyncProducer(bootstrapServers, config)
	if err != nil {
		global.ESALog.Error("Kafka producer init err", zap.String("err", err.Error()))
		return nil
	}

	go func(producer sarama.AsyncProducer) {
		global.ESALog.Info("start monitor kafka producer status...")
		for {
			select {
			case success, ok := <-producer.Successes():
				if ok {
					global.ESALog.Info("Kafka producer success", zap.Any("msg", success))
				}
			case errors, ok := <-producer.Errors():
				if ok {
					global.ESALog.Error("Kafka producer err", zap.Any("err", errors))
				}
			}
		}
	}(producer)

	return producer
}
