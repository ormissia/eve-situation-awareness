package model

import (
	"fmt"
	"time"

	"go.uber.org/zap"
	"gorm.io/gorm"

	"aeon/global"
	"aeon/main/esa-web/model/evebasic"
	"aeon/utils"
)

type SolarSystemKill struct {
	SolarSystemId   int64  `json:"solar_system_id" form:"solar_system_id" gorm:"column:solar_system_id"`
	KillQuantity    int64  `json:"kill_quantity" form:"kill_quantity" gorm:"column:kill_quantity"`
	KillValue       int64  `json:"kill_value" form:"kill_value" gorm:"column:kill_value"`
	Dt              string `json:"dt" form:"dt" gorm:"column:dt"`
	SolarSystemName string `json:"solar_system_name" form:"-" gorm:"-"`
	CreateTime      string `json:"create_time" form:"create_time" gorm:"column:create_time"`
}

func (SolarSystemKill) TableName() string {
	return "solar_system_kill_statistical"
}

func (s *SolarSystemKill) AfterFind(tx *gorm.DB) (err error) {
	ss := new(evebasic.SolarSystem)
	ss.SolarSystemID = s.SolarSystemId
	if err = ss.SelectInfoById(); err != nil {
		global.ESALog.Error("find solar system from cache err", zap.Any("err", err))
	}
	s.SolarSystemName = ss.SolarSystemName
	return
}

func (s *SolarSystemKill) SelectSolarSystem(params BaseParams) (result []SolarSystemKill, err error) {
	db := global.ESAMySqlESA.Model(s)

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

	if s.SolarSystemId != 0 {
		db.Where("solar_system_id = ?", s.SolarSystemId)
	}

	db.Order("dt")

	err = db.Select("dt, solar_system_id, sum(kill_quantity) as kill_quantity, sum(kill_value) as kill_value").
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

func (s *SolarSystemKill) SelectSolarSystemSum(params BaseParams) (result []SolarSystemKill, err error) {
	db := global.ESAMySqlESA.Model(s)

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

	orderField := ""
	orderType := ""

	if params.OrderField == "kill_value" {
		orderField = params.OrderField
	} else if params.OrderField == "kill_quantity" {
		orderField = params.OrderField
	}
	if params.OrderType == "" || params.OrderType == "asc" {
		orderType = params.OrderType
	} else if params.OrderType == "desc" {
		orderType = params.OrderType
	}

	db.Order(orderField + " " + orderType)

	db.Group("solar_system_id")

	db.Limit(10)

	err = db.Select("solar_system_id, sum(kill_quantity) as kill_quantity, sum(kill_value) as kill_value").
		Find(&result).Error
	return
}
