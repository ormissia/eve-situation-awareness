package middleware

import (
	"strings"

	"github.com/gin-gonic/gin"

	"esa-go-service/global"
	"esa-go-service/main/esa-admin/model/request"
	"esa-go-service/main/esa-admin/model/response"
	"esa-go-service/main/esa-admin/service"
	"esa-go-service/utils"
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
		waitUse := claims.(*request.CustomClaims)
		// 获取请求的URI并去除Query参数
		reqURL := c.Request.URL.RequestURI()
		obj := ""
		if urls := strings.Split(reqURL, "?"); len(urls) > 0 {
			obj = urls[0]
		}
		// 获取请求方法
		act := c.Request.Method
		// 获取用户的角色
		sub := waitUse.AuthorityId

		e := service.InitCasbinEnforcer()
		// 判断策略中是否存在
		success, _ := e.Enforce(sub, obj, act)
		if global.ESAConfig.ESAEnv == "dev" || success {
			c.Next()
		} else {
			response.ErrorResponse(c, utils.ErrPermissionDenied)
			c.Abort()
			return
		}
	}
}
