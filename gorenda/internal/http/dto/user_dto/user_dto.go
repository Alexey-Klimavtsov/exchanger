package user_dto

import (
	"database/sql"
	"github.com/asaipov/gorenda/internal/it/model/user_model"
	"time"
)

type UserRequestDto struct {
	FirstName      string    `json:"firstName"      binding:"required"`
	LastName       string    `json:"lastName"       binding:"required"`
	Surname        *string   `json:"surname"`
	RightsCategory *string   `json:"rightsCategory"`
	Birthday       time.Time `json:"birthday"       binding:"required"`
}

type UserResponseDto struct {
	ID             int64     `json:"id"`
	FirstName      string    `json:"firstName"`
	LastName       string    `json:"lastName"`
	Surname        *string   `json:"surname"`
	RightsCategory *string   `json:"rightsCategory"`
	Birthday       time.Time `json:"birthday"`
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
}

func nullStringToPtr(ns sql.NullString) *string {
	if ns.Valid {
		return &ns.String
	}
	return nil
}

func UserToResponseDto(u *user_model.UserModel) *UserResponseDto {
	return &UserResponseDto{
		ID:             u.ID,
		FirstName:      u.FirstName,
		LastName:       u.LastName,
		Surname:        nullStringToPtr(u.Surname),
		RightsCategory: nullStringToPtr(u.RightsCategory),
		Birthday:       u.Birthday,
		CreatedAt:      u.CreatedAt,
		UpdatedAt:      u.UpdatedAt,
	}
}
