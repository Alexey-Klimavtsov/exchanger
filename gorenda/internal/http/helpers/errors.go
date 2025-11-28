package helpers

import (
	"errors"
	"github.com/asaipov/gorenda/internal/service/booking_service"
	"github.com/asaipov/gorenda/internal/service/car_service"
	"github.com/asaipov/gorenda/internal/service/driver_license_service"
	"github.com/asaipov/gorenda/internal/service/user_service"
	"net/http"
)

var (
	ErrBinding   = errors.New("invalid input")
	ErrReadingId = errors.New("invalid id")
)

var ErrorCodeMap = map[error]int{
	// Car service
	car_service.ErrInvalidInput: http.StatusBadRequest,
	car_service.ErrNotFound:     http.StatusNotFound,
	car_service.ErrDataReading:  http.StatusBadRequest,
	car_service.ErrUpdate:       http.StatusBadRequest,
	car_service.ErrCreate:       http.StatusBadRequest,
	car_service.ErrDelete:       http.StatusBadRequest,

	driver_license_service.ErrInvalidInput: http.StatusBadRequest,
	driver_license_service.ErrNotFound:     http.StatusNotFound,
	driver_license_service.ErrDataReading:  http.StatusBadRequest,
	driver_license_service.ErrUpdate:       http.StatusBadRequest,
	driver_license_service.ErrCreate:       http.StatusBadRequest,
	driver_license_service.ErrDelete:       http.StatusBadRequest,

	user_service.ErrInvalidInput: http.StatusBadRequest,
	user_service.ErrNotFound:     http.StatusNotFound,
	user_service.ErrDataReading:  http.StatusBadRequest,
	user_service.ErrUpdate:       http.StatusBadRequest,
	user_service.ErrCreate:       http.StatusBadRequest,
	user_service.ErrDelete:       http.StatusBadRequest,

	booking_service.ErrInvalidInput: http.StatusBadRequest,
	booking_service.ErrNotFound:     http.StatusNotFound,
	booking_service.ErrCreate:       http.StatusBadRequest,

	ErrBinding:   http.StatusBadRequest,
	ErrReadingId: http.StatusBadRequest,
}
