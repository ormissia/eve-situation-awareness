package initialize

import (
	"github.com/gin-gonic/gin"

	router "admin/api/router"
	"admin/middleware"
)

func Routers() {
	r := gin.Default()
	r.Use(middleware.Cors())

	r.Use(middleware.Cors())

	publicGroup := r.Group("/admin")
	{
		router.InitSystemRouter(publicGroup)
	}

	err := r.Run(":8080")
	if err != nil {
		panic(err)
	}
}
