package service

import (
	"fmt"
	"time"
	"weather-go/cache"
	"weather-go/model"
	"weather-go/provider"
)

type Service struct{
cache  *cache.Cache
cacheTTL  time.Duration
}

func New(cache *cache.Cache,ttl time.Duration) *Service{
return &Service{
cache:  cache,
cacheTTL: ttl,
}
}

var cities = map[string][2]float64{
	"almaty": {43.2567, 76.9286},
	"astana": {51.1694, 71.4491},
}

func(s*Service) GetWeather(city string) (model.Weather, error) {
 	if cached,ok:=s.cache.Get(city);ok{
		return cached.(model.Weather),nil
	}

	coords, ok := cities[city]
	if!ok {
		return model.Weather{}, fmt.Errorf("unknown city")
	}


	openMeteo, err := provider.GetOpenMeteo(coords[0], coords[1])
	if err != nil {
		return model.Weather{}, err
	}

	wttr, err := provider.GetWttr(city)
	if err != nil {
		return model.Weather{}, err
	}

	result:= model.Weather{
		City:      city,
		OpenMeteo: openMeteo,
		Wttr:      wttr,
	}

	s.cache.Set(city,result,s.cacheTTL)
	return result,nil 
}
