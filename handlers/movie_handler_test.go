package handlers

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
	"movie_app/config"
	"movie_app/models/entity"
	"movie_app/service/mocks"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestBaseHandler_CreateMovie(t *testing.T) {

	mockMovieService := mocks.NewBaseService(t)
	mockMovieService.On("CreateMovie", mock.Anything, mock.Anything).Return(nil)

	movieHandler := NewBaseHandler(mockMovieService, config.Configuration{})

	rec := httptest.NewRecorder()
	gin.SetMode(gin.TestMode)
	context, _ := gin.CreateTestContext(rec)

	movie := entity.Movie{
		Model:    gorm.Model{},
		Title:    "Test Movie",
		Language: "EN",
		Duration: "120",
		Genre:    "Adventure",
		Shows:    nil,
	}
	movieJson, err := json.Marshal(movie)
	assert.Nil(t, err)

	context.Request = httptest.NewRequest(http.MethodPost, "/movie", bytes.NewReader(movieJson))
	movieHandler.CreateMovie(context)

	assert.Equal(t, http.StatusOK, rec.Code)
}
