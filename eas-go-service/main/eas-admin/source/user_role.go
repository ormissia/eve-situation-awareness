package source

import (
	"gorm.io/gorm"

	"eas-go-service/global"
	"eas-go-service/main/eas-admin/model"
)

var UserRole = new(userRole)

type userRole struct{}

var userRoles = []model.UserRole{
	{1, 2115581995},
}

func (a *userRole) Init() error {
	return global.EASMySql.Transaction(func(db *gorm.DB) error {
		if err := db.Create(&userRoles).Error; err != nil {
			// 遇到错误时回滚事务
			return err
		}
		global.EASLog.Info("表初始化成功")
		return nil
	})
}
