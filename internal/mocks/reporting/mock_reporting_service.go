// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/eneskzlcn/catbyte-test-task/internal/reporting (interfaces: ReportingService)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	reporting "github.com/eneskzlcn/catbyte-test-task/internal/reporting"
	gomock "github.com/golang/mock/gomock"
)

// MockReportingService is a mock of ReportingService interface.
type MockReportingService struct {
	ctrl     *gomock.Controller
	recorder *MockReportingServiceMockRecorder
}

// MockReportingServiceMockRecorder is the mock recorder for MockReportingService.
type MockReportingServiceMockRecorder struct {
	mock *MockReportingService
}

// NewMockReportingService creates a new mock instance.
func NewMockReportingService(ctrl *gomock.Controller) *MockReportingService {
	mock := &MockReportingService{ctrl: ctrl}
	mock.recorder = &MockReportingServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockReportingService) EXPECT() *MockReportingServiceMockRecorder {
	return m.recorder
}

// Report mocks base method.
func (m *MockReportingService) Report(arg0, arg1 string) ([]reporting.Report, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Report", arg0, arg1)
	ret0, _ := ret[0].([]reporting.Report)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Report indicates an expected call of Report.
func (mr *MockReportingServiceMockRecorder) Report(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Report", reflect.TypeOf((*MockReportingService)(nil).Report), arg0, arg1)
}