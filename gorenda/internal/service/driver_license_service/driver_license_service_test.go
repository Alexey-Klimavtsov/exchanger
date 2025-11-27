package driver_license_service_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/asaipov/gorenda/internal/it/model/driver_license_model"
	"github.com/asaipov/gorenda/internal/service/driver_license_service"
)

type fakeDLRepo struct {
	licenses []*driver_license_model.DriverLicenseModel
}

func (r *fakeDLRepo) CreateLicense(ctx context.Context, dl *driver_license_model.DriverLicenseModel) (*driver_license_model.DriverLicenseModel, error) {
	dl.ID = int64(len(r.licenses) + 1)
	r.licenses = append(r.licenses, dl)
	return dl, nil
}

func (r *fakeDLRepo) UpdateLicense(ctx context.Context, dl *driver_license_model.DriverLicenseModel, id int64) (*driver_license_model.DriverLicenseModel, error) {
	if id <= 0 {
		return nil, errors.New("not found")
	}
	return dl, nil
}

func TestDriverLicenseService_CreateLicense(t *testing.T) {

	repo := &fakeDLRepo{}
	s := driver_license_service.NewDriverLicenseService(repo)

	now := time.Now().UTC()

	tests := []struct {
		name    string
		input   *driver_license_service.DriverLicenseInput
		wantErr bool
	}{
		{
			name: "empty number",
			input: &driver_license_service.DriverLicenseInput{
				UserID:    1,
				Number:    "",
				IssuedAt:  now.AddDate(-1, 0, 0),
				ExpiresAt: now.AddDate(1, 0, 0),
			},
			wantErr: true,
		},
		{
			name: "empty issuedAt",
			input: &driver_license_service.DriverLicenseInput{
				UserID:    1,
				Number:    "12345",
				IssuedAt:  time.Time{},
				ExpiresAt: now.AddDate(1, 0, 0),
			},
			wantErr: true,
		},
		{
			name: "empty expiresAt",
			input: &driver_license_service.DriverLicenseInput{
				UserID:    1,
				Number:    "12345",
				IssuedAt:  now.AddDate(-1, 0, 0),
				ExpiresAt: time.Time{},
			},
			wantErr: true,
		},
		{
			name: "expiresAt before issuedAt",
			input: &driver_license_service.DriverLicenseInput{
				UserID:    1,
				Number:    "12345",
				IssuedAt:  now,
				ExpiresAt: now.AddDate(-1, 0, 0),
			},
			wantErr: true,
		},
		{
			name: "issuedAt in the future",
			input: &driver_license_service.DriverLicenseInput{
				UserID:    1,
				Number:    "12345",
				IssuedAt:  now.AddDate(0, 0, 1),
				ExpiresAt: now.AddDate(1, 0, 0),
			},
			wantErr: true,
		},
		{
			name: "valid input",
			input: &driver_license_service.DriverLicenseInput{
				UserID:    1,
				Number:    "98765",
				IssuedAt:  now.AddDate(-2, 0, 0),
				ExpiresAt: now.AddDate(2, 0, 0),
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := s.CreateLicense(context.Background(), tt.input)

			if tt.wantErr && err == nil {
				t.Errorf("expected error, got nil")
			}
			if !tt.wantErr && err != nil {
				t.Errorf("unexpected error: %v", err)
			}
		})
	}
}

func TestDriverLicenseService_UpdateLicense(t *testing.T) {

	repo := &fakeDLRepo{}
	s := driver_license_service.NewDriverLicenseService(repo)

	now := time.Now().UTC()

	tests := []struct {
		name    string
		input   *driver_license_service.DriverLicenseInput
		id      int64
		wantErr bool
	}{
		{
			name: "empty id",
			input: &driver_license_service.DriverLicenseInput{
				UserID:    1,
				Number:    "12345",
				IssuedAt:  now,
				ExpiresAt: now.AddDate(1, 0, 0),
			},
			id:      0,
			wantErr: true,
		},
		{
			name: "empty number",
			input: &driver_license_service.DriverLicenseInput{
				UserID:    1,
				Number:    "",
				IssuedAt:  now.AddDate(-1, 0, 0),
				ExpiresAt: now.AddDate(1, 0, 0),
			},
			id:      1,
			wantErr: true,
		},
		{
			name: "expiresAt before issuedAt",
			input: &driver_license_service.DriverLicenseInput{
				UserID:    1,
				Number:    "12345",
				IssuedAt:  now,
				ExpiresAt: now.AddDate(-1, 0, 0),
			},
			id:      1,
			wantErr: true,
		},
		{
			name: "valid update",
			input: &driver_license_service.DriverLicenseInput{
				UserID:    1,
				Number:    "99999",
				IssuedAt:  now.AddDate(-2, 0, 0),
				ExpiresAt: now.AddDate(3, 0, 0),
			},
			id:      1,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			_, err := s.UpdateLicense(context.Background(), tt.input, tt.id)

			if tt.wantErr && err == nil {
				t.Errorf("expected error, got nil")
			}

			if !tt.wantErr && err != nil {
				t.Errorf("unexpected error: %v", err)
			}
		})
	}
}
