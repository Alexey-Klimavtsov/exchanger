package car_model_test

import (
	"database/sql"
	car_model "github.com/asaipov/gorenda/internal/it/model/car_model"
	"testing"
	"time"
)

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
				Model:       "Машина",
				Year:        time.Time{},
				RentalPrice: 2000,
				ImageUrl:    sql.NullString{String: "ссылка", Valid: true},
			},
			testName: "Empty brand",
		},
		{
			needErr: true,
			value: car_model.CarModel{
				Brand:       "Машина",
				Model:       "",
				Year:        time.Time{},
				RentalPrice: 2000,
				ImageUrl:    sql.NullString{String: "ссылка", Valid: true},
			},
			testName: "Empty model",
		},
		{
			needErr: true,
			value: car_model.CarModel{
				Brand:       "",
				Model:       "",
				Year:        time.Time{},
				RentalPrice: 0,
				ImageUrl:    sql.NullString{Valid: false}, // это NULL
			},
			testName: "All empty",
		},
		{
			needErr: true,
			value: car_model.CarModel{
				Brand:       "BMW",
				Model:       "X5",
				Year:        time.Date(1800, 1, 1, 0, 0, 0, 0, time.UTC),
				RentalPrice: 3000,
				ImageUrl:    sql.NullString{String: "url", Valid: true},
			},
			testName: "Year < 1900",
		},
		{
			needErr: true,
			value: car_model.CarModel{
				Brand:       "BMW",
				Model:       "X5",
				Year:        time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
				RentalPrice: -500,
				ImageUrl:    sql.NullString{String: "url", Valid: true},
			},
			testName: "Negative price",
		},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			err := tt.value.Validate()
			if tt.needErr && err == nil {
				t.Errorf("expected error but got nil")
			}
			if !tt.needErr && err != nil {
				t.Errorf("did not expect error, but got: %v", err)
			}
		})
	}
}
