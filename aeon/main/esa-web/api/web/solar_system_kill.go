package web

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"aeon/global"
	"aeon/main/esa-web/model"
	"aeon/utils"
)

func SearchSolarSystemKill(c *gin.Context) {
	var param model.BaseParams
	if err := c.ShouldBind(&param); err != nil {
		global.ESALog.Error("param should bind BaseParams err", zap.Any("err", err))
		model.ErrorResponse(c, utils.ErrCodeMissingParamError)
		return
	}

	var solarSystemKill model.SolarSystemKill
	if err := c.ShouldBind(&solarSystemKill); err != nil {
		global.ESALog.Error("param should bind solarSystemKill err", zap.Any("err", err))
		model.ErrorResponse(c, utils.ErrCodeParamShouldBindError)
		return
	}

	global.ESALog.Info("params", zap.Any("BaseParams", param), zap.Any("SolarSystemKill", solarSystemKill))

	solarSystemKills, err := solarSystemKill.SelectSolarSystem(param)
	if err != nil {
		global.ESALog.Error("select mysql err", zap.Any("err", err))
		model.ErrorResponse(c, utils.ErrCodeMySQLError)
		return
	}

	var result model.ChartFormatData
	err = result.Convert("dt", []string{"kill_quantity", "kill_value"}, solarSystemKills)
	if err != nil {
		global.ESALog.Error("convert chart format data err", zap.Any("err", err))
		model.ErrorResponse(c, utils.ErrServerError)
		return
	}

	model.SuccessResponse(c, result)
}
