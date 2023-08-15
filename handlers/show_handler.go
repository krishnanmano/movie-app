package handlers

import (
	"context"
	"movie_app/models/entity"
	appError "movie_app/models/error"
	"movie_app/models/response"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func (h *BaseHandler) CreateShow(c *gin.Context) {
	var show entity.Show
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(h.config.Context.Timeout)*time.Second)
	defer cancel()

	if err := c.ShouldBindJSON(&show); err != nil {
		response.APIResponse(c, err.Error(), http.StatusBadRequest, "Error")
		return
	}

	err := h.BaseSvc.CreateShow(ctx, &show)
	if err != nil {
		response.APIResponse(c, err.Error(), http.StatusInternalServerError, "Error creating show")
		return
	}

	response.APIResponse(c, show, 200, "show Created")
}

func (h *BaseHandler) GetAllShows(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(h.config.Context.Timeout)*time.Second)
	defer cancel()

	movies, err := h.BaseSvc.GetShows(ctx, "")
	if err != nil {
		response.APIResponse(c, err.Error(), http.StatusInternalServerError, "Error fetching show")
		return
	}

	if len(movies) == 0 {
		response.APIResponse(c, nil, http.StatusNotFound, "show not found")
		return
	}

	response.APIResponse(c, movies, 200, "List of show")
}

func (h *BaseHandler) UpdateShow(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(h.config.Context.Timeout)*time.Second)
	defer cancel()

	var show entity.Show
	showId := c.Param("showId")

	if err := c.ShouldBindJSON(&show); err != nil {
		response.APIResponse(c, err.Error(), http.StatusBadRequest, "Error")
		return
	}

	err := h.BaseSvc.UpdateShow(ctx, &show, showId)
	if err != nil {
		if strings.Contains(err.Error(), appError.ErrorNotFound) {
			response.APIResponse(c, err.Error(), http.StatusNotFound, "Error updating show")
			return
		}
		response.APIResponse(c, err.Error(), http.StatusInternalServerError, "Error updating show")
		return
	}

	response.APIResponse(c, show, 200, "Updated show")
}
