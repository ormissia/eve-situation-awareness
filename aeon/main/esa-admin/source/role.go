package source

import (
	"time"

	"gorm.io/gorm"

	"aeon/global"
	model2 "aeon/main/esa-admin/model"
)

var Role = new(role)

type role struct{}

var roles = []model2.Role{
	{ESABase: model2.ESABase{ID: 2115581995, CreateTime: time.Now(), UpdateTime: time.Now()}, Rolename: "超级管理员"},
}

func (a *role) Init() error {
	return global.ESAMySql.Transaction(func(db *gorm.DB) error {
		if db.Where("id IN ?", []int{1}).Find(&[]model2.Role{}).RowsAffected == 1 {
			global.ESALog.Warn("表的初始数据已存在")
			return nil
		}
		if err := db.Create(&roles).Error; err != nil {
			// 遇到错误时回滚事务
			return err
		}
		global.ESALog.Info("表初始化成功")
		return nil
	})
}
