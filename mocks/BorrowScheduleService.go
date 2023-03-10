// Code generated by mockery v2.16.0. DO NOT EDIT.

package mocks

import (
	form "cosmart-library/entity/form"

	mock "github.com/stretchr/testify/mock"

	model "cosmart-library/entity/model"
)

// BorrowScheduleService is an autogenerated mock type for the BorrowScheduleService type
type BorrowScheduleService struct {
	mock.Mock
}

// CreateSchedule provides a mock function with given fields: _a0
func (_m *BorrowScheduleService) CreateSchedule(_a0 form.FormCreateSchedule) (model.BorrowSchedule, error) {
	ret := _m.Called(_a0)

	var r0 model.BorrowSchedule
	if rf, ok := ret.Get(0).(func(form.FormCreateSchedule) model.BorrowSchedule); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(model.BorrowSchedule)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(form.FormCreateSchedule) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ReturnBook provides a mock function with given fields: _a0
func (_m *BorrowScheduleService) ReturnBook(_a0 form.FormReturnBook) (model.BorrowSchedule, error) {
	ret := _m.Called(_a0)

	var r0 model.BorrowSchedule
	if rf, ok := ret.Get(0).(func(form.FormReturnBook) model.BorrowSchedule); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(model.BorrowSchedule)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(form.FormReturnBook) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewBorrowScheduleService interface {
	mock.TestingT
	Cleanup(func())
}

// NewBorrowScheduleService creates a new instance of BorrowScheduleService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewBorrowScheduleService(t mockConstructorTestingTNewBorrowScheduleService) *BorrowScheduleService {
	mock := &BorrowScheduleService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
