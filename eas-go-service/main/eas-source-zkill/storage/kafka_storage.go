package storage

import (
	"github.com/Shopify/sarama"
	"go.uber.org/zap"

	"eas-go-service/global"
)

type Kafka struct {
	producer sarama.SyncProducer
}

func (k *Kafka) Save(msg []byte) {
	global.EASLog.Info("kafka producer receive msg", zap.Any("msg", msg))
	kafkaMsg := sarama.ProducerMessage{
		Topic: global.EASConfig.Kafka.Topic,
		Value: sarama.ByteEncoder(msg),
	}
	message, offset, err := global.EASKafka.Producer.SendMessage(&kafkaMsg)
	if err != nil {
		global.EASLog.Error("kafka producer err", zap.Any("err", err))
	}
	global.EASLog.Info("kafka produce msg success", zap.Any("message", message), zap.Any("offset", offset))
}
