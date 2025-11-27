package user_service

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/asaipov/gorenda/internal/it/model/user_model"
)

type fakeUserRepo struct {
	users []*user_model.UserModel
}

func (r *fakeUserRepo) CreateUser(ctx context.Context, u *user_model.UserModel) (*user_model.UserModel, error) {
	u.ID = int64(len(r.users) + 1)
	r.users = append(r.users, u)
	return u, nil
}

func (r *fakeUserRepo) UpdateUser(ctx context.Context, u *user_model.UserModel, id int64) (*user_model.UserModel, error) {
	if id <= 0 {
		return nil, errors.New("not found")
	}
	return u, nil
}

func (r *fakeUserRepo) GetUserById(ctx context.Context, id int64) (*user_model.UserModel, error) {
	if id <= 0 || int(id) > len(r.users) {
		return nil, errors.New("not found")
	}
	return r.users[id-1], nil
}

func (r *fakeUserRepo) DeleteUser(ctx context.Context, id int64) (int64, error) {
	if id <= 0 || int(id) > len(r.users) {
		return 0, errors.New("not found")
	}
	return id, nil
}

func TestUserService_CreateUser(t *testing.T) {
	repo := &fakeUserRepo{}
	s := NewUserService(repo)

	now := time.Now().AddDate(-20, 0, 0) // age = 20

	tests := []struct {
		name    string
		input   *CreateUserInput
		wantErr bool
	}{
		{"empty firstName", &CreateUserInput{FirstName: "", LastName: "Ivanov", Birthday: now}, true},
		{"empty lastName", &CreateUserInput{FirstName: "Ivan", LastName: "", Birthday: now}, true},
		{"under 18", &CreateUserInput{FirstName: "Ivan", LastName: "Ivanov", Birthday: time.Now()}, true},
		{"valid user", &CreateUserInput{FirstName: "Ivan", LastName: "Ivanov", Birthday: now}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := s.CreateUser(context.Background(), tt.input)

			if tt.wantErr && err == nil {
				t.Errorf("expected error, got nil")
			}
			if !tt.wantErr && err != nil {
				t.Errorf("unexpected error: %v", err)
			}
		})
	}
}

func TestUserService_UpdateUser(t *testing.T) {
	repo := &fakeUserRepo{}
	s := NewUserService(repo)
	validBirthday := time.Now().AddDate(-25, 0, 0)

	tests := []struct {
		name    string
		id      int64
		input   *CreateUserInput
		wantErr bool
	}{
		{"negative id", -1, &CreateUserInput{FirstName: "Ivan", LastName: "Ivanov", Birthday: validBirthday}, true},
		{"empty firstName", 1, &CreateUserInput{FirstName: "", LastName: "Ivanov", Birthday: validBirthday}, true},
		{"empty lastName", 1, &CreateUserInput{FirstName: "Ivan", LastName: "", Birthday: validBirthday}, true},
		{"under 18", 1, &CreateUserInput{FirstName: "Ivan", LastName: "Ivanov", Birthday: time.Now()}, true},
		{"valid update", 1, &CreateUserInput{FirstName: "Ivan", LastName: "Ivanov", Birthday: validBirthday}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := s.UpdateUser(context.Background(), tt.input, tt.id)

			if tt.wantErr && err == nil {
				t.Errorf("expected error, got nil")
			}
			if !tt.wantErr && err != nil {
				t.Errorf("unexpected error: %v", err)
			}
		})
	}
}

func TestUserService_GetUserById(t *testing.T) {
	repo := &fakeUserRepo{}
	repo.users = append(repo.users, &user_model.UserModel{ID: 1, FirstName: "Ivan"})

	s := NewUserService(repo)

	tests := []struct {
		name    string
		id      int64
		wantErr bool
	}{
		{"valid id", 1, false},
		{"zero id", 0, true},
		{"not found", 10, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := s.GetUserById(context.Background(), tt.id)

			if tt.wantErr && err == nil {
				t.Errorf("expected error, got nil")
			}
			if !tt.wantErr && err != nil {
				t.Errorf("unexpected error: %v", err)
			}
		})
	}
}

func TestUserService_DeleteUser(t *testing.T) {
	repo := &fakeUserRepo{}
	repo.users = append(repo.users, &user_model.UserModel{ID: 1})

	s := NewUserService(repo)

	tests := []struct {
		name    string
		id      int64
		wantErr bool
	}{
		{"valid id", 1, false},
		{"zero id", 0, true},
		{"not found", 20, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := s.DeleteUser(context.Background(), tt.id)

			if tt.wantErr && err == nil {
				t.Errorf("expected error, got nil")
			}
			if !tt.wantErr && err != nil {
				t.Errorf("unexpected error: %v", err)
			}
		})
	}
}
