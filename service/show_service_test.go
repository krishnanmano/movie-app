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

type ShowServiceTestSuite struct {
	suite.Suite
	ctx         context.Context
	mockCtrl    *gomock.Controller
	ShowRepo    *mocks.MockISqliteRepository
	ShowService IBaseService
}

func TestShowServiceTestSuite(t *testing.T) {
	suite.Run(t, new(ShowServiceTestSuite))
}

func (t *ShowServiceTestSuite) SetupTest() {
	gin.SetMode(gin.TestMode)
	t.mockCtrl = gomock.NewController(t.T())
	t.ctx = context.WithValue(context.Background(), "logger", nil)
	t.ShowRepo = mocks.NewMockISqliteRepository(t.mockCtrl)
	t.ShowService = NewBaseService(t.ShowRepo)
}

func (t *ShowServiceTestSuite) TearDownTest() {
	t.mockCtrl.Finish()
}

func (t *ShowServiceTestSuite) TestCreateShow() {

	seats := []entity.Seat{{
		TheatreId: 123,
	},
	}
	show := entity.Show{
		TheatreId: 123,
	}
	t.ShowRepo.EXPECT().CreateShow(gomock.Any(), gomock.Any()).Return(nil)
	t.ShowRepo.EXPECT().GetSeats(gomock.Any(), gomock.Any()).Return(seats, nil)
	t.ShowRepo.EXPECT().CreateTickets(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
	err := t.ShowService.CreateShow(t.ctx, &show)
	assert.Nil(t.T(), err)

}

func (t *ShowServiceTestSuite) TestGetShows() {
	shows := []entity.Show{{
		TheatreId: 123,
	},
	}
	t.ShowRepo.EXPECT().GetShows(gomock.Any()).Return(shows, nil)
	_, err := t.ShowService.GetShows(t.ctx, "sample-show")
	assert.Nil(t.T(), err)

}

func (t *ShowServiceTestSuite) TestUpdateShows() {
	show := entity.Show{
		TheatreId: 123,
	}
	t.ShowRepo.EXPECT().UpdateShow(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
	err := t.ShowService.UpdateShow(t.ctx, &show, "123")
	assert.Nil(t.T(), err)
}
