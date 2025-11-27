package main

import (
	"errors"
	"github.com/asaipov/gorenda/internal/app/server"
	"github.com/asaipov/gorenda/internal/infra/db/sqlite"
	"github.com/asaipov/gorenda/internal/repo/sqlite/car_repo"
	"github.com/asaipov/gorenda/internal/repo/sqlite/driver_license_repo"
	"github.com/asaipov/gorenda/internal/repo/sqlite/user_repo"
	"github.com/asaipov/gorenda/internal/service/car_service"
	"github.com/asaipov/gorenda/internal/service/driver_license_service"
	"github.com/asaipov/gorenda/internal/service/user_service"
	"log"
	"net/http"
	"time"
)

func main() {
	dbc, err := sqlite.Open("./gorenda.db")
	if err != nil {
		log.Fatal("Не удалось открыть базу:", err)
	}
	defer dbc.Close()

	if pingErr := dbc.Ping(); pingErr != nil {
		log.Fatal("База не отвечает:", pingErr)
	}

	carRepo := car_repo.NewCarRepo(dbc)
	carService := car_service.NewCarService(carRepo)
	driverLicenseRepo := driver_license_repo.NewDriverLicenseRepo(dbc)
	driverLicenseService := driver_license_service.NewDriverLicenseService(driverLicenseRepo)
	userRepo := user_repo.NewUserRepo(dbc)
	userService := user_service.NewUserService(userRepo)

	newServer := server.NewServer(carService, driverLicenseService, userService)

	log.Println("Подключено")

	httpServer := &http.Server{
		Addr:         ":8080",
		Handler:      newServer.Router(),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	log.Println("listen :8080")
	if listenErr := httpServer.ListenAndServe(); listenErr != nil && !errors.Is(listenErr, http.ErrServerClosed) {
		log.Fatalf("Server error: %v", listenErr)
	}

}
