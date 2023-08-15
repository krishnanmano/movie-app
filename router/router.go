package router

import (
	"github.com/gin-gonic/gin"
	"movie_app/config"
	"movie_app/gorm_client"
	"movie_app/handlers"
	"movie_app/repository"
	"movie_app/service"
)

func SetUpRouter(config config.Configuration) *gin.Engine {
	router := gin.Default()

	gormClient := gorm_client.NewGormDB()
	sqliteRepository := repository.NewSqliteRepository(gormClient)

	baseService := service.NewBaseService(sqliteRepository)
	baseHandler := handlers.NewBaseHandler(baseService, config)
	movie := router.Group("/movie")
	{
		movie.POST("", baseHandler.CreateMovie)
		movie.GET("", baseHandler.GetMovies)
		movie.PUT("/:movieId", baseHandler.UpdateMovie)
		movie.DELETE("/:movieId", baseHandler.DeleteMovie)
	}

	theatre := router.Group("/theatre")
	{
		theatre.POST("", baseHandler.CreateTheatre)
		theatre.GET("", baseHandler.GetTheatres)
	}

	show := router.Group("/show")
	{
		show.POST("", baseHandler.CreateShow)
		show.GET("", baseHandler.GetAllShows)
		show.PUT("/:showId", baseHandler.UpdateShow)
	}

	booking := router.Group("/booking")
	{
		booking.POST("", baseHandler.BookShow)
		booking.GET("", baseHandler.GetBooking)
	}

	ticket := router.Group("/ticket")
	{
		ticket.GET("", baseHandler.GetSeatMatrix)
	}

	return router
}
