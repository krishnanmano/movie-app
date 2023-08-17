package handlers

import (
	"context"
	"movie_app/models/entity"
	"movie_app/models/response"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func (h *baseHandler) CreateTheatre(c *gin.Context) {
	var theatre entity.Theatre
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(h.config.Context.Timeout)*time.Second)
	defer cancel()

	if err := c.ShouldBindJSON(&theatre); err != nil {
		response.APIResponse(c, err.Error(), http.StatusBadRequest, "Error")
		return
	}

	err := h.BaseSvc.CreateTheatre(ctx, &theatre)
	if err != nil {
		response.APIResponse(c, err.Error(), http.StatusInternalServerError, "Error creating theatre")
		return
	}

	response.APIResponse(c, theatre, 200, "Theatre Created")
}

func (h *baseHandler) GetTheatres(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(h.config.Context.Timeout)*time.Second)
	defer cancel()
	var theatres []entity.Theatre
	var err error
	if movieId := c.Query("movieId"); movieId != "" {
		theatres, err = h.BaseSvc.DisplayTheatreByMovies(ctx, movieId)
	} else {
		theatreName := c.Query("theatreName")
		theatres, err = h.BaseSvc.GetTheatres(ctx, theatreName)
	}

	if err != nil {
		response.APIResponse(c, err.Error(), http.StatusInternalServerError, "Error fetching theatre")
		return
	}

	if len(theatres) == 0 {
		response.APIResponse(c, nil, http.StatusNotFound, "Theatre not found")
		return
	}

	response.APIResponse(c, theatres, 200, "List of Theatre")
}
