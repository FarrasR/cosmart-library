package service_test

import (
	"cosmart-library/mocks"
	"cosmart-library/service"

	"github.com/stretchr/testify/suite"
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
