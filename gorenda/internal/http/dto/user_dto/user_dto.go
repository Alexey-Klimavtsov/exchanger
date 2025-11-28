package user_dto

import (
	"github.com/asaipov/gorenda/internal/http/dto/driver_license_dto"
	"github.com/asaipov/gorenda/internal/it/model/user_model"
	"github.com/asaipov/gorenda/internal/service/user_service"
	"time"
)

type UserRequestDto struct {
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Surname   *string   `json:"surname"`
	Email     string    `json:"email"`
	Birthday  time.Time `json:"birthday"`
}

type UserResponseDto struct {
	ID             int64                                          `json:"id"`
	FirstName      string                                         `json:"firstName"`
	LastName       string                                         `json:"lastName"`
	Surname        *string                                        `json:"surname"`
	Email          string                                         `json:"email"`
	RightsCategory []*driver_license_dto.DriverLicenseResponseDto `json:"rightsCategory"`
	Birthday       time.Time                                      `json:"birthday"`
	CreatedAt      time.Time                                      `json:"createdAt"`
	UpdatedAt      *time.Time                                     `json:"updatedAt"`
}

func DtoToInput(dto *UserRequestDto) *user_service.CreateUserInput {
	return &user_service.CreateUserInput{
		FirstName: dto.FirstName,
		LastName:  dto.LastName,
		Email:     dto.Email,
		Surname:   dto.Surname,
		Birthday:  dto.Birthday,
	}
}

func UserToResponseDto(u *user_model.UserModel) *UserResponseDto {
	var userRights []*driver_license_dto.DriverLicenseResponseDto

	for _, model := range u.RightsCategory {
		userRights = append(userRights, driver_license_dto.DriverLicenseModelToDto(model))
	}

	return &UserResponseDto{
		ID:             u.ID,
		FirstName:      u.FirstName,
		LastName:       u.LastName,
		Surname:        u.Surname,
		Email:          u.Email,
		RightsCategory: userRights,
		Birthday:       u.Birthday,
		CreatedAt:      u.CreatedAt,
		UpdatedAt:      u.UpdatedAt,
	}
}
