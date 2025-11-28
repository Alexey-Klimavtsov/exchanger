package car_dto

import (
	"github.com/asaipov/gorenda/internal/it/model/car_model"
	"github.com/asaipov/gorenda/internal/service/car_service"
	"time"
)

type CreateCarRequest struct {
	Brand       string    `json:"brand"`
	Model       string    `json:"model"`
	Year        time.Time `json:"year"`
	RentalPrice int64     `json:"rentalPrice"`
	ImageUrl    *string   `json:"imageUrl"`
}

type CarResponseDto struct {
	ID          int64      `json:"id"`
	Brand       string     `json:"brand"`
	Model       string     `json:"model"`
	Year        time.Time  `json:"year"`
	RentalPrice int64      `json:"rentalPrice"`
	ImageUrl    *string    `json:"imageUrl"`
	CreatedAt   time.Time  `json:"createdAt"`
	UpdatedAt   *time.Time `json:"updatedAt"`
}

func CarToResponseDto(c *car_model.CarModel) *CarResponseDto {
	return &CarResponseDto{
		ID:          c.ID,
		Brand:       c.Brand,
		Model:       c.Model,
		Year:        c.Year,
		RentalPrice: c.RentalPrice,
		ImageUrl:    c.ImageUrl,
		CreatedAt:   c.CreatedAt,
		UpdatedAt:   c.UpdatedAt,
	}
}

func DtoToInput(dto *CreateCarRequest) *car_service.CreateCarInput {
	return &car_service.CreateCarInput{
		Brand:       dto.Brand,
		Model:       dto.Model,
		Year:        dto.Year,
		RentalPrice: dto.RentalPrice,
		ImageUrl:    dto.ImageUrl,
	}
}
