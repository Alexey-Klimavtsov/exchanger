package car_service

import (
	"context"
	"github.com/asaipov/gorenda/internal/it/model/car_model"
)

type CarRepo interface {
	CreateNewCar(ctx context.Context, car *car_model.CarModel) (*car_model.CarModel, error)
	UpdateCar(ctx context.Context, car *car_model.CarModel, id int64) (*car_model.CarModel, error)
	GetCarById(ctx context.Context, id int64) (*car_model.CarModel, error)
	DeleteCar(ctx context.Context, id int64) (deletedId int64, err error)
	Exists(ctx context.Context, id int64) (bool, error)
}
