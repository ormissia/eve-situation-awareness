package web

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"aeon/global"
	"aeon/main/esa-web/calculate"
	"aeon/main/esa-web/model"
	"aeon/utils"
)

func SearchCharacterKill(c *gin.Context) {
	var param model.BaseParams
	if err := c.ShouldBind(&param); err != nil {
		global.ESALog.Error("param should bind BaseParams err", zap.Any("err", err))
		model.ErrorResponse(c, utils.ErrCodeMissingParamError)
		return
	}

	var characterKill model.CharacterKill
	if err := c.ShouldBind(&characterKill); err != nil {
		global.ESALog.Error("param should bind characterKill err", zap.Any("err", err))
		model.ErrorResponse(c, utils.ErrCodeParamShouldBindError)
		return
	}

	global.ESALog.Info("params", zap.Any("BaseParams", param), zap.Any("characterKill", characterKill))

	characterKills, err := characterKill.SelectCharacterKill(param)
	if err != nil {
		global.ESALog.Error("select mysql err", zap.Any("err", err))
		model.ErrorResponse(c, utils.ErrCodeMySQLError)
		return
	}

	for i, _ := range characterKills {
		//
		labelMR := calculate.MRUnit{
			Param: characterKills[i].Labels,
		}

		if err := labelMR.Map().Reduce().Error; err != nil {
			global.ESALog.Error("mr err", zap.Any("err", err))
			model.ErrorResponse(c, utils.ErrCodeMySQLError)
			return
		}
		labels := make([]string, 0)
		for i, l := range labelMR.Result {
			labels = append(labels, l)
			if i != len(labelMR.Result)-1 {
				labels = append(labels, ",")
			}
		}
		characterKills[i].Labels = utils.StringSliceBuilder(labels)

		//
		shipTypeMR := calculate.MRUnit{
			Param: characterKills[i].ShipTypes,
		}

		if err := shipTypeMR.Map().Reduce().Error; err != nil {
			global.ESALog.Error("mr err", zap.Any("err", err))
			model.ErrorResponse(c, utils.ErrCodeMySQLError)
			return
		}
		shipTypes := make([]string, 0)
		for i, l := range shipTypeMR.Result {
			shipTypes = append(shipTypes, l)
			if i != len(shipTypeMR.Result)-1 {
				shipTypes = append(shipTypes, ",")
			}
		}
		characterKills[i].ShipTypes = utils.StringSliceBuilder(shipTypes)

		//
		solarSystemMR := calculate.MRUnit{
			Param: characterKills[i].SolarSystems,
		}

		if err := solarSystemMR.Map().Reduce().Error; err != nil {
			global.ESALog.Error("mr err", zap.Any("err", err))
			model.ErrorResponse(c, utils.ErrCodeMySQLError)
			return
		}
		solarSystems := make([]string, 0)
		for i, l := range solarSystemMR.Result {
			solarSystems = append(solarSystems, l)
			if i != len(solarSystemMR.Result)-1 {
				solarSystems = append(solarSystems, ",")
			}
		}
		characterKills[i].SolarSystems = utils.StringSliceBuilder(solarSystems)
	}

	var result model.ChartFormatData
	err = result.Convert("dt", []string{"kill_quantity", "kill_value", "labels", "ship_types", "solar_systems"}, characterKills)
	if err != nil {
		global.ESALog.Error("convert chart format data err", zap.Any("err", err))
		model.ErrorResponse(c, utils.ErrServerError)
		return
	}

	model.SuccessResponse(c, result)
}
