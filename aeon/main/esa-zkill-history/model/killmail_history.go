package model

import (
	"time"

	"aeon/global"
)

type KillmailHash struct {
	BelongDay    int       `json:"belong_day"`
	KillmailId   int       `json:"killmail_id"`
	KillmailHash string    `json:"killmail_hash"`
	CreateTime   time.Time `json:"create_time"`
}

func (KillmailHash) TableName() string {
	return "killmail_hash"
}

func (k *KillmailHash) BatchSave(data []KillmailHash) (err error) {
	db := global.ESAMySqlESA.Model(k)

	return db.CreateInBatches(data, 1000).Error
}
