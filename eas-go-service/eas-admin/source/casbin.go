package source

import (
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"gorm.io/gorm"

	"eas-go-service/global"
)

var Casbin = new(casbin)

type casbin struct{}

/*
g, alice, user_admin_group

g2, data1, system_request
g2, data2, system_request

p, user_admin_group, system_request, get
p, user_admin_group, system_request, post
*/

var casbins = []gormadapter.CasbinRule{
	// 角色分组
	{Ptype: "g", V0: "2115581995", V1: "user_admin_group"},
	// 资源分组
	{Ptype: "g2", V0: "/admin/user/info", V1: "user_get_request"},
	// 组-资源 操作权限
	{Ptype: "p", V0: "user_admin_group", V1: "user_get_request", V2: "GET"},
}

func (a *casbin) Init() error {
	_ = global.EASMySql.AutoMigrate(gormadapter.CasbinRule{})
	return global.EASMySql.Transaction(func(tx *gorm.DB) error {
		if tx.Find(&[]gormadapter.CasbinRule{}).RowsAffected != 0 {
			global.EASLog.Warn("表的初始数据已存在")
			return nil
		}
		if err := tx.Create(&casbins).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		global.EASLog.Info("表初始化成功")
		return nil
	})
}
