package service

import (
	"cosmart-library/entity/form"
	"cosmart-library/entity/model"
	"cosmart-library/repository"
)

type BorrowScheduleService interface {
	CreateSchedule(form form.FormCreateSchedule) (model.BorrowSchedule, error)
	ReturnBook(form form.FormReturnBook) (model.BorrowSchedule, error)
}

type borrowScheduleService struct {
	BorrowScheduleRepository repository.BorrowScheduleRepository
}

func NewBorrowScheduleService(BorrowScheduleRepository repository.BorrowScheduleRepository) BorrowScheduleService {
	return &borrowScheduleService{
		BorrowScheduleRepository: BorrowScheduleRepository,
	}
}

func (s *borrowScheduleService) CreateSchedule(form form.FormCreateSchedule) (model.BorrowSchedule, error) {
	dueTime := form.PickupTime.AddDate(0, 0, 7)

	schedule := model.BorrowSchedule{
		Name:       form.Name,
		BookId:     form.BookId,
		PickupTime: &form.PickupTime,
		DueTime:    &dueTime,
	}

	result, err := s.BorrowScheduleRepository.Create(schedule)
	if err != nil {
		return model.BorrowSchedule{}, err
	}
	return result, nil
}

func (s *borrowScheduleService) ReturnBook(form form.FormReturnBook) (model.BorrowSchedule, error) {
	schedule, err := s.BorrowScheduleRepository.FindOne(form.ScheduleId)
	if err != nil {
		return model.BorrowSchedule{}, err
	}

	schedule.ReturnTime = &form.ReturnTime

	result, err := s.BorrowScheduleRepository.Update(schedule)
	if err != nil {
		return model.BorrowSchedule{}, err
	}

	return result, nil
}
