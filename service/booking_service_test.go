package service

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"movie_app/mocks"
	"movie_app/models/entity"
	"movie_app/models/request"
	"testing"
)

type BookingServiceTestSuite struct {
	suite.Suite
	ctx            context.Context
	mockCtrl       *gomock.Controller
	BookingRepo    *mocks.MockISqliteRepository
	BookingService IBaseService
}

func TestBookingServiceTestSuite(t *testing.T) {
	suite.Run(t, new(BookingServiceTestSuite))
}

func (t *BookingServiceTestSuite) SetupTest() {
	gin.SetMode(gin.TestMode)
	t.mockCtrl = gomock.NewController(t.T())
	t.ctx = context.WithValue(context.Background(), "logger", nil)
	t.BookingRepo = mocks.NewMockISqliteRepository(t.mockCtrl)
	t.BookingService = NewBaseService(t.BookingRepo)
}

func (t *BookingServiceTestSuite) TearDownTest() {
	t.mockCtrl.Finish()
}

func (t *BookingServiceTestSuite) TestBookingService_BookShow() {

	tickets := []entity.Ticket{{
		Price:  500,
		SeatId: 12,
	},
	}

	BookingRequest := request.BookingRequest{UserName: "Sample-user"}

	t.BookingRepo.EXPECT().GetAvailableTickets(gomock.Any(), gomock.Any()).Return(nil)
	t.BookingRepo.EXPECT().GetTickets(gomock.Any(), gomock.Any(), gomock.Any()).Return(tickets, nil)
	t.BookingRepo.EXPECT().CreateBooking(gomock.Any(), gomock.Any()).Return(nil)
	t.BookingRepo.EXPECT().CreateTicketBooking(gomock.Any(), gomock.Any()).Return(nil)
	booking, err := t.BookingService.BookShow(t.ctx, BookingRequest)
	assert.Equal(t.T(), booking.UserName, BookingRequest.UserName)
	assert.Nil(t.T(), err)
}

func (t *BookingServiceTestSuite) TestBookingService_GetBooking() {
	Booking := entity.Booking{UserName: "sample-user"}
	t.BookingRepo.EXPECT().GetBooking(gomock.Any(), gomock.Any()).Return(Booking, nil)
	actualBooking, err := t.BookingService.GetBooking(t.ctx, "123")
	assert.Equal(t.T(), actualBooking.UserName, Booking.UserName)
	assert.Nil(t.T(), err)
}
