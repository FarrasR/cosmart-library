package service_test

import (
	"cosmart-library/entity/form"
	"cosmart-library/entity/model"
	"cosmart-library/mocks"
	"cosmart-library/service"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
)

type BorrowServiceSuite struct {
	suite.Suite
	service service.BorrowScheduleService
	repo    *mocks.BorrowScheduleRepository
}

func (s *BorrowServiceSuite) SetupTest() {
	s.repo = mocks.NewBorrowScheduleRepository(s.T())
	s.service = service.NewBorrowScheduleService(
		s.repo,
	)
}

func (s *BorrowServiceSuite) TestCreateScheduleSuccess() {
	pickupTime := time.Now()
	dueTime := pickupTime.AddDate(0, 0, 7)

	form := form.FormCreateSchedule{
		Name:       "bung messi",
		BookId:     1,
		PickupTime: pickupTime,
	}

	expectedSchedule := model.BorrowSchedule{
		Name:       form.Name,
		BookId:     form.BookId,
		PickupTime: &form.PickupTime,
		DueTime:    &dueTime,
	}

	s.repo.On("Create", expectedSchedule).Return(expectedSchedule, nil)

	resultSchedule, err := s.service.CreateSchedule(form)
	assert := assert.New(s.T())
	assert.Nil(err)
	assert.Equal(resultSchedule, expectedSchedule)
}

func (s *BorrowServiceSuite) TestCreateScheduleFailed() {
	pickupTime := time.Now()
	dueTime := pickupTime.AddDate(0, 0, 7)

	form := form.FormCreateSchedule{
		Name:       "bung messi",
		BookId:     1,
		PickupTime: pickupTime,
	}

	expectedSchedule := model.BorrowSchedule{
		Name:       form.Name,
		BookId:     form.BookId,
		PickupTime: &pickupTime,
		DueTime:    &dueTime,
	}

	s.repo.On("Create", expectedSchedule).Return(model.BorrowSchedule{}, errors.New("error"))

	resultSchedule, err := s.service.CreateSchedule(form)
	assert := assert.New(s.T())

	assert.NotNil(err)
	assert.Equal(resultSchedule, model.BorrowSchedule{})
}

func (s *BorrowServiceSuite) TestReturnBookSuccess() {
	pickupTime := time.Now()
	dueTime := pickupTime.AddDate(0, 0, 7)
	returnTime := pickupTime.AddDate(0, 0, 3)

	form := form.FormReturnBook{
		ScheduleId: 1,
		ReturnTime: returnTime,
	}

	initialSchedule := model.BorrowSchedule{
		Model: gorm.Model{

			ID: 1,
		},
		Name:       "bung messi",
		BookId:     1,
		PickupTime: &pickupTime,
		DueTime:    &dueTime,
	}

	expectedSchedule := initialSchedule
	expectedSchedule.ReturnTime = &returnTime

	s.repo.On("FindOne", form.ScheduleId).Return(initialSchedule, nil)
	s.repo.On("Update", expectedSchedule).Return(expectedSchedule, nil)

	resultSchedule, err := s.service.ReturnBook(form)

	assert := assert.New(s.T())
	assert.Nil(err)
	assert.Equal(resultSchedule, expectedSchedule)
}

func (s *BorrowServiceSuite) TestReturnBookFailedOnFindOne() {
	returnTime := time.Now().AddDate(0, 0, 3)

	form := form.FormReturnBook{
		ScheduleId: 1,
		ReturnTime: returnTime,
	}

	s.repo.On("FindOne", form.ScheduleId).Return(model.BorrowSchedule{}, errors.New("error"))

	resultSchedule, err := s.service.ReturnBook(form)

	assert := assert.New(s.T())

	s.repo.AssertNotCalled(s.T(), "Update")
	assert.NotNil(err)
	assert.Equal(resultSchedule, model.BorrowSchedule{})
}

func TestBorrowService(t *testing.T) {
	suite.Run(t, new(BorrowServiceSuite))
}
