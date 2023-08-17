package handlers

import (
	"movie_app/config"
	"movie_app/service"

	"github.com/gin-gonic/gin"
)

type baseHandler struct {
	BaseSvc service.BaseService
	config  config.Configuration
}

type BaseHandler interface {
	CreateMovie(c *gin.Context)
	GetMovies(c *gin.Context)
	UpdateMovie(c *gin.Context)
	DeleteMovie(c *gin.Context)
	CreateTheatre(c *gin.Context)
	GetTheatres(c *gin.Context)
	CreateShow(c *gin.Context)
	GetAllShows(c *gin.Context)
	UpdateShow(c *gin.Context)
	BookShow(c *gin.Context)
	GetBooking(c *gin.Context)
	GetSeatMatrix(c *gin.Context)
}

func NewBaseHandler(baseService service.BaseService, config config.Configuration) BaseHandler {
	return &baseHandler{
		BaseSvc: baseService,
		config:  config,
	}
}
