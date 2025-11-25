package main

import (
	"errors"
	"github.com/asaipov/gorenda/internal/app/server"
	"github.com/asaipov/gorenda/internal/infra/db/sqlite"
	"github.com/asaipov/gorenda/internal/repo/sqlite/car_repo"
	"github.com/asaipov/gorenda/internal/service/car_service"
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

	carRepo := car_repo.NewCarRepo(dbc)
	carService := car_service.NewCarService(carRepo)

	newServer := server.NewServer(carService)

	if pingErr := dbc.Ping(); pingErr != nil {
		log.Fatal("База не отвечает:", pingErr)
	}

	log.Println("Подключено")

	httpServer := &http.Server{
		Addr:         ":8081",
		Handler:      newServer.Router(),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	log.Println("listen :8081")
	if listenErr := httpServer.ListenAndServe(); listenErr != nil && !errors.Is(listenErr, http.ErrServerClosed) {
		log.Fatalf("Server error: %v", listenErr)
	}

}
