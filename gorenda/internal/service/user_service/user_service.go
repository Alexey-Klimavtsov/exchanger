package user_service

import (
	"context"
	"fmt"
	"github.com/asaipov/gorenda/internal/it/model/user_model"
	"github.com/asaipov/gorenda/internal/service/driver_license_service"
	"time"
)

type CreateUserInput struct {
	FirstName string
	LastName  string
	Email     string
	Surname   *string
	Birthday  time.Time
}

func userInputToModel(in *CreateUserInput) (*user_model.UserModel, error) {
	user := &user_model.UserModel{
		FirstName: in.FirstName,
		LastName:  in.LastName,
		Email:     in.Email,
		Surname:   in.Surname,
		Birthday:  in.Birthday,
	}

	return user, user.Validate()
}

type UserRepo interface {
	CreateUser(ctx context.Context, u *user_model.UserModel) (*user_model.UserModel, error)
	UpdateUser(ctx context.Context, u *user_model.UserModel, id int64) (*user_model.UserModel, error)
	GetUserById(ctx context.Context, id int64) (*user_model.UserModel, error)
	DeleteUser(ctx context.Context, id int64) (int64, error)
	Exists(ctx context.Context, id int64) (bool, error)
}

type UserService interface {
	CreateUser(ctx context.Context, in *CreateUserInput) (*user_model.UserModel, error)
	UpdateUser(ctx context.Context, in *CreateUserInput, id int64) (*user_model.UserModel, error)
	GetUserById(ctx context.Context, id int64) (*user_model.UserModel, error)
	DeleteUser(ctx context.Context, id int64) (int64, error)
}

type userService struct {
	repo              UserRepo
	driverLicenseRepo driver_license_service.DriverLicenseRepo
}

func NewUserService(r UserRepo, dlR driver_license_service.DriverLicenseRepo) UserService {
	return &userService{repo: r, driverLicenseRepo: dlR}
}

func (s *userService) CreateUser(ctx context.Context, in *CreateUserInput) (*user_model.UserModel, error) {
	user, err := userInputToModel(in)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrInvalidInput, err)
	}

	return s.repo.CreateUser(ctx, user)
}

func (s *userService) UpdateUser(ctx context.Context, in *CreateUserInput, id int64) (*user_model.UserModel, error) {
	if id <= 0 {
		return nil, fmt.Errorf("%w: id is negative %v", ErrInvalidInput, id)
	}

	user, err := userInputToModel(in)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrInvalidInput, err)
	}

	return s.repo.UpdateUser(ctx, user, id)
}

func (s *userService) GetUserById(ctx context.Context, id int64) (*user_model.UserModel, error) {
	if id <= 0 {
		return nil, fmt.Errorf("%w: id is negative %v", ErrInvalidInput, id)
	}
	user, err := s.repo.GetUserById(ctx, id)
	if err != nil {
		return nil, err
	}
	licenses, getErr := s.driverLicenseRepo.GetLicensesByUserId(ctx, user.ID)
	if getErr != nil {
		return nil, getErr
	}

	user.RightsCategory = licenses

	return user, nil
}

func (s *userService) DeleteUser(ctx context.Context, id int64) (int64, error) {
	if id <= 0 {
		return 0, fmt.Errorf("%w: id is negative %v", ErrInvalidInput, id)
	}

	return s.repo.DeleteUser(ctx, id)
}
