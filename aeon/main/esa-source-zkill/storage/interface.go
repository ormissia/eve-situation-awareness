package storage

import (
	"errors"

	"go.uber.org/zap"

	"aeon/global"
)

type Storage interface {
	Save(msg []byte)
}

const (
	KafkaStorage = "Kafka"
)

var (
	TypeErr = errors.New("unknown storage type")
)

func Factory(storageType string) (storage Storage, err error) {
	if storageType == KafkaStorage {
		return &Kafka{producer: global.ESAKafka.Producer}, nil
	} else {
		global.ESALog.Info("create storage failed", zap.Any("err", TypeErr))
		return nil, TypeErr
	}
}
