// Code generated by mockery v2.32.4. DO NOT EDIT.

package mocks

import (
	gin "github.com/gin-gonic/gin"

	mock "github.com/stretchr/testify/mock"
)

// BaseHandler is an autogenerated mock type for the BaseHandler type
type BaseHandler struct {
	mock.Mock
}

// BookShow provides a mock function with given fields: c
func (_m *BaseHandler) BookShow(c *gin.Context) {
	_m.Called(c)
}

// CreateMovie provides a mock function with given fields: c
func (_m *BaseHandler) CreateMovie(c *gin.Context) {
	_m.Called(c)
}

// CreateShow provides a mock function with given fields: c
func (_m *BaseHandler) CreateShow(c *gin.Context) {
	_m.Called(c)
}

// CreateTheatre provides a mock function with given fields: c
func (_m *BaseHandler) CreateTheatre(c *gin.Context) {
	_m.Called(c)
}

// DeleteMovie provides a mock function with given fields: c
func (_m *BaseHandler) DeleteMovie(c *gin.Context) {
	_m.Called(c)
}

// GetAllShows provides a mock function with given fields: c
func (_m *BaseHandler) GetAllShows(c *gin.Context) {
	_m.Called(c)
}

// GetBooking provides a mock function with given fields: c
func (_m *BaseHandler) GetBooking(c *gin.Context) {
	_m.Called(c)
}

// GetMovies provides a mock function with given fields: c
func (_m *BaseHandler) GetMovies(c *gin.Context) {
	_m.Called(c)
}

// GetSeatMatrix provides a mock function with given fields: c
func (_m *BaseHandler) GetSeatMatrix(c *gin.Context) {
	_m.Called(c)
}

// GetTheatres provides a mock function with given fields: c
func (_m *BaseHandler) GetTheatres(c *gin.Context) {
	_m.Called(c)
}

// UpdateMovie provides a mock function with given fields: c
func (_m *BaseHandler) UpdateMovie(c *gin.Context) {
	_m.Called(c)
}

// UpdateShow provides a mock function with given fields: c
func (_m *BaseHandler) UpdateShow(c *gin.Context) {
	_m.Called(c)
}

// NewBaseHandler creates a new instance of BaseHandler. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewBaseHandler(t interface {
	mock.TestingT
	Cleanup(func())
}) *BaseHandler {
	mock := &BaseHandler{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
