package provider

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type wttrResponse struct {
	CurrentCondition []struct {
		TempC string `json:"temp_C"`
	} `json:"current_condition"`
}

func GetWttr(city string) (float64, error) {
	resp, err := http.Get("https://wttr.in/" + city + "?format=j1")
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	var data wttrResponse
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return 0, err
	}

	return strconv.ParseFloat(data.CurrentCondition[0].TempC, 64)
}
