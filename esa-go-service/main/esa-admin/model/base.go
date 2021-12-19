package model

import (
	"time"

	"gorm.io/gorm"
)

type ESABase struct {
	ID         uint           `json:"id" gorm:"primary key"` // 主键ID
	CreateTime time.Time      `json:"create_time"`           // 创建时间
	UpdateTime time.Time      `json:"update_time"`           // 更新时间
	DeleteTime gorm.DeletedAt `json:"-" gorm:"index"`        // 删除时间
}
