// Code generated by MockGen. DO NOT EDIT.
// Source: gorm_client/gormDB.go

// Package mocks is a generated GoMock package.
package mocks

import (
	gorm_client "movie_app/gorm_client"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	gorm "gorm.io/gorm"
)

// MockIGormDB is a mock of IGormDB interface.
type MockIGormDB struct {
	ctrl     *gomock.Controller
	recorder *MockIGormDBMockRecorder
}

// MockIGormDBMockRecorder is the mock recorder for MockIGormDB.
type MockIGormDBMockRecorder struct {
	mock *MockIGormDB
}

// NewMockIGormDB creates a new mock instance.
func NewMockIGormDB(ctrl *gomock.Controller) *MockIGormDB {
	mock := &MockIGormDB{ctrl: ctrl}
	mock.recorder = &MockIGormDBMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIGormDB) EXPECT() *MockIGormDBMockRecorder {
	return m.recorder
}

// AutoMigrate mocks base method.
func (m *MockIGormDB) AutoMigrate(dst ...interface{}) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range dst {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "AutoMigrate", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// AutoMigrate indicates an expected call of AutoMigrate.
func (mr *MockIGormDBMockRecorder) AutoMigrate(dst ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AutoMigrate", reflect.TypeOf((*MockIGormDB)(nil).AutoMigrate), dst...)
}

// Connect mocks base method.
func (m *MockIGormDB) Connect() gorm_client.IGormDB {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Connect")
	ret0, _ := ret[0].(gorm_client.IGormDB)
	return ret0
}

// Connect indicates an expected call of Connect.
func (mr *MockIGormDBMockRecorder) Connect() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Connect", reflect.TypeOf((*MockIGormDB)(nil).Connect))
}

// Create mocks base method.
func (m *MockIGormDB) Create(value interface{}) *gorm.DB {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", value)
	ret0, _ := ret[0].(*gorm.DB)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockIGormDBMockRecorder) Create(value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockIGormDB)(nil).Create), value)
}

// Find mocks base method.
func (m *MockIGormDB) Find(out interface{}, where ...interface{}) *gorm.DB {
	m.ctrl.T.Helper()
	varargs := []interface{}{out}
	for _, a := range where {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Find", varargs...)
	ret0, _ := ret[0].(*gorm.DB)
	return ret0
}

// Find indicates an expected call of Find.
func (mr *MockIGormDBMockRecorder) Find(out interface{}, where ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{out}, where...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockIGormDB)(nil).Find), varargs...)
}

// Model mocks base method.
func (m *MockIGormDB) Model(value interface{}) *gorm.DB {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Model", value)
	ret0, _ := ret[0].(*gorm.DB)
	return ret0
}

// Model indicates an expected call of Model.
func (mr *MockIGormDBMockRecorder) Model(value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Model", reflect.TypeOf((*MockIGormDB)(nil).Model), value)
}

// Select mocks base method.
func (m *MockIGormDB) Select(query interface{}, args ...interface{}) *gorm.DB {
	m.ctrl.T.Helper()
	varargs := []interface{}{query}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Select", varargs...)
	ret0, _ := ret[0].(*gorm.DB)
	return ret0
}

// Select indicates an expected call of Select.
func (mr *MockIGormDBMockRecorder) Select(query interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{query}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Select", reflect.TypeOf((*MockIGormDB)(nil).Select), varargs...)
}

// Where mocks base method.
func (m *MockIGormDB) Where(query interface{}, args ...interface{}) *gorm.DB {
	m.ctrl.T.Helper()
	varargs := []interface{}{query}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Where", varargs...)
	ret0, _ := ret[0].(*gorm.DB)
	return ret0
}

// Where indicates an expected call of Where.
func (mr *MockIGormDBMockRecorder) Where(query interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{query}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Where", reflect.TypeOf((*MockIGormDB)(nil).Where), varargs...)
}