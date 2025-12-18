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

type WeatherService interface {
	Today(city, unit string) (model.TodayWeather, error)
	Weekly(city, unit string) (model.WeeklyWeather, error)
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


func (s *Service) Today(city, unit string) (model.TodayWeather, error) {
	w, err := s.GetWeather(city)
	if err != nil {
		return model.TodayWeather{}, err
	}

	temp := w.OpenMeteo
	if unit == "fahrenheit" {
		temp = temp*1.8 + 32
	}

	return model.TodayWeather{
		Temperature: temp,
		Description: "Clear",
		Unit:        unit,
	}, nil
}

func (s *Service) Weekly(city, unit string) (model.WeeklyWeather, error) {
	_, err := s.GetWeather(city)
	if err != nil {
		return model.WeeklyWeather{}, err
	}

	days := []model.DayWeather{
	    {Day: "Mon", Temperature: 20},
		{Day: "Tue", Temperature: 21},
		{Day: "Wed", Temperature: 19},
		{Day: "Thu", Temperature: 18},
		{Day: "Fri", Temperature: 22},
		{Day: "Sat", Temperature: 23},
		{Day: "Sun", Temperature: 21},

	}

	if unit == "fahrenheit" {
		for i := range days {
			days[i].Temperature = days[i].Temperature*1.8 + 32
		}
	}

	return model.WeeklyWeather{
		Days: days,
		Unit: unit,
	}, nil
}
