package handlers

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"movie_app/config"
	"movie_app/mocks"
	"movie_app/models/entity"
	"movie_app/models/request"
	"net/http"
	"net/http/httptest"
	"testing"
)

type BookingHandlerTestSuite struct {
	suite.Suite
	mockCtrl       *gomock.Controller
	context        *gin.Context
	engine         *gin.Engine
	recorder       *httptest.ResponseRecorder
	bookingSvc     *mocks.MockIBaseService
	config         config.Configuration
	bookingHandler IBaseHandler
}

func TestBookingHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(BookingHandlerTestSuite))
}

func (t *BookingHandlerTestSuite) SetupTest() {
	gin.SetMode(gin.TestMode)
	t.mockCtrl = gomock.NewController(t.T())
	t.recorder = httptest.NewRecorder()
	t.context, t.engine = gin.CreateTestContext(t.recorder)
	t.config = config.Configuration{}
	t.bookingSvc = mocks.NewMockIBaseService(t.mockCtrl)
	t.bookingHandler = NewBaseHandler(t.bookingSvc, t.config)
}

func (t *BookingHandlerTestSuite) TearDownTest() {
	t.mockCtrl.Finish()
	t.recorder.Flush()
}

func (t *BookingHandlerTestSuite) TestBookingHandler_BookShow() {
	bookingRequest := request.BookingRequest{
		UserName: "sample-user",
		EmailId:  "sample-user@gmail.com",
		Seats:    []uint{1, 2},
		ShowId:   0,
	}
	booking := entity.Booking{
		UserName:       "sample-user",
		EmailId:        "sample-user@gmail.com",
		TicketBookings: nil,
	}
	jsonValue, _ := json.Marshal(bookingRequest)
	t.bookingSvc.EXPECT().BookShow(gomock.Any(), gomock.Any()).Return(booking, nil)
	t.engine.POST("/booking", t.bookingHandler.BookShow)
	req, _ := http.NewRequest(
		http.MethodPost,
		"/booking",
		bytes.NewBuffer(jsonValue))

	t.engine.ServeHTTP(t.recorder, req)
	assert.Equal(t.T(), 200, t.recorder.Code)
	assert.NotNil(t.T(), t.recorder.Body)

}

func (t *BookingHandlerTestSuite) TestBookingHandler_GetBooking() {
	var bookingId uint = 345
	booking := entity.Booking{
		UserName: "sample-user",
		EmailId:  "sample-user@gmail.com",
		TicketBookings: []entity.TicketBooking{
			{
				BookingId: bookingId,
			},
		},
	}

	t.bookingSvc.EXPECT().GetBooking(gomock.Any(), gomock.Any()).Return(booking, nil)
	t.engine.GET("/booking", t.bookingHandler.GetBooking)
	req, _ := http.NewRequest(
		http.MethodGet,
		"/booking?bookingId=1",
		nil)

	t.engine.ServeHTTP(t.recorder, req)
	assert.Equal(t.T(), 200, t.recorder.Code)

}
