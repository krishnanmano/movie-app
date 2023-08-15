package service

import (
	"context"
	"movie_app/models/entity"
	"movie_app/models/request"
	"strconv"
)

func (baseSvc *BaseService) BookShow(ctx context.Context, bookingRequest request.BookingRequest) (entity.Booking, error) {
	err := baseSvc.sqliteRepository.GetAvailableTickets(ctx, bookingRequest.ShowId)
	if err != nil {
		return entity.Booking{}, err
	}

	tickets, err := baseSvc.sqliteRepository.GetTickets(ctx, bookingRequest.Seats, bookingRequest.ShowId)
	if err != nil {
		return entity.Booking{}, err
	}

	booking := entity.Booking{UserName: bookingRequest.UserName, EmailId: bookingRequest.EmailId}
	err = baseSvc.sqliteRepository.CreateBooking(ctx, &booking)
	if err != nil {
		return entity.Booking{}, err
	}

	for _, ticket := range tickets {
		ticketBooking := entity.TicketBooking{TicketId: ticket.ID, BookingId: booking.ID}
		err = baseSvc.sqliteRepository.CreateTicketBooking(ctx, &ticketBooking)
		if err != nil {
			return entity.Booking{}, err
		}
		booking.TicketBookings = append(booking.TicketBookings, ticketBooking)
	}

	return booking, nil
}

func (baseSvc *BaseService) GetBooking(ctx context.Context, bookingId string) (entity.Booking, error) {
	bookingIdInt, _ := strconv.Atoi(bookingId)
	booking, err := baseSvc.sqliteRepository.GetBooking(ctx, uint(bookingIdInt))
	if err != nil {
		return entity.Booking{}, err
	}
	return booking, nil
}
