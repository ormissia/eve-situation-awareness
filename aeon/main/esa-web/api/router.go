package api

import (
	"github.com/gin-gonic/gin"

	"aeon/main/esa-web/api/v1"
	"aeon/main/esa-web/api/web"
	"aeon/main/esa-web/middleware"
)

func Routers() (r *gin.Engine) {
	r = gin.Default()

	r.GET("/info", v1.Info)

	// TODO 限流
	webApi := r.Group("/web")
	webApi.Use(middleware.Cors())

	eveBasicInfo := webApi.Group("/basic")
	{
		eveBasicInfo.GET("/solar_system_fuzzy", web.SolarSystemFuzzySearch)
	}

	analysisApi := webApi.Group("/analysis")
	{
		analysisApi.GET("/solar_system_kill", web.SearchSolarSystemKill)
		analysisApi.GET("/solar_system_kill_order", web.SearchSolarSystemKillOrder)

		analysisApi.GET("/character_kill", web.SearchCharacterKill)
	}

	return r
}
