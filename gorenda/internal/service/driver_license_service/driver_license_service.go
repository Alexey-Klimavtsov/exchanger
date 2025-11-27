package driver_license_service

import (
	"context"
	"fmt"
	"github.com/asaipov/gorenda/internal/it/model/driver_license_model"
	"time"
)

type DriverLicenseInput struct {
	UserID    int64     `json:"userID"`
	Number    string    `json:"number"`
	IssuedAt  time.Time `json:"issuedAt"`
	ExpiresAt time.Time `json:"expiresAt"`
}

func mapInputToModel(input *DriverLicenseInput) (*driver_license_model.DriverLicenseModel, error) {
	driverLicense := &driver_license_model.DriverLicenseModel{
		UserID:    input.UserID,
		Number:    input.Number,
		IssuedAt:  input.IssuedAt,
		ExpiresAt: input.ExpiresAt,
	}
	return driverLicense, driverLicense.Validate()
}

type DriverLicenseService interface {
	CreateLicense(ctx context.Context, dl *DriverLicenseInput) (*driver_license_model.DriverLicenseModel, error)
	UpdateLicense(ctx context.Context, dl *DriverLicenseInput, id int64) (*driver_license_model.DriverLicenseModel, error)
}

type DriverLicenseRepo interface {
	CreateLicense(ctx context.Context, dl *driver_license_model.DriverLicenseModel) (*driver_license_model.DriverLicenseModel, error)
	UpdateLicense(ctx context.Context, dl *driver_license_model.DriverLicenseModel, id int64) (*driver_license_model.DriverLicenseModel, error)
}

type driverLicenseService struct {
	repo DriverLicenseRepo
}

func NewDriverLicenseService(r DriverLicenseRepo) DriverLicenseService {
	return &driverLicenseService{repo: r}
}

func (s *driverLicenseService) CreateLicense(ctx context.Context, dl *DriverLicenseInput) (*driver_license_model.DriverLicenseModel, error) {
	dlModel, err := mapInputToModel(dl)

	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrInvalidInput, err)
	}

	return s.repo.CreateLicense(ctx, dlModel)
}

func (s *driverLicenseService) UpdateLicense(ctx context.Context, dl *DriverLicenseInput, id int64) (*driver_license_model.DriverLicenseModel, error) {
	if id <= 0 {
		return nil, fmt.Errorf("%w: %v", ErrInvalidInput, id)
	}

	dlModel, err := mapInputToModel(dl)

	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrInvalidInput, err)
	}

	return s.repo.UpdateLicense(ctx, dlModel, id)
}
