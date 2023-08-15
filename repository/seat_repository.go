package repository

import (
	"context"
	"errors"
	"fmt"
	"log"
	"movie_app/models/entity"
)

func (sql SqliteRepository) CreateSeats(ctx context.Context, theatre entity.Theatre) error {
	rows := 65 + theatre.TotalRows
	for i := 65; i < int(rows); i++ {
		for j := 1; j <= int(theatre.TotalColumns); j++ {
			seatNo := fmt.Sprintf("%c%v", i, j)
			seat := entity.Seat{SeatNo: seatNo, TheatreId: uint(theatre.ID)}
			if err := sql.sqliteClient.Create(&seat).Error; err != nil {
				log.Println("Error creating seats", err)
				return err
			}
		}
	}
	return nil
}

func (sql SqliteRepository) GetSeats(ctx context.Context, theatreId uint) ([]entity.Seat, error) {
	seats := []entity.Seat{}
	if err := sql.sqliteClient.Select("id").Where("theatre_id = ?", theatreId).Find(&seats).Error; err != nil {
		log.Println("Error fetching seats", err)
		return []entity.Seat{}, err
	}
	if len(seats) == 0 {
		return []entity.Seat{}, errors.New("no seats found for the theatre")
	}

	return seats, nil
}
