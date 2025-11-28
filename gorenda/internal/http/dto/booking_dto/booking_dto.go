package booking_dto

import (
	"github.com/asaipov/gorenda/internal/it/model/booking_model"
	"github.com/asaipov/gorenda/internal/service/booking_service"
	"time"
)

type BookingRequestDto struct {
	UserID      int64     `json:"userId"`
	CarID       int64     `json:"carId"`
	DateFrom    time.Time `json:"dateFrom"`
	DateTo      time.Time `json:"dateTo"`
	PricePerDay int64     `json:"pricePerDay"`
}

type BookingResponseDto struct {
	ID          int64     `json:"id"`
	UserID      int64     `json:"userId"`
	CarID       int64     `json:"carId"`
	DateFrom    time.Time `json:"dateFrom"`
	DateTo      time.Time `json:"dateTo"`
	PricePerDay int64     `json:"pricePerDay"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"createdAt"`
}

func DtoToInput(dto *BookingRequestDto) *booking_service.CreateBookingInput {
	return &booking_service.CreateBookingInput{
		UserID:      dto.UserID,
		CarID:       dto.CarID,
		DateFrom:    dto.DateFrom,
		DateTo:      dto.DateTo,
		PricePerDay: dto.PricePerDay,
	}
}

func BookingToResponseDto(b *booking_model.BookingModel) *BookingResponseDto {
	return &BookingResponseDto{
		ID:          b.ID,
		UserID:      b.UserID,
		CarID:       b.CarID,
		DateFrom:    b.DateFrom,
		DateTo:      b.DateTo,
		PricePerDay: b.PricePerDay,
		Status:      b.Status,
		CreatedAt:   b.CreatedAt,
	}
}
