package api

import (
	"github.com/gin-gonic/gin"

	"eas-go-service/eas-admin/api/routers"
	"eas-go-service/eas-admin/middleware"
)

func Routers() (r *gin.Engine) {
	r = gin.Default()
	baseGroup := r.Group("admin")
	baseGroup.Use(middleware.Cors())

	publicGroup := baseGroup.Group("")
	{
		routers.InitSystemRouter(publicGroup)
		routers.InitUserBaseRouter(publicGroup)
	}

	privateGroup := baseGroup.Group("")
	privateGroup.Use(middleware.JWT()).Use(middleware.Casbin())
	{
		routers.InitUserRouter(privateGroup)
	}

	return
}
