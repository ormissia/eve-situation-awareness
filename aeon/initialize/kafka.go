package initialize

import (
	"context"
	"log"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"
	"time"

	"github.com/Shopify/sarama"
	"go.uber.org/zap"

	"aeon/global"
)

var once sync.Once

func init() {
	once.Do(func() {
		if global.ESAKafka == nil {
			global.ESAKafka = new(global.KafkaClient)
		}
	})
}

type Consumer interface {
	sarama.ConsumerGroupHandler
	SetReady(*chan bool)
	GetReady() *chan bool
}

func KafkaConsume(consumer Consumer) {
	config := sarama.NewConfig()
	config.Version = sarama.V3_0_0_0
	config.Consumer.Return.Errors = true
	config.Consumer.Offsets.Initial = sarama.OffsetNewest
	config.Consumer.Offsets.AutoCommit.Enable = true
	config.Consumer.Offsets.AutoCommit.Interval = time.Second
	config.Consumer.Group.Rebalance.Strategy = sarama.BalanceStrategyRoundRobin

	client, err := sarama.NewConsumerGroup(strings.Split(global.ESAConfig.KafkaOut.Path, ","), global.ESAConfig.KafkaOut.Group, config)
	if err != nil {
		global.ESALog.Error("Kafka consumer init err", zap.String("err", err.Error()))
	}

	ctx, cancel := context.WithCancel(context.Background())
	wg := &sync.WaitGroup{}
	wg.Add(1)

	go func() {
		defer wg.Done()
		for {
			// `Consume` should be called inside an infinite loop, when a
			// server-side rebalance happens, the consumer session will need to be
			// recreated to get the new claims
			if err := client.Consume(ctx, strings.Split(global.ESAConfig.KafkaOut.Topic, ","), consumer); err != nil {
				log.Panicf("Error from consumer: %v", err)
			}
			// check if context was cancelled, signaling that the consumer should stop
			if ctx.Err() != nil {
				return
			}
			cb := make(chan bool)
			consumer.SetReady(&cb)
		}
	}()

	<-*consumer.GetReady() // Await till the consumer has been set up
	global.ESALog.Info("Sarama consumer up and running!...")

	sigterm := make(chan os.Signal, 1)
	signal.Notify(sigterm, syscall.SIGINT, syscall.SIGTERM)
	select {
	case <-ctx.Done():
		global.ESALog.Info("terminating: context cancelled")
	case <-sigterm:
		global.ESALog.Info("terminating: via signal")
	}
	cancel()
	wg.Wait()
	if err = client.Close(); err != nil {
		global.ESALog.Error("Error closing client", zap.Any("err", err))
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

	return producer
}
