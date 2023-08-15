package repository

import (
	"context"
	"log"
	"movie_app/models/entity"
)

func (sql SqliteRepository) CreateTheatre(ctx context.Context, theatre *entity.Theatre) error {
	if err := sql.sqliteClient.Create(&theatre).Error; err != nil {
		log.Println("Error creating theatre", err.Error())
		return err
	}
	return nil
}

func (sql SqliteRepository) GetTheatres(ctx context.Context) ([]entity.Theatre, error) {
	theatres := []entity.Theatre{}
	if err := sql.sqliteClient.Find(&theatres).Error; err != nil {
		log.Println("Error fetching theatres", err.Error())
		return []entity.Theatre{}, err
	}
	return theatres, nil
}

func (sql SqliteRepository) GetTheatre(ctx context.Context, theatreName string) ([]entity.Theatre, error) {
	theatres := []entity.Theatre{}
	if err := sql.sqliteClient.Where("name = ?", theatreName).Find(&theatres).Error; err != nil {
		log.Println("Error fetching theatre", err.Error())
		return []entity.Theatre{}, err
	}
	return theatres, nil
}

func (sql SqliteRepository) DisplayTheatreByMovies(ctx context.Context, movieId int) ([]entity.Theatre, error) {
	var theatres []entity.Theatre

	if err := sql.sqliteClient.Model(&entity.Theatre{}).Joins("JOIN shows ON theatres.id = shows.theatre_id").Where("movie_id = ?", movieId).Find(&theatres).Error; err != nil {
		log.Println("Error", err.Error())
		return []entity.Theatre{}, err
	}
	return theatres, nil
}
