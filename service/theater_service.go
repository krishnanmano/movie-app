package service

import (
	"context"
	"movie_app/models/entity"
	"strconv"
)

func (baseSvc *BaseService) CreateTheatre(ctx context.Context, theatre *entity.Theatre) error {

	err := baseSvc.sqliteRepository.CreateTheatre(ctx, theatre)
	if err != nil {
		return err
	}
	if err := baseSvc.sqliteRepository.CreateSeats(ctx, *theatre); err != nil {
		return err
	}
	return nil
}

func (baseSvc *BaseService) GetTheatres(ctx context.Context, theatreName string) ([]entity.Theatre, error) {
	var theatres []entity.Theatre
	var err error
	if theatreName != "" {
		theatres, err = baseSvc.sqliteRepository.GetTheatre(ctx, theatreName)
	} else {
		theatres, err = baseSvc.sqliteRepository.GetTheatres(ctx)
	}

	if err != nil {
		return []entity.Theatre{}, err
	}
	return theatres, nil
}

func (baseSvc *BaseService) DisplayTheatreByMovies(ctx context.Context, movieId string) ([]entity.Theatre, error) {
	movieIdInt, _ := strconv.Atoi(movieId)
	theatres, err := baseSvc.sqliteRepository.DisplayTheatreByMovies(ctx, movieIdInt)

	if err != nil {
		return []entity.Theatre{}, err
	}
	return theatres, nil
}
