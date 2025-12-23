package service

import "weather-go/model"

type OpenMeteoService struct {
	apiKey string
}

func NewOpenMeteo(apiKey string) *OpenMeteoService {
	return &OpenMeteoService{apiKey: apiKey}
}

// 👇 РЕАЛИЗАЦИЯ интерфейса
func (o *OpenMeteoService) Today(city, unit string) (model.DayWeather, error) {
	// запрос к API
	return model.DayWeather{
		City: city,
		Temp: 25,
		Unit: unit,
	}, nil
}

func (o *OpenMeteoService) Weekly(city, unit string) (model.WeeklyWeather, error) {
	return model.WeeklyWeather{
		City: city,
	}, nil
}
