// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/kubernetes-sigs/minibroker/pkg/helm (interfaces: ConfigProvider,ConfigInitializer,ConfigInitializerProvider)

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	helm "github.com/kubernetes-sigs/minibroker/pkg/helm"
	action "helm.sh/helm/v3/pkg/action"
	genericclioptions "k8s.io/cli-runtime/pkg/genericclioptions"
	reflect "reflect"
)

// MockConfigProvider is a mock of ConfigProvider interface
type MockConfigProvider struct {
	ctrl     *gomock.Controller
	recorder *MockConfigProviderMockRecorder
}

// MockConfigProviderMockRecorder is the mock recorder for MockConfigProvider
type MockConfigProviderMockRecorder struct {
	mock *MockConfigProvider
}

// NewMockConfigProvider creates a new mock instance
func NewMockConfigProvider(ctrl *gomock.Controller) *MockConfigProvider {
	mock := &MockConfigProvider{ctrl: ctrl}
	mock.recorder = &MockConfigProviderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockConfigProvider) EXPECT() *MockConfigProviderMockRecorder {
	return m.recorder
}

// Provide mocks base method
func (m *MockConfigProvider) Provide(arg0 string) (interface{}, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Provide", arg0)
	ret0, _ := ret[0].(interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Provide indicates an expected call of Provide
func (mr *MockConfigProviderMockRecorder) Provide(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Provide", reflect.TypeOf((*MockConfigProvider)(nil).Provide), arg0)
}

// MockConfigInitializer is a mock of ConfigInitializer interface
type MockConfigInitializer struct {
	ctrl     *gomock.Controller
	recorder *MockConfigInitializerMockRecorder
}

// MockConfigInitializerMockRecorder is the mock recorder for MockConfigInitializer
type MockConfigInitializerMockRecorder struct {
	mock *MockConfigInitializer
}

// NewMockConfigInitializer creates a new mock instance
func NewMockConfigInitializer(ctrl *gomock.Controller) *MockConfigInitializer {
	mock := &MockConfigInitializer{ctrl: ctrl}
	mock.recorder = &MockConfigInitializerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockConfigInitializer) EXPECT() *MockConfigInitializerMockRecorder {
	return m.recorder
}

// Init mocks base method
func (m *MockConfigInitializer) Init(arg0 genericclioptions.RESTClientGetter, arg1, arg2 string, arg3 action.DebugLog) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Init", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// Init indicates an expected call of Init
func (mr *MockConfigInitializerMockRecorder) Init(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Init", reflect.TypeOf((*MockConfigInitializer)(nil).Init), arg0, arg1, arg2, arg3)
}

// MockConfigInitializerProvider is a mock of ConfigInitializerProvider interface
type MockConfigInitializerProvider struct {
	ctrl     *gomock.Controller
	recorder *MockConfigInitializerProviderMockRecorder
}

// MockConfigInitializerProviderMockRecorder is the mock recorder for MockConfigInitializerProvider
type MockConfigInitializerProviderMockRecorder struct {
	mock *MockConfigInitializerProvider
}

// NewMockConfigInitializerProvider creates a new mock instance
func NewMockConfigInitializerProvider(ctrl *gomock.Controller) *MockConfigInitializerProvider {
	mock := &MockConfigInitializerProvider{ctrl: ctrl}
	mock.recorder = &MockConfigInitializerProviderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockConfigInitializerProvider) EXPECT() *MockConfigInitializerProviderMockRecorder {
	return m.recorder
}

// Provide mocks base method
func (m *MockConfigInitializerProvider) Provide() helm.ConfigInitializer {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Provide")
	ret0, _ := ret[0].(helm.ConfigInitializer)
	return ret0
}

// Provide indicates an expected call of Provide
func (mr *MockConfigInitializerProviderMockRecorder) Provide() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Provide", reflect.TypeOf((*MockConfigInitializerProvider)(nil).Provide))
}