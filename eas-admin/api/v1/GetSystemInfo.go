package v1

import (
	"github.com/gin-gonic/gin"

	"admin/model/response"
)

func Ping(c *gin.Context) {
	response.SuccessResponse(c, "pong")
}
