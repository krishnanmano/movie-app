package repository

import (
	"context"
	"errors"
	"log"
	"movie_app/models/entity"
)

func (sql SqliteRepository) CreateShow(ctx context.Context, show *entity.Show) error {
	if err := sql.sqliteClient.Create(show).Error; err != nil {
		log.Println("Error creating show", err)
		return err
	}

	return nil
}

func (sql SqliteRepository) GetShows(ctx context.Context) ([]entity.Show, error) {
	shows := []entity.Show{}
	if err := sql.sqliteClient.Find(&shows).Error; err != nil {
		log.Println("Error fetching shows", err)
		return []entity.Show{}, err
	}

	return shows, nil
}

func (sql SqliteRepository) UpdateShow(ctx context.Context, show *entity.Show, showId int) error {
	db := sql.sqliteClient.Model(&entity.Show{}).Where("id = ?", showId).Updates(show)
	show.ID = uint(showId)
	if err := db.Error; err != nil {
		log.Println("Error updating shows", err)
		return err
	}
	if db.RowsAffected == 0 {
		log.Println("Show not found")
		return errors.New("show not found")
	}

	return nil
}
