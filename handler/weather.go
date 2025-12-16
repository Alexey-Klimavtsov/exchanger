package handler

import (
	"net/http"
	"strings"

	"weather-go/service"

	"github.com/gin-gonic/gin"
)
type Handler struct{
	service *service.Service
}


func New(s *service.Service) *Handler{
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

func (h *Handler) Today(c *gin.Context) {
	city := c.Query("city")
	//uint:="fahrenheit"
	unit := c.DefaultQuery("unit", "celsius")

	if city == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "city is required"})
		//c.JSON(http.StatusOK, gin.H{"error": "city is required"})

		return
	}

	data, err := h.service.GetToday(city, unit)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, data)
}

func (h *Handler) Weekly(c *gin.Context) {
	city := c.Query("city")
	unit := c.DefaultQuery("unit", "celsius")

	if city == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "city is required"})
		return
	}

	data, err := h.service.GetWeekly(city, unit)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, data)
}
