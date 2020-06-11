// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/kubernetes-sigs/minibroker/pkg/helm (interfaces: ChartLoader,ChartHelmClientProvider)

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	helm "github.com/kubernetes-sigs/minibroker/pkg/helm"
	chart "helm.sh/helm/v3/pkg/chart"
	reflect "reflect"
)

// MockChartLoader is a mock of ChartLoader interface
type MockChartLoader struct {
	ctrl     *gomock.Controller
	recorder *MockChartLoaderMockRecorder
}

// MockChartLoaderMockRecorder is the mock recorder for MockChartLoader
type MockChartLoaderMockRecorder struct {
	mock *MockChartLoader
}

// NewMockChartLoader creates a new mock instance
func NewMockChartLoader(ctrl *gomock.Controller) *MockChartLoader {
	mock := &MockChartLoader{ctrl: ctrl}
	mock.recorder = &MockChartLoaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockChartLoader) EXPECT() *MockChartLoaderMockRecorder {
	return m.recorder
}

// Load mocks base method
func (m *MockChartLoader) Load(arg0 string) (*chart.Chart, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Load", arg0)
	ret0, _ := ret[0].(*chart.Chart)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Load indicates an expected call of Load
func (mr *MockChartLoaderMockRecorder) Load(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Load", reflect.TypeOf((*MockChartLoader)(nil).Load), arg0)
}

// MockChartHelmClientProvider is a mock of ChartHelmClientProvider interface
type MockChartHelmClientProvider struct {
	ctrl     *gomock.Controller
	recorder *MockChartHelmClientProviderMockRecorder
}

// MockChartHelmClientProviderMockRecorder is the mock recorder for MockChartHelmClientProvider
type MockChartHelmClientProviderMockRecorder struct {
	mock *MockChartHelmClientProvider
}

// NewMockChartHelmClientProvider creates a new mock instance
func NewMockChartHelmClientProvider(ctrl *gomock.Controller) *MockChartHelmClientProvider {
	mock := &MockChartHelmClientProvider{ctrl: ctrl}
	mock.recorder = &MockChartHelmClientProviderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockChartHelmClientProvider) EXPECT() *MockChartHelmClientProviderMockRecorder {
	return m.recorder
}

// ProvideInstaller mocks base method
func (m *MockChartHelmClientProvider) ProvideInstaller(arg0, arg1 string) (helm.HelmClientInstallRunner, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProvideInstaller", arg0, arg1)
	ret0, _ := ret[0].(helm.HelmClientInstallRunner)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ProvideInstaller indicates an expected call of ProvideInstaller
func (mr *MockChartHelmClientProviderMockRecorder) ProvideInstaller(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProvideInstaller", reflect.TypeOf((*MockChartHelmClientProvider)(nil).ProvideInstaller), arg0, arg1)
}

// ProvideUninstaller mocks base method
func (m *MockChartHelmClientProvider) ProvideUninstaller(arg0 string) (helm.HelmClientUninstallRunner, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProvideUninstaller", arg0)
	ret0, _ := ret[0].(helm.HelmClientUninstallRunner)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ProvideUninstaller indicates an expected call of ProvideUninstaller
func (mr *MockChartHelmClientProviderMockRecorder) ProvideUninstaller(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProvideUninstaller", reflect.TypeOf((*MockChartHelmClientProvider)(nil).ProvideUninstaller), arg0)
}
