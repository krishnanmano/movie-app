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

func (h *BaseHandler) CreateMovie(c *gin.Context) {
	var movie entity.Movie
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(h.config.Context.Timeout)*time.Second)
	defer cancel()

	if err := c.ShouldBindJSON(&movie); err != nil {
		response.APIResponse(c, err.Error(), http.StatusBadRequest, "Error")
		return
	}

	err := h.BaseSvc.CreateMovie(ctx, &movie)
	if err != nil {
		response.APIResponse(c, err.Error(), http.StatusInternalServerError, "Error creating Movie")
		return
	}

	response.APIResponse(c, movie, 200, "Movie Created")
}

func (h *BaseHandler) GetMovies(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(h.config.Context.Timeout)*time.Second)
	defer cancel()
	var movies []entity.Movie
	var err error
	if theatreId := c.Query("theatreId"); theatreId != "" {
		movies, err = h.BaseSvc.DisplayMoviesByTheatre(ctx, theatreId)
	} else {
		movieId := c.Query("movieId")
		movies, err = h.BaseSvc.GetMovies(ctx, movieId)
	}

	if err != nil {
		response.APIResponse(c, err.Error(), http.StatusInternalServerError, "Error fetching movies")
		return
	}

	if len(movies) == 0 {
		response.APIResponse(c, nil, http.StatusNotFound, "Movie not found")
		return
	}

	response.APIResponse(c, movies, 200, "List of movies")
}

func (h *BaseHandler) UpdateMovie(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(h.config.Context.Timeout)*time.Second)
	defer cancel()

	var movie entity.Movie
	movieId := c.Param("movieId")

	if err := c.ShouldBindJSON(&movie); err != nil {
		response.APIResponse(c, err.Error(), http.StatusBadRequest, "Error")
		return
	}

	err := h.BaseSvc.UpdateMovie(ctx, &movie, movieId)
	if err != nil {
		if strings.Contains(err.Error(), appError.ErrorNotFound) {
			response.APIResponse(c, err.Error(), http.StatusNotFound, "Error updating movie")
			return
		}
		response.APIResponse(c, err.Error(), http.StatusInternalServerError, "Error updating movie")
		return
	}

	response.APIResponse(c, movie, 200, "Updated movie")
}

func (h *BaseHandler) DeleteMovie(c *gin.Context) {
	movieId := c.Param("movieId")
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(h.config.Context.Timeout)*time.Second)
	defer cancel()
	err := h.BaseSvc.DeleteMovie(ctx, movieId)
	if err != nil {
		if strings.Contains(err.Error(), appError.ErrorNotFound) {
			response.APIResponse(c, err.Error(), http.StatusNotFound, "Error deleting movie")
			return
		}
		response.APIResponse(c, err.Error(), http.StatusInternalServerError, "Error deleting movie")
		return
	}
	response.APIResponse(c, nil, http.StatusOK, "Successfully deleted movie")
}
