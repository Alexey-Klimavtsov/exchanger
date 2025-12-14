package model

type Weather struct {
	City      string  `json:"city"`
	OpenMeteo float64 `json:"open_meteo"`
	Wttr      float64 `json:"wttr"`
}
