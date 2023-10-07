// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/usecase/interface/user.go

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	models "github.com/aarathyaadhiv/ecommerce-fashionsture-cleanarch.git/pkg/utils/models"
	gomock "github.com/golang/mock/gomock"
)

// MockUserUseCase is a mock of UserUseCase interface.
type MockUserUseCase struct {
	ctrl     *gomock.Controller
	recorder *MockUserUseCaseMockRecorder
}

// MockUserUseCaseMockRecorder is the mock recorder for MockUserUseCase.
type MockUserUseCaseMockRecorder struct {
	mock *MockUserUseCase
}

// NewMockUserUseCase creates a new mock instance.
func NewMockUserUseCase(ctrl *gomock.Controller) *MockUserUseCase {
	mock := &MockUserUseCase{ctrl: ctrl}
	mock.recorder = &MockUserUseCaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserUseCase) EXPECT() *MockUserUseCaseMockRecorder {
	return m.recorder
}

// AddAddress mocks base method.
func (m *MockUserUseCase) AddAddress(address models.ShowAddress, userId uint) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddAddress", address, userId)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddAddress indicates an expected call of AddAddress.
func (mr *MockUserUseCaseMockRecorder) AddAddress(address, userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddAddress", reflect.TypeOf((*MockUserUseCase)(nil).AddAddress), address, userId)
}

// Checkout mocks base method.
func (m *MockUserUseCase) Checkout(id uint) (models.Checkout, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Checkout", id)
	ret0, _ := ret[0].(models.Checkout)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Checkout indicates an expected call of Checkout.
func (mr *MockUserUseCaseMockRecorder) Checkout(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Checkout", reflect.TypeOf((*MockUserUseCase)(nil).Checkout), id)
}

// ForgotPassword mocks base method.
func (m *MockUserUseCase) ForgotPassword(forgot models.Forgot) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ForgotPassword", forgot)
	ret0, _ := ret[0].(error)
	return ret0
}

// ForgotPassword indicates an expected call of ForgotPassword.
func (mr *MockUserUseCaseMockRecorder) ForgotPassword(forgot interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ForgotPassword", reflect.TypeOf((*MockUserUseCase)(nil).ForgotPassword), forgot)
}

// ResetPassword mocks base method.
func (m *MockUserUseCase) ResetPassword(id uint, password string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ResetPassword", id, password)
	ret0, _ := ret[0].(error)
	return ret0
}

// ResetPassword indicates an expected call of ResetPassword.
func (mr *MockUserUseCaseMockRecorder) ResetPassword(id, password interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResetPassword", reflect.TypeOf((*MockUserUseCase)(nil).ResetPassword), id, password)
}

// ShowAddress mocks base method.
func (m *MockUserUseCase) ShowAddress(id uint, page, count string) ([]models.ShowAddress, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ShowAddress", id, page, count)
	ret0, _ := ret[0].([]models.ShowAddress)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ShowAddress indicates an expected call of ShowAddress.
func (mr *MockUserUseCaseMockRecorder) ShowAddress(id, page, count interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ShowAddress", reflect.TypeOf((*MockUserUseCase)(nil).ShowAddress), id, page, count)
}

// ShowDetails mocks base method.
func (m *MockUserUseCase) ShowDetails(id uint) (models.UserDetails, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ShowDetails", id)
	ret0, _ := ret[0].(models.UserDetails)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ShowDetails indicates an expected call of ShowDetails.
func (mr *MockUserUseCaseMockRecorder) ShowDetails(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ShowDetails", reflect.TypeOf((*MockUserUseCase)(nil).ShowDetails), id)
}

// UpdateAddress mocks base method.
func (m *MockUserUseCase) UpdateAddress(address models.ShowAddress, addressId string, userId uint) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAddress", address, addressId, userId)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateAddress indicates an expected call of UpdateAddress.
func (mr *MockUserUseCaseMockRecorder) UpdateAddress(address, addressId, userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAddress", reflect.TypeOf((*MockUserUseCase)(nil).UpdateAddress), address, addressId, userId)
}

// UpdateUserDetails mocks base method.
func (m *MockUserUseCase) UpdateUserDetails(userId uint, userdetails models.UserUpdate) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUserDetails", userId, userdetails)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateUserDetails indicates an expected call of UpdateUserDetails.
func (mr *MockUserUseCaseMockRecorder) UpdateUserDetails(userId, userdetails interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUserDetails", reflect.TypeOf((*MockUserUseCase)(nil).UpdateUserDetails), userId, userdetails)
}

// UserLogin mocks base method.
func (m *MockUserUseCase) UserLogin(user models.UserLogin) (models.TokenResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UserLogin", user)
	ret0, _ := ret[0].(models.TokenResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UserLogin indicates an expected call of UserLogin.
func (mr *MockUserUseCaseMockRecorder) UserLogin(user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UserLogin", reflect.TypeOf((*MockUserUseCase)(nil).UserLogin), user)
}

// UserSignUp mocks base method.
func (m *MockUserUseCase) UserSignUp(user models.UserSignUp) (models.TokenResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UserSignUp", user)
	ret0, _ := ret[0].(models.TokenResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UserSignUp indicates an expected call of UserSignUp.
func (mr *MockUserUseCaseMockRecorder) UserSignUp(user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UserSignUp", reflect.TypeOf((*MockUserUseCase)(nil).UserSignUp), user)
}

// VerifyResetOtp mocks base method.
func (m *MockUserUseCase) VerifyResetOtp(data models.ForgotVerify) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VerifyResetOtp", data)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// VerifyResetOtp indicates an expected call of VerifyResetOtp.
func (mr *MockUserUseCaseMockRecorder) VerifyResetOtp(data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VerifyResetOtp", reflect.TypeOf((*MockUserUseCase)(nil).VerifyResetOtp), data)
}