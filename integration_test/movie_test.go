package integration_test

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm/logger"
	"movie_app/config"
	"movie_app/gorm_client"
	"movie_app/handlers"
	"movie_app/models/entity"
	"movie_app/models/response"
	"movie_app/repository"
	"movie_app/service"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

var sampleMovie = entity.Movie{
	Title:    "Sample-movie",
	Language: "English",
	Duration: "2:00:00",
	Genre:    "Comedy",
	Shows:    nil,
}

type MovieHandlerTestSuite struct {
	suite.Suite
	context      *gin.Context
	engine       *gin.Engine
	recorder     *httptest.ResponseRecorder
	gormClient   gorm_client.IGormDB
	repo         repository.ISqliteRepository
	movieSvc     service.BaseService
	config       config.Configuration
	movieHandler handlers.BaseHandler
}

func TestMovieHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(MovieHandlerTestSuite))
}

func (t *MovieHandlerTestSuite) SetupTest() {
	gin.SetMode(gin.TestMode)
	t.recorder = httptest.NewRecorder()
	t.context, t.engine = gin.CreateTestContext(t.recorder)
	t.config = config.Configuration{
		Debug: true,
	}
	gorm_client.LogLevel = logger.Silent
	t.gormClient = gorm_client.NewGormDB()
	t.gormClient.Model(entity.Movie{})
	t.repo = repository.NewSqliteRepository(t.gormClient)
	t.movieSvc = service.NewBaseService(t.repo)
	t.movieHandler = handlers.NewBaseHandler(t.movieSvc, t.config)
}

func (t *MovieHandlerTestSuite) TearDownTest() {
	t.recorder.Flush()
	//goland:noinspection SqlResolve
	t.gormClient.Model(entity.Movie{}).Exec("DELETE FROM movies WHERE 1=1;")
}

func (t *MovieHandlerTestSuite) TearDownSuite() {
	err := os.Remove("./movieApp.db")
	assert.Nil(t.T(), err)
}

func (t *MovieHandlerTestSuite) TestMovieHandler_CreateMovie() {
	jsonValue, _ := json.Marshal(sampleMovie)
	t.engine.POST("/movie", t.movieHandler.CreateMovie)
	req := httptest.NewRequest(
		http.MethodPost,
		"/movie",
		bytes.NewBuffer(jsonValue))

	t.engine.ServeHTTP(t.recorder, req)
	assert.Equal(t.T(), 200, t.recorder.Code)
	assert.NotNil(t.T(), t.recorder.Body)
}

func (t *MovieHandlerTestSuite) TestMovieHandler_GetMovies() {
	// === Data population
	err := t.movieSvc.CreateMovie(context.Background(), &sampleMovie)
	assert.Nil(t.T(), err)
	//

	t.engine.GET("/movie", t.movieHandler.GetMovies)
	req := httptest.NewRequest(
		http.MethodGet,
		"/movie", nil)

	t.engine.ServeHTTP(t.recorder, req)
	assert.Equal(t.T(), 200, t.recorder.Code)
	assert.NotNil(t.T(), t.recorder.Body)

	var resp response.HttpResp[[]entity.Movie]
	err = json.Unmarshal(t.recorder.Body.Bytes(), &resp)
	assert.Nil(t.T(), err)
	assert.Equal(t.T(), 1, len(resp.Data))
}

func (t *MovieHandlerTestSuite) TestMovieHandler_UpdateMovie() {
	// === Data population
	err := t.movieSvc.CreateMovie(context.Background(), &sampleMovie)
	assert.Nil(t.T(), err)
	//

	sampleMovie := entity.Movie{
		Title:    "Sample-movie",
		Language: "English",
		Duration: "2:50:00",
		Genre:    "Comedy",
		Shows:    nil,
	}
	jsonValue, _ := json.Marshal(sampleMovie)
	t.engine.PUT("/movie/:movieId", t.movieHandler.UpdateMovie)
	req := httptest.NewRequest(
		http.MethodPut,
		"/movie/1",
		bytes.NewBuffer(jsonValue))

	t.engine.ServeHTTP(t.recorder, req)
	assert.Equal(t.T(), 200, t.recorder.Code)
	assert.NotNil(t.T(), t.recorder.Body)
}

func (t *MovieHandlerTestSuite) TestMovieHandler_DeleteMovie() {
	// === Data population
	err := t.movieSvc.CreateMovie(context.Background(), &sampleMovie)
	assert.Nil(t.T(), err)
	//

	t.engine.DELETE("/movie/:movieId", t.movieHandler.DeleteMovie)
	req := httptest.NewRequest(
		http.MethodDelete,
		"/movie/1", nil)

	t.engine.ServeHTTP(t.recorder, req)
	assert.Equal(t.T(), 200, t.recorder.Code)
	assert.NotNil(t.T(), t.recorder.Body)
}
