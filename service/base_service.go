package service

import (
	"context"
	"movie_app/models/entity"
	"movie_app/repository"

	"movie_app/models/request"
)

type BaseService struct {
	sqliteRepository repository.ISqliteRepository
}

func NewBaseService(sqliteRepository repository.ISqliteRepository) IBaseService {
	return &BaseService{
		sqliteRepository: sqliteRepository,
	}
}

type IBaseService interface {
	CreateMovie(ctx context.Context, movie *entity.Movie) error
	GetMovies(ctx context.Context, movieId string) ([]entity.Movie, error)
	UpdateMovie(ctx context.Context, movie *entity.Movie, movieId string) error
	DeleteMovie(ctx context.Context, movieId string) error
	DisplayMoviesByTheatre(ctx context.Context, theatreId string) ([]entity.Movie, error)
	CreateTheatre(ctx context.Context, theatre *entity.Theatre) error
	GetTheatres(ctx context.Context, theatreName string) ([]entity.Theatre, error)
	DisplayTheatreByMovies(ctx context.Context, movieId string) ([]entity.Theatre, error)
	CreateShow(ctx context.Context, show *entity.Show) error
	GetShows(ctx context.Context, showName string) ([]entity.Show, error)
	UpdateShow(ctx context.Context, show *entity.Show, showId string) error
	BookShow(ctx context.Context, booking request.BookingRequest) (entity.Booking, error)
	GetBooking(ctx context.Context, bookingId string) (entity.Booking, error)
	GetTicketsByShowId(ctx context.Context, showId string) ([]entity.Ticket, error)
}
