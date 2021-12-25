package v1

import (
	"github.com/gin-gonic/gin"

	"aeon/main/esa-web/model"
)

func Info(c *gin.Context) {
	model.SuccessResponse(c, "esa-web")
}
