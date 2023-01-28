package repository_test

import (
	"cosmart-library/database"
	"cosmart-library/entity/model"
	"cosmart-library/repository"
	"database/sql"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type BorrowScheduleRepositorySuite struct {
	suite.Suite
	repo    repository.BorrowScheduleRepository
	sqlmock sqlmock.Sqlmock
}

func (s *BorrowScheduleRepositorySuite) SetupTest() {
	var err error
	var conn *sql.DB

	conn, s.sqlmock, err = sqlmock.New()
	require.NoError(s.T(), err)

	DB, err := gorm.Open(mysql.New(mysql.Config{
		Conn:                      conn,
		DriverName:                "mysql",
		SkipInitializeWithVersion: true,
	}), &gorm.Config{})
	require.NoError(s.T(), err)

	instance := database.NewDatabaseInstance(DB)
	s.repo = repository.NewBorrowScheduleRepository(
		instance,
	)
}

func (s *BorrowScheduleRepositorySuite) TestCreate() {
	pickupDate := time.Now()
	dueDate := pickupDate.AddDate(0, 0, 7)

	var expectedSchedule = model.BorrowSchedule{
		Name:       "bung messi",
		BookId:     1,
		PickupTime: &pickupDate,
		DueTime:    &dueDate,
	}

	s.sqlmock.ExpectBegin()
	s.sqlmock.ExpectExec("INSERT INTO `borrow_schedules`").
		WillReturnResult(sqlmock.NewResult(1, 1))
	s.sqlmock.ExpectCommit()

	resultSchedule, err := s.repo.Create(expectedSchedule)
	assert := assert.New(s.T())
	assert.Nil(err)
	assert.Equal(resultSchedule.Name, expectedSchedule.Name)
	assert.Equal(resultSchedule.BookId, expectedSchedule.BookId)
	assert.Equal(resultSchedule.PickupTime, expectedSchedule.PickupTime)
	assert.Equal(resultSchedule.DueTime, expectedSchedule.DueTime)
}

func (s *BorrowScheduleRepositorySuite) TestUpdate() {
	pickupDate := time.Now()
	dueDate := pickupDate.AddDate(0, 0, 7)

	var expectedSchedule = model.BorrowSchedule{
		Model: gorm.Model{
			ID: 1,
		},
		Name:       "bung messi",
		BookId:     1,
		PickupTime: &pickupDate,
		DueTime:    &dueDate,
		ReturnTime: &dueDate,
	}
	s.sqlmock.ExpectBegin()
	s.sqlmock.ExpectExec("UPDATE `borrow_schedules`").
		WillReturnResult(sqlmock.NewResult(1, 1))
	s.sqlmock.ExpectCommit()

	resultSchedule, err := s.repo.Update(expectedSchedule)
	assert := assert.New(s.T())
	assert.Nil(err)
	assert.Equal(resultSchedule.Name, expectedSchedule.Name)
	assert.Equal(resultSchedule.BookId, expectedSchedule.BookId)
	assert.Equal(resultSchedule.PickupTime, expectedSchedule.PickupTime)
	assert.Equal(resultSchedule.DueTime, expectedSchedule.DueTime)
}

func (s *BorrowScheduleRepositorySuite) TestFindOne() {
	pickupDate := time.Now()
	dueDate := pickupDate.AddDate(0, 0, 7)

	var expectedSchedule = model.BorrowSchedule{
		Model: gorm.Model{
			ID: 1,
		},
		Name:       "bung messi",
		BookId:     1,
		PickupTime: &pickupDate,
		DueTime:    &dueDate,
		ReturnTime: &dueDate,
	}

	s.sqlmock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `borrow_schedules` WHERE `borrow_schedules`.`id` = ? AND `borrow_schedules`.`deleted_at` IS NULL ORDER BY `borrow_schedules`.`id` LIMIT 1")).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "book_id", "pickup_time", "due_time", "return_time"}).
			AddRow(expectedSchedule.ID,
				expectedSchedule.Name,
				expectedSchedule.BookId,
				expectedSchedule.PickupTime,
				expectedSchedule.DueTime,
				expectedSchedule.ReturnTime))

	resultSchedule, err := s.repo.FindOne(1)
	assert := assert.New(s.T())

	assert.Nil(err)
	assert.Equal(resultSchedule.ID, expectedSchedule.ID)
	assert.Equal(resultSchedule.Name, expectedSchedule.Name)
	assert.Equal(resultSchedule.BookId, expectedSchedule.BookId)
	assert.Equal(resultSchedule.PickupTime, expectedSchedule.PickupTime)
	assert.Equal(resultSchedule.DueTime, expectedSchedule.DueTime)
	assert.Equal(resultSchedule.ReturnTime, expectedSchedule.ReturnTime)
}

func TestBorrowScheduleRepo(t *testing.T) {
	suite.Run(t, new(BorrowScheduleRepositorySuite))
}
