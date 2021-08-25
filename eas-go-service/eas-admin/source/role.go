package source

import (
	"time"

	"gorm.io/gorm"

	"eas-go-service/eas-admin/model"
	"eas-go-service/global"
)

var Role = new(role)

type role struct{}

var roles = []model.Role{
	{EASBase: model.EASBase{ID: 2115581995, CreateTime: time.Now(), UpdateTime: time.Now()}, Rolename: "超级管理员"},
}

func (a *role) Init() error {
	return global.EASMySql.Transaction(func(db *gorm.DB) error {
		if db.Where("id IN ?", []int{1}).Find(&[]model.Role{}).RowsAffected == 1 {
			global.EASLog.Warn("表的初始数据已存在")
			return nil
		}
		if err := db.Create(&roles).Error; err != nil {
			// 遇到错误时回滚事务
			return err
		}
		global.EASLog.Info("表初始化成功")
		return nil
	})
}
