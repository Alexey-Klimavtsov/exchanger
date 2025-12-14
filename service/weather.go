package service

import (
	"fmt"

	"weather-go/model"
	"weather-go/provider"
)

var cities = map[string][2]float64{
	"almaty": {43.2567, 76.9286},
	"astana": {51.1694, 71.4491},
}

func GetWeather(city string) (model.Weather, error) {
	coords, ok := cities[city]
	if !ok {
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

	return model.Weather{
		City:      city,
		OpenMeteo: openMeteo,
		Wttr:      wttr,
	}, nil
}
