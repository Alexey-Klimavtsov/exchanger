package service

import "weather-go/model"

type OpenMeteoService struct {
	apiKey string
}

func NewOpenMeteo(apiKey string) *OpenMeteoService {
	return &OpenMeteoService{apiKey: apiKey}
}

// Today реализует интерфейс WeatherService
func (o *OpenMeteoService) Today(city, unit string) (model.TodayWeather, error) {
	// TODO: реализовать запрос к API
	return model.TodayWeather{
		Temperature: 25,
		Description: "Clear",
		Unit:        unit,
	}, nil
}

func (o *OpenMeteoService) Weekly(city, unit string) (model.WeeklyWeather, error) {
	return model.WeeklyWeather{}, nil
}
