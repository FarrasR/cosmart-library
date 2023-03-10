// Code generated by mockery v2.16.0. DO NOT EDIT.

package mocks

import (
	form "cosmart-library/entity/form"

	mock "github.com/stretchr/testify/mock"

	model "cosmart-library/entity/model"
)

// BookService is an autogenerated mock type for the BookService type
type BookService struct {
	mock.Mock
}

// CreateBook provides a mock function with given fields: _a0
func (_m *BookService) CreateBook(_a0 form.FormCreateBook) (model.Book, error) {
	ret := _m.Called(_a0)

	var r0 model.Book
	if rf, ok := ret.Get(0).(func(form.FormCreateBook) model.Book); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(model.Book)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(form.FormCreateBook) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetBookById provides a mock function with given fields: bookId
func (_m *BookService) GetBookById(bookId int) (model.Book, error) {
	ret := _m.Called(bookId)

	var r0 model.Book
	if rf, ok := ret.Get(0).(func(int) model.Book); ok {
		r0 = rf(bookId)
	} else {
		r0 = ret.Get(0).(model.Book)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(bookId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetBooks provides a mock function with given fields: _a0
func (_m *BookService) GetBooks(_a0 form.FormGetBooks) ([]model.Book, error) {
	ret := _m.Called(_a0)

	var r0 []model.Book
	if rf, ok := ret.Get(0).(func(form.FormGetBooks) []model.Book); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.Book)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(form.FormGetBooks) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewBookService interface {
	mock.TestingT
	Cleanup(func())
}

// NewBookService creates a new instance of BookService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewBookService(t mockConstructorTestingTNewBookService) *BookService {
	mock := &BookService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
