package cli

import (
	"fmt"
	"weather-go/model"
	"weather-go/service"
	"weather-go/util"
)

type CLI struct {
	weather service.WeatherService
}

func New(w service.WeatherService) *CLI {
	return &CLI{weather: w}
}

func (c *CLI) Run() {

	weekly, err := c.weather.Weekly("almaty", "celsius")
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	temps := util.Map(weekly.Days, func(d model.DayWeather) float64 {
		return d.Temperature
	})

	avg := util.Sum(temps) / float64(len(temps))

	hot := util.Filter(weekly.Days, func(d model.DayWeather) bool {
		return d.Temperature > 20
	})

	fmt.Println("Город: Almaty")
	fmt.Println("Средняя температура:", avg)

	fmt.Println("Тёплые дни:")
	for _, d := range hot {
		fmt.Printf("- %s: %.1f°C\n", d.Day, d.Temperature)
	}

}
