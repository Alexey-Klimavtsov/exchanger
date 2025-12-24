package provider

import (
	"encoding/json"
	"fmt"
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

	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var data wttrResponse
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return 0, err
	}

	if len(data.CurrentCondition) == 0 {
		return 0, fmt.Errorf("no weather data available")
	}

	return strconv.ParseFloat(data.CurrentCondition[0].TempC, 64)
}
