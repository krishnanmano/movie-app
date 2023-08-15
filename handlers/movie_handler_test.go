package handlers

import (
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"movie_app/config"
	"movie_app/mocks"
	"movie_app/models/entity"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
)

type MovieHandlerTestSuite struct {
	suite.Suite
	mockCtrl     *gomock.Controller
	context      *gin.Context
	engine       *gin.Engine
	recorder     *httptest.ResponseRecorder
	movieSvc     *mocks.MockIBaseService
	config       config.Configuration
	movieHandler IBaseHandler
}

func TestMovieHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(MovieHandlerTestSuite))
}

func (t *MovieHandlerTestSuite) SetupTest() {
	gin.SetMode(gin.TestMode)
	t.mockCtrl = gomock.NewController(t.T())
	t.recorder = httptest.NewRecorder()
	t.context, t.engine = gin.CreateTestContext(t.recorder)
	t.config = config.Configuration{
		Debug: true,
	}
	t.movieSvc = mocks.NewMockIBaseService(t.mockCtrl)
	t.movieHandler = NewBaseHandler(t.movieSvc, t.config)
}

func (t *MovieHandlerTestSuite) TearDownTest() {
	t.mockCtrl.Finish()
	t.recorder.Flush()
}

func (t *MovieHandlerTestSuite) TestMovieHandler_CreateMovie() {

	sampleMovie := entity.Movie{
		Title:    "Sample-movie",
		Language: "English",
		Duration: "2:00:00",
		Genre:    "Comedy",
		Shows:    nil,
	}
	jsonValue, _ := json.Marshal(sampleMovie)
	t.movieSvc.EXPECT().CreateMovie(gomock.Any(), gomock.Any()).Return(nil)
	t.engine.POST("/movie", t.movieHandler.CreateMovie)
	req, _ := http.NewRequest(
		http.MethodPost,
		"/movie",
		bytes.NewBuffer(jsonValue))

	t.engine.ServeHTTP(t.recorder, req)
	assert.Equal(t.T(), 200, t.recorder.Code)
	assert.NotNil(t.T(), t.recorder.Body)

}

func (t *MovieHandlerTestSuite) TestMovieHandler_GetMovies() {

	movies := []entity.Movie{
		{
			Title:    "Sample-movie",
			Language: "English",
			Duration: "2:00:00",
			Genre:    "Comedy",
			Shows:    nil,
		},
	}

	t.movieSvc.EXPECT().GetMovies(gomock.Any(), gomock.Any()).Return(movies, nil)
	t.engine.GET("/movie", t.movieHandler.GetMovies)
	req, _ := http.NewRequest(
		http.MethodGet,
		"/movie", nil)

	t.engine.ServeHTTP(t.recorder, req)
	assert.Equal(t.T(), 200, t.recorder.Code)
	assert.NotNil(t.T(), t.recorder.Body)

}

func (t *MovieHandlerTestSuite) TestMovieHandler_UpdateMovie() {

	sampleMovie := entity.Movie{
		Title:    "Sample-movie",
		Language: "English",
		Duration: "2:00:00",
		Genre:    "Comedy",
		Shows:    nil,
	}
	jsonValue, _ := json.Marshal(sampleMovie)
	t.movieSvc.EXPECT().UpdateMovie(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
	t.engine.PUT("/movie/:movieName", t.movieHandler.UpdateMovie)
	req, _ := http.NewRequest(
		http.MethodPut,
		"/movie/:movieName",
		bytes.NewBuffer(jsonValue))

	t.engine.ServeHTTP(t.recorder, req)
	assert.Equal(t.T(), 200, t.recorder.Code)
	assert.NotNil(t.T(), t.recorder.Body)

}

func (t *MovieHandlerTestSuite) TestMovieHandler_DeleteMovie() {

	t.movieSvc.EXPECT().DeleteMovie(gomock.Any(), gomock.Any()).Return(nil)
	t.engine.DELETE("/movie/:movieName", t.movieHandler.DeleteMovie)
	req, _ := http.NewRequest(
		http.MethodDelete,
		"/movie/:movieName", nil)

	t.engine.ServeHTTP(t.recorder, req)
	assert.Equal(t.T(), 200, t.recorder.Code)
	assert.NotNil(t.T(), t.recorder.Body)

}
