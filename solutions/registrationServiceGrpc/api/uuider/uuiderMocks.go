// Code generated by MockGen. DO NOT EDIT.
// Source: uuider.go

// Package uuider is a generated GoMock package.
package uuider

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockUuidGenerator is a mock of UuidGenerator interface.
type MockUuidGenerator struct {
	ctrl     *gomock.Controller
	recorder *MockUuidGeneratorMockRecorder
}

// MockUuidGeneratorMockRecorder is the mock recorder for MockUuidGenerator.
type MockUuidGeneratorMockRecorder struct {
	mock *MockUuidGenerator
}

// NewMockUuidGenerator creates a new mock instance.
func NewMockUuidGenerator(ctrl *gomock.Controller) *MockUuidGenerator {
	mock := &MockUuidGenerator{ctrl: ctrl}
	mock.recorder = &MockUuidGeneratorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUuidGenerator) EXPECT() *MockUuidGeneratorMockRecorder {
	return m.recorder
}

// GenerateUuid mocks base method.
func (m *MockUuidGenerator) GenerateUuid() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateUuid")
	ret0, _ := ret[0].(string)
	return ret0
}

// GenerateUuid indicates an expected call of GenerateUuid.
func (mr *MockUuidGeneratorMockRecorder) GenerateUuid() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateUuid", reflect.TypeOf((*MockUuidGenerator)(nil).GenerateUuid))
}