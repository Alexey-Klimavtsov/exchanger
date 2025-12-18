package main

import (

	"flag"
	"fmt"
	"weather-go/cli"

	"weather-go/handler"
	"weather-go/middleware"
"weather-go/cache"
"weather-go/config"
"weather-go/service"
	"github.com/gin-gonic/gin"


	swaggerfiles "github.com/swaggo/files" // сам сваггер 
	ginSwagger "github.com/swaggo/gin-swagger" // пакет для Gin 
	docs "weather-go/docs" 
)

// @title Weather API
// @version 1.0
// @description API для получения прогноза погоды
// @host localhost:8080
// @BasePath /
func main() {

	cliMode:=flag.Bool("cli",false,"Run as CLI")
flag.Parse()
	
	docs.SwaggerInfo.Title = "Weather API"
    docs.SwaggerInfo.Description = "API для получения прогноза погоды"
    docs.SwaggerInfo.Version = "1.0"
    docs.SwaggerInfo.Host = "localhost:8080"
    docs.SwaggerInfo.BasePath = "/"

	cfg:=config.Load()
	c:=cache.New()
	svc:=service.New(c,cfg.CacheTTL)

	if *cliMode{
		fmt.Println("CLI режим (Погода)")
		cliApp:=cli.New(svc)
		cliApp.Run()
		return
	}
	h:=handler.New(svc)

	r := gin.New()
    r.Use(
		gin.Recovery(),
		middleware.Logger(),
		middleware.NewLimiter(cfg.RateLimit,cfg.RateWindow),
	)
      
      	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
		r.GET("/weather", h.Weather)
		r.GET("/today", h.Today)
        r.GET("/weekly", h.Weekly)
	

	r.Run(cfg.Port)


}
