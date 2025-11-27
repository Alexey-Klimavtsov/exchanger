package server

import (
	"context"
	"github.com/asaipov/gorenda/internal/app/cache"
	"github.com/asaipov/gorenda/internal/http/handlers/car_handlers"
	"github.com/asaipov/gorenda/internal/http/handlers/driver_license_handlers"
	"github.com/asaipov/gorenda/internal/http/handlers/user_handlers"
	"github.com/asaipov/gorenda/internal/service/car_service"
	"github.com/asaipov/gorenda/internal/service/driver_license_service"
	"github.com/asaipov/gorenda/internal/service/user_service"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
	"time"
)

type Server struct {
	carsService          car_service.CarService
	driverLicenseService driver_license_service.DriverLicenseService
	userService          user_service.UserService
	router               *gin.Engine
}

func NewServer(carsService car_service.CarService, driverLicenseService driver_license_service.DriverLicenseService, userService user_service.UserService) *Server {
	s := &Server{carsService: carsService, driverLicenseService: driverLicenseService, userService: userService}
	s.setupRouter()
	return s
}

func (s *Server) Router() *gin.Engine {
	return s.router
}

func (s *Server) setupRouter() {
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(withTimeout(30 * time.Second))

	carsHandlers := car_handlers.NewCarHandlers(s.carsService)
	driverLicenseHandlers := driver_license_handlers.NewDriverLicenseHandlers(s.driverLicenseService)
	userHandlers := user_handlers.NewUserHandlers(s.userService)

	middlewareCache := cache.NewCacheMiddleware(30 * time.Second)

	v1 := r.Group("/v1")
	{
		cars := v1.Group("/cars")
		cars.GET("/:id", middlewareCache.Middleware(), carsHandlers.GetCarById)
		cars.POST("", carsHandlers.CreateNewCar)
		cars.DELETE("/:id", carsHandlers.DeleteCar)
		cars.PATCH("/:id", carsHandlers.UpdateCar)
	}

	{
		license := v1.Group("/license")
		license.POST("", driverLicenseHandlers.Create)
		license.PATCH("/:id", driverLicenseHandlers.Update)
	}
	{
		user := v1.Group("user")
		user.GET("/:id", userHandlers.GetById)
		user.POST("", userHandlers.Create)
		user.PATCH("/:id", userHandlers.Update)
		user.DELETE("/:id", userHandlers.Delete)
	}
	s.router = r
	s.router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

}

func withTimeout(d time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(c.Request.Context(), d)
		defer cancel()

		c.Request = c.Request.WithContext(ctx)

		done := make(chan struct{})

		go func() {
			c.Next()
			close(done)
		}()

		select {
		case <-done:
			return
		case <-ctx.Done():
			c.Abort()
			c.JSON(http.StatusRequestTimeout, gin.H{
				"error": "Request timeout",
				"code":  "request_timeout",
			})
			return
		}
	}
}
