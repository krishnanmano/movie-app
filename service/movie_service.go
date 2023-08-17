package service

import (
	"context"
	"movie_app/models/entity"
	"strconv"
)

func (baseSvc *baseService) CreateMovie(ctx context.Context, movie *entity.Movie) error {

	err := baseSvc.sqliteRepository.CreateMovie(ctx, movie)
	if err != nil {
		return err
	}
	return nil
}

func (baseSvc *baseService) GetMovies(ctx context.Context, movieId string) ([]entity.Movie, error) {
	var movies []entity.Movie
	var movie entity.Movie

	var err error
	if movieId != "" {
		movie, err = baseSvc.sqliteRepository.GetMovie(ctx, movieId)
		movies = append(movies, movie)
	} else {
		movies, err = baseSvc.sqliteRepository.GetMovies(ctx)
	}

	if err != nil {
		return []entity.Movie{}, err
	}
	return movies, nil
}

func (baseSvc *baseService) UpdateMovie(ctx context.Context, movie *entity.Movie, movieId string) error {
	movieIdInt, _ := strconv.Atoi(movieId)
	err := baseSvc.sqliteRepository.UpdateMovie(ctx, movie, movieIdInt)
	if err != nil {
		return err
	}
	return nil
}

func (baseSvc *baseService) DeleteMovie(ctx context.Context, movieId string) error {
	movieIdInt, _ := strconv.Atoi(movieId)
	err := baseSvc.sqliteRepository.DeleteMovie(ctx, movieIdInt)
	if err != nil {
		return err
	}
	return nil
}

func (baseSvc *baseService) DisplayMoviesByTheatre(ctx context.Context, theatreId string) ([]entity.Movie, error) {
	theatreIdInt, _ := strconv.Atoi(theatreId)
	movies, err := baseSvc.sqliteRepository.DisplayMoviesByTheatre(ctx, theatreIdInt)

	if err != nil {
		return []entity.Movie{}, err
	}
	return movies, nil
}
