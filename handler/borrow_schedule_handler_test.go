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
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type BorrowScheduleHandlerSuite struct {
	suite.Suite
	handler *handler.BorrowScheduleHandler
	service *mocks.BorrowScheduleService
	engine  *gin.Engine
}

func (s *BorrowScheduleHandlerSuite) SetupSuite() {
	s.service = mocks.NewBorrowScheduleService(s.T())
	s.handler = handler.NewBorrowScheduleHandler(s.service)
	s.engine = router.BuildHandler(s.handler)
}

func (s *BorrowScheduleHandlerSuite) AfterTest(suiteName, testName string) {
	for _, call := range s.service.ExpectedCalls {
		call.Unset()
	}
}

func (s *BorrowScheduleHandlerSuite) TestPostCreateBorrowScheduleSuccess() {
	pickupTime := time.Date(2023, 1, 1, 1, 1, 1, 1, time.UTC)
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

	s.service.On("CreateSchedule", form).Return(expectedSchedule, nil)

	body, _ := json.Marshal(form)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/schedule-borrow", bytes.NewReader(body))
	s.engine.ServeHTTP(w, req)

	assert := assert.New(s.T())
	assert.Equal(http.StatusCreated, w.Code)
	assert.Equal("{\"success\":true,\"message\":\"Borrow Schedule Created Successfully\",\"data\":{\"ID\":0,\"CreatedAt\":\"0001-01-01T00:00:00Z\",\"UpdatedAt\":\"0001-01-01T00:00:00Z\",\"DeletedAt\":null,\"Name\":\"bung messi\",\"BookId\":1,\"PickupTime\":\"2023-01-01T01:01:01.000000001Z\",\"DueTime\":\"2023-01-08T01:01:01.000000001Z\",\"ReturnTime\":null}}",
		w.Body.String())
}

func (s *BorrowScheduleHandlerSuite) TestPostCreateBorrowScheduleInvalidParameter() {
	form := struct {
		Name       string `json:"name"`
		BookId     int    `json:"book_id"`
		PickupTime int    `json:"pickup_time"`
	}{
		Name:       "bung messi",
		BookId:     1,
		PickupTime: 2,
	}

	body, _ := json.Marshal(form)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/schedule-borrow", bytes.NewReader(body))
	s.engine.ServeHTTP(w, req)

	assert := assert.New(s.T())
	assert.Equal(http.StatusBadRequest, w.Code)
	assert.Equal("{\"success\":false,\"message\":\"Invalid parameter\"}",
		w.Body.String())
}

func (s *BorrowScheduleHandlerSuite) TestPostCreateBorrowScheduleFailedService() {
	pickupTime := time.Date(2023, 1, 1, 1, 1, 1, 1, time.UTC)

	form := form.FormCreateSchedule{
		Name:       "bung messi",
		BookId:     1,
		PickupTime: pickupTime,
	}

	s.service.On("CreateSchedule", form).Return(model.BorrowSchedule{}, errors.New("error"))

	body, _ := json.Marshal(form)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/schedule-borrow", bytes.NewReader(body))
	s.engine.ServeHTTP(w, req)

	assert := assert.New(s.T())
	assert.Equal(http.StatusBadRequest, w.Code)
	assert.Equal("{\"success\":false,\"message\":\"error\"}",
		w.Body.String())
}

func (s *BorrowScheduleHandlerSuite) TestPostReturnBookSuccess() {
	pickupTime := time.Date(2023, 1, 1, 1, 1, 1, 1, time.UTC)
	dueTime := pickupTime.AddDate(0, 0, 7)
	returnTime := pickupTime.AddDate(0, 0, 7)

	form := form.FormReturnBook{
		ScheduleId: 1,
		ReturnTime: returnTime,
	}

	expectedSchedule := model.BorrowSchedule{
		Name:       "bung messi",
		BookId:     1,
		PickupTime: &pickupTime,
		DueTime:    &dueTime,
		ReturnTime: &returnTime,
	}

	s.service.On("ReturnBook", form).Return(expectedSchedule, nil)

	body, _ := json.Marshal(form)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/schedule-return", bytes.NewReader(body))
	s.engine.ServeHTTP(w, req)

	assert := assert.New(s.T())
	assert.Equal(http.StatusOK, w.Code)
	assert.Equal("{\"success\":true,\"message\":\"Success\",\"data\":{\"ID\":0,\"CreatedAt\":\"0001-01-01T00:00:00Z\",\"UpdatedAt\":\"0001-01-01T00:00:00Z\",\"DeletedAt\":null,\"Name\":\"bung messi\",\"BookId\":1,\"PickupTime\":\"2023-01-01T01:01:01.000000001Z\",\"DueTime\":\"2023-01-08T01:01:01.000000001Z\",\"ReturnTime\":\"2023-01-08T01:01:01.000000001Z\"}}",
		w.Body.String())
}

func (s *BorrowScheduleHandlerSuite) TestPostReturnBookInvalidParameter() {
	form := struct {
		ScheduleId int `json:"schedule_id"`
		ReturnTime int `json:"return_time"`
	}{
		ScheduleId: 1,
		ReturnTime: 1,
	}

	body, _ := json.Marshal(form)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/schedule-return", bytes.NewReader(body))
	s.engine.ServeHTTP(w, req)

	assert := assert.New(s.T())
	assert.Equal(http.StatusBadRequest, w.Code)
	assert.Equal("{\"success\":false,\"message\":\"Invalid parameter\"}",
		w.Body.String())
}

func (s *BorrowScheduleHandlerSuite) TestPostReturnBookFailedService() {
	returnTime := time.Date(2023, 1, 1, 1, 1, 1, 1, time.UTC)

	form := form.FormReturnBook{
		ScheduleId: 1,
		ReturnTime: returnTime,
	}
	s.service.On("ReturnBook", form).Return(model.BorrowSchedule{}, errors.New("error"))

	body, _ := json.Marshal(form)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/schedule-return", bytes.NewReader(body))
	s.engine.ServeHTTP(w, req)

	assert := assert.New(s.T())
	assert.Equal(http.StatusBadRequest, w.Code)
	assert.Equal("{\"success\":false,\"message\":\"error\"}",
		w.Body.String())
}

func TestBorrowScheduleHandler(t *testing.T) {
	suite.Run(t, new(BorrowScheduleHandlerSuite))
}
