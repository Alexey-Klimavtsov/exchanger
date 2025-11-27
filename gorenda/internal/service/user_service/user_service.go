package user_service

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/asaipov/gorenda/internal/it/model/user_model"
	"time"
)

type CreateUserInput struct {
	FirstName string
	LastName  string
	Email     string
	Surname   sql.NullString
	Birthday  time.Time
}

func userInputToModel(input *CreateUserInput) (*user_model.UserModel, error) {
	user := &user_model.UserModel{
		FirstName: input.FirstName,
		LastName:  input.LastName,
		Email:     input.Email,
		Surname:   input.Surname,
		Birthday:  input.Birthday,
	}

	return user, user.Validate()
}

func sqlNull(s string) sql.NullString {
	if s == "" {
		return sql.NullString{Valid: false}
	}
	return sql.NullString{String: s, Valid: true}
}

type UserRepo interface {
	CreateUser(ctx context.Context, u *user_model.UserModel) (*user_model.UserModel, error)
	UpdateUser(ctx context.Context, u *user_model.UserModel, id int64) (*user_model.UserModel, error)
	GetUserById(ctx context.Context, id int64) (*user_model.UserModel, error)
	DeleteUser(ctx context.Context, id int64) (int64, error)
}

type UserService interface {
	CreateUser(ctx context.Context, in *CreateUserInput) (*user_model.UserModel, error)
	UpdateUser(ctx context.Context, in *CreateUserInput, id int64) (*user_model.UserModel, error)
	GetUserById(ctx context.Context, id int64) (*user_model.UserModel, error)
	DeleteUser(ctx context.Context, id int64) (int64, error)
}

type userService struct {
	repo UserRepo
}

func NewUserService(r UserRepo) UserService {
	return &userService{repo: r}
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

	return s.repo.GetUserById(ctx, id)
}

func (s *userService) DeleteUser(ctx context.Context, id int64) (int64, error) {
	if id <= 0 {
		return 0, fmt.Errorf("%w: id is negative %v", ErrInvalidInput, id)
	}

	return s.repo.DeleteUser(ctx, id)
}
