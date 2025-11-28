package user_handlers

import (
	"fmt"
	"github.com/asaipov/gorenda/internal/http/dto/user_dto"
	"github.com/asaipov/gorenda/internal/http/helpers"
	"github.com/asaipov/gorenda/internal/service/user_service"
	"github.com/gin-gonic/gin"
)

type UserHandlers struct {
	s user_service.UserService
}

func NewUserHandlers(s user_service.UserService) *UserHandlers {
	return &UserHandlers{s: s}
}

func (h *UserHandlers) Create(c *gin.Context) {
	var req user_dto.UserRequestDto

	err := c.ShouldBindJSON(&req)
	if err != nil {
		helpers.HandleError(c, fmt.Errorf("%w", helpers.ErrBinding))
		return
	}

	input := user_dto.DtoToInput(&req)
	user, createErr := h.s.CreateUser(c, input)
	if createErr != nil {
		helpers.HandleError(c, createErr)
		return
	}

	helpers.WriteCreated(c, user_dto.UserToResponseDto(user))
}

func (h *UserHandlers) Delete(c *gin.Context) {
	id, getErr := helpers.GetIdFromQuery(c)
	if getErr != nil {
		helpers.HandleError(c, fmt.Errorf("%w", helpers.ErrReadingId))
		return
	}

	userId, deleteErr := h.s.DeleteUser(c, id)
	if deleteErr != nil {
		helpers.HandleError(c, deleteErr)
		return
	}

	helpers.WriteOK(c, userId)
}

func (h *UserHandlers) Update(c *gin.Context) {
	var req user_dto.UserRequestDto
	id, getErr := helpers.GetIdFromQuery(c)
	if getErr != nil {
		helpers.HandleError(c, fmt.Errorf("%w", helpers.ErrReadingId))
		return
	}

	err := c.ShouldBindJSON(&req)
	if err != nil {
		helpers.HandleError(c, fmt.Errorf("%w", helpers.ErrBinding))
		return
	}

	input := user_dto.DtoToInput(&req)
	user, updateErr := h.s.UpdateUser(c, input, id)
	if updateErr != nil {
		helpers.HandleError(c, updateErr)
		return
	}

	helpers.WriteOK(c, user_dto.UserToResponseDto(user))
}

func (h *UserHandlers) GetById(c *gin.Context) {
	id, getErr := helpers.GetIdFromQuery(c)
	if getErr != nil {
		helpers.HandleError(c, fmt.Errorf("%w", helpers.ErrReadingId))
		return
	}

	user, err := h.s.GetUserById(c, id)
	if err != nil {
		helpers.HandleError(c, err)
		return
	}

	helpers.WriteOK(c, user_dto.UserToResponseDto(user))
}
