// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/eneskzlcn/catbyte-test-task/internal/message (interfaces: RabbitMQClient)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockRabbitMQClient is a mock of RabbitMQClient interface.
type MockRabbitMQClient struct {
	ctrl     *gomock.Controller
	recorder *MockRabbitMQClientMockRecorder
}

// MockRabbitMQClientMockRecorder is the mock recorder for MockRabbitMQClient.
type MockRabbitMQClientMockRecorder struct {
	mock *MockRabbitMQClient
}

// NewMockRabbitMQClient creates a new mock instance.
func NewMockRabbitMQClient(ctrl *gomock.Controller) *MockRabbitMQClient {
	mock := &MockRabbitMQClient{ctrl: ctrl}
	mock.recorder = &MockRabbitMQClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRabbitMQClient) EXPECT() *MockRabbitMQClientMockRecorder {
	return m.recorder
}

// PushMessage mocks base method.
func (m *MockRabbitMQClient) PushMessage(arg0 []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PushMessage", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// PushMessage indicates an expected call of PushMessage.
func (mr *MockRabbitMQClientMockRecorder) PushMessage(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PushMessage", reflect.TypeOf((*MockRabbitMQClient)(nil).PushMessage), arg0)
}
