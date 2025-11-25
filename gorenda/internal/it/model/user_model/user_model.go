package user_model

import (
	"database/sql"
	"fmt"
	"time"
)

type UserModel struct {
	ID             int64          `json:"id"`
	FirstName      string         `json:"firstName"`
	LastName       string         `json:"lastName"`
	Surname        sql.NullString `json:"surname"`
	IsAdmin        bool           `json:"isAdmin"`
	RightsCategory sql.NullString `json:"rightsCategory"`
	Birthday       time.Time      `json:"birthday"`
	CreatedAt      time.Time      `json:"createdAt"`
	UpdatedAt      time.Time      `json:"updatedAt"`
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
