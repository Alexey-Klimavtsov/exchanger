package user_model_test

import (
	"testing"
	"time"

	"github.com/asaipov/gorenda/internal/it/model/user_model" // поправь путь при необходимости
)

type TestsTable struct {
	name    string
	value   user_model.UserModel
	wantErr bool
}

func TestUserModel_Validate(t *testing.T) {
	now := time.Now()

	tests := []TestsTable{
		{
			name:    "Empty first name",
			wantErr: true,
			value: user_model.UserModel{
				FirstName: "",
				LastName:  "Иванов",
				Birthday:  now.AddDate(-20, 0, 0),
			},
		},
		{
			name:    "Empty last name",
			wantErr: true,
			value: user_model.UserModel{
				FirstName: "Иван",
				LastName:  "",
				Birthday:  now.AddDate(-20, 0, 0),
			},
		},
		{
			name:    "Birthday younger than 18",
			wantErr: true,
			value: user_model.UserModel{
				FirstName: "Иван",
				LastName:  "Иванов",
				Birthday:  now.AddDate(-17, 0, 0),
			},
		},
		{
			name:    "Valid user",
			wantErr: false,
			value: user_model.UserModel{
				FirstName: "Иван",
				LastName:  "Иванов",
				Birthday:  now.AddDate(-25, 0, 0),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.value.Validate()

			if tt.wantErr && err == nil {
				t.Errorf("expected error, got nil")
			}
			if !tt.wantErr && err != nil {
				t.Errorf("did not expect error, got %v", err)
			}
		})
	}
}
