// Code generated by MockGen. DO NOT EDIT.
// Source: service/base_service.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	entity "movie_app/models/entity"
	request "movie_app/models/request"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockIBaseService is a mock of IBaseService interface.
type MockIBaseService struct {
	ctrl     *gomock.Controller
	recorder *MockIBaseServiceMockRecorder
}

// MockIBaseServiceMockRecorder is the mock recorder for MockIBaseService.
type MockIBaseServiceMockRecorder struct {
	mock *MockIBaseService
}

// NewMockIBaseService creates a new mock instance.
func NewMockIBaseService(ctrl *gomock.Controller) *MockIBaseService {
	mock := &MockIBaseService{ctrl: ctrl}
	mock.recorder = &MockIBaseServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIBaseService) EXPECT() *MockIBaseServiceMockRecorder {
	return m.recorder
}

// BookShow mocks base method.
func (m *MockIBaseService) BookShow(ctx context.Context, booking request.BookingRequest) (entity.Booking, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BookShow", ctx, booking)
	ret0, _ := ret[0].(entity.Booking)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BookShow indicates an expected call of BookShow.
func (mr *MockIBaseServiceMockRecorder) BookShow(ctx, booking interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BookShow", reflect.TypeOf((*MockIBaseService)(nil).BookShow), ctx, booking)
}

// CreateMovie mocks base method.
func (m *MockIBaseService) CreateMovie(ctx context.Context, movie *entity.Movie) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateMovie", ctx, movie)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateMovie indicates an expected call of CreateMovie.
func (mr *MockIBaseServiceMockRecorder) CreateMovie(ctx, movie interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateMovie", reflect.TypeOf((*MockIBaseService)(nil).CreateMovie), ctx, movie)
}

// CreateShow mocks base method.
func (m *MockIBaseService) CreateShow(ctx context.Context, show *entity.Show) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateShow", ctx, show)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateShow indicates an expected call of CreateShow.
func (mr *MockIBaseServiceMockRecorder) CreateShow(ctx, show interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateShow", reflect.TypeOf((*MockIBaseService)(nil).CreateShow), ctx, show)
}

// CreateTheatre mocks base method.
func (m *MockIBaseService) CreateTheatre(ctx context.Context, theatre *entity.Theatre) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTheatre", ctx, theatre)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateTheatre indicates an expected call of CreateTheatre.
func (mr *MockIBaseServiceMockRecorder) CreateTheatre(ctx, theatre interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTheatre", reflect.TypeOf((*MockIBaseService)(nil).CreateTheatre), ctx, theatre)
}

// DeleteMovie mocks base method.
func (m *MockIBaseService) DeleteMovie(ctx context.Context, movieId string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteMovie", ctx, movieId)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteMovie indicates an expected call of DeleteMovie.
func (mr *MockIBaseServiceMockRecorder) DeleteMovie(ctx, movieId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteMovie", reflect.TypeOf((*MockIBaseService)(nil).DeleteMovie), ctx, movieId)
}

// DisplayMoviesByTheatre mocks base method.
func (m *MockIBaseService) DisplayMoviesByTheatre(ctx context.Context, theatreId string) ([]entity.Movie, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DisplayMoviesByTheatre", ctx, theatreId)
	ret0, _ := ret[0].([]entity.Movie)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DisplayMoviesByTheatre indicates an expected call of DisplayMoviesByTheatre.
func (mr *MockIBaseServiceMockRecorder) DisplayMoviesByTheatre(ctx, theatreId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DisplayMoviesByTheatre", reflect.TypeOf((*MockIBaseService)(nil).DisplayMoviesByTheatre), ctx, theatreId)
}

// DisplayTheatreByMovies mocks base method.
func (m *MockIBaseService) DisplayTheatreByMovies(ctx context.Context, movieId string) ([]entity.Theatre, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DisplayTheatreByMovies", ctx, movieId)
	ret0, _ := ret[0].([]entity.Theatre)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DisplayTheatreByMovies indicates an expected call of DisplayTheatreByMovies.
func (mr *MockIBaseServiceMockRecorder) DisplayTheatreByMovies(ctx, movieId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DisplayTheatreByMovies", reflect.TypeOf((*MockIBaseService)(nil).DisplayTheatreByMovies), ctx, movieId)
}

// GetBooking mocks base method.
func (m *MockIBaseService) GetBooking(ctx context.Context, bookingId string) (entity.Booking, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBooking", ctx, bookingId)
	ret0, _ := ret[0].(entity.Booking)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBooking indicates an expected call of GetBooking.
func (mr *MockIBaseServiceMockRecorder) GetBooking(ctx, bookingId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBooking", reflect.TypeOf((*MockIBaseService)(nil).GetBooking), ctx, bookingId)
}

// GetMovies mocks base method.
func (m *MockIBaseService) GetMovies(ctx context.Context, movieId string) ([]entity.Movie, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMovies", ctx, movieId)
	ret0, _ := ret[0].([]entity.Movie)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMovies indicates an expected call of GetMovies.
func (mr *MockIBaseServiceMockRecorder) GetMovies(ctx, movieId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMovies", reflect.TypeOf((*MockIBaseService)(nil).GetMovies), ctx, movieId)
}

// GetShows mocks base method.
func (m *MockIBaseService) GetShows(ctx context.Context, showName string) ([]entity.Show, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetShows", ctx, showName)
	ret0, _ := ret[0].([]entity.Show)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetShows indicates an expected call of GetShows.
func (mr *MockIBaseServiceMockRecorder) GetShows(ctx, showName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetShows", reflect.TypeOf((*MockIBaseService)(nil).GetShows), ctx, showName)
}

// GetTheatres mocks base method.
func (m *MockIBaseService) GetTheatres(ctx context.Context, theatreName string) ([]entity.Theatre, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTheatres", ctx, theatreName)
	ret0, _ := ret[0].([]entity.Theatre)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTheatres indicates an expected call of GetTheatres.
func (mr *MockIBaseServiceMockRecorder) GetTheatres(ctx, theatreName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTheatres", reflect.TypeOf((*MockIBaseService)(nil).GetTheatres), ctx, theatreName)
}

// GetTicketsByShowId mocks base method.
func (m *MockIBaseService) GetTicketsByShowId(ctx context.Context, showId string) ([]entity.Ticket, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTicketsByShowId", ctx, showId)
	ret0, _ := ret[0].([]entity.Ticket)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTicketsByShowId indicates an expected call of GetTicketsByShowId.
func (mr *MockIBaseServiceMockRecorder) GetTicketsByShowId(ctx, showId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTicketsByShowId", reflect.TypeOf((*MockIBaseService)(nil).GetTicketsByShowId), ctx, showId)
}

// UpdateMovie mocks base method.
func (m *MockIBaseService) UpdateMovie(ctx context.Context, movie *entity.Movie, movieId string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateMovie", ctx, movie, movieId)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateMovie indicates an expected call of UpdateMovie.
func (mr *MockIBaseServiceMockRecorder) UpdateMovie(ctx, movie, movieId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateMovie", reflect.TypeOf((*MockIBaseService)(nil).UpdateMovie), ctx, movie, movieId)
}

// UpdateShow mocks base method.
func (m *MockIBaseService) UpdateShow(ctx context.Context, show *entity.Show, showId string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateShow", ctx, show, showId)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateShow indicates an expected call of UpdateShow.
func (mr *MockIBaseServiceMockRecorder) UpdateShow(ctx, show, showId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateShow", reflect.TypeOf((*MockIBaseService)(nil).UpdateShow), ctx, show, showId)
}