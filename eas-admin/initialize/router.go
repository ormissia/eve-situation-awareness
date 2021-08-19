package initialize

import (
	"github.com/gin-gonic/gin"

	router "eas-admin/api/router"
	"eas-admin/middleware"
)

func Routers() (r *gin.Engine) {
	r = gin.Default()
	baseGroup := r.Group("admin")
	baseGroup.Use(middleware.Cors())

	publicGroup := baseGroup.Group("")
	{
		router.InitSystemRouter(publicGroup)
		router.InitUserBaseRouter(publicGroup)
	}

	privateGroup := baseGroup.Group("")
	privateGroup.Use(middleware.JWT()).Use(middleware.Casbin())
	{
		router.InitUserRouter(privateGroup)
	}

	return
}
