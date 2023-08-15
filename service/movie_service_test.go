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

type MovieServiceTestSuite struct {
	suite.Suite
	ctx          context.Context
	mockCtrl     *gomock.Controller
	movieRepo    *mocks.MockISqliteRepository
	movieService IBaseService
}

func TestMovieHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(MovieServiceTestSuite))
}

func (t *MovieServiceTestSuite) SetupTest() {
	gin.SetMode(gin.TestMode)
	t.mockCtrl = gomock.NewController(t.T())
	t.ctx = context.WithValue(context.Background(), "logger", nil)
	t.movieRepo = mocks.NewMockISqliteRepository(t.mockCtrl)
	t.movieService = NewBaseService(t.movieRepo)
}

func (t *MovieServiceTestSuite) TearDownTest() {
	t.mockCtrl.Finish()
}

func (t *MovieServiceTestSuite) TestMovieService_CreateMovie() {
	movie := entity.Movie{
		Title:    "Sample-movie",
		Language: "English",
		Duration: "2:00:00",
		Genre:    "Comedy",
		Shows:    nil,
	}
	t.movieRepo.EXPECT().CreateMovie(gomock.Any(), gomock.Any()).Return(nil)
	err := t.movieService.CreateMovie(t.ctx, &movie)
	assert.Nil(t.T(), err)
}

func (t *MovieServiceTestSuite) TestMovieService_GetMovies() {
	movies := []entity.Movie{
		{
			Title:    "Sample-movie",
			Language: "English",
			Duration: "2:00:00",
			Genre:    "Comedy",
			Shows:    nil,
		},
	}
	t.movieRepo.EXPECT().GetMovie(gomock.Any(), gomock.Any()).Return(movies[0], nil)
	actualMovies, err := t.movieService.GetMovies(t.ctx, "sample-movie")
	assert.Equal(t.T(), actualMovies, movies)
	assert.Nil(t.T(), err)
}

func (t *MovieServiceTestSuite) TestMovieService_DeleteMovie() {
	t.movieRepo.EXPECT().DeleteMovie(gomock.Any(), gomock.Any()).Return(nil)
	actualErr := t.movieService.DeleteMovie(t.ctx, "sample-movie")
	assert.Equal(t.T(), nil, actualErr)
}

func (t *MovieServiceTestSuite) TestMovieService_UpdateMovie() {
	movie := entity.Movie{
		Title:    "Sample-movie",
		Language: "English",
		Duration: "2:00:00",
		Genre:    "Comedy",
		Shows:    nil,
	}
	t.movieRepo.EXPECT().UpdateMovie(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
	actualErr := t.movieService.UpdateMovie(t.ctx, &movie, "sample-movie")
	assert.Equal(t.T(), nil, actualErr)
}
