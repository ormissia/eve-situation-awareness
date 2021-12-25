package api

import (
	"github.com/gin-gonic/gin"

	"aeon/main/esa-web/api/web"
	"aeon/main/esa-web/middleware"
)

func Routers() (r *gin.Engine) {
	r = gin.Default()

	webApi := r.Group("/web")
	webApi.Use(middleware.Cors())
	{
		webApi.GET("/solar_system_kill", web.SearchSolarSystemKill)
	}

	return r
}
