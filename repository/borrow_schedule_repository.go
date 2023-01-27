package repository

import (
	"cosmart-library/database"
	"cosmart-library/entity/model"
)

type BorrowScheduleRepository interface {
	Create(schedule model.BorrowSchedule) (model.BorrowSchedule, error)
	Update(schedule model.BorrowSchedule) (model.BorrowSchedule, error)
	FindOne(id int) (model.BorrowSchedule, error)
}

type borrowScheduleRepository struct {
	DatabaseInstance database.DatabaseInstance
}

func NewBorrowScheduleRepository(DatabaseInstance database.DatabaseInstance) BorrowScheduleRepository {
	return &borrowScheduleRepository{
		DatabaseInstance: DatabaseInstance,
	}
}

func (r *borrowScheduleRepository) Create(schedule model.BorrowSchedule) (model.BorrowSchedule, error) {
	if result := r.DatabaseInstance.GetConn().Create(&schedule); result.Error != nil {
		return model.BorrowSchedule{}, result.Error
	}

	return schedule, nil
}

func (r *borrowScheduleRepository) Update(schedule model.BorrowSchedule) (model.BorrowSchedule, error) {
	if result := r.DatabaseInstance.GetConn().Save(&schedule); result.Error != nil {
		return model.BorrowSchedule{}, result.Error
	}

	return schedule, nil
}

func (r *borrowScheduleRepository) FindOne(id int) (model.BorrowSchedule, error) {
	var schedule model.BorrowSchedule

	if result := r.DatabaseInstance.GetConn().First(&schedule, id); result.Error != nil {
		return model.BorrowSchedule{}, result.Error
	}

	return schedule, nil
}
