// Code generated by mockery v2.16.0. DO NOT EDIT.

package mocks

import (
	model "cosmart-library/entity/model"

	mock "github.com/stretchr/testify/mock"
)

// BookRepository is an autogenerated mock type for the BookRepository type
type BookRepository struct {
	mock.Mock
}

// Create provides a mock function with given fields: book
func (_m *BookRepository) Create(book model.Book) (model.Book, error) {
	ret := _m.Called(book)

	var r0 model.Book
	if rf, ok := ret.Get(0).(func(model.Book) model.Book); ok {
		r0 = rf(book)
	} else {
		r0 = ret.Get(0).(model.Book)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(model.Book) error); ok {
		r1 = rf(book)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Find provides a mock function with given fields: limit, offset, genre
func (_m *BookRepository) Find(limit int, offset int, genre string) ([]model.Book, error) {
	ret := _m.Called(limit, offset, genre)

	var r0 []model.Book
	if rf, ok := ret.Get(0).(func(int, int, string) []model.Book); ok {
		r0 = rf(limit, offset, genre)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.Book)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int, int, string) error); ok {
		r1 = rf(limit, offset, genre)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindOne provides a mock function with given fields: id
func (_m *BookRepository) FindOne(id int) (model.Book, error) {
	ret := _m.Called(id)

	var r0 model.Book
	if rf, ok := ret.Get(0).(func(int) model.Book); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(model.Book)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewBookRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewBookRepository creates a new instance of BookRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewBookRepository(t mockConstructorTestingTNewBookRepository) *BookRepository {
	mock := &BookRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
