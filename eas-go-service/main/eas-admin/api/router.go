package api

import (
	"github.com/gin-gonic/gin"

	routers2 "eas-go-service/main/eas-admin/api/routers"
	middleware2 "eas-go-service/main/eas-admin/middleware"
)

func Routers() (r *gin.Engine) {
	r = gin.Default()
	baseGroup := r.Group("admin")
	baseGroup.Use(middleware2.Cors())

	publicGroup := baseGroup.Group("")
	{
		routers2.InitSystemRouter(publicGroup)
		routers2.InitUserBaseRouter(publicGroup)
	}

	privateGroup := baseGroup.Group("")
	privateGroup.Use(middleware2.JWT(),middleware2.Casbin())
	{
		routers2.InitRoleRouter(privateGroup)
		routers2.InitUserRouter(privateGroup)
	}

	return
}
