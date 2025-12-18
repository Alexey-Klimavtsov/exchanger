package cli

import (
	"fmt"

	"weather-go/service"
)

type CLI struct {
	weather service.WeatherService
}

func New(w service.WeatherService) *CLI {
	return &CLI{weather: w}
}

func (c *CLI) Run() {
	var city string

	fmt.Print("Введите город: ")
	fmt.Scanln(&city)

	if city == "" {
		city = "almaty"
	}

	result, err := c.weather.Today(city, "celsius")
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	fmt.Printf("🌤 Погода в %s: %.1f °C\n",
		city,
		result.Temperature,
	)
}
