package service

import (
	"context"
	"movie_app/models/entity"
	"strconv"
)

func (baseSvc *BaseService) CreateShow(ctx context.Context, show *entity.Show) error {

	err := baseSvc.sqliteRepository.CreateShow(ctx, show)
	if err != nil {
		return err
	}

	seats, err := baseSvc.sqliteRepository.GetSeats(ctx, show.TheatreId)
	if err != nil {
		return err
	}

	err = baseSvc.sqliteRepository.CreateTickets(ctx, seats, show.ID)
	if err != nil {
		return err
	}

	return nil
}

func (baseSvc *BaseService) GetShows(ctx context.Context, showName string) ([]entity.Show, error) {
	var shows []entity.Show
	var err error

	shows, err = baseSvc.sqliteRepository.GetShows(ctx)
	if err != nil {
		return []entity.Show{}, err
	}
	return shows, nil
}

func (baseSvc *BaseService) UpdateShow(ctx context.Context, show *entity.Show, showId string) error {
	showIdInt, _ := strconv.Atoi(showId)
	err := baseSvc.sqliteRepository.UpdateShow(ctx, show, showIdInt)
	if err != nil {
		return err
	}

	return nil
}
