package service

import (
	"context"
	"movie_app/models/entity"
	"strconv"
)

func (baseSvc *BaseService) GetTicketsByShowId(ctx context.Context, showId string) ([]entity.Ticket, error) {
	var err error
	showIdInt, _ := strconv.Atoi(showId)
	tickets, err := baseSvc.sqliteRepository.GetTicketsByShowId(ctx, uint(showIdInt))

	if err != nil {
		return []entity.Ticket{}, err
	}
	return tickets, nil
}
