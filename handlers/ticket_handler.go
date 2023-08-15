package handlers

import (
	"context"
	"movie_app/models/response"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func (h *BaseHandler) GetSeatMatrix(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(h.config.Context.Timeout)*time.Second)
	defer cancel()
	showId := c.Query("showId")
	if showId == "" {
		response.APIResponse(c, "invalid showId", http.StatusBadRequest, "invalid showId")
		return
	}

	seats, err := h.BaseSvc.GetTicketsByShowId(ctx, showId)
	if err != nil {
		response.APIResponse(c, err.Error(), http.StatusInternalServerError, "Error fetching seats")
		return
	}

	if len(seats) == 0 {
		response.APIResponse(c, nil, http.StatusNotFound, "seats not found for the show")
		return
	}

	response.APIResponse(c, seats, 200, "Seat matrix")
}
