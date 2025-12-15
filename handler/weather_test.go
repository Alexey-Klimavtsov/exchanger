package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"
"time"
	"weather-go/service"
"weather-go/cache"
	
	"github.com/gin-gonic/gin"
)

func TestWeatherHandler_OK(t *testing.T) {
	gin.SetMode(gin.TestMode)

	c := cache.New()
		svc := service.New(c, 5*time.Minute)
		h := New(svc)

	r := gin.New()
	r.GET("/weather", h.Weather)

	req, _ := http.NewRequest(http.MethodGet, "/weather?city=almaty",nil)
		w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK{
		t.Fatalf("expected status 200, got %d", w.Code)
	}
}


func TestWeatherHandler_NoCity(t *testing.T) {
	gin.SetMode(gin.TestMode)

	c := cache.New()
	svc := service.New(c, 5*time.Minute)
	h := New(svc)

	r := gin.New()
	r.GET("/weather", h.Weather)

	req, _ := http.NewRequest(http.MethodGet, "/weather", nil)
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Fatalf("expected status 400, got %d", w.Code)
	}
}