package storage

import (
	"go.uber.org/zap"

	"eas-go-service/global"
)

// Kafka TODO
type Kafka struct {
	Name string
}

func (k *Kafka) Save(msg interface{}) {
	global.EASLog.Info("kafka receive msg", zap.Any("msg", msg))
}



