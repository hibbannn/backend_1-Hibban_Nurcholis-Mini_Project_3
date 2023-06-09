// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/hibbannn/backend_1-Hibban_Nurcholis-Mini_Project_2/repositories (interfaces: UserRepositoryInterface)

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	entities "github.com/hibbannn/backend_1-Hibban_Nurcholis-Mini_Project_2/entities"
	reflect "reflect"
)

// MockUserRepositoryInterface is a mock of UserRepositoryInterface interface
type MockUserRepositoryInterface struct {
	ctrl     *gomock.Controller
	recorder *MockUserRepositoryInterfaceMockRecorder
}

// MockUserRepositoryInterfaceMockRecorder is the mock recorder for MockUserRepositoryInterface
type MockUserRepositoryInterfaceMockRecorder struct {
	mock *MockUserRepositoryInterface
}

// NewMockUserRepositoryInterface creates a new mock instance
func NewMockUserRepositoryInterface(ctrl *gomock.Controller) *MockUserRepositoryInterface {
	mock := &MockUserRepositoryInterface{ctrl: ctrl}
	mock.recorder = &MockUserRepositoryInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockUserRepositoryInterface) EXPECT() *MockUserRepositoryInterfaceMockRecorder {
	return m.recorder
}

// CreateUser mocks base method
func (m *MockUserRepositoryInterface) CreateUser(arg0 *entities.User) (*entities.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", arg0)
	ret0, _ := ret[0].(*entities.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser
func (mr *MockUserRepositoryInterfaceMockRecorder) CreateUser(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockUserRepositoryInterface)(nil).CreateUser), arg0)
}

// DeleteUserById mocks base method
func (m *MockUserRepositoryInterface) DeleteUserById(arg0 uint) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUserById", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteUserById indicates an expected call of DeleteUserById
func (mr *MockUserRepositoryInterfaceMockRecorder) DeleteUserById(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUserById", reflect.TypeOf((*MockUserRepositoryInterface)(nil).DeleteUserById), arg0)
}

// GetUserById mocks base method
func (m *MockUserRepositoryInterface) GetUserById(arg0 uint) (*entities.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserById", arg0)
	ret0, _ := ret[0].(*entities.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserById indicates an expected call of GetUserById
func (mr *MockUserRepositoryInterfaceMockRecorder) GetUserById(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserById", reflect.TypeOf((*MockUserRepositoryInterface)(nil).GetUserById), arg0)
}

// UpdateUserById mocks base method
func (m *MockUserRepositoryInterface) UpdateUserById(arg0 uint, arg1 *entities.User) (*entities.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUserById", arg0, arg1)
	ret0, _ := ret[0].(*entities.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateUserById indicates an expected call of UpdateUserById
func (mr *MockUserRepositoryInterfaceMockRecorder) UpdateUserById(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUserById", reflect.TypeOf((*MockUserRepositoryInterface)(nil).UpdateUserById), arg0, arg1)
}