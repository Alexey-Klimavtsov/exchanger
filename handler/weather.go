package handler

import (
	"net/http"
	"strings"

	"weather-go/service"

	"github.com/gin-gonic/gin"
)

func Weather(c *gin.Context) {
	city := strings.ToLower(c.Query("city"))
	if city == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "city is required",
		})
		return
	}

	result, err := service.GetWeather(city)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, result)
}
