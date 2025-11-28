package booking_service

import (
	"context"
	"testing"
	"time"

	"github.com/asaipov/gorenda/internal/it/model/booking_model"
)

type fakeBookingRepo struct {
	bookings []*booking_model.BookingModel
}

func (r *fakeBookingRepo) CreateBooking(ctx context.Context, b *booking_model.BookingModel) (*booking_model.BookingModel, error) {
	b.ID = int64(len(r.bookings) + 1)
	r.bookings = append(r.bookings, b)
	return b, nil
}

func (r *fakeBookingRepo) GetBookings(ctx context.Context) ([]*booking_model.BookingModel, error) {
	return r.bookings, nil
}

func (r *fakeBookingRepo) IsCarAvailable(ctx context.Context, carID int64, from, to time.Time) (bool, error) {

	for _, b := range r.bookings {
		if b.CarID == carID {
			if !(to.Before(b.DateFrom) || from.After(b.DateTo)) {
				return false, nil
			}
		}
	}

	return true, nil
}

func TestBookingService_CreateBooking(t *testing.T) {
	repo := &fakeBookingRepo{}
	s := NewBookingService(repo)

	now := time.Now().AddDate(0, 0, 1)

	tests := []struct {
		name    string
		input   *CreateBookingInput
		wantErr bool
	}{
		{
			name: "date in past",
			input: &CreateBookingInput{
				UserID:      1,
				CarID:       1,
				DateFrom:    time.Now().AddDate(0, 0, -5),
				DateTo:      time.Now().AddDate(0, 0, -3),
				PricePerDay: 100,
			},
			wantErr: true,
		},
		{
			name: "more than 28 days",
			input: &CreateBookingInput{
				UserID:      1,
				CarID:       1,
				DateFrom:    now,
				DateTo:      now.AddDate(0, 0, 40),
				PricePerDay: 100,
			},
			wantErr: true,
		},
		{
			name: "car unavailable",
			input: &CreateBookingInput{
				UserID:      1,
				CarID:       1,
				DateFrom:    now,
				DateTo:      now.AddDate(0, 0, 2),
				PricePerDay: 100,
			},
			wantErr: true,
		},
	}

	_, _ = s.CreateBooking(context.Background(), tests[2].input)

	unavailable := CreateBookingInput{
		UserID:      2,
		CarID:       1,
		DateFrom:    now,
		DateTo:      now.AddDate(0, 0, 1),
		PricePerDay: 100,
	}

	t.Run("unavailable car case", func(t *testing.T) {
		_, err := s.CreateBooking(context.Background(), &unavailable)
		if err == nil {
			t.Errorf("expected error, got nil")
		}
	})

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			_, err := s.CreateBooking(context.Background(), tt.input)
			if tt.wantErr && err == nil {
				t.Errorf("expected error, got nil")
			}
			if !tt.wantErr && err != nil {
				t.Errorf("unexpected error: %v", err)
			}
		})
	}
}

func TestBookingService_GetAllBookings(t *testing.T) {
	repo := &fakeBookingRepo{}
	s := NewBookingService(repo)

	now := time.Now().AddDate(0, 0, 1)

	_, _ = s.CreateBooking(context.Background(), &CreateBookingInput{
		UserID:      1,
		CarID:       2,
		DateFrom:    now,
		DateTo:      now.AddDate(0, 0, 2),
		PricePerDay: 200,
	})

	list, err := s.GetAllBookings(context.Background())
	if err != nil {
		t.Errorf("unexpected error %v", err)
	}

	if len(list) != 1 {
		t.Errorf("expected 1 booking, got %d", len(list))
	}
}
