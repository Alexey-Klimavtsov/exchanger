package car_model_test

import (
	"github.com/asaipov/gorenda/internal/it/model/car_model"
	"testing"
	"time"
)

func strPtr(s string) *string { return &s }

type TestsTable struct {
	needErr  bool
	value    car_model.CarModel
	testName string
}

func TestCarModel_Validate(t *testing.T) {
	tests := []TestsTable{
		{
			needErr: true,
			value: car_model.CarModel{
				Brand:       "",
				Model:       "X5",
				Year:        time.Now(),
				RentalPrice: 100,
				ImageUrl:    strPtr("url"),
			},
			testName: "Empty brand",
		},
		{
			needErr: true,
			value: car_model.CarModel{
				Brand:       "BMW",
				Model:       "",
				Year:        time.Now(),
				RentalPrice: 100,
			},
			testName: "Empty model",
		},
		{
			needErr: true,
			value: car_model.CarModel{
				Brand:       "BMW",
				Model:       "X5",
				Year:        time.Date(1800, 1, 1, 0, 0, 0, 0, time.UTC),
				RentalPrice: 100,
			},
			testName: "Year < 1900",
		},
		{
			needErr: true,
			value: car_model.CarModel{
				Brand:       "BMW",
				Model:       "X5",
				Year:        time.Now(),
				RentalPrice: -100,
			},
			testName: "Negative price",
		},
		{
			needErr: false,
			value: car_model.CarModel{
				Brand:       "BMW",
				Model:       "X5",
				Year:        time.Now(),
				RentalPrice: 100,
				ImageUrl:    nil,
			},
			testName: "Valid car",
		},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			err := tt.value.Validate()

			if tt.needErr && err == nil {
				t.Errorf("expected error but got nil")
			}

			if !tt.needErr && err != nil {
				t.Errorf("unexpected error: %v", err)
			}
		})
	}
}
