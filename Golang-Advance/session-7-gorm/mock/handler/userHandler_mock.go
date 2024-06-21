// Code generated by MockGen. DO NOT EDIT.
// Source: ./handler/userHandler.go

// Package mock_handler is a generated GoMock package.
package mock_handler

import (
	reflect "reflect"

	gin "github.com/gin-gonic/gin"
	gomock "github.com/golang/mock/gomock"
)

// MockIUserHandler is a mock of IUserHandler interface.
type MockIUserHandler struct {
	ctrl     *gomock.Controller
	recorder *MockIUserHandlerMockRecorder
}

// MockIUserHandlerMockRecorder is the mock recorder for MockIUserHandler.
type MockIUserHandlerMockRecorder struct {
	mock *MockIUserHandler
}

// NewMockIUserHandler creates a new mock instance.
func NewMockIUserHandler(ctrl *gomock.Controller) *MockIUserHandler {
	mock := &MockIUserHandler{ctrl: ctrl}
	mock.recorder = &MockIUserHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIUserHandler) EXPECT() *MockIUserHandlerMockRecorder {
	return m.recorder
}

// CreateUser mocks base method.
func (m *MockIUserHandler) CreateUser(c *gin.Context) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "CreateUser", c)
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockIUserHandlerMockRecorder) CreateUser(c interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockIUserHandler)(nil).CreateUser), c)
}

// DeleteUser mocks base method.
func (m *MockIUserHandler) DeleteUser(c *gin.Context) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "DeleteUser", c)
}

// DeleteUser indicates an expected call of DeleteUser.
func (mr *MockIUserHandlerMockRecorder) DeleteUser(c interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUser", reflect.TypeOf((*MockIUserHandler)(nil).DeleteUser), c)
}

// GetAllUsers mocks base method.
func (m *MockIUserHandler) GetAllUsers(c *gin.Context) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "GetAllUsers", c)
}

// GetAllUsers indicates an expected call of GetAllUsers.
func (mr *MockIUserHandlerMockRecorder) GetAllUsers(c interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllUsers", reflect.TypeOf((*MockIUserHandler)(nil).GetAllUsers), c)
}

// GetUser mocks base method.
func (m *MockIUserHandler) GetUser(c *gin.Context) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "GetUser", c)
}

// GetUser indicates an expected call of GetUser.
func (mr *MockIUserHandlerMockRecorder) GetUser(c interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUser", reflect.TypeOf((*MockIUserHandler)(nil).GetUser), c)
}

// UpdateUser mocks base method.
func (m *MockIUserHandler) UpdateUser(c *gin.Context) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "UpdateUser", c)
}

// UpdateUser indicates an expected call of UpdateUser.
func (mr *MockIUserHandlerMockRecorder) UpdateUser(c interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUser", reflect.TypeOf((*MockIUserHandler)(nil).UpdateUser), c)
}
