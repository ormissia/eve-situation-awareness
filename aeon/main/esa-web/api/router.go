package api

import (
	"github.com/gin-gonic/gin"

	"aeon/main/esa-web/api/web"
)

func Routers() (r *gin.Engine) {
	r = gin.Default()

	webApi := r.Group("/web")
	{
		webApi.GET("/solar_system_kill", web.SearchSolarSystemKill)
	}

	return r
}
