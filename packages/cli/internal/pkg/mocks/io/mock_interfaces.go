// Code generated by MockGen. DO NOT EDIT.
// Source: ./internal/pkg/mocks/io/interfaces.go

// Package iomocks is a generated GoMock package.
package iomocks

import (
	fs "io/fs"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockOS is a mock of OS interface.
type MockOS struct {
	ctrl     *gomock.Controller
	recorder *MockOSMockRecorder
}

// MockOSMockRecorder is the mock recorder for MockOS.
type MockOSMockRecorder struct {
	mock *MockOS
}

// NewMockOS creates a new mock instance.
func NewMockOS(ctrl *gomock.Controller) *MockOS {
	mock := &MockOS{ctrl: ctrl}
	mock.recorder = &MockOSMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOS) EXPECT() *MockOSMockRecorder {
	return m.recorder
}

// Chdir mocks base method.
func (m *MockOS) Chdir(dir string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Chdir", dir)
	ret0, _ := ret[0].(error)
	return ret0
}

// Chdir indicates an expected call of Chdir.
func (mr *MockOSMockRecorder) Chdir(dir interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Chdir", reflect.TypeOf((*MockOS)(nil).Chdir), dir)
}

// MkdirTemp mocks base method.
func (m *MockOS) MkdirTemp(dir, pattern string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MkdirTemp", dir, pattern)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MkdirTemp indicates an expected call of MkdirTemp.
func (mr *MockOSMockRecorder) MkdirTemp(dir, pattern interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MkdirTemp", reflect.TypeOf((*MockOS)(nil).MkdirTemp), dir, pattern)
}

// Remove mocks base method.
func (m *MockOS) Remove(name string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Remove", name)
	ret0, _ := ret[0].(error)
	return ret0
}

// Remove indicates an expected call of Remove.
func (mr *MockOSMockRecorder) Remove(name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Remove", reflect.TypeOf((*MockOS)(nil).Remove), name)
}

// RemoveAll mocks base method.
func (m *MockOS) RemoveAll(path string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveAll", path)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveAll indicates an expected call of RemoveAll.
func (mr *MockOSMockRecorder) RemoveAll(path interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveAll", reflect.TypeOf((*MockOS)(nil).RemoveAll), path)
}

// UserHomeDir mocks base method.
func (m *MockOS) UserHomeDir() (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UserHomeDir")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UserHomeDir indicates an expected call of UserHomeDir.
func (mr *MockOSMockRecorder) UserHomeDir() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UserHomeDir", reflect.TypeOf((*MockOS)(nil).UserHomeDir))
}

// MockZip is a mock of Zip interface.
type MockZip struct {
	ctrl     *gomock.Controller
	recorder *MockZipMockRecorder
}

// MockZipMockRecorder is the mock recorder for MockZip.
type MockZipMockRecorder struct {
	mock *MockZip
}

// NewMockZip creates a new mock instance.
func NewMockZip(ctrl *gomock.Controller) *MockZip {
	mock := &MockZip{ctrl: ctrl}
	mock.recorder = &MockZipMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockZip) EXPECT() *MockZipMockRecorder {
	return m.recorder
}

// CompressToTmp mocks base method.
func (m *MockZip) CompressToTmp(srcPath string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CompressToTmp", srcPath)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CompressToTmp indicates an expected call of CompressToTmp.
func (mr *MockZipMockRecorder) CompressToTmp(srcPath interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CompressToTmp", reflect.TypeOf((*MockZip)(nil).CompressToTmp), srcPath)
}

// MockTmp is a mock of Tmp interface.
type MockTmp struct {
	ctrl     *gomock.Controller
	recorder *MockTmpMockRecorder
}

// MockTmpMockRecorder is the mock recorder for MockTmp.
type MockTmpMockRecorder struct {
	mock *MockTmp
}

// NewMockTmp creates a new mock instance.
func NewMockTmp(ctrl *gomock.Controller) *MockTmp {
	mock := &MockTmp{ctrl: ctrl}
	mock.recorder = &MockTmpMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTmp) EXPECT() *MockTmpMockRecorder {
	return m.recorder
}

// Write mocks base method.
func (m *MockTmp) Write(namePattern, content string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Write", namePattern, content)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Write indicates an expected call of Write.
func (mr *MockTmpMockRecorder) Write(namePattern, content interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Write", reflect.TypeOf((*MockTmp)(nil).Write), namePattern, content)
}

// MockFileReader is a mock of FileReader interface.
type MockFileReader struct {
	ctrl     *gomock.Controller
	recorder *MockFileReaderMockRecorder
}

// MockFileReaderMockRecorder is the mock recorder for MockFileReader.
type MockFileReaderMockRecorder struct {
	mock *MockFileReader
}

// NewMockFileReader creates a new mock instance.
func NewMockFileReader(ctrl *gomock.Controller) *MockFileReader {
	mock := &MockFileReader{ctrl: ctrl}
	mock.recorder = &MockFileReaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFileReader) EXPECT() *MockFileReaderMockRecorder {
	return m.recorder
}

// ReadFile mocks base method.
func (m *MockFileReader) ReadFile(arg0 string) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadFile", arg0)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadFile indicates an expected call of ReadFile.
func (mr *MockFileReaderMockRecorder) ReadFile(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadFile", reflect.TypeOf((*MockFileReader)(nil).ReadFile), arg0)
}

// MockFileWriter is a mock of FileWriter interface.
type MockFileWriter struct {
	ctrl     *gomock.Controller
	recorder *MockFileWriterMockRecorder
}

// MockFileWriterMockRecorder is the mock recorder for MockFileWriter.
type MockFileWriterMockRecorder struct {
	mock *MockFileWriter
}

// NewMockFileWriter creates a new mock instance.
func NewMockFileWriter(ctrl *gomock.Controller) *MockFileWriter {
	mock := &MockFileWriter{ctrl: ctrl}
	mock.recorder = &MockFileWriterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFileWriter) EXPECT() *MockFileWriterMockRecorder {
	return m.recorder
}

// WriteFile mocks base method.
func (m *MockFileWriter) WriteFile(filename string, data []byte, perm fs.FileMode) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WriteFile", filename, data, perm)
	ret0, _ := ret[0].(error)
	return ret0
}

// WriteFile indicates an expected call of WriteFile.
func (mr *MockFileWriterMockRecorder) WriteFile(filename, data, perm interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WriteFile", reflect.TypeOf((*MockFileWriter)(nil).WriteFile), filename, data, perm)
}
