package model

import (
	"time"

	"gorm.io/gorm"
)

type EASBase struct {
	ID         uint           `gorm:"primarykey"` // 主键ID
	CreateTime time.Time      // 创建时间
	UpdateTime time.Time      // 更新时间
	DeleteTime gorm.DeletedAt `gorm:"index" json:"-"` // 删除时间
}
