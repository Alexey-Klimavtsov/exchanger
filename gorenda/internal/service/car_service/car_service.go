package car_service

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/asaipov/gorenda/internal/it/model/car_model"
	"time"
)

type CreateCarInput struct {
	Brand       string
	Model       string
	Year        time.Time
	RentalPrice int64
	ImageUrl    sql.NullString
}

func carToModel(input *CreateCarInput) (*car_model.CarModel, error) {
	car := &car_model.CarModel{
		Brand:       input.Brand,
		Model:       input.Model,
		Year:        input.Year,
		RentalPrice: input.RentalPrice,
		ImageUrl:    input.ImageUrl,
	}
	return car, car.Validate()
}

type CarService interface {
	CreateNewCar(ctx context.Context, car *CreateCarInput) (*car_model.CarModel, error)
	UpdateCar(ctx context.Context, car *CreateCarInput, id int64) (*car_model.CarModel, error)
	GetCarById(ctx context.Context, id int64) (*car_model.CarModel, error)
	DeleteCar(ctx context.Context, id int64) (deletedId int64, err error)
}

type carService struct {
	repo CarRepo
}

func NewCarService(r CarRepo) CarService {
	return &carService{repo: r}
}

func (s *carService) CreateNewCar(ctx context.Context, car *CreateCarInput) (*car_model.CarModel, error) {
	c, err := carToModel(car)

	if err != nil {
		return nil, fmt.Errorf("%w", ErrInvalidInput)
	}

	return s.repo.CreateNewCar(ctx, c)
}

func (s *carService) UpdateCar(ctx context.Context, car *CreateCarInput, id int64) (*car_model.CarModel, error) {
	c, err := carToModel(car)

	if err != nil {
		return nil, fmt.Errorf("%w", ErrInvalidInput)
	}

	return s.repo.UpdateCar(ctx, c, id)
}

func (s *carService) GetCarById(ctx context.Context, id int64) (*car_model.CarModel, error) {
	if id <= 0 {
		return nil, fmt.Errorf("%w: id is negative", ErrInvalidInput)
	}

	return s.repo.GetCarById(ctx, id)
}

func (s *carService) DeleteCar(ctx context.Context, id int64) (deletedId int64, err error) {
	if id <= 0 {
		return 0, fmt.Errorf("%w: id is negative", ErrInvalidInput)
	}

	return s.repo.DeleteCar(ctx, id)
}
