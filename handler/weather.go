package handler

import (
	"net/http"
	"strings"

	"weather-go/service"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *service.Service
}

func New(s *service.Service) *Handler {
	return &Handler{service: s}
}

func (h *Handler) Weather(c *gin.Context) {
	city := strings.ToLower(c.Query("city"))
	if city == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "city is required",
		})
		return
	}

	result, err := h.service.GetWeather(city)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, result)
}

// Today godoc
// @Summary Прогноз погоды на сегодня
// @Description Возвращает погоду на текущий день
// @Tags Weather
// @Param city query string true "Город" default(almaty)
// @Param unit query string false "Единицы измерения" Enums(celsius, fahrenheit) default(celsius)
// @Success 200 {object} model.TodayWeather
// @Failure 400 {object} map[string]string
// @Router /today [get]
func (h *Handler) Today(c *gin.Context) {
	city := strings.ToLower(c.Query("city"))
	unit := c.DefaultQuery("unit", "celsius")

	if city == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "city is required"})
		return
	}

	data, err := h.service.Today(city, unit)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, data)
}

// Weekly godoc
// @Summary Получить прогноз на неделю
// @Description Возвращает прогноз погоды на 7 дней
// @Tags Weather
// @Param city query string true "Название города" default(almaty)
// @Param unit query string false "Единицы измерения" Enums(celsius, fahrenheit) default(celsius)
// @Success 200 {object} model.WeeklyWeather
// @Failure 400 {object} map[string]string
// @Router /weekly [get]
func (h *Handler) Weekly(c *gin.Context) {
	city := strings.ToLower(c.Query("city"))
	unit := c.DefaultQuery("unit", "celsius")

	if city == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "city is required"})
		return
	}

	data, err := h.service.Weekly(city, unit)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, data)
}
