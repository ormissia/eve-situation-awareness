package router

import (
	"github.com/gin-gonic/gin"

	v1 "eas-admin/api/v1"
)

func InitSystemRouter(r *gin.RouterGroup) {
	systemGroup := r.Group("/system")
	{
		systemGroup.GET("/ping", v1.Ping)
		systemGroup.POST("/initDB", v1.InitMySQL)
	}
}
