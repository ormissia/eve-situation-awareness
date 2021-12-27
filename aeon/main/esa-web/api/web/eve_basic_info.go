package web

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"aeon/global"
	"aeon/main/esa-web/model"
	"aeon/main/esa-web/model/eve_basic"
	"aeon/utils"
)

func SolarSystemFuzzySearch(c *gin.Context) {
	var solarSystem eve_basic.SolarSystem
	if err := c.ShouldBind(&solarSystem); err != nil {
		global.ESALog.Error("param should bind BaseParams err", zap.Any("err", err))
		model.ErrorResponse(c, utils.ErrCodeMissingParamError)
		return
	}

	global.ESALog.Info("params", zap.Any("SolarSystem", solarSystem))

	if len(solarSystem.SolarSystemName) <= 1 {
		model.SuccessResponse(c, []eve_basic.SolarSystem{})
		return
	}

	// TODO 星系名称大小写

	solarSystems, err := solarSystem.SelectSolarSystem()
	if err != nil {
		global.ESALog.Error("select mysql err", zap.Any("err", err))
		model.ErrorResponse(c, utils.ErrCodeMySQLError)
		return
	}

	model.SuccessResponse(c, solarSystems)
}
