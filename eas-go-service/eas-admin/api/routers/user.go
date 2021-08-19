package routers

import (
	"github.com/gin-gonic/gin"

	"eas-go-service/eas-admin/api/v1"
)

func InitUserBaseRouter(Router *gin.RouterGroup) {
	baseRouter := Router.Group("base")
	{
		baseRouter.POST("login", v1.Login)
	}
}

func InitUserRouter(r *gin.RouterGroup) {
	userGroup := r.Group("user")
	{
		userGroup.GET("info", v1.GetUserInfo)
	}
}
