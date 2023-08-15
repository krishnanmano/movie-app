package handlers

import (
	"movie_app/config"
	"movie_app/service"

	"github.com/gin-gonic/gin"
)

type BaseHandler struct {
	BaseSvc service.IBaseService
	config  config.Configuration
}

type IBaseHandler interface {
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

func NewBaseHandler(baseService service.IBaseService, config config.Configuration) IBaseHandler {
	return &BaseHandler{
		BaseSvc: baseService,
		config:  config,
	}
}
