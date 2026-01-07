package exchanger

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Rate float64
type apiResponse map[string]any

func FetchRate(ctx context.Context, from, to string) (Rate, error) {
	resultChan := make(chan Rate)
	errorChan := make(chan error)

	go func() {
		// rate := fakeRequest(from, to)
		rate, err := getRate(ctx, from, to)
		if err != nil {
			errorChan <- err
			return
		}
		resultChan <- rate
	}()

	select {
	case err := <-errorChan:
		return 0, err
	case rate := <-resultChan:
		return rate, nil
	case <-ctx.Done():
		return 0, ctx.Err()
	}
}

func getRate(ctx context.Context, from, to string) (Rate, error) {
	url := "https://latest.currency-api.pages.dev/v1/currencies/" + from + ".json"

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)

	if err != nil {
		return 0, err
	}

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return 0, err
	}

	if resp.StatusCode != http.StatusOK {
		return 0, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}

	var res apiResponse
	err = json.Unmarshal(body, &res)
	if err != nil {
		return 0, fmt.Errorf("can't parse response. Error: %w", err)
	}

	rates, ok := res[from].(map[string]any)
	if !ok {
		return 0, fmt.Errorf("incorrect data from server. Error: %w", err)
	}

	rate, ok := rates[to].(float64)
	if !ok {
		return 0, fmt.Errorf("incorrect data from server. Error: %w", err)
	}

	return Rate(rate), nil
}
