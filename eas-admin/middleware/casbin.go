package middleware

import (
	"fmt"
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"sync"

	"admin/global"
	"admin/model/request"
	"admin/model/response"
	"admin/utils"
)

var (
	syncedEnforcer *casbin.SyncedEnforcer
	once           sync.Once
)

// Casbin 鉴权
func Casbin() gin.HandlerFunc {
	return func(c *gin.Context) {
		claims, exists := c.Get("claims")
		if !exists {
			response.ErrorResponse(c, utils.ErrPermissionDenied)
			c.Abort()
			return
		}
		fmt.Println(claims)
		waitUse := claims.(*request.CustomClaims)
		// 获取请求的URI
		obj := c.Request.URL.RequestURI()
		// 获取请求方法
		act := c.Request.Method
		// 获取用户的角色
		sub := waitUse.AuthorityId

		e := initCasbinEnforcer()
		// 判断策略中是否存在
		success, _ := e.Enforce(sub, obj, act)
		if global.EASConfig.EASEnv == "dev" || success {
			c.Next()
		} else {
			response.ErrorResponse(c, utils.ErrPermissionDenied)
			c.Abort()
			return
		}
	}
}

func initCasbinEnforcer() *casbin.SyncedEnforcer {
	once.Do(func() {
		a, err := gormadapter.NewAdapterByDB(global.EASMySql)
		if err != nil {
			global.EASLog.Error("Casbin load data to mysql failed:", zap.String("err:", err.Error()))
			return
		}
		syncedEnforcer, err = casbin.NewSyncedEnforcer(global.EASConfig.Casbin.ModelPath, a)
		if err != nil {
			global.EASLog.Error("Casbin syncedEnforcer failed:", zap.String("err:", err.Error()))
			return
		}
		//TODO
		//syncedEnforcer.AddFunction("ParamsMatch", casbinService.ParamsMatchFunc)
	})
	_ = syncedEnforcer.LoadPolicy()
	return syncedEnforcer
}
