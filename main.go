package main

import (
	"weather-go/handler"
	"weather-go/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()

	r.Use(
		gin.Recovery(),
		middleware.Logger(),
	)

		r.GET("/weather", handler.Weather)
	

	r.Run(":8080")
}
