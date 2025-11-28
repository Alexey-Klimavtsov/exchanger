package driver_license_handlers

import (
	"fmt"
	"github.com/asaipov/gorenda/internal/http/dto/driver_license_dto"
	"github.com/asaipov/gorenda/internal/http/helpers"
	"github.com/asaipov/gorenda/internal/service/driver_license_service"
	"github.com/gin-gonic/gin"
)

type DriverLicenseHandlers struct {
	s driver_license_service.DriverLicenseService
}

func NewDriverLicenseHandlers(s driver_license_service.DriverLicenseService) *DriverLicenseHandlers {
	return &DriverLicenseHandlers{s: s}
}

func (h *DriverLicenseHandlers) Create(c *gin.Context) {
	var req driver_license_dto.DriverLicenseRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		helpers.HandleError(c, fmt.Errorf("%w: %v", helpers.ErrBinding, req))
		return
	}

	input := driver_license_dto.DtoToInput(&req)
	dl, createErr := h.s.CreateLicense(c, input)
	if createErr != nil {
		helpers.HandleError(c, createErr)
		return
	}
	helpers.WriteCreated(c, driver_license_dto.DriverLicenseModelToDto(dl))
}

func (h *DriverLicenseHandlers) Update(c *gin.Context) {
	var req driver_license_dto.DriverLicenseRequest
	id, getErr := helpers.GetIdFromQuery(c)
	if getErr != nil {
		helpers.HandleError(c, fmt.Errorf("%w", helpers.ErrReadingId))
		return
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		helpers.HandleError(c, fmt.Errorf("%w: %v", helpers.ErrBinding, req))
		return
	}

	input := driver_license_dto.DtoToInput(&req)
	dl, updateErr := h.s.UpdateLicense(c, input, id)
	if updateErr != nil {
		helpers.HandleError(c, updateErr)
		return
	}
	helpers.WriteOK(c, driver_license_dto.DriverLicenseModelToDto(dl))
}
