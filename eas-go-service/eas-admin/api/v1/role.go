package v1

import (
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"eas-go-service/eas-admin/model"
	"eas-go-service/eas-admin/model/response"
	"eas-go-service/global"
	"eas-go-service/utils"
)

func CreateRule(c *gin.Context) {
	param := new(model.Role)
	_ = c.ShouldBind(&param)

	// TODO 参数校验
	now := time.Now()
	param.CreateTime = now
	param.UpdateTime = now
	if param.ParentRoleId != 0 {
		// 判断该ID是否存在
		roles, err := param.Select([]uint{param.ParentRoleId}, "")
		if err != nil {
			response.ErrorResponseCustom(c, utils.ErrCodeMySQLError, err.Error())
			return
		}
		if len(roles) != 1 {
			response.ErrorResponseCustom(c, utils.ErrCodeParamError, "不存在该ID的父角色")
			return
		}
	}
	if err := param.Creat(); err != nil {
		global.EASLog.Error("Create role failed", zap.Any("err", err))
		response.ErrorResponseCustom(c, utils.ErrCodeMySQLError, "角色代码重复或者其他错误")
		return
	}
	response.SuccessResponse(c, param)
}
