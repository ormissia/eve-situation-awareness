package service

import (
	"sync"

	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"go.uber.org/zap"

	"esa-go-service/global"
)

var (
	syncedEnforcer *casbin.SyncedEnforcer
	once           sync.Once
)

func InitCasbinEnforcer() *casbin.SyncedEnforcer {
	once.Do(func() {
		a, err := gormadapter.NewAdapterByDB(global.ESAMySql)
		if err != nil {
			global.ESALog.Error("Casbin load data to mysql failed:", zap.String("err", err.Error()))
			return
		}
		syncedEnforcer, err = casbin.NewSyncedEnforcer(global.ESAConfig.Casbin.ModelPath, a)
		if err != nil {
			global.ESALog.Error("Casbin syncedEnforcer failed:", zap.String("err", err.Error()))
			return
		}
		// TODO
		// syncedEnforcer.AddFunction("ParamsMatch", casbinService.ParamsMatchFunc)
	})
	_ = syncedEnforcer.LoadPolicy()
	return syncedEnforcer
}
