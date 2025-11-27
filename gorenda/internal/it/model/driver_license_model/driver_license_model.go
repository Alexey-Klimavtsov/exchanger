package driver_license_model

import (
	"database/sql"
	"fmt"
	"time"
)

type DriverLicenseModel struct {
	ID        int64        `json:"id"`
	UserID    int64        `json:"userId"`
	Number    string       `json:"number"`
	IssuedAt  time.Time    `json:"issuedAt"`
	ExpiresAt time.Time    `json:"expiresAt"`
	CreatedAt time.Time    `json:"createdAt"`
	UpdatedAt sql.NullTime `json:"updatedAt"`
}

func (m *DriverLicenseModel) Validate() error {
	if m.UserID <= 0 {
		return fmt.Errorf("id is required")
	}
	if m.Number == "" {
		return fmt.Errorf("license number is required")
	}

	if m.IssuedAt.IsZero() {
		return fmt.Errorf("issuedAt is required")
	}

	if m.ExpiresAt.IsZero() {
		return fmt.Errorf("expiresAt is required")
	}

	if !m.ExpiresAt.After(m.IssuedAt) {
		return fmt.Errorf("expiresAt must be after issuedAt")
	}

	if m.IssuedAt.After(time.Now()) {
		return fmt.Errorf("issuedAt cannot be in the future")
	}

	return nil
}
