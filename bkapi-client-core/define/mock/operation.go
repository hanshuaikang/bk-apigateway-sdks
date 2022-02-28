// Code generated by MockGen. DO NOT EDIT.
// Source: operation.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	io "io"
	http "net/http"
	reflect "reflect"

	define "github.com/TencentBlueKing/bk-apigateway-sdks/bkapi-client-core/define"
	gomock "github.com/golang/mock/gomock"
)

// MockOperation is a mock of Operation interface.
type MockOperation struct {
	ctrl     *gomock.Controller
	recorder *MockOperationMockRecorder
}

// MockOperationMockRecorder is the mock recorder for MockOperation.
type MockOperationMockRecorder struct {
	mock *MockOperation
}

// NewMockOperation creates a new mock instance.
func NewMockOperation(ctrl *gomock.Controller) *MockOperation {
	mock := &MockOperation{ctrl: ctrl}
	mock.recorder = &MockOperationMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOperation) EXPECT() *MockOperationMockRecorder {
	return m.recorder
}

// Apply mocks base method.
func (m *MockOperation) Apply(opts ...define.OperationOption) define.Operation {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Apply", varargs...)
	ret0, _ := ret[0].(define.Operation)
	return ret0
}

// Apply indicates an expected call of Apply.
func (mr *MockOperationMockRecorder) Apply(opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Apply", reflect.TypeOf((*MockOperation)(nil).Apply), opts...)
}

// Request mocks base method.
func (m *MockOperation) Request() (*http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Request")
	ret0, _ := ret[0].(*http.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Request indicates an expected call of Request.
func (mr *MockOperationMockRecorder) Request() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Request", reflect.TypeOf((*MockOperation)(nil).Request))
}

// SetBody mocks base method.
func (m *MockOperation) SetBody(data interface{}) define.Operation {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetBody", data)
	ret0, _ := ret[0].(define.Operation)
	return ret0
}

// SetBody indicates an expected call of SetBody.
func (mr *MockOperationMockRecorder) SetBody(data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetBody", reflect.TypeOf((*MockOperation)(nil).SetBody), data)
}

// SetBodyProvider mocks base method.
func (m *MockOperation) SetBodyProvider(bodyProvider define.BodyProvider) define.Operation {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetBodyProvider", bodyProvider)
	ret0, _ := ret[0].(define.Operation)
	return ret0
}

// SetBodyProvider indicates an expected call of SetBodyProvider.
func (mr *MockOperationMockRecorder) SetBodyProvider(bodyProvider interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetBodyProvider", reflect.TypeOf((*MockOperation)(nil).SetBodyProvider), bodyProvider)
}

// SetBodyReader mocks base method.
func (m *MockOperation) SetBodyReader(body io.Reader) define.Operation {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetBodyReader", body)
	ret0, _ := ret[0].(define.Operation)
	return ret0
}

// SetBodyReader indicates an expected call of SetBodyReader.
func (mr *MockOperationMockRecorder) SetBodyReader(body interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetBodyReader", reflect.TypeOf((*MockOperation)(nil).SetBodyReader), body)
}

// SetContentLength mocks base method.
func (m *MockOperation) SetContentLength(contentLength int64) define.Operation {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetContentLength", contentLength)
	ret0, _ := ret[0].(define.Operation)
	return ret0
}

// SetContentLength indicates an expected call of SetContentLength.
func (mr *MockOperationMockRecorder) SetContentLength(contentLength interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetContentLength", reflect.TypeOf((*MockOperation)(nil).SetContentLength), contentLength)
}

// SetContentType mocks base method.
func (m *MockOperation) SetContentType(contentType string) define.Operation {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetContentType", contentType)
	ret0, _ := ret[0].(define.Operation)
	return ret0
}

// SetContentType indicates an expected call of SetContentType.
func (mr *MockOperationMockRecorder) SetContentType(contentType interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetContentType", reflect.TypeOf((*MockOperation)(nil).SetContentType), contentType)
}

// SetContext mocks base method.
func (m *MockOperation) SetContext(ctx context.Context) define.Operation {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetContext", ctx)
	ret0, _ := ret[0].(define.Operation)
	return ret0
}

// SetContext indicates an expected call of SetContext.
func (mr *MockOperationMockRecorder) SetContext(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetContext", reflect.TypeOf((*MockOperation)(nil).SetContext), ctx)
}

// SetHeaders mocks base method.
func (m *MockOperation) SetHeaders(headers map[string]string) define.Operation {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetHeaders", headers)
	ret0, _ := ret[0].(define.Operation)
	return ret0
}

// SetHeaders indicates an expected call of SetHeaders.
func (mr *MockOperationMockRecorder) SetHeaders(headers interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetHeaders", reflect.TypeOf((*MockOperation)(nil).SetHeaders), headers)
}

// SetPathParams mocks base method.
func (m *MockOperation) SetPathParams(params map[string]string) define.Operation {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetPathParams", params)
	ret0, _ := ret[0].(define.Operation)
	return ret0
}

// SetPathParams indicates an expected call of SetPathParams.
func (mr *MockOperationMockRecorder) SetPathParams(params interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetPathParams", reflect.TypeOf((*MockOperation)(nil).SetPathParams), params)
}

// SetQueryParams mocks base method.
func (m *MockOperation) SetQueryParams(params map[string]string) define.Operation {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetQueryParams", params)
	ret0, _ := ret[0].(define.Operation)
	return ret0
}

// SetQueryParams indicates an expected call of SetQueryParams.
func (mr *MockOperationMockRecorder) SetQueryParams(params interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetQueryParams", reflect.TypeOf((*MockOperation)(nil).SetQueryParams), params)
}

// SetResult mocks base method.
func (m *MockOperation) SetResult(result interface{}) define.Operation {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetResult", result)
	ret0, _ := ret[0].(define.Operation)
	return ret0
}

// SetResult indicates an expected call of SetResult.
func (mr *MockOperationMockRecorder) SetResult(result interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetResult", reflect.TypeOf((*MockOperation)(nil).SetResult), result)
}

// SetResultProvider mocks base method.
func (m *MockOperation) SetResultProvider(provider func(*http.Response, interface{}) error) define.Operation {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetResultProvider", provider)
	ret0, _ := ret[0].(define.Operation)
	return ret0
}

// SetResultProvider indicates an expected call of SetResultProvider.
func (mr *MockOperationMockRecorder) SetResultProvider(provider interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetResultProvider", reflect.TypeOf((*MockOperation)(nil).SetResultProvider), provider)
}

// MockOperationOption is a mock of OperationOption interface.
type MockOperationOption struct {
	ctrl     *gomock.Controller
	recorder *MockOperationOptionMockRecorder
}

// MockOperationOptionMockRecorder is the mock recorder for MockOperationOption.
type MockOperationOptionMockRecorder struct {
	mock *MockOperationOption
}

// NewMockOperationOption creates a new mock instance.
func NewMockOperationOption(ctrl *gomock.Controller) *MockOperationOption {
	mock := &MockOperationOption{ctrl: ctrl}
	mock.recorder = &MockOperationOptionMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOperationOption) EXPECT() *MockOperationOptionMockRecorder {
	return m.recorder
}

// ApplyToOperation mocks base method.
func (m *MockOperationOption) ApplyToOperation(operation define.Operation) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ApplyToOperation", operation)
	ret0, _ := ret[0].(error)
	return ret0
}

// ApplyToOperation indicates an expected call of ApplyToOperation.
func (mr *MockOperationOptionMockRecorder) ApplyToOperation(operation interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ApplyToOperation", reflect.TypeOf((*MockOperationOption)(nil).ApplyToOperation), operation)
}
