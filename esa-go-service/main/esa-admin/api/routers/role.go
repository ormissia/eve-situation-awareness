package routers

import (
	"github.com/gin-gonic/gin"

	"esa-go-service/main/esa-admin/api/v1"
)

func InitRoleRouter(r *gin.RouterGroup) {
	systemGroup := r.Group("/role")
	{
		systemGroup.POST("", v1.CreateRole)
		systemGroup.GET("", v1.SearchRole)
		systemGroup.PUT("", v1.UpdateRole)
	}
}
