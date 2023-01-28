package repository_test

import (
	"cosmart-library/database"
	"cosmart-library/entity/model"
	"cosmart-library/repository"
	"database/sql"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type BookRepositorySuite struct {
	suite.Suite
	repo    repository.BookRepository
	sqlmock sqlmock.Sqlmock
}

func (s *BookRepositorySuite) SetupSuite() {
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
	s.repo = repository.NewBookRepository(
		instance,
	)
}

func (s *BookRepositorySuite) TestFindOneSuccess() {
	var expectedBook = model.Book{
		Model: gorm.Model{
			ID: 1,
		},
		Title:   "test",
		Author:  "lorem bin ipsum",
		Edition: 1,
	}

	s.sqlmock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `books` WHERE `books`.`id` = ? AND `books`.`deleted_at` IS NULL ORDER BY `books`.`id` LIMIT 1")).
		WillReturnRows(sqlmock.NewRows([]string{"id", "title", "author", "edition"}).
			AddRow(expectedBook.ID, expectedBook.Title, expectedBook.Author, expectedBook.Edition))

	resultBook, err := s.repo.FindOne(1)

	assert := assert.New(s.T())

	assert.Nil(err)
	assert.Equal(resultBook, expectedBook, "book should be equal")
}

func (s *BookRepositorySuite) TestFindOneFail() {
	s.sqlmock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `books` WHERE `books`.`id` = ? AND `books`.`deleted_at` IS NULL ORDER BY `books`.`id` LIMIT 1")).
		WillReturnRows(sqlmock.NewRows([]string{"id", "title", "author", "edition"}))

	result, err := s.repo.FindOne(1)
	assert := assert.New(s.T())

	assert.NotNil(err)
	assert.Equal(result, model.Book{})
}

func (s *BookRepositorySuite) TestFindSuccess() {
	var expectedBooks = []model.Book{
		{
			Model: gorm.Model{
				ID: 1,
			},
			Title:   "test",
			Author:  "lorem bin ipsum",
			Edition: 1,
		}, {
			Model: gorm.Model{
				ID: 2,
			},
			Title:   "test 2",
			Author:  "ipsum bin dolor",
			Edition: 2,
		},
	}

	s.sqlmock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `books` WHERE `books`.`deleted_at` IS NULL LIMIT 0 OFFSET 10")).
		WillReturnRows(sqlmock.NewRows([]string{"id", "title", "author", "edition"}).
			AddRow(expectedBooks[0].ID, expectedBooks[0].Title, expectedBooks[0].Author, expectedBooks[0].Edition).
			AddRow(expectedBooks[1].ID, expectedBooks[1].Title, expectedBooks[1].Author, expectedBooks[1].Edition))

	resultBooks, err := s.repo.Find(0, 10)

	assert := assert.New(s.T())
	assert.Nil(err)
	assert.Equal(resultBooks[0], expectedBooks[0])
	assert.Equal(resultBooks[1], expectedBooks[1])
}

func (s *BookRepositorySuite) TestCreate() {
	var expectedBook = model.Book{
		Title:   "test",
		Author:  "lorem bin ipsum",
		Edition: 1,
	}

	s.sqlmock.ExpectBegin()
	s.sqlmock.ExpectExec("INSERT INTO `books`").
		WillReturnResult(sqlmock.NewResult(1, 1))
	s.sqlmock.ExpectCommit()

	resultBook, err := s.repo.Create(expectedBook)
	assert := assert.New(s.T())
	assert.Nil(err)
	assert.Equal(resultBook.Author, expectedBook.Author)
	assert.Equal(resultBook.Title, expectedBook.Title)
	assert.Equal(resultBook.Edition, expectedBook.Edition)
}

func TestBookRepo(t *testing.T) {
	suite.Run(t, new(BookRepositorySuite))
}
