package main

import (
	"flag"
	"fmt"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"weather-go/cache"
	"weather-go/cli"
	"weather-go/config"
	_ "weather-go/docs" // для инициализации Swagger
	"weather-go/handler"
	"weather-go/middleware"
	"weather-go/service"
)

// @title Weather API
// @version 1.0
// @description API для получения прогноза погоды
// @host localhost:8080
// @BasePath /
func main() {

	cliMode := flag.Bool("cli", false, "Run as CLI")
	flag.Parse()

	cfg := config.Load()
	c := cache.New()
	svc := service.New(c, cfg.CacheTTL)

	if *cliMode {
		fmt.Println("CLI режим (Погода)")
		cliApp := cli.New(svc)
		cliApp.Run()
		return
	}
	weatherHandler := handler.New(svc)

	r := gin.New()
	r.Use(
		middleware.Recovery(),
		middleware.Logger(),
		middleware.NewLimiter(cfg.RateLimit, cfg.RateWindow),
	)

	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.GET("/weather", weatherHandler.Weather)
	r.GET("/today", weatherHandler.Today)
	r.GET("/weekly", weatherHandler.Weekly)

	r.Run(cfg.Port)
}
