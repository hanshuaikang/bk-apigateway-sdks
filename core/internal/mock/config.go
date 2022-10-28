// Code generated by MockGen. DO NOT EDIT.
// Source: config.go

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	define "github.com/TencentBlueKing/bk-apigateway-sdks/core/define"
	logging "github.com/TencentBlueKing/gopkg/logging"
	gomock "github.com/golang/mock/gomock"
)

// MockClientConfig is a mock of ClientConfig interface.
type MockClientConfig struct {
	ctrl     *gomock.Controller
	recorder *MockClientConfigMockRecorder
}

// MockClientConfigMockRecorder is the mock recorder for MockClientConfig.
type MockClientConfigMockRecorder struct {
	mock *MockClientConfig
}

// NewMockClientConfig creates a new mock instance.
func NewMockClientConfig(ctrl *gomock.Controller) *MockClientConfig {
	mock := &MockClientConfig{ctrl: ctrl}
	mock.recorder = &MockClientConfigMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClientConfig) EXPECT() *MockClientConfigMockRecorder {
	return m.recorder
}

// GetAuthorizationHeaders mocks base method.
func (m *MockClientConfig) GetAuthorizationHeaders() map[string]string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAuthorizationHeaders")
	ret0, _ := ret[0].(map[string]string)
	return ret0
}

// GetAuthorizationHeaders indicates an expected call of GetAuthorizationHeaders.
func (mr *MockClientConfigMockRecorder) GetAuthorizationHeaders() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAuthorizationHeaders", reflect.TypeOf((*MockClientConfig)(nil).GetAuthorizationHeaders))
}

// GetLogger mocks base method.
func (m *MockClientConfig) GetLogger() logging.Logger {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLogger")
	ret0, _ := ret[0].(logging.Logger)
	return ret0
}

// GetLogger indicates an expected call of GetLogger.
func (mr *MockClientConfigMockRecorder) GetLogger() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLogger", reflect.TypeOf((*MockClientConfig)(nil).GetLogger))
}

// GetUrl mocks base method.
func (m *MockClientConfig) GetUrl() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUrl")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetUrl indicates an expected call of GetUrl.
func (mr *MockClientConfigMockRecorder) GetUrl() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUrl", reflect.TypeOf((*MockClientConfig)(nil).GetUrl))
}

// MockClientConfigProvider is a mock of ClientConfigProvider interface.
type MockClientConfigProvider struct {
	ctrl     *gomock.Controller
	recorder *MockClientConfigProviderMockRecorder
}

// MockClientConfigProviderMockRecorder is the mock recorder for MockClientConfigProvider.
type MockClientConfigProviderMockRecorder struct {
	mock *MockClientConfigProvider
}

// NewMockClientConfigProvider creates a new mock instance.
func NewMockClientConfigProvider(ctrl *gomock.Controller) *MockClientConfigProvider {
	mock := &MockClientConfigProvider{ctrl: ctrl}
	mock.recorder = &MockClientConfigProviderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClientConfigProvider) EXPECT() *MockClientConfigProviderMockRecorder {
	return m.recorder
}

// ProvideConfig mocks base method.
func (m *MockClientConfigProvider) ProvideConfig(apiName string) define.ClientConfig {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProvideConfig", apiName)
	ret0, _ := ret[0].(define.ClientConfig)
	return ret0
}

// ProvideConfig indicates an expected call of ProvideConfig.
func (mr *MockClientConfigProviderMockRecorder) ProvideConfig(apiName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProvideConfig", reflect.TypeOf((*MockClientConfigProvider)(nil).ProvideConfig), apiName)
}

// MockOperationConfig is a mock of OperationConfig interface.
type MockOperationConfig struct {
	ctrl     *gomock.Controller
	recorder *MockOperationConfigMockRecorder
}

// MockOperationConfigMockRecorder is the mock recorder for MockOperationConfig.
type MockOperationConfigMockRecorder struct {
	mock *MockOperationConfig
}

// NewMockOperationConfig creates a new mock instance.
func NewMockOperationConfig(ctrl *gomock.Controller) *MockOperationConfig {
	mock := &MockOperationConfig{ctrl: ctrl}
	mock.recorder = &MockOperationConfigMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOperationConfig) EXPECT() *MockOperationConfigMockRecorder {
	return m.recorder
}

// GetMethod mocks base method.
func (m *MockOperationConfig) GetMethod() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMethod")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetMethod indicates an expected call of GetMethod.
func (mr *MockOperationConfigMockRecorder) GetMethod() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMethod", reflect.TypeOf((*MockOperationConfig)(nil).GetMethod))
}

// GetName mocks base method.
func (m *MockOperationConfig) GetName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetName")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetName indicates an expected call of GetName.
func (mr *MockOperationConfigMockRecorder) GetName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetName", reflect.TypeOf((*MockOperationConfig)(nil).GetName))
}

// GetPath mocks base method.
func (m *MockOperationConfig) GetPath() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPath")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetPath indicates an expected call of GetPath.
func (mr *MockOperationConfigMockRecorder) GetPath() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPath", reflect.TypeOf((*MockOperationConfig)(nil).GetPath))
}

// MockOperationConfigProvider is a mock of OperationConfigProvider interface.
type MockOperationConfigProvider struct {
	ctrl     *gomock.Controller
	recorder *MockOperationConfigProviderMockRecorder
}

// MockOperationConfigProviderMockRecorder is the mock recorder for MockOperationConfigProvider.
type MockOperationConfigProviderMockRecorder struct {
	mock *MockOperationConfigProvider
}

// NewMockOperationConfigProvider creates a new mock instance.
func NewMockOperationConfigProvider(ctrl *gomock.Controller) *MockOperationConfigProvider {
	mock := &MockOperationConfigProvider{ctrl: ctrl}
	mock.recorder = &MockOperationConfigProviderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOperationConfigProvider) EXPECT() *MockOperationConfigProviderMockRecorder {
	return m.recorder
}

// ProvideConfig mocks base method.
func (m *MockOperationConfigProvider) ProvideConfig() define.OperationConfig {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProvideConfig")
	ret0, _ := ret[0].(define.OperationConfig)
	return ret0
}

// ProvideConfig indicates an expected call of ProvideConfig.
func (mr *MockOperationConfigProviderMockRecorder) ProvideConfig() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProvideConfig", reflect.TypeOf((*MockOperationConfigProvider)(nil).ProvideConfig))
}
