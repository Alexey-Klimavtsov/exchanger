
package service

import "weather-go/model"

type FakeWeatherService struct {}

func (f FakeWeatherService) Today(city, unit string) (model.DayWeather, error) {
	return model.DayWeather{
		City: city,
		Temp: 20,
		Unit: unit,
	}, nil
}

func (f FakeWeatherService) Weekly(city, unit string) (model.WeeklyWeather, error) {
	return model.WeeklyWeather{}, nil
}
