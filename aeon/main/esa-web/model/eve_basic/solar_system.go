package eve_basic

import "aeon/global"

type SolarSystem struct {
	RegionID        int    `json:"region_id" form:"region_id" gorm:"column:regionID"`
	ConstellationID int    `json:"constellation_id" form:"constellation_id" gorm:"column:constellationID"`
	SolarSystemID   int    `json:"solar_system_id" form:"solar_system_id" gorm:"column:solarSystemID"`
	SolarSystemName string `json:"solar_system_name" form:"solar_system_name" gorm:"column:solarSystemName"`
}

func (SolarSystem) TableName() string {
	return "mapSolarSystems"
}

func (s *SolarSystem) SelectSolarSystem() (result []SolarSystem, err error) {
	db := global.ESAMySqlBasic.Model(s)

	db.Where("solarSystemName like ?", "%"+s.SolarSystemName+"%")
	err = db.Find(&result).Error
	return
}
