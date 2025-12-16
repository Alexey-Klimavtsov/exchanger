package main

import (
	"weather-go/handler"
	"weather-go/middleware"
"weather-go/cache"
"weather-go/config"
"weather-go/service"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg:=config.Load()

	c:=cache.New()
	svc:=service.New(c,cfg.CacheTTL)
	h:=handler.New(svc)

	r := gin.New()
    r.Use(
		gin.Recovery(),
		middleware.Logger(),
		middleware.NewLimiter(cfg.RateLimit,cfg.RateWindow),
	)

		r.GET("/weather", h.Weather)

		r.GET("/today", h.Today)
        r.GET("/weekly", h.Weekly)
	

	r.Run(cfg.Port)

}
