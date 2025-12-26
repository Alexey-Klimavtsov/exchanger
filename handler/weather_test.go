package handler

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/gin-gonic/gin"

	"weather-go/cache"
	"weather-go/service"
)

func TestWeatherHandler_OK(t *testing.T) {
	gin.SetMode(gin.TestMode)

	c := cache.New()
	svc := service.New(c, 5*time.Minute)
	h := New(svc)

	r := gin.New()
	r.GET("/weather", h.Weather)

	req, _ := http.NewRequest(http.MethodGet, "/weather?city=almaty", nil)
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
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

func setupRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)

	c := cache.New()
	svc := service.New(c, time.Minute)
	h := New(svc)

	r := gin.New()
	r.GET("/today", h.Today)
	r.GET("/weekly", h.Weekly)

	return r
}
func TestToday_OK(t *testing.T) {
	r := setupRouter()

	req, _ := http.NewRequest("GET", "/today?city=almaty", nil)
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", w.Code)
	}

	if !strings.Contains(w.Body.String(), "temperature") {
		t.Fatal("response does not contain temperature")
	}
}
func TestToday_NoCity(t *testing.T) {
	r := setupRouter()

	req, _ := http.NewRequest("GET", "/today", nil)
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Fatalf("expected status 400, got %d", w.Code)
	}

	if !strings.Contains(w.Body.String(), "city") {
		t.Fatal("expected error about city")
	}
}
func TestToday_DefaultUnit(t *testing.T) {
	r := setupRouter()

	req, _ := http.NewRequest("GET", "/today?city=almaty", nil)
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	if !strings.Contains(w.Body.String(), "celsius") {
		t.Fatal("expected default unit celsius")
	}
}
func TestWeekly_Fahrenheit(t *testing.T) {
	r := setupRouter()

	req, _ := http.NewRequest("GET", "/weekly?city=almaty&unit=fahrenheit", nil)
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", w.Code)
	}

	if !strings.Contains(w.Body.String(), "fahrenheit") {
		t.Fatal("expected unit fahrenheit")
	}
}
