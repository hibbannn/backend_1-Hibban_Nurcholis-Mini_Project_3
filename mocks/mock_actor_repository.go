// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/hibbannn/backend_1-Hibban_Nurcholis-Mini_Project_2/repositories (interfaces: AdminRepositoryInterface)

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	entities "github.com/hibbannn/backend_1-Hibban_Nurcholis-Mini_Project_2/entities"
	reflect "reflect"
)

// MockAdminRepositoryInterface is a mock of AdminRepositoryInterface interface
type MockAdminRepositoryInterface struct {
	ctrl     *gomock.Controller
	recorder *MockAdminRepositoryInterfaceMockRecorder
}

// MockAdminRepositoryInterfaceMockRecorder is the mock recorder for MockAdminRepositoryInterface
type MockAdminRepositoryInterfaceMockRecorder struct {
	mock *MockAdminRepositoryInterface
}

// NewMockAdminRepositoryInterface creates a new mock instance
func NewMockAdminRepositoryInterface(ctrl *gomock.Controller) *MockAdminRepositoryInterface {
	mock := &MockAdminRepositoryInterface{ctrl: ctrl}
	mock.recorder = &MockAdminRepositoryInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAdminRepositoryInterface) EXPECT() *MockAdminRepositoryInterfaceMockRecorder {
	return m.recorder
}

// CreateApproval mocks base method
func (m *MockAdminRepositoryInterface) CreateApproval(arg0 uint) (*entities.Approval, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateApproval", arg0)
	ret0, _ := ret[0].(*entities.Approval)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateApproval indicates an expected call of CreateApproval
func (mr *MockAdminRepositoryInterfaceMockRecorder) CreateApproval(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateApproval", reflect.TypeOf((*MockAdminRepositoryInterface)(nil).CreateApproval), arg0)
}

// CreateUser mocks base method
func (m *MockAdminRepositoryInterface) CreateUser(arg0 *entities.User) (*entities.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", arg0)
	ret0, _ := ret[0].(*entities.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser
func (mr *MockAdminRepositoryInterfaceMockRecorder) CreateUser(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockAdminRepositoryInterface)(nil).CreateUser), arg0)
}

// DeleteUserById mocks base method
func (m *MockAdminRepositoryInterface) DeleteUserById(arg0 uint, arg1 *entities.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUserById", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteUserById indicates an expected call of DeleteUserById
func (mr *MockAdminRepositoryInterfaceMockRecorder) DeleteUserById(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUserById", reflect.TypeOf((*MockAdminRepositoryInterface)(nil).DeleteUserById), arg0, arg1)
}

// FetchUsersFromAPI mocks base method
func (m *MockAdminRepositoryInterface) FetchUsersFromAPI() ([]*entities.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchUsersFromAPI")
	ret0, _ := ret[0].([]*entities.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchUsersFromAPI indicates an expected call of FetchUsersFromAPI
func (mr *MockAdminRepositoryInterfaceMockRecorder) FetchUsersFromAPI() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchUsersFromAPI", reflect.TypeOf((*MockAdminRepositoryInterface)(nil).FetchUsersFromAPI))
}

// GetAllUsers mocks base method
func (m *MockAdminRepositoryInterface) GetAllUsers(arg0, arg1, arg2 string, arg3, arg4 int) ([]*entities.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllUsers", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].([]*entities.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllUsers indicates an expected call of GetAllUsers
func (mr *MockAdminRepositoryInterfaceMockRecorder) GetAllUsers(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllUsers", reflect.TypeOf((*MockAdminRepositoryInterface)(nil).GetAllUsers), arg0, arg1, arg2, arg3, arg4)
}

// GetApprovalById mocks base method
func (m *MockAdminRepositoryInterface) GetApprovalById(arg0 uint) (*entities.Approval, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetApprovalById", arg0)
	ret0, _ := ret[0].(*entities.Approval)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetApprovalById indicates an expected call of GetApprovalById
func (mr *MockAdminRepositoryInterfaceMockRecorder) GetApprovalById(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetApprovalById", reflect.TypeOf((*MockAdminRepositoryInterface)(nil).GetApprovalById), arg0)
}

// GetUserByEmail mocks base method
func (m *MockAdminRepositoryInterface) GetUserByEmail(arg0 string) (*entities.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByEmail", arg0)
	ret0, _ := ret[0].(*entities.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByEmail indicates an expected call of GetUserByEmail
func (mr *MockAdminRepositoryInterfaceMockRecorder) GetUserByEmail(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByEmail", reflect.TypeOf((*MockAdminRepositoryInterface)(nil).GetUserByEmail), arg0)
}

// GetUserById mocks base method
func (m *MockAdminRepositoryInterface) GetUserById(arg0 uint) (*entities.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserById", arg0)
	ret0, _ := ret[0].(*entities.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserById indicates an expected call of GetUserById
func (mr *MockAdminRepositoryInterfaceMockRecorder) GetUserById(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserById", reflect.TypeOf((*MockAdminRepositoryInterface)(nil).GetUserById), arg0)
}

// LoginAdmin mocks base method
func (m *MockAdminRepositoryInterface) LoginAdmin(arg0 string) (*entities.Actor, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LoginAdmin", arg0)
	ret0, _ := ret[0].(*entities.Actor)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LoginAdmin indicates an expected call of LoginAdmin
func (mr *MockAdminRepositoryInterfaceMockRecorder) LoginAdmin(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoginAdmin", reflect.TypeOf((*MockAdminRepositoryInterface)(nil).LoginAdmin), arg0)
}

// RegisterAdmin mocks base method
func (m *MockAdminRepositoryInterface) RegisterAdmin(arg0 *entities.Actor) (*entities.Actor, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegisterAdmin", arg0)
	ret0, _ := ret[0].(*entities.Actor)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RegisterAdmin indicates an expected call of RegisterAdmin
func (mr *MockAdminRepositoryInterfaceMockRecorder) RegisterAdmin(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterAdmin", reflect.TypeOf((*MockAdminRepositoryInterface)(nil).RegisterAdmin), arg0)
}

// SaveUsersFromAPI mocks base method
func (m *MockAdminRepositoryInterface) SaveUsersFromAPI(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveUsersFromAPI", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveUsersFromAPI indicates an expected call of SaveUsersFromAPI
func (mr *MockAdminRepositoryInterfaceMockRecorder) SaveUsersFromAPI(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveUsersFromAPI", reflect.TypeOf((*MockAdminRepositoryInterface)(nil).SaveUsersFromAPI), arg0)
}
