package storage

import (
	"time"

	"github.com/Shopify/sarama"
	"go.uber.org/zap"

	"eas-go-service/global"
)

type Kafka struct {
	producer sarama.AsyncProducer
}

func (k *Kafka) Save(msg []byte) {
	global.EASLog.Info("kafka producer receive msg", zap.String("msg", string(msg)))
	kafkaMsg := &sarama.ProducerMessage{
		Topic:     global.EASConfig.Kafka.Topic,
		Value:     sarama.ByteEncoder(msg),
		Timestamp: time.Now(),
	}
	global.EASKafka.Producer.Input() <- kafkaMsg
}
