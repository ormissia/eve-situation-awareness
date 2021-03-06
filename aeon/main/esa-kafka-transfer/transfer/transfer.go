package transfer

import (
	"time"

	"github.com/Shopify/sarama"
	"go.uber.org/zap"

	"aeon/global"
)

type ZkillConsumer struct {
	ready chan bool
}

func (z *ZkillConsumer) SetReady(bc *chan bool) {
	z.ready = *bc
}

func (z *ZkillConsumer) GetReady() *chan bool {
	return &z.ready
}

// Setup is run at the beginning of a new session, before ConsumeClaim
func (z *ZkillConsumer) Setup(sarama.ConsumerGroupSession) error {
	// Mark the consumer as ready
	close(z.ready)
	return nil
}

// Cleanup is run at the end of a session, once all ConsumeClaim goroutines have exited
func (z *ZkillConsumer) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

// ConsumeClaim must start a consumer loop of ConsumerGroupClaim's Messages().
func (z *ZkillConsumer) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	// NOTE:
	// Do not move the code below to a goroutine.
	// The `ConsumeClaim` itself is called within a goroutine, see:
	// https://github.com/Shopify/sarama/blob/main/consumer_group.go#L27-L29
	for message := range claim.Messages() {
		global.ESALog.Info("Message claimed: ", zap.Any("timestamp", message.Timestamp), zap.Any("topic", message.Topic), zap.Any("value", string(message.Value)))
		sendMsg(message.Value)
		session.MarkMessage(message, "")
	}
	return nil
}

func sendMsg(msg []byte) {
	kafkaMsg := &sarama.ProducerMessage{
		Topic:     global.ESAConfig.KafkaIn.Topic,
		Value:     sarama.ByteEncoder(msg),
		Timestamp: time.Now(),
	}
	for {
		global.ESAKafka.Producer.Input() <- kafkaMsg
		select {
		case success, ok := <-global.ESAKafka.Producer.Successes():
			if ok {
				global.ESALog.Info("Kafka producer success", zap.Any("msg", success))
				return
			}
		case errors, ok := <-global.ESAKafka.Producer.Errors():
			if ok {
				global.ESALog.Error("Kafka producer err", zap.Any("err", errors))
			}
		}
	}
}
