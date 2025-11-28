package driver_license_dto

import (
	"github.com/asaipov/gorenda/internal/it/model/driver_license_model"
	"github.com/asaipov/gorenda/internal/service/driver_license_service"
	"time"
)

type DriverLicenseRequest struct {
	UserID    int64     `json:"userId"`
	Number    string    `json:"number"`
	IssuedAt  time.Time `json:"issuedAt"`
	ExpiresAt time.Time `json:"expiresAt"`
}

type DriverLicenseResponseDto struct {
	ID        int64      `json:"id"`
	UserID    int64      `json:"userId"`
	Number    string     `json:"number"`
	IssuedAt  time.Time  `json:"issuedAt"`
	ExpiresAt time.Time  `json:"expiresAt"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt *time.Time `json:"updatedAt"`
}

func DtoToInput(dto *DriverLicenseRequest) *driver_license_service.DriverLicenseInput {
	return &driver_license_service.DriverLicenseInput{
		UserID:    dto.UserID,
		Number:    dto.Number,
		IssuedAt:  dto.IssuedAt,
		ExpiresAt: dto.ExpiresAt,
	}
}

func DriverLicenseModelToDto(model *driver_license_model.DriverLicenseModel) *DriverLicenseResponseDto {
	return &DriverLicenseResponseDto{
		ID:        model.ID,
		UserID:    model.UserID,
		Number:    model.Number,
		IssuedAt:  model.IssuedAt,
		ExpiresAt: model.ExpiresAt,
		CreatedAt: model.CreatedAt,
		UpdatedAt: model.UpdatedAt,
	}
}
