// Code generated by MockGen. DO NOT EDIT.
// Source: ./businesslogic/blacklistreason.go

// Package mock_businesslogic is a generated GoMock package.
package mock_businesslogic

import (
	"github.com/DancesportSoftware/das/businesslogic"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockIPartnershipRequestBlacklistReasonRepository is a mock of IPartnershipRequestBlacklistReasonRepository interface
type MockIPartnershipRequestBlacklistReasonRepository struct {
	ctrl     *gomock.Controller
	recorder *MockIPartnershipRequestBlacklistReasonRepositoryMockRecorder
}

// MockIPartnershipRequestBlacklistReasonRepositoryMockRecorder is the mock recorder for MockIPartnershipRequestBlacklistReasonRepository
type MockIPartnershipRequestBlacklistReasonRepositoryMockRecorder struct {
	mock *MockIPartnershipRequestBlacklistReasonRepository
}

// NewMockIPartnershipRequestBlacklistReasonRepository creates a new mock instance
func NewMockIPartnershipRequestBlacklistReasonRepository(ctrl *gomock.Controller) *MockIPartnershipRequestBlacklistReasonRepository {
	mock := &MockIPartnershipRequestBlacklistReasonRepository{ctrl: ctrl}
	mock.recorder = &MockIPartnershipRequestBlacklistReasonRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIPartnershipRequestBlacklistReasonRepository) EXPECT() *MockIPartnershipRequestBlacklistReasonRepositoryMockRecorder {
	return m.recorder
}

// GetPartnershipRequestBlacklistReasons mocks base method
func (m *MockIPartnershipRequestBlacklistReasonRepository) GetPartnershipRequestBlacklistReasons() ([]businesslogic.PartnershipRequestBlacklistReason, error) {
	ret := m.ctrl.Call(m, "GetPartnershipRequestBlacklistReasons")
	ret0, _ := ret[0].([]businesslogic.PartnershipRequestBlacklistReason)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPartnershipRequestBlacklistReasons indicates an expected call of GetPartnershipRequestBlacklistReasons
func (mr *MockIPartnershipRequestBlacklistReasonRepositoryMockRecorder) GetPartnershipRequestBlacklistReasons() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPartnershipRequestBlacklistReasons", reflect.TypeOf((*MockIPartnershipRequestBlacklistReasonRepository)(nil).GetPartnershipRequestBlacklistReasons))
}
