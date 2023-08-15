package entity

import (
	"time"

	"gorm.io/gorm"
)

type Movie struct {
	gorm.Model
	Title    string `json:"title" gorm:"unique"`
	Language string `json:"language"`
	Duration string `json:"duration"`
	Genre    string `json:"genre"`
	Shows    []Show `json:"shows,omitempty"`
}

type Show struct {
	gorm.Model
	TheatreId uint      `json:"theatre_id" `
	MovieId   uint      `json:"movie_id"`
	CreatedAt time.Time `json:"created_at"`
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
	Tickets   []Ticket  `json:"seats,omitempty"`
}

type Theatre struct {
	gorm.Model
	Name         string `json:"name" gorm:"index:idx_theatre,unique"`
	Address      string `json:"address"  gorm:"index:idx_theatre,unique"`
	TotalColumns int32  `json:"total_columns"`
	TotalRows    int32  `json:"total_rows"`
	Shows        []Show `json:"shows,omitempty"`
	Seats        []Seat `json:"seats,omitempty"`
}

type Seat struct {
	gorm.Model
	TheatreId uint   `json:"theatre_id,omitempty" gorm:"index:idx_seat,unique"`
	SeatNo    string `json:"seat_no" gorm:"index:idx_seat,unique"`

	Tickets []Ticket `json:"tickets,omitempty"`
}

type Booking struct {
	gorm.Model
	UserName       string          `json:"user_name"`
	EmailId        string          `json:"email_id"`
	TicketBookings []TicketBooking `json:"ticket_bookings,omitempty"`
}

type Ticket struct {
	gorm.Model
	Price         int32         `json:"price"`
	SeatId        uint          `json:"seat_id"`
	ShowId        uint          `json:"show_id"`
	TicketBooking TicketBooking `json:"ticket_booking,omitempty"`
}

type TicketBooking struct {
	gorm.Model
	TicketId  uint `json:"ticket_id" gorm:"unique"`
	BookingId uint `json:"booking_id"`
}
