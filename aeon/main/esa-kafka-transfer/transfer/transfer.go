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
		go func(message *sarama.ConsumerMessage) {
			global.ESALog.Info("Message claimed: ", zap.Any("timestamp", message.Timestamp), zap.Any("topic", message.Topic), zap.Any("value", string(message.Value)))
			kafkaMsg := &sarama.ProducerMessage{
				Topic:     global.ESAConfig.KafkaIn.Topic,
				Value:     sarama.ByteEncoder(message.Value),
				Timestamp: time.Now(),
			}
			global.ESAKafka.Producer.Input() <- kafkaMsg

			session.MarkMessage(message, "")
		}(message)
	}
	return nil
}
