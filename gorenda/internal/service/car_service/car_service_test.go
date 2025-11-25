package car_service

import (
	"context"
	"database/sql"
	"errors"
	"testing"
	"time"

	"github.com/asaipov/gorenda/internal/it/model/car_model"
)

type fakeCarRepo struct {
	cars []*car_model.CarModel
}

func (r *fakeCarRepo) CreateNewCar(ctx context.Context, car *car_model.CarModel) (*car_model.CarModel, error) {
	car.ID = int64(len(r.cars) + 1)
	r.cars = append(r.cars, car)
	return car, nil
}

func (r *fakeCarRepo) UpdateCar(ctx context.Context, car *car_model.CarModel, id int64) (*car_model.CarModel, error) {
	if id <= 0 {
		return nil, errors.New("not found")
	}
	return car, nil
}

func (r *fakeCarRepo) GetCarById(ctx context.Context, id int64) (*car_model.CarModel, error) {
	if id <= 0 || int(id) > len(r.cars) {
		return nil, errors.New("not found")
	}
	return r.cars[id-1], nil
}

func (r *fakeCarRepo) DeleteCar(ctx context.Context, id int64) (int64, error) {
	if id <= 0 || int(id) > len(r.cars) {
		return 0, errors.New("not found")
	}
	return id, nil
}

func TestCarService_CreateNewCar(t *testing.T) {
	repo := &fakeCarRepo{}
	s := NewCarService(repo)

	tests := []struct {
		name    string
		input   *CreateCarInput
		wantErr bool
	}{
		{
			name: "empty brand",
			input: &CreateCarInput{
				Brand:       "",
				Model:       "X5",
				Year:        time.Now(),
				RentalPrice: 100,
				ImageUrl:    sql.NullString{String: "url", Valid: true},
			},
			wantErr: true,
		},
		{
			name: "empty model",
			input: &CreateCarInput{
				Brand:       "BMW",
				Model:       "",
				Year:        time.Now(),
				RentalPrice: 100,
				ImageUrl:    sql.NullString{String: "url", Valid: true},
			},
			wantErr: true,
		},
		{
			name: "year < 1900",
			input: &CreateCarInput{
				Brand:       "BMW",
				Model:       "X5",
				Year:        time.Date(1800, 1, 1, 0, 0, 0, 0, time.UTC),
				RentalPrice: 100,
				ImageUrl:    sql.NullString{String: "url", Valid: true},
			},
			wantErr: true,
		},
		{
			name: "negative price",
			input: &CreateCarInput{
				Brand:       "BMW",
				Model:       "X5",
				Year:        time.Now(),
				RentalPrice: -10,
				ImageUrl:    sql.NullString{String: "url", Valid: true},
			},
			wantErr: true,
		},
		{
			name: "valid input",
			input: &CreateCarInput{
				Brand:       "BMW",
				Model:       "X5",
				Year:        time.Now(),
				RentalPrice: 100,
				ImageUrl:    sql.NullString{String: "url", Valid: true},
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := s.CreateNewCar(context.Background(), tt.input)

			if tt.wantErr && err == nil {
				t.Errorf("expected error, got nil")
			}
			if !tt.wantErr && err != nil {
				t.Errorf("unexpected error: %v", err)
			}
		})
	}
}

func TestCarService_UpdateCar(t *testing.T) {
	repo := &fakeCarRepo{}
	s := NewCarService(repo)

	tests := []struct {
		name    string
		input   *CreateCarInput
		wantErr bool
		id      int64
	}{
		{
			name: "empty brand",
			input: &CreateCarInput{
				Brand:       "",
				Model:       "X5",
				Year:        time.Now(),
				RentalPrice: 100,
				ImageUrl:    sql.NullString{String: "url", Valid: true},
			},

			wantErr: true,
			id:      1,
		},
		{
			name: "empty model",
			input: &CreateCarInput{
				Brand:       "BMW",
				Model:       "",
				Year:        time.Now(),
				RentalPrice: 100,
				ImageUrl:    sql.NullString{String: "url", Valid: true},
			},
			wantErr: true,
			id:      1,
		},
		{
			name: "year < 1900",
			input: &CreateCarInput{
				Brand:       "BMW",
				Model:       "X5",
				Year:        time.Date(1800, 1, 1, 0, 0, 0, 0, time.UTC),
				RentalPrice: 100,
				ImageUrl:    sql.NullString{String: "url", Valid: true},
			},
			wantErr: true,
			id:      1,
		},
		{
			name: "negative price",
			input: &CreateCarInput{
				Brand:       "BMW",
				Model:       "X5",
				Year:        time.Now(),
				RentalPrice: -10,
				ImageUrl:    sql.NullString{String: "url", Valid: true},
			},
			wantErr: true,
			id:      1,
		},
		{
			name: "valid input",
			input: &CreateCarInput{
				Brand:       "BMW",
				Model:       "X5",
				Year:        time.Now(),
				RentalPrice: 100,
				ImageUrl:    sql.NullString{String: "url", Valid: true},
			},
			wantErr: false,
			id:      1,
		},
		{
			name: "empty id",
			input: &CreateCarInput{
				Brand:       "BMW",
				Model:       "X5",
				Year:        time.Now(),
				RentalPrice: 100,
				ImageUrl:    sql.NullString{String: "url", Valid: true},
			},
			wantErr: true,
			id:      -1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := s.UpdateCar(context.Background(), tt.input, tt.id)

			if tt.wantErr && err == nil {
				t.Errorf("expected error, got nil")
			}
			if !tt.wantErr && err != nil {
				t.Errorf("unexpected error: %v", err)
			}
		})
	}
}

func TestCarService_GetCarById(t *testing.T) {
	repo := &fakeCarRepo{}
	repo.cars = append(repo.cars, &car_model.CarModel{ID: 1, Brand: "BMW"})

	s := NewCarService(repo)

	tests := []struct {
		name    string
		id      int64
		wantErr bool
	}{
		{"valid id", 1, false},
		{"invalid id", 0, true},
		{"not found", 10, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := s.GetCarById(context.Background(), tt.id)

			if tt.wantErr && err == nil {
				t.Errorf("expected error, got nil")
			}
			if !tt.wantErr && err != nil {
				t.Errorf("unexpected error: %v", err)
			}
		})
	}
}

func TestCarService_DeleteCar(t *testing.T) {
	repo := &fakeCarRepo{}
	repo.cars = append(repo.cars, &car_model.CarModel{ID: 1})

	s := NewCarService(repo)

	tests := []struct {
		name    string
		id      int64
		wantErr bool
	}{
		{"valid id", 1, false},
		{"invalid id", 0, true},
		{"not found", 10, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := s.DeleteCar(context.Background(), tt.id)

			if tt.wantErr && err == nil {
				t.Errorf("expected error, got nil")
			}
			if !tt.wantErr && err != nil {
				t.Errorf("unexpected error: %v", err)
			}
		})
	}
}
