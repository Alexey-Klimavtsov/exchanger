package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"weather-go/handler"

	"github.com/gin-gonic/gin"
)

func TestWeatherHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)

	r := gin.New()
	r.GET("/weather", handler.Weather)

	req, _ := http.NewRequest(
		http.MethodGet,
		"/weather?city=almaty",
		nil,
	)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", w.Code)
	}
}
