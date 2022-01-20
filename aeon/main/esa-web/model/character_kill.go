package model

import (
	"fmt"
	"time"

	"aeon/global"
	"aeon/utils"
)

type CharacterKill struct {
	CharacterId  int    `json:"character_id"`
	KillQuantity int    `json:"kill_quantity"`
	KillValue    int    `json:"kill_value"`
	FinalShoot   int    `json:"final_shoot"`
	Dt           string `json:"dt"`
	Labels       string `json:"labels"`
	ShipTypes    string `json:"ship_types"`
	SolarSystems string `json:"solar_systems"`
	CreateTime   string `json:"create_time"`
}

func (CharacterKill) TableName() string {
	return "character_kill_statistical"
}

func (c *CharacterKill) SelectCharacterKill(params BaseParams) (result []CharacterKill, err error) {
	db := global.ESAMySqlESA.Model(c)

	if params.StartTimeStamp != 0 {
		start := time.UnixMilli(params.StartTimeStamp)
		startStr := start.Format(utils.DTTimeFormat)
		db.Where("dt >= ?", startStr)
	}

	if params.EndTimeStamp != 0 {
		end := time.UnixMilli(params.EndTimeStamp)
		endStr := end.Format(utils.DTTimeFormat)
		db.Where("dt <= ?", endStr)
	}

	dtFormat := ""

	switch params.TimeType {
	case utils.Hour, utils.Day, utils.Month, utils.Year:
		dtFormat = fmt.Sprintf("date_format(dt, '%s')", utils.SQLGroupFormatMap[params.TimeType])
	default:
		// 默认按照小时分组
		global.ESALog.Warn("invalid or empty time_type")
		dtFormat = fmt.Sprintf("date_format(dt, '%s')", utils.SQLGroupFormatMap[utils.Hour])
	}
	db.Group(dtFormat)

	if params.PageNo != 0 {
		db.Offset((params.PageNo - 1) * params.PageSize)
	}

	if params.PageSize != 0 {
		db.Limit(params.PageSize)
	} else {
		// 默认限制8760条结果
		db.Limit(8760)
	}

	if c.CharacterId != 0 {
		db.Where("character_id = ?", c.CharacterId)
	}

	db.Order("dt")

	err = db.Select("dt, character_id, sum(kill_quantity) as kill_quantity, sum(kill_value) as kill_value, group_concat(labels) as labels, group_concat(ship_types) as ship_types, group_concat(solar_systems) as solar_systems").
		Scan(&result).Error

	timeFormat, ok := utils.ResultDTFormatMap[params.TimeType]
	if ok {
		for i, killInfo := range result {
			dtTime, err := time.Parse(utils.DTTimeFormat, killInfo.Dt)
			if err != nil {
				return nil, err
			}
			result[i].Dt = dtTime.Format(timeFormat)
		}
	}

	return
}
