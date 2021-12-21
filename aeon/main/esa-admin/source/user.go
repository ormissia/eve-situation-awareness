package source

import (
	"time"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"

	"aeon/global"
	model2 "aeon/main/esa-admin/model"
)

var User = new(user)

type user struct{}

var admins = []model2.User{
	{ESABase: model2.ESABase{ID: 1, CreateTime: time.Now(), UpdateTime: time.Now()}, UUID: uuid.NewV4(), Username: "admin", Password: "e10adc3949ba59abbe56e057f20f883e", Nickname: "超级管理员", HeaderImg: "https://imageserver.eveonline.com/Character/2115581995_1024.jpg", RoleId: "2115581995"},
}

func (a *user) Init() error {
	return global.ESAMySql.Transaction(func(db *gorm.DB) error {
		if db.Where("id IN ?", []int{1}).Find(&[]model2.User{}).RowsAffected == 1 {
			global.ESALog.Warn("表的初始数据已存在")
			return nil
		}
		if err := db.Create(&admins).Error; err != nil {
			// 遇到错误时回滚事务
			return err
		}
		global.ESALog.Info("表初始化成功")
		return nil
	})
}
