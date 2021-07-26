package router

import (
	"github.com/gin-gonic/gin"

	v1 "admin/api/v1"
)

func InitUserRouter(r *gin.RouterGroup) {
	userGroup := r.Group("user")
	{
		userGroup.GET("info", v1.GetUserInfo)
	}
}
