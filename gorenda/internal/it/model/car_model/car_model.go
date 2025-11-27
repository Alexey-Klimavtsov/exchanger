package car_model

import (
	"database/sql"
	"fmt"
	"time"
)

type CarModel struct {
	ID          int64          `json:"id"`
	Brand       string         `json:"brand"`
	Model       string         `json:"model"`
	Year        time.Time      `json:"year"`
	RentalPrice int64          `json:"rentalPrice"`
	ImageUrl    sql.NullString `json:"imageUrl"`
	CreatedAt   time.Time      `json:"createdAt"`
	UpdatedAt   sql.NullTime   `json:"updatedAt"`
}

func (m *CarModel) Validate() error {
	if m.Brand == "" {
		return fmt.Errorf("brand is required")
	}
	if m.Model == "" {
		return fmt.Errorf("model is required")
	}
	if m.Year.Year() < 1900 {
		return fmt.Errorf("car is too old")
	}
	if m.RentalPrice <= 0 {
		return fmt.Errorf("negative price")
	}
	return nil
}
