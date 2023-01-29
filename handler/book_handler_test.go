package handler_test

import (
	"bytes"
	"cosmart-library/entity/form"
	"cosmart-library/entity/model"
	"cosmart-library/handler"
	"cosmart-library/mocks"
	"cosmart-library/router"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
)

type BookHandlerSuite struct {
	suite.Suite
	handler *handler.BookHandler
	service *mocks.BookService
	engine  *gin.Engine
}

func (s *BookHandlerSuite) SetupSuite() {
	s.service = mocks.NewBookService(s.T())
	s.handler = handler.NewBookHandler(s.service)
	s.engine = router.BuildHandler(s.handler)
}

func (s *BookHandlerSuite) AfterTest(suiteName, testName string) {
	for _, call := range s.service.ExpectedCalls {
		call.Unset()
	}
}

func (s *BookHandlerSuite) TestGetBooksFail() {
	form := form.FormGetBooks{
		Limit:  10,
		Offset: 0,
		Genre:  "test",
	}

	s.service.On("GetBooks", form).Return([]model.Book{}, errors.New("error"))

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/books?limit=10&offset=0&genre=test", nil)
	s.engine.ServeHTTP(w, req)

	assert := assert.New(s.T())
	assert.Equal(http.StatusBadRequest, w.Code)
	assert.Equal("{\"success\":false,\"message\":\"error\"}", w.Body.String())
}

func (s *BookHandlerSuite) TestGetBooksSuccess() {
	form := form.FormGetBooks{
		Limit:  10,
		Offset: 0,
		Genre:  "test",
	}

	var expectedBooks = []model.Book{
		{
			Model: gorm.Model{
				ID: 1,
			},
			Title:   "testing for golang",
			Author:  "lorem bin ipsum",
			Edition: 1,
			Genre:   "test",
		}, {
			Model: gorm.Model{
				ID: 2,
			},
			Title:  "testing for golang 2",
			Author: "ipsum bin dolor",
			Genre:  "test",
		},
	}

	s.service.On("GetBooks", form).Return(expectedBooks, nil)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/books?limit=10&offset=0&genre=test", nil)
	s.engine.ServeHTTP(w, req)

	assert := assert.New(s.T())
	assert.Equal(http.StatusOK, w.Code)
	assert.Equal("{\"success\":true,\"message\":\"Success\",\"data\":[{\"ID\":1,\"CreatedAt\":\"0001-01-01T00:00:00Z\",\"UpdatedAt\":\"0001-01-01T00:00:00Z\",\"DeletedAt\":null,\"Title\":\"testing for golang\",\"Author\":\"lorem bin ipsum\",\"Edition\":1,\"Genre\":\"test\"},{\"ID\":2,\"CreatedAt\":\"0001-01-01T00:00:00Z\",\"UpdatedAt\":\"0001-01-01T00:00:00Z\",\"DeletedAt\":null,\"Title\":\"testing for golang 2\",\"Author\":\"ipsum bin dolor\",\"Edition\":0,\"Genre\":\"test\"}]}",
		w.Body.String())
}

func (s *BookHandlerSuite) TestGetBookByIdSuccess() {
	var expectedBook = model.Book{
		Model: gorm.Model{
			ID: 1,
		},
		Title:   "testing for golang",
		Author:  "lorem bin ipsum",
		Edition: 1,
		Genre:   "test",
	}

	s.service.On("GetBookById", 1).Return(expectedBook, nil)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/books/1", nil)
	s.engine.ServeHTTP(w, req)

	assert := assert.New(s.T())
	assert.Equal(http.StatusOK, w.Code)
	assert.Equal("{\"success\":true,\"message\":\"Success\",\"data\":{\"ID\":1,\"CreatedAt\":\"0001-01-01T00:00:00Z\",\"UpdatedAt\":\"0001-01-01T00:00:00Z\",\"DeletedAt\":null,\"Title\":\"testing for golang\",\"Author\":\"lorem bin ipsum\",\"Edition\":1,\"Genre\":\"test\"}}",
		w.Body.String())
}

func (s *BookHandlerSuite) TestGetBookByIdFailedWrongParameter() {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/books/asoiaf", nil)
	s.engine.ServeHTTP(w, req)

	assert := assert.New(s.T())
	assert.Equal(http.StatusBadRequest, w.Code)
	assert.Equal("{\"success\":false,\"message\":\"Invalid parameter\"}",
		w.Body.String())
}

func (s *BookHandlerSuite) TestGetBookByIdFailedService() {
	s.service.On("GetBookById", 1).Return(model.Book{}, errors.New("error"))

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/books/1", nil)
	s.engine.ServeHTTP(w, req)

	assert := assert.New(s.T())
	assert.Equal(http.StatusBadRequest, w.Code)
	assert.Equal("{\"success\":false,\"message\":\"error\"}",
		w.Body.String())
}

func (s *BookHandlerSuite) TestPostBookSuccess() {
	form := form.FormCreateBook{
		Title:   "testing for golang",
		Author:  "lorem bin ipsum",
		Edition: 1,
		Genre:   "test",
	}

	expectedBook := model.Book{
		Model: gorm.Model{
			ID: 1,
		},
		Title:   "testing for golang",
		Author:  "lorem bin ipsum",
		Edition: 1,
		Genre:   "test",
	}

	s.service.On("CreateBook", form).Return(expectedBook, nil)

	body, _ := json.Marshal(form)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/books", bytes.NewReader(body))
	s.engine.ServeHTTP(w, req)

	assert := assert.New(s.T())
	assert.Equal(http.StatusCreated, w.Code)
	assert.Equal("{\"success\":true,\"message\":\"Book created successfully\",\"data\":{\"ID\":1,\"CreatedAt\":\"0001-01-01T00:00:00Z\",\"UpdatedAt\":\"0001-01-01T00:00:00Z\",\"DeletedAt\":null,\"Title\":\"testing for golang\",\"Author\":\"lorem bin ipsum\",\"Edition\":1,\"Genre\":\"test\"}}",
		w.Body.String())
}

func (s *BookHandlerSuite) TestPostBookFailedWrongParameter() {
	form := struct {
		Title   string `json:"title"`
		Author  int    `json:"author"`
		Edition int    `json:"edition"`
		Genre   int    `json:"genre"`
	}{
		Title:   "wrong title",
		Author:  1,
		Edition: 1,
		Genre:   1,
	}

	body, _ := json.Marshal(form)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/books", bytes.NewReader(body))
	s.engine.ServeHTTP(w, req)

	assert := assert.New(s.T())
	assert.Equal(http.StatusBadRequest, w.Code)
	assert.Equal("{\"success\":false,\"message\":\"Invalid parameter\"}",
		w.Body.String())
}

func (s *BookHandlerSuite) TestPostBookFailedService() {

	form := form.FormCreateBook{
		Title:   "testing for golang",
		Author:  "lorem bin ipsum",
		Edition: 1,
		Genre:   "test",
	}

	s.service.On("CreateBook", form).Return(model.Book{}, errors.New("error"))

	body, _ := json.Marshal(form)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/books", bytes.NewReader(body))
	s.engine.ServeHTTP(w, req)

	assert := assert.New(s.T())
	assert.Equal(http.StatusBadRequest, w.Code)
	assert.Equal("{\"success\":false,\"message\":\"error\"}",
		w.Body.String())
}

func TestBookHandler(t *testing.T) {
	suite.Run(t, new(BookHandlerSuite))
}
