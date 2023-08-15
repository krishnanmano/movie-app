package repository

import (
	"context"
	"log"
	"movie_app/gorm_client"
	"movie_app/models/entity"
)

type SqliteRepository struct {
	sqliteClient gorm_client.IGormDB
}

type ISqliteRepository interface {
	CreateMovie(ctx context.Context, movie *entity.Movie) error
	GetMovies(ctx context.Context) ([]entity.Movie, error)
	GetMovie(ctx context.Context, movieId string) (entity.Movie, error)
	UpdateMovie(ctx context.Context, movie *entity.Movie, movieId int) error
	DeleteMovie(ctx context.Context, movieId int) error
	DisplayMoviesByTheatre(ctx context.Context, theatreId int) ([]entity.Movie, error)
	CreateTheatre(ctx context.Context, theatre *entity.Theatre) error
	GetTheatres(ctx context.Context) ([]entity.Theatre, error)
	GetTheatre(ctx context.Context, theatreName string) ([]entity.Theatre, error)
	DisplayTheatreByMovies(ctx context.Context, movieId int) ([]entity.Theatre, error)
	CreateShow(ctx context.Context, show *entity.Show) error
	GetShows(ctx context.Context) ([]entity.Show, error)
	UpdateShow(ctx context.Context, show *entity.Show, showId int) error
	CreateBooking(ctx context.Context, booking *entity.Booking) error
	GetBooking(ctx context.Context, bookingId uint) (entity.Booking, error)
	CreateSeats(ctx context.Context, theatre entity.Theatre) error
	GetSeats(ctx context.Context, theatreId uint) ([]entity.Seat, error)
	CreateTickets(ctx context.Context, seats []entity.Seat, showId uint) error
	GetTickets(ctx context.Context, seats []uint, showId uint) ([]entity.Ticket, error)
	GetTicketsByShowId(ctx context.Context, showId uint) ([]entity.Ticket, error)
	GetAvailableTickets(ctx context.Context, showId uint) error
	CreateTicketBooking(ctx context.Context, ticketbooking *entity.TicketBooking) error
}

func NewSqliteRepository(sqliteClient gorm_client.IGormDB) ISqliteRepository {

	if err := sqliteClient.AutoMigrate(
		&entity.Movie{},
		&entity.Theatre{},
		&entity.Show{},
		&entity.Seat{},
		&entity.Booking{},
		&entity.Ticket{},
		&entity.TicketBooking{},
	); err != nil {
		log.Fatalf("error auto migrating models: %v", err)
	}
	return &SqliteRepository{
		sqliteClient: sqliteClient,
	}
}
