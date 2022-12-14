// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/eneskzlcn/catbyte-test-task/internal/reporting (interfaces: RedisClient)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockRedisClient is a mock of RedisClient interface.
type MockRedisClient struct {
	ctrl     *gomock.Controller
	recorder *MockRedisClientMockRecorder
}

// MockRedisClientMockRecorder is the mock recorder for MockRedisClient.
type MockRedisClientMockRecorder struct {
	mock *MockRedisClient
}

// NewMockRedisClient creates a new mock instance.
func NewMockRedisClient(ctrl *gomock.Controller) *MockRedisClient {
	mock := &MockRedisClient{ctrl: ctrl}
	mock.recorder = &MockRedisClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRedisClient) EXPECT() *MockRedisClientMockRecorder {
	return m.recorder
}

// GetArray mocks base method.
func (m *MockRedisClient) GetArray(arg0 string, arg1 *[]string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetArray", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// GetArray indicates an expected call of GetArray.
func (mr *MockRedisClientMockRecorder) GetArray(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetArray", reflect.TypeOf((*MockRedisClient)(nil).GetArray), arg0, arg1)
}
