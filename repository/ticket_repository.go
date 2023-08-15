package repository

import (
	"context"
	"errors"
	"log"
	"movie_app/models/entity"
)

func (sql SqliteRepository) CreateTickets(ctx context.Context, seats []entity.Seat, showId uint) error {
	for _, seat := range seats {
		ticket := entity.Ticket{Price: 500, SeatId: seat.ID, ShowId: showId}
		if err := sql.sqliteClient.Create(&ticket).Error; err != nil {
			return err
		}
	}
	return nil
}

func (sql SqliteRepository) GetTickets(ctx context.Context, seats []uint, showId uint) ([]entity.Ticket, error) {
	tickets := []entity.Ticket{}
	if err := sql.sqliteClient.Where("id IN (?) AND show_id = ?", seats, showId).Find(&tickets).Error; err != nil {
		log.Println("Error fetching tickets", err)
		return []entity.Ticket{}, err
	}
	return tickets, nil
}

func (sql SqliteRepository) GetAvailableTickets(ctx context.Context, showId uint) error {
	var totalTickets int64
	var ticketBookings []entity.TicketBooking
	if err := sql.sqliteClient.Model([]entity.Ticket{}).
		Where("show_id = ?", showId).Count(&totalTickets).Error; err != nil {
		log.Println("Error fetching tickets", err)
		return err
	}

	if err := sql.sqliteClient.Model(&entity.TicketBooking{}).
		Joins("JOIN tickets ON ticket_bookings.ticket_id = tickets.id").
		Where("show_id = ?", showId).Find(&ticketBookings).Error; err != nil {
		log.Println("Error fetching ticketBookings", err)
		return err
	}

	if len(ticketBookings) == int(totalTickets) {
		return errors.New("houseful all seats are booked")
	}

	return nil
}

func (sql SqliteRepository) GetTicketsByShowId(ctx context.Context, showId uint) ([]entity.Ticket, error) {
	var tickets []entity.Ticket
	if err := sql.sqliteClient.Model(&entity.Ticket{}).
		Joins("JOIN ticket_bookings ON ticket_bookings.ticket_id = tickets.id").
		Where("show_id = ?", showId).Find(&tickets).Error; err != nil {
		log.Println("Error fetching tickets", err)
		return []entity.Ticket{}, err
	}

	for i := 0; i < len(tickets); i++ {
		var ticketBooking entity.TicketBooking
		err := sql.sqliteClient.Model(&entity.TicketBooking{}).Debug().
			Where("ticket_id = ?", tickets[i].ID).Find(&ticketBooking).Error
		if err != nil {
			return nil, err
		}
		tickets[i].TicketBooking = ticketBooking
	}

	return tickets, nil
}
