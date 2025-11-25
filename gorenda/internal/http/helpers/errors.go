package helpers

import (
	"database/sql"
	"github.com/asaipov/gorenda/internal/repo/helpers"
	"github.com/asaipov/gorenda/internal/service/car_service"
	"net/http"
)

var ErrorCodeMap = map[error]int{
	car_service.ErrInvalidInput: http.StatusBadRequest,
	car_service.ErrNotFound:     http.StatusNotFound,

	helpers.ErrDuplicate:   http.StatusBadRequest,
	helpers.ErrForeignKey:  http.StatusBadRequest,
	helpers.ErrNotNull:     http.StatusBadRequest,
	helpers.ErrCheckFailed: http.StatusBadRequest,

	sql.ErrNoRows: http.StatusNotFound,
}
