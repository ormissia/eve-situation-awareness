package api

import (
	"github.com/gin-gonic/gin"

	v1 "aeon/main/esa-web/api/v1"
	"aeon/main/esa-web/api/web"
	"aeon/main/esa-web/middleware"
)

func Routers() (r *gin.Engine) {
	r = gin.Default()

	r.GET("/info", v1.Info)

	webApi := r.Group("/web")
	webApi.Use(middleware.Cors())
	{
		webApi.GET("/solar_system_kill", web.SearchSolarSystemKill)
	}

	return r
}
