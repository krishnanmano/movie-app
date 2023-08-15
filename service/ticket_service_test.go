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

type TicketServiceTestSuite struct {
	suite.Suite
	ctx           context.Context
	mockCtrl      *gomock.Controller
	TicketRepo    *mocks.MockISqliteRepository
	TicketService IBaseService
}

func TestTicketServiceTestSuite(t *testing.T) {
	suite.Run(t, new(TicketServiceTestSuite))
}

func (t *TicketServiceTestSuite) SetupTest() {
	gin.SetMode(gin.TestMode)
	t.mockCtrl = gomock.NewController(t.T())
	t.ctx = context.WithValue(context.Background(), "logger", nil)
	t.TicketRepo = mocks.NewMockISqliteRepository(t.mockCtrl)
	t.TicketService = NewBaseService(t.TicketRepo)
}

func (t *TicketServiceTestSuite) TearDownTest() {
	t.mockCtrl.Finish()
}

func (t *TicketServiceTestSuite) TestTicketService_GetTicketsByShowId() {
	tickets := []entity.Ticket{{
		Price:  500,
		SeatId: 12,
	},
	}
	t.TicketRepo.EXPECT().GetTicketsByShowId(gomock.Any(), gomock.Any()).Return(tickets, nil)
	_, err := t.TicketService.GetTicketsByShowId(t.ctx, "123")
	assert.Nil(t.T(), err)
}
