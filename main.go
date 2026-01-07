package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/Alexey-Klimavtsov/exchanger/exchanger"
)

type Rate float64

func main() {

	http.HandleFunc("/rate", rateHandler)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func rateHandler(w http.ResponseWriter, r *http.Request) {
	delay := 3 * time.Second
	ctx, cancel := context.WithTimeout(r.Context(), delay)
	defer cancel()

	from := getQueryParameter("from", r)
	to := getQueryParameter("to", r)

	rate, err := exchanger.FetchRate(ctx, from, to)

	if err != nil {
		switch {
		case errors.Is(err, context.DeadlineExceeded):
			fmt.Println("timeout error: ", err)
			http.Error(w, "timeout error", http.StatusRequestTimeout)
			return
		case errors.Is(err, context.Canceled):
			fmt.Println("Canceled: ", err)
			http.Error(w, "canceled", http.StatusRequestTimeout)
			return
		default:
			http.Error(w, "Server error", http.StatusInternalServerError)
			return
		}
	}

	data, _ := json.Marshal(rate)
	w.Write(data)

}

func getQueryParameter(key string, r *http.Request) string {
	p := r.URL.Query().Get(key)
	return p
}
