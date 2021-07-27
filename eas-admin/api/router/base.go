package router

import (
	"github.com/gin-gonic/gin"

	v1 "admin/api/v1"
)

func InitBaseRouter(Router *gin.RouterGroup) {
	baseRouter := Router.Group("base")
	{
		baseRouter.POST("login", v1.Login)
	}
}
