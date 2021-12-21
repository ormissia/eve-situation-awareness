package source

import (
	"gorm.io/gorm"

	"aeon/global"
	"aeon/main/esa-admin/model"
)

var UserRole = new(userRole)

type userRole struct{}

var userRoles = []model.UserRole{
	{1, 2115581995},
}

func (a *userRole) Init() error {
	return global.ESAMySql.Transaction(func(db *gorm.DB) error {
		if err := db.Create(&userRoles).Error; err != nil {
			// 遇到错误时回滚事务
			return err
		}
		global.ESALog.Info("表初始化成功")
		return nil
	})
}
