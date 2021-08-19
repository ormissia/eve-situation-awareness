package source

import (
	"time"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"

	model2 "eas-go-service/eas-admin/model"
	"eas-go-service/global"
)

var User = new(user)

type user struct{}

var admins = []model2.User{
	{EASBase: model2.EASBase{ID: 1, CreateTime: time.Now(), UpdateTime: time.Now()}, UUID: uuid.NewV4(), Username: "admin", Password: "e10adc3949ba59abbe56e057f20f883e", NickName: "超级管理员", HeaderImg: "https://imageserver.eveonline.com/Character/2115581995_1024.jpg", AuthorityId: "2115581995"},
}

func (a *user) Init() error {
	return global.EASMySql.Transaction(func(db *gorm.DB) error {
		if db.Where("id IN ?", []int{1, 2}).Find(&[]model2.User{}).RowsAffected == 2 {
			global.EASLog.Warn("表的初始数据已存在")
			return nil
		}
		if err := db.Create(&admins).Error; err != nil {
			// 遇到错误时回滚事务
			return err
		}
		global.EASLog.Info("表初始化成功")
		return nil
	})
}
