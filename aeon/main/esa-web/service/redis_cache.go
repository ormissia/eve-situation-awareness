package service

import (
	"aeon/main/esa-web/model/eve_basic"
)

func InitSolarSystemRedisCache() (err error) {
	solarSystems, err := new(eve_basic.SolarSystem).SelectSolarSystem()
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
