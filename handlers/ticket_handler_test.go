package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"movie_app/config"
	"movie_app/mocks"
	"movie_app/models/entity"
	"net/http"
	"net/http/httptest"
	"testing"
)

type TicketHandlerTestSuite struct {
	suite.Suite
	mockCtrl      *gomock.Controller
	context       *gin.Context
	engine        *gin.Engine
	recorder      *httptest.ResponseRecorder
	ticketSvc     *mocks.MockIBaseService
	config        config.Configuration
	ticketHandler IBaseHandler
}

func TestTicketHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(TicketHandlerTestSuite))
}

func (t *TicketHandlerTestSuite) SetupTest() {
	gin.SetMode(gin.TestMode)
	t.mockCtrl = gomock.NewController(t.T())
	t.recorder = httptest.NewRecorder()
	t.context, t.engine = gin.CreateTestContext(t.recorder)
	t.config = config.Configuration{}
	t.ticketSvc = mocks.NewMockIBaseService(t.mockCtrl)
	t.ticketHandler = NewBaseHandler(t.ticketSvc, t.config)
}

func (t *TicketHandlerTestSuite) TearDownTest() {
	t.mockCtrl.Finish()
	t.recorder.Flush()
}

func (t *TicketHandlerTestSuite) TestTicketHandler_GetSeatMatrix() {

	seats := []entity.Ticket{{
		SeatId: 5,
	},
	}

	t.ticketSvc.EXPECT().GetTicketsByShowId(gomock.Any(), gomock.Any()).Return(seats, nil)
	t.engine.GET("/ticket", t.ticketHandler.GetSeatMatrix)
	req, _ := http.NewRequest(
		http.MethodGet,
		"/ticket",
		nil)

	t.engine.ServeHTTP(t.recorder, req)
	assert.Equal(t.T(), 200, t.recorder.Code)
	assert.NotNil(t.T(), t.recorder.Body)
}
