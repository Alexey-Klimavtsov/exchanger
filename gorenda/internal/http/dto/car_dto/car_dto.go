package car_dto

import (
	"database/sql"
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
	var imageUrl *string
	var updatedAt *time.Time
	if c.ImageUrl.Valid {
		imageUrl = &c.ImageUrl.String
	} else {
		imageUrl = nil
	}
	if c.UpdatedAt.Valid {
		updatedAt = &c.UpdatedAt.Time
	} else {
		updatedAt = nil
	}

	return &CarResponseDto{
		ID:          c.ID,
		Brand:       c.Brand,
		Model:       c.Model,
		Year:        c.Year,
		RentalPrice: c.RentalPrice,
		ImageUrl:    imageUrl,
		CreatedAt:   c.CreatedAt,
		UpdatedAt:   updatedAt,
	}
}

func DtoToInput(dto *CreateCarRequest) *car_service.CreateCarInput {
	var imageUrl sql.NullString
	if dto.ImageUrl != nil {
		imageUrl = sql.NullString{
			String: *dto.ImageUrl,
			Valid:  true,
		}
	} else {
		imageUrl = sql.NullString{
			Valid: false,
		}
	}

	return &car_service.CreateCarInput{
		Brand:       dto.Brand,
		Model:       dto.Model,
		Year:        dto.Year,
		RentalPrice: dto.RentalPrice,
		ImageUrl:    imageUrl,
	}
}
