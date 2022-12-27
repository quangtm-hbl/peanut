// Code generated by MockGen. DO NOT EDIT.
// Source: repository/book.go

// Package mock is a generated GoMock package.
package mock

import (
	domain "peanut/domain"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockBookRepo is a mock of BookRepo interface.
type MockBookRepo struct {
	ctrl     *gomock.Controller
	recorder *MockBookRepoMockRecorder
}

// MockBookRepoMockRecorder is the mock recorder for MockBookRepo.
type MockBookRepoMockRecorder struct {
	mock *MockBookRepo
}

// NewMockBookRepo creates a new mock instance.
func NewMockBookRepo(ctrl *gomock.Controller) *MockBookRepo {
	mock := &MockBookRepo{ctrl: ctrl}
	mock.recorder = &MockBookRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBookRepo) EXPECT() *MockBookRepoMockRecorder {
	return m.recorder
}

// CreateBook mocks base method.
func (m *MockBookRepo) CreateBook(b domain.Book) (*domain.Book, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateBook", b)
	ret0, _ := ret[0].(*domain.Book)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateBook indicates an expected call of CreateBook.
func (mr *MockBookRepoMockRecorder) CreateBook(b interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateBook", reflect.TypeOf((*MockBookRepo)(nil).CreateBook), b)
}

// DeleteBook mocks base method.
func (m *MockBookRepo) DeleteBook(id int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteBook", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteBook indicates an expected call of DeleteBook.
func (mr *MockBookRepoMockRecorder) DeleteBook(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteBook", reflect.TypeOf((*MockBookRepo)(nil).DeleteBook), id)
}

// GetBookById mocks base method.
func (m *MockBookRepo) GetBookById(id int) (*domain.Book, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBookById", id)
	ret0, _ := ret[0].(*domain.Book)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBookById indicates an expected call of GetBookById.
func (mr *MockBookRepoMockRecorder) GetBookById(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBookById", reflect.TypeOf((*MockBookRepo)(nil).GetBookById), id)
}

// GetBooks mocks base method.
func (m *MockBookRepo) GetBooks() ([]domain.Book, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBooks")
	ret0, _ := ret[0].([]domain.Book)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBooks indicates an expected call of GetBooks.
func (mr *MockBookRepoMockRecorder) GetBooks() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBooks", reflect.TypeOf((*MockBookRepo)(nil).GetBooks))
}

// UpdateBook mocks base method.
func (m *MockBookRepo) UpdateBook(updatingBook domain.Book, id int) (*domain.Book, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateBook", updatingBook, id)
	ret0, _ := ret[0].(*domain.Book)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateBook indicates an expected call of UpdateBook.
func (mr *MockBookRepoMockRecorder) UpdateBook(updatingBook, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateBook", reflect.TypeOf((*MockBookRepo)(nil).UpdateBook), updatingBook, id)
}
