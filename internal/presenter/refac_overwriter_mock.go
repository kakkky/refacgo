// Code generated by MockGen. DO NOT EDIT.
// Source: ./refac_overwriter_interface.go
//
// Generated by this command:
//
//	mockgen -package=presenter -source=./refac_overwriter_interface.go -destination=./refac_overwriter_mock.go
//

// Package presenter is a generated GoMock package.
package presenter

import (
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockRefacOverWriter is a mock of RefacOverWriter interface.
type MockRefacOverWriter struct {
	ctrl     *gomock.Controller
	recorder *MockRefacOverWriterMockRecorder
	isgomock struct{}
}

// MockRefacOverWriterMockRecorder is the mock recorder for MockRefacOverWriter.
type MockRefacOverWriterMockRecorder struct {
	mock *MockRefacOverWriter
}

// NewMockRefacOverWriter creates a new mock instance.
func NewMockRefacOverWriter(ctrl *gomock.Controller) *MockRefacOverWriter {
	mock := &MockRefacOverWriter{ctrl: ctrl}
	mock.recorder = &MockRefacOverWriterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRefacOverWriter) EXPECT() *MockRefacOverWriterMockRecorder {
	return m.recorder
}

// OverWrite mocks base method.
func (m *MockRefacOverWriter) OverWrite(filename, src string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OverWrite", filename, src)
	ret0, _ := ret[0].(error)
	return ret0
}

// OverWrite indicates an expected call of OverWrite.
func (mr *MockRefacOverWriterMockRecorder) OverWrite(filename, src any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OverWrite", reflect.TypeOf((*MockRefacOverWriter)(nil).OverWrite), filename, src)
}

// OverWriteWithHeaderComment mocks base method.
func (m *MockRefacOverWriter) OverWriteWithHeaderComment(filename, src string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OverWriteWithHeaderComment", filename, src)
	ret0, _ := ret[0].(error)
	return ret0
}

// OverWriteWithHeaderComment indicates an expected call of OverWriteWithHeaderComment.
func (mr *MockRefacOverWriterMockRecorder) OverWriteWithHeaderComment(filename, src any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OverWriteWithHeaderComment", reflect.TypeOf((*MockRefacOverWriter)(nil).OverWriteWithHeaderComment), filename, src)
}
