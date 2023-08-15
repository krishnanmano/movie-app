package repository

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
	"movie_app/mocks"
	"movie_app/models/entity"
	"net/http/httptest"
	"testing"
)

type BookingRepoTestSuite struct {
	suite.Suite
	mockCtrl          *gomock.Controller
	context           *gin.Context
	engine            *gin.Engine
	recorder          *httptest.ResponseRecorder
	mockDB            *mocks.MockIGormDB
	bookingRepository ISqliteRepository
}

func TestTicketHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(BookingRepoTestSuite))
}

func (t *BookingRepoTestSuite) SetupTest() {
	gin.SetMode(gin.TestMode)
	t.mockCtrl = gomock.NewController(t.T())
	t.recorder = httptest.NewRecorder()
	t.context, t.engine = gin.CreateTestContext(t.recorder)
	t.mockDB = mocks.NewMockIGormDB(t.mockCtrl)
	t.mockDB.EXPECT().AutoMigrate(gomock.Any()).Return(nil)
	t.bookingRepository = NewSqliteRepository(t.mockDB)
}

func (t *BookingRepoTestSuite) TearDownTest() {
	t.mockCtrl.Finish()
	t.recorder.Flush()
}

func (t *BookingRepoTestSuite) TestBookingRepository_CreateBooking_Success() {
	testGorm := gorm.DB{Error: nil}
	t.mockDB.EXPECT().Create(gomock.Any()).Return(&testGorm)
	err := t.bookingRepository.CreateBooking(t.context, &entity.Booking{})
	assert.Nil(t.T(), err)
}

func (t *BookingRepoTestSuite) TestBookingRepository_CreateBooking_Error() {
	testGorm := gorm.DB{Error: errors.New("test-error")}
	t.mockDB.EXPECT().Create(gomock.Any()).Return(&testGorm)
	err := t.bookingRepository.CreateBooking(t.context, &entity.Booking{})
	assert.NotNil(t.T(), err)
}

func (t *BookingRepoTestSuite) TestBookingRepository_CreateTicketBooking_Success() {
	testGorm := gorm.DB{Error: nil}
	t.mockDB.EXPECT().Create(gomock.Any()).Return(&testGorm)
	err := t.bookingRepository.CreateTicketBooking(t.context, &entity.TicketBooking{})
	assert.Nil(t.T(), err)
}

func (t *BookingRepoTestSuite) TestBookingRepository_CreateTicketBooking_Error() {
	testGorm := gorm.DB{Error: errors.New("test-error")}
	t.mockDB.EXPECT().Create(gomock.Any()).Return(&testGorm)
	err := t.bookingRepository.CreateTicketBooking(t.context, &entity.TicketBooking{})
	assert.NotNil(t.T(), err)
}

//func (t *BookingRepoTestSuite) TestBookingRepository_GetBooking_Success() {
//	testGorm := gorm.DB{Error: nil}
//	t.mockDB.EXPECT().Where(gomock.Any(), gomock.Any()).Return(&testGorm)
//	t.mockDB.EXPECT().Find(gomock.Any()).Return(&testGorm)
//	_, err := t.bookingRepository.GetBooking(t.context, uint(123))
//	assert.Nil(t.T(), err)
//}
