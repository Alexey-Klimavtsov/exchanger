package booking_handlers

import (
	"fmt"
	"github.com/asaipov/gorenda/internal/http/dto/booking_dto"
	"github.com/asaipov/gorenda/internal/http/helpers"
	"github.com/asaipov/gorenda/internal/service/booking_service"
	"github.com/gin-gonic/gin"
)

type BookingHandlers struct {
	s booking_service.BookingService
}

func NewBookingHandlers(s booking_service.BookingService) *BookingHandlers {
	return &BookingHandlers{s: s}
}

func (h *BookingHandlers) Create(c *gin.Context) {
	var req booking_dto.BookingRequestDto

	if err := c.ShouldBindJSON(&req); err != nil {
		helpers.HandleError(c, fmt.Errorf("%w", helpers.ErrBinding))
		return
	}

	input := booking_dto.DtoToInput(&req)

	booking, err := h.s.CreateBooking(c, input)
	if err != nil {
		helpers.HandleError(c, err)
		return
	}

	helpers.WriteCreated(c, booking_dto.BookingToResponseDto(booking))
}

func (h *BookingHandlers) GetAll(c *gin.Context) {
	bookings, err := h.s.GetAllBookings(c)
	if err != nil {
		helpers.HandleError(c, err)
		return
	}

	resp := make([]*booking_dto.BookingResponseDto, 0, len(bookings))
	for _, b := range bookings {
		resp = append(resp, booking_dto.BookingToResponseDto(b))
	}

	helpers.WriteOK(c, resp)
}
