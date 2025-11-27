package driver_license_model_test

import (
	"github.com/asaipov/gorenda/internal/it/model/driver_license_model"
	"testing"
	"time"
)

type TestsTable struct {
	needErr  bool
	value    driver_license_model.DriverLicenseModel
	testName string
}

func TestDriverLicenseModel_Validate(t *testing.T) {

	now := time.Now().UTC()

	tests := []TestsTable{
		{
			needErr:  true,
			testName: "Empty license number",
			value: driver_license_model.DriverLicenseModel{
				Number:    "",
				IssuedAt:  now.AddDate(-1, 0, 0),
				ExpiresAt: now.AddDate(1, 0, 0),
			},
		},
		{
			needErr:  true,
			testName: "Empty issuedAt",
			value: driver_license_model.DriverLicenseModel{
				Number:    "123456",
				IssuedAt:  time.Time{},
				ExpiresAt: now.AddDate(1, 0, 0),
			},
		},
		{
			needErr:  true,
			testName: "Empty expiresAt",
			value: driver_license_model.DriverLicenseModel{
				Number:    "123456",
				IssuedAt:  now.AddDate(-1, 0, 0),
				ExpiresAt: time.Time{},
			},
		},
		{
			needErr:  true,
			testName: "ExpiresAt before IssuedAt",
			value: driver_license_model.DriverLicenseModel{
				Number:    "123456",
				IssuedAt:  now,
				ExpiresAt: now.AddDate(-1, 0, 0),
			},
		},
		{
			needErr:  true,
			testName: "IssuedAt in the future",
			value: driver_license_model.DriverLicenseModel{
				Number:    "123456",
				IssuedAt:  now.AddDate(0, 0, 1),
				ExpiresAt: now.AddDate(1, 0, 0),
			},
		},
		{
			needErr:  false,
			testName: "Valid license",
			value: driver_license_model.DriverLicenseModel{
				Number:    "987654",
				IssuedAt:  now.AddDate(-2, 0, 0),
				ExpiresAt: now.AddDate(2, 0, 0),
			},
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
