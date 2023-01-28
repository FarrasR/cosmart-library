package service_test

import (
	"cosmart-library/entity/form"
	"cosmart-library/entity/model"
	"cosmart-library/service"
	"errors"

	"cosmart-library/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
)

type BookServiceSuite struct {
	suite.Suite
	service service.BookService
	repo    *mocks.BookRepository
}

func (s *BookServiceSuite) SetupTest() {
	s.repo = mocks.NewBookRepository(s.T())
	s.service = service.NewBookService(
		s.repo,
	)
}

func (s *BookServiceSuite) TestGetBookByIdSuccess() {
	var expectedBook = model.Book{
		Model: gorm.Model{
			ID: 1,
		},
		Title:   "test",
		Author:  "lorem bin ipsum",
		Edition: 1,
	}

	s.repo.On("FindOne", 1).Return(expectedBook, nil)

	resultBook, err := s.service.GetBookById(1)
	assert := assert.New(s.T())
	assert.Nil(err)
	assert.Equal(resultBook, expectedBook)
}

func (s *BookServiceSuite) TestGetBookByIdFailed() {
	s.repo.On("FindOne", 1).Return(model.Book{}, gorm.ErrRecordNotFound)

	resultBook, err := s.service.GetBookById(1)

	assert := assert.New(s.T())
	assert.NotNil(err)
	assert.Equal(resultBook, model.Book{})
}

func (s *BookServiceSuite) TestGetBooksSuccess() {
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

	s.repo.On("Find", 10, 0).Return(expectedBooks, nil)

	resultBooks, err := s.service.GetBooks(10, 0)

	assert := assert.New(s.T())
	assert.Nil(err)
	assert.Equal(resultBooks, expectedBooks)
}

func (s *BookServiceSuite) TestGetBooksFailed() {
	s.repo.On("Find", 10, 0).Return([]model.Book{}, errors.New("something went wrong"))

	resultBooks, err := s.service.GetBooks(10, 0)

	assert := assert.New(s.T())
	assert.NotNil(err)
	assert.Equal(resultBooks, []model.Book{})
}

func (s *BookServiceSuite) TestCreateBookSuccess() {
	var form = form.FormCreateBook{
		Title:   "test",
		Author:  "lorem bin ipsum",
		Edition: 1,
	}

	var expectedBook = model.Book{
		Title:   "test",
		Author:  "lorem bin ipsum",
		Edition: 1,
	}

	s.repo.On("Create", expectedBook).Return(expectedBook, nil)

	resultBook, err := s.service.CreateBook(form)
	assert := assert.New(s.T())
	assert.Nil(err)
	assert.Equal(resultBook, expectedBook)
}

func (s *BookServiceSuite) TestCreateBookFailed() {
	var form = form.FormCreateBook{
		Title:   "test",
		Author:  "lorem bin ipsum",
		Edition: 1,
	}

	var expectedBook = model.Book{
		Title:   "test",
		Author:  "lorem bin ipsum",
		Edition: 1,
	}

	s.repo.On("Create", expectedBook).Return(model.Book{}, errors.New("something went wrong"))

	resultBook, err := s.service.CreateBook(form)
	assert := assert.New(s.T())
	assert.NotNil(err)
	assert.Equal(resultBook, model.Book{})
}

func TestBookService(t *testing.T) {
	suite.Run(t, new(BookServiceSuite))
}
