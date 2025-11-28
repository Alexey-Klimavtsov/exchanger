package user_model

import (
	"fmt"
	"github.com/asaipov/gorenda/internal/it/model/driver_license_model"
	"time"
)

type UserModel struct {
	ID             int64                                      `json:"id"`
	FirstName      string                                     `json:"firstName"`
	LastName       string                                     `json:"lastName"`
	Email          string                                     `json:"email"`
	Surname        *string                                    `json:"surname"`
	IsAdmin        bool                                       `json:"isAdmin"`
	RightsCategory []*driver_license_model.DriverLicenseModel `json:"rightsCategory"`
	Birthday       time.Time                                  `json:"birthday"`
	DeletedAt      time.Time                                  `json:"deletedAt"`
	CreatedAt      time.Time                                  `json:"createdAt"`
	UpdatedAt      *time.Time                                 `json:"updatedAt"`
}

func (m *UserModel) Validate() error {
	if m.FirstName == "" {
		return fmt.Errorf("firstName is required")
	}
	if m.LastName == "" {
		return fmt.Errorf("lastName is required")
	}
	if time.Now().Year()-m.Birthday.Year() < 18 {
		return fmt.Errorf("user must be at least 18 years old")
	}
	return nil
}
