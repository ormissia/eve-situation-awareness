package eve_basic

import (
	"context"
	"encoding/json"
	"errors"
	"strconv"
	"time"

	"go.uber.org/zap"

	"aeon/global"
)

var ErrInvalidSolarSystemID = errors.New("invalid id to generate redis key")
var ErrInvalidSolarSystem = errors.New("invalid solar system can not be cache to redis")

// SolarSystem
// Redis 缓存设计：每一条星系信息作为redis中的一个key下的数据
type SolarSystem struct {
	RegionID        int64  `json:"region_id" form:"region_id" gorm:"column:regionID"`
	ConstellationID int64  `json:"constellation_id" form:"constellation_id" gorm:"column:constellationID"`
	SolarSystemID   int64  `json:"solar_system_id" form:"solar_system_id" gorm:"column:solarSystemID"`
	SolarSystemName string `json:"solar_system_name" form:"solar_system_name" gorm:"column:solarSystemName"`
}

func (s *SolarSystem) GetRedisKey() (key string, err error) {
	return global.ESAConfig.Server.ServiceName + ":" + "solar-system-info" + ":" + strconv.FormatInt(s.SolarSystemID, 10), nil
}

func (s *SolarSystem) GetRedisExpiration() (duration time.Duration) {
	return time.Hour * 72
}

func (s *SolarSystem) CacheToRedis() (err error) {
	if s.SolarSystemID == 0 || s.SolarSystemName == "" {
		return ErrInvalidSolarSystem
	}
	bytes, err := json.Marshal(s)
	if err != nil {
		return
	}

	redisKey, err := s.GetRedisKey()
	if err != nil {
		return err
	}
	statusCmd := global.ESARedis.Set(context.Background(), redisKey, string(bytes), s.GetRedisExpiration())
	err = statusCmd.Err()
	return
}

func (s *SolarSystem) GetFromRedis() (err error) {
	if s.SolarSystemID == 0 {
		return ErrInvalidSolarSystemID
	}

	redisKey, err := s.GetRedisKey()
	if err != nil {
		return err
	}

	stringCmd := global.ESARedis.Get(context.Background(), redisKey)
	if stringCmd.Err() != nil {
		return stringCmd.Err()
	}

	if err = json.Unmarshal([]byte(stringCmd.Val()), s); err != nil {
		return err
	}

	return nil
}

func (SolarSystem) TableName() string {
	return "mapSolarSystems"
}

func (s *SolarSystem) SelectSolarSystem() (result []SolarSystem, err error) {
	db := global.ESAMySqlBasic.Model(s)

	if s.SolarSystemName != "" {
		db.Where("solarSystemName like ?", "%"+s.SolarSystemName+"%")
	}
	err = db.Find(&result).Error
	return
}

func (s *SolarSystem) SelectInfoById() (err error) {
	if s.SolarSystemID == 0 {
		return ErrInvalidSolarSystemID
	}
	// 从redis中查询
	if err := s.GetFromRedis(); err == nil {
		return nil
	}
	global.ESALog.Warn("not found from redis", zap.Any("key", s.GetFromRedis()))

	// 从mysql中查询
	db := global.ESAMySqlBasic.Model(s)

	if s.SolarSystemID != 0 {
		db.Where("solarSystemId = ?", s.SolarSystemID)
	}
	if err = db.Find(&s).Error; err != nil {
		return err
	}

	// 将缓存写入redis
	if err = s.CacheToRedis(); err != nil {
		return
	}
	return
}
