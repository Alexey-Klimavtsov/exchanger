package model

type TodayWeather struct {
	Temperature   float64  `json:"temperature"`
	Description string `json:"description"`
	Unit      string `json:"unit"`
}

type WeeklyWeather struct{
	Days []DayWeather `json:"days"`
	Unit string `json:"unit"`
}

type DayWeather struct{
	Day  string`json:"day"`
	Temperature float64 `json:"temperature"`
}

type Weather struct {
	City      string  `json:"city"`
	OpenMeteo float64 `json:"open_meteo"`
	Wttr      float64 `json:"wttr"`
}