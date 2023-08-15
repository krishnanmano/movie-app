package repository

import (
	"context"
	"log"
	"movie_app/models/entity"
)

func (sql SqliteRepository) CreateBooking(ctx context.Context, booking *entity.Booking) error {

	db := sql.sqliteClient.Create(booking)
	if db.Error != nil {
		log.Println("Error creating booking: ", db.Error)
		return db.Error
	}

	return nil
}

func (sql SqliteRepository) CreateTicketBooking(ctx context.Context, ticketbooking *entity.TicketBooking) error {

	db := sql.sqliteClient.Create(ticketbooking)
	if db.Error != nil {
		log.Println("Error creating ticketBooking: ", db.Error)
		return db.Error
	}

	return nil
}

func (sql SqliteRepository) GetBooking(ctx context.Context, bookingId uint) (entity.Booking, error) {
	var booking entity.Booking
	var ticketBookings []entity.TicketBooking
	db := sql.sqliteClient.Where("id = ?", bookingId).Find(&booking)
	if db.Error != nil {
		log.Println("Error fetching booking: ", db.Error)
		return entity.Booking{}, db.Error
	}
	if err := sql.sqliteClient.Where("booking_id = ?", bookingId).Find(&ticketBookings).Error; err != nil {
		log.Println("Error fetching ticketBooking: ", err)
	}
	booking.TicketBookings = append(booking.TicketBookings, ticketBookings...)
	return booking, nil
}
