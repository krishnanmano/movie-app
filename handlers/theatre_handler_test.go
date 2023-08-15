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
	"net/http"
	"net/http/httptest"
	"testing"
)

type TheatreHandlerTestSuite struct {
	suite.Suite
	mockCtrl       *gomock.Controller
	context        *gin.Context
	engine         *gin.Engine
	recorder       *httptest.ResponseRecorder
	theatreSvc     *mocks.MockIBaseService
	config         config.Configuration
	theatreHandler IBaseHandler
}

func TestTheatreHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(TheatreHandlerTestSuite))
}

func (t *TheatreHandlerTestSuite) SetupTest() {
	gin.SetMode(gin.TestMode)
	t.mockCtrl = gomock.NewController(t.T())
	t.recorder = httptest.NewRecorder()
	t.context, t.engine = gin.CreateTestContext(t.recorder)
	t.config = config.Configuration{}
	t.theatreSvc = mocks.NewMockIBaseService(t.mockCtrl)
	t.theatreHandler = NewBaseHandler(t.theatreSvc, t.config)
}

func (t *TheatreHandlerTestSuite) TearDownTest() {
	t.mockCtrl.Finish()
	t.recorder.Flush()
}

func (t *TheatreHandlerTestSuite) TestTheatreHandler_CreateTheatre() {

	Theatre := entity.Theatre{
		Name:         "sample-theatre",
		Address:      "sample-address",
		TotalColumns: 5,
		TotalRows:    5,
		Shows:        nil,
		Seats:        nil,
	}

	jsonValue, _ := json.Marshal(Theatre)
	t.theatreSvc.EXPECT().CreateTheatre(gomock.Any(), gomock.Any()).Return(nil)
	t.engine.POST("/theatre", t.theatreHandler.CreateTheatre)
	req, _ := http.NewRequest(
		http.MethodPost,
		"/theatre",
		bytes.NewBuffer(jsonValue))

	t.engine.ServeHTTP(t.recorder, req)
	assert.Equal(t.T(), 200, t.recorder.Code)
	assert.NotNil(t.T(), t.recorder.Body)
}

func (t *TheatreHandlerTestSuite) TestTheatreHandler_GetTheatres() {

	Theatres := []entity.Theatre{{
		Name:         "sample-theatre",
		Address:      "sample-address",
		TotalColumns: 5,
		TotalRows:    5,
		Shows:        nil,
		Seats:        nil,
	},
	}

	t.theatreSvc.EXPECT().GetTheatres(gomock.Any(), gomock.Any()).Return(Theatres, nil)
	t.engine.GET("/theatre", t.theatreHandler.GetTheatres)
	req, _ := http.NewRequest(
		http.MethodGet,
		"/theatre",
		nil)

	t.engine.ServeHTTP(t.recorder, req)
	assert.Equal(t.T(), 200, t.recorder.Code)
	assert.NotNil(t.T(), t.recorder.Body)
}
