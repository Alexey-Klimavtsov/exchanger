package booking_service

import (
	"context"
	"fmt"
	"github.com/asaipov/gorenda/internal/service/car_service"
	"github.com/asaipov/gorenda/internal/service/user_service"
	"time"

	"github.com/asaipov/gorenda/internal/it/model/booking_model"
)

var (
	ErrInvalidInput = fmt.Errorf("invalid input")
	ErrNotFound     = fmt.Errorf("booking not found")
	ErrUnavailable  = fmt.Errorf("car is not available for these dates")
	ErrCreate       = fmt.Errorf("cannot create booking")
	ErrReading      = fmt.Errorf("cannot read booking")
	ErrInternal     = fmt.Errorf("server booking internal error")
)

type CreateBookingInput struct {
	UserID      int64
	CarID       int64
	DateFrom    time.Time
	DateTo      time.Time
	PricePerDay int64
}

func inputToModel(in *CreateBookingInput) (*booking_model.BookingModel, error) {
	b := &booking_model.BookingModel{
		UserID:      in.UserID,
		CarID:       in.CarID,
		DateFrom:    in.DateFrom,
		DateTo:      in.DateTo,
		PricePerDay: in.PricePerDay,
		Status:      "done",
	}

	return b, b.Validate()
}

type BookingRepo interface {
	CreateBooking(ctx context.Context, b *booking_model.BookingModel) (*booking_model.BookingModel, error)
	GetBookings(ctx context.Context) ([]*booking_model.BookingModel, error)
	IsCarAvailable(ctx context.Context, carID int64, from, to time.Time) (bool, error)
}

type BookingService interface {
	CreateBooking(ctx context.Context, in *CreateBookingInput) (*booking_model.BookingModel, error)
	GetAllBookings(ctx context.Context) ([]*booking_model.BookingModel, error)
}

type bookingService struct {
	bookingRepo BookingRepo
	userRepo    user_service.UserRepo
	carRepo     car_service.CarRepo
}

func NewBookingService(bookingRepo BookingRepo, userRepo user_service.UserRepo, carRepo car_service.CarRepo) BookingService {
	return &bookingService{bookingRepo: bookingRepo, userRepo: userRepo, carRepo: carRepo}
}

func (s *bookingService) CreateBooking(ctx context.Context, in *CreateBookingInput) (*booking_model.BookingModel, error) {
	b, err := inputToModel(in)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrInvalidInput, err)
	}

	userNotExists, userExistsErr := s.userRepo.Exists(ctx, b.UserID)
	if userExistsErr != nil {
		return nil, userExistsErr
	}
	if userNotExists {
		return nil, user_service.ErrNotFound
	}

	carNotExists, carExistsErr := s.carRepo.Exists(ctx, b.CarID)
	if carExistsErr != nil {
		return nil, carExistsErr
	}
	if carNotExists {
		return nil, car_service.ErrNotFound
	}

	available, err := s.bookingRepo.IsCarAvailable(ctx, b.CarID, b.DateFrom, b.DateTo)
	if err != nil {
		return nil, ErrReading
	}
	if !available {
		return nil, ErrUnavailable
	}

	return s.bookingRepo.CreateBooking(ctx, b)
}

func (s *bookingService) GetAllBookings(ctx context.Context) ([]*booking_model.BookingModel, error) {
	list, err := s.bookingRepo.GetBookings(ctx)
	if err != nil {
		return nil, ErrReading
	}
	return list, nil
}
