package service

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"movie_app/mocks"
	"movie_app/models/entity"
	"testing"
)

type TheaterServiceTestSuite struct {
	suite.Suite
	ctx            context.Context
	mockCtrl       *gomock.Controller
	TheaterRepo    *mocks.MockISqliteRepository
	TheaterService IBaseService
}

func TestTheaterServiceTestSuite(t *testing.T) {
	suite.Run(t, new(TheaterServiceTestSuite))
}

func (t *TheaterServiceTestSuite) SetupTest() {
	gin.SetMode(gin.TestMode)
	t.mockCtrl = gomock.NewController(t.T())
	t.ctx = context.WithValue(context.Background(), "logger", nil)
	t.TheaterRepo = mocks.NewMockISqliteRepository(t.mockCtrl)
	t.TheaterService = NewBaseService(t.TheaterRepo)
}

func (t *TheaterServiceTestSuite) TearDownTest() {
	t.mockCtrl.Finish()
}
func (t *TheaterServiceTestSuite) TestTheaterService_CreateTheatre() {

	theater := entity.Theatre{Name: "INOX"}
	t.TheaterRepo.EXPECT().CreateTheatre(gomock.Any(), gomock.Any()).Return(nil)
	t.TheaterRepo.EXPECT().CreateSeats(gomock.Any(), gomock.Any()).Return(nil)
	err := t.TheaterService.CreateTheatre(t.ctx, &theater)
	assert.Nil(t.T(), err)

}

func (t *TheaterServiceTestSuite) TestTheaterService_GetTheatres() {

	theaters := []entity.Theatre{{Name: "INOX"}}
	t.TheaterRepo.EXPECT().GetTheatre(gomock.Any(), gomock.Any()).Return(theaters, nil)
	//t.TheaterRepo.EXPECT().GetTheatres(gomock.Any()).Return(theaters, nil)
	_, err := t.TheaterService.GetTheatres(t.ctx, "abc")
	assert.Nil(t.T(), err)

}
