package storage

import (
	"errors"

	"go.uber.org/zap"

	"eas-go-service/global"
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
		return &Kafka{producer: global.EASKafka.Producer}, nil
	} else {
		global.EASLog.Info("create storage failed", zap.Any("err", TypeErr))
		return nil, TypeErr
	}
}
