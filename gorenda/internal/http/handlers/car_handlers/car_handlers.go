package car_handlers

import (
	"fmt"
	"github.com/asaipov/gorenda/internal/http/dto/car_dto"
	helpers "github.com/asaipov/gorenda/internal/http/helpers"

	"github.com/asaipov/gorenda/internal/service/car_service"
	"github.com/gin-gonic/gin"
)

type CarHandlers struct {
	s car_service.CarService
}

func NewCarHandlers(cs car_service.CarService) *CarHandlers {
	return &CarHandlers{s: cs}
}

func (h *CarHandlers) CreateNewCar(c *gin.Context) {
	var req car_dto.CreateCarRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		helpers.HandleError(c, fmt.Errorf("%w %v", helpers.ErrBinding, req))
		return
	}

	input := car_dto.DtoToInput(&req)

	car, createErr := h.s.CreateNewCar(c, input)
	if createErr != nil {
		helpers.HandleError(c, createErr)
		return
	}
	helpers.WriteCreated(c, car_dto.CarToResponseDto(car))
}

func (h *CarHandlers) UpdateCar(c *gin.Context) {
	queryId, err := helpers.GetIdFromQuery(c)
	if err != nil {
		helpers.HandleError(c, fmt.Errorf("%w", helpers.ErrReadingId))
		return
	}

	var req car_dto.CreateCarRequest
	if bindErr := c.ShouldBindJSON(&req); bindErr != nil {
		helpers.HandleError(c, fmt.Errorf("%w %v", helpers.ErrBinding, req))
		return
	}

	input := car_dto.DtoToInput(&req)

	car, err := h.s.UpdateCar(c, input, queryId)
	if err != nil {
		helpers.HandleError(c, err)
		return
	}

	helpers.WriteOK(c, car_dto.CarToResponseDto(car))
}

func (h *CarHandlers) DeleteCar(c *gin.Context) {
	queryId, err := helpers.GetIdFromQuery(c)
	if err != nil {
		helpers.HandleError(c, fmt.Errorf("%w", helpers.ErrReadingId))
		return
	}

	carId, deleteErr := h.s.DeleteCar(c, queryId)
	if deleteErr != nil {
		helpers.HandleError(c, deleteErr)
		return
	}

	helpers.WriteOK(c, carId)
}

func (h *CarHandlers) GetCarById(c *gin.Context) {
	queryId, err := helpers.GetIdFromQuery(c)
	if err != nil {
		helpers.HandleError(c, fmt.Errorf("%w", helpers.ErrReadingId))
		return
	}

	car, getErr := h.s.GetCarById(c, queryId)
	if getErr != nil {
		helpers.HandleError(c, getErr)
		return
	}
	helpers.WriteOK(c, car_dto.CarToResponseDto(car))
}
