package service

import (
	"aeon/main/esa-web/model/evebasic"
)

func InitSolarSystemRedisCache() (err error) {
	solarSystems, err := new(evebasic.SolarSystem).SelectSolarSystem()
	if err != nil {
		return err
	}

	for _, solarSystem := range solarSystems {
		if err := solarSystem.CacheToRedis(); err != nil {
			return err
		}
	}
	return
}
