package handlers

import (
	"context"
	appError "movie_app/models/error"
	"movie_app/models/request"
	"movie_app/models/response"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func (h *BaseHandler) BookShow(c *gin.Context) {
	var bookingRequest request.BookingRequest
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(h.config.Context.Timeout)*time.Second)
	defer cancel()

	if err := c.ShouldBindJSON(&bookingRequest); err != nil {
		response.APIResponse(c, err.Error(), http.StatusBadRequest, "Error")
		return
	}

	bookings, err := h.BaseSvc.BookShow(ctx, bookingRequest)
	if err != nil {
		if strings.Contains(err.Error(), appError.HouseFullError) {
			response.APIResponse(c, err.Error(), http.StatusBadRequest, "Booking failed")
			return
		}

		if strings.Contains(err.Error(), appError.ValidationError) {
			response.APIResponse(c, err.Error(), http.StatusBadRequest, "Seat is occupied")
			return
		}

		response.APIResponse(c, err.Error(), http.StatusInternalServerError, "Booking failed")
		return
	}

	response.APIResponse(c, bookings, 200, "Successfully booked")
}

func (h *BaseHandler) GetBooking(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(h.config.Context.Timeout)*time.Second)
	defer cancel()

	bookingId := c.Query("bookingId")
	if bookingId == "" {
		response.APIResponse(c, "invalid booking id", http.StatusBadRequest, "invalid booking id")
		return
	}

	booking, err := h.BaseSvc.GetBooking(ctx, bookingId)
	if err != nil {
		response.APIResponse(c, err.Error(), http.StatusInternalServerError, "Error fetching booking")
		return
	}

	response.APIResponse(c, booking, 200, "Booking fetched")
}
