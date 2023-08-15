package repository

import (
	"context"
	"errors"
	"log"
	"movie_app/models/entity"
)

func (sql SqliteRepository) CreateMovie(ctx context.Context, movie *entity.Movie) error {
	if err := sql.sqliteClient.Create(movie).Error; err != nil {
		log.Println("Error creating movie", err.Error())
		return err
	}
	return nil
}

func (sql SqliteRepository) GetMovies(ctx context.Context) ([]entity.Movie, error) {
	movies := []entity.Movie{}
	if err := sql.sqliteClient.Find(&movies).Error; err != nil {
		log.Println("Error fetching movies", err.Error())
		return []entity.Movie{}, err
	}
	return movies, nil
}

func (sql SqliteRepository) GetMovie(ctx context.Context, movieId string) (entity.Movie, error) {
	movie := entity.Movie{}
	db := sql.sqliteClient.Where("id = ?", movieId).Find(&movie)
	if db.Error != nil {
		log.Println("Error fetching movie: ", db.Error)
		return entity.Movie{}, db.Error
	}
	return movie, nil
}

func (sql SqliteRepository) UpdateMovie(ctx context.Context, movie *entity.Movie, movieId int) error {
	db := sql.sqliteClient.Model(&entity.Movie{}).Where("id = ?", movieId).Updates(movie)
	movie.ID = uint(movieId)
	if err := db.Error; err != nil {
		log.Println("Error updating movie", err)
		return err
	}

	if db.RowsAffected == 0 {
		log.Println("Movie not found")
		return errors.New("movie not found")
	}

	return nil
}

func (sql SqliteRepository) DeleteMovie(ctx context.Context, movieId int) error {
	db := sql.sqliteClient.Where("id = ?", movieId).Delete(&entity.Movie{})
	if db.Error != nil {
		log.Println("Error deleting movie", db.Error)
		return db.Error
	}

	if db.RowsAffected == 0 {
		log.Println("Movie not found")
		return errors.New("movie not found")
	}

	return nil
}

func (sql SqliteRepository) DisplayMoviesByTheatre(ctx context.Context, theatreId int) ([]entity.Movie, error) {
	var movies []entity.Movie

	if err := sql.sqliteClient.Model(&entity.Movie{}).Joins("JOIN shows ON movies.id = shows.movie_id").Where("theatre_id = ?", theatreId).Find(&movies).Error; err != nil {
		log.Println("Error fetching movie", err)
		return []entity.Movie{}, err
	}
	return movies, nil
}
