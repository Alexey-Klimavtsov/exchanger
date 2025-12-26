package service

import "weather-go/model"

type FakeWeatherService struct{}

func (f FakeWeatherService) Today(city, unit string) (model.TodayWeather, error) {
	return model.TodayWeather{
		Temperature: 20,
		Description: "Clear",
		Unit:        unit,
	}, nil
}

func (f FakeWeatherService) Weekly(city, unit string) (model.WeeklyWeather, error) {
	return model.WeeklyWeather{}, nil
}
