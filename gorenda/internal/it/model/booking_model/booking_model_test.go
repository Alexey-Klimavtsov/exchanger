package booking_model_test

import (
	"github.com/asaipov/gorenda/internal/it/model/booking_model"
	"testing"
	"time"
)

type TestCase struct {
	name    string
	value   booking_model.BookingModel
	wantErr bool
}

func TestBookingModel_Validate(t *testing.T) {
	now := time.Now().UTC()
	tomorrow := now.AddDate(0, 0, 1)
	afterWeek := now.AddDate(0, 0, 7)
	after30Days := now.AddDate(0, 0, 30)

	tests := []TestCase{
		{
			name: "empty userID",
			value: booking_model.BookingModel{
				UserID:      0,
				CarID:       1,
				DateFrom:    tomorrow,
				DateTo:      afterWeek,
				PricePerDay: 100,
			},
			wantErr: true,
		},
		{
			name: "empty carID",
			value: booking_model.BookingModel{
				UserID:      1,
				CarID:       0,
				DateFrom:    tomorrow,
				DateTo:      afterWeek,
				PricePerDay: 100,
			},
			wantErr: true,
		},
		{
			name: "dateFrom in past",
			value: booking_model.BookingModel{
				UserID:      1,
				CarID:       1,
				DateFrom:    now.AddDate(0, 0, -1),
				DateTo:      afterWeek,
				PricePerDay: 100,
			},
			wantErr: true,
		},
		{
			name: "dateTo before dateFrom",
			value: booking_model.BookingModel{
				UserID:      1,
				CarID:       1,
				DateFrom:    tomorrow,
				DateTo:      now,
				PricePerDay: 100,
			},
			wantErr: true,
		},
		{
			name: "booking longer than 28 days",
			value: booking_model.BookingModel{
				UserID:      1,
				CarID:       1,
				DateFrom:    tomorrow,
				DateTo:      after30Days,
				PricePerDay: 100,
			},
			wantErr: true,
		},
		{
			name: "valid booking",
			value: booking_model.BookingModel{
				UserID:      1,
				CarID:       1,
				DateFrom:    tomorrow,
				DateTo:      afterWeek,
				PricePerDay: 100,
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.value.Validate()

			if tt.wantErr && err == nil {
				t.Errorf("expected error, got nil")
			}
			if !tt.wantErr && err != nil {
				t.Errorf("unexpected error: %v", err)
			}
		})
	}
}
