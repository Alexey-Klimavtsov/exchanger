package booking_model

import (
	"fmt"
	"time"
)

type BookingModel struct {
	ID          int64     `json:"id"`
	UserID      int64     `json:"userId"`
	CarID       int64     `json:"carId"`
	DateFrom    time.Time `json:"dateFrom"`
	DateTo      time.Time `json:"dateTo"`
	PricePerDay int64     `json:"pricePerDay"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"createdAt"`
}

func (m *BookingModel) Validate() error {

	if m.UserID <= 0 {
		return fmt.Errorf("userID is required")
	}
	if m.CarID <= 0 {
		return fmt.Errorf("carID is required")
	}
	if m.DateFrom.Before(time.Now().AddDate(0, 0, -1)) {
		return fmt.Errorf("cannot book in the past")
	}
	if m.DateTo.Before(m.DateFrom) {
		return fmt.Errorf("dateTo must be after dateFrom")
	}

	if m.DateTo.Sub(m.DateFrom).Hours()/24 > 28 {
		return fmt.Errorf("cannot book more than 28 days")
	}

	return nil
}
