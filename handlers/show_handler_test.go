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
	"time"
)

type ShowHandlerTestSuite struct {
	suite.Suite
	mockCtrl    *gomock.Controller
	context     *gin.Context
	engine      *gin.Engine
	recorder    *httptest.ResponseRecorder
	showSvc     *mocks.MockIBaseService
	config      config.Configuration
	showHandler IBaseHandler
}

func TestShowHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(ShowHandlerTestSuite))
}

func (t *ShowHandlerTestSuite) SetupTest() {
	gin.SetMode(gin.TestMode)
	t.mockCtrl = gomock.NewController(t.T())
	t.recorder = httptest.NewRecorder()
	t.context, t.engine = gin.CreateTestContext(t.recorder)
	t.config = config.Configuration{}
	t.showSvc = mocks.NewMockIBaseService(t.mockCtrl)
	t.showHandler = NewBaseHandler(t.showSvc, t.config)
}

func (t *ShowHandlerTestSuite) TearDownTest() {
	t.mockCtrl.Finish()
	t.recorder.Flush()
}

func (t *ShowHandlerTestSuite) TestShowHandler_CreateShow() {

	show := entity.Show{
		TheatreId: 1,
		MovieId:   1,
		CreatedAt: time.Time{},
		StartTime: time.Time{},
		EndTime:   time.Time{},
		Tickets:   nil,
	}

	jsonValue, _ := json.Marshal(show)
	t.showSvc.EXPECT().CreateShow(gomock.Any(), gomock.Any()).Return(nil)
	t.engine.POST("/show", t.showHandler.CreateShow)
	req, _ := http.NewRequest(
		http.MethodPost,
		"/show",
		bytes.NewBuffer(jsonValue))

	t.engine.ServeHTTP(t.recorder, req)
	assert.Equal(t.T(), 200, t.recorder.Code)
	assert.NotNil(t.T(), t.recorder.Body)
}

func (t *ShowHandlerTestSuite) TestShowHandler_GetAllShow() {

	show := []entity.Show{{
		TheatreId: 1,
		MovieId:   1,
		CreatedAt: time.Time{},
		StartTime: time.Time{},
		EndTime:   time.Time{},
		Tickets:   nil,
	},
	}
	t.showSvc.EXPECT().GetShows(gomock.Any(), gomock.Any()).Return(show, nil)
	t.engine.GET("/show", t.showHandler.GetAllShows)
	req, _ := http.NewRequest(
		http.MethodGet,
		"/show",
		nil)

	t.engine.ServeHTTP(t.recorder, req)
	assert.Equal(t.T(), 200, t.recorder.Code)
	assert.NotNil(t.T(), t.recorder.Body)
}

func (t *ShowHandlerTestSuite) TestShowHandler_UpdateShow() {

	show := entity.Show{
		TheatreId: 1,
		MovieId:   1,
		CreatedAt: time.Time{},
		StartTime: time.Time{},
		EndTime:   time.Time{},
		Tickets:   nil,
	}

	jsonValue, _ := json.Marshal(show)
	t.showSvc.EXPECT().UpdateShow(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
	t.engine.PUT("/show/:showId", t.showHandler.UpdateShow)
	req, _ := http.NewRequest(
		http.MethodPut,
		"/show/:showId",
		bytes.NewBuffer(jsonValue))

	t.engine.ServeHTTP(t.recorder, req)
	assert.Equal(t.T(), 200, t.recorder.Code)
	assert.NotNil(t.T(), t.recorder.Body)
}
