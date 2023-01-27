package handler

import (
	"cosmart-library/entity/form"
	"cosmart-library/entity/response"
	"cosmart-library/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type BorrowScheduleHandler struct {
	BorrowScheduleServices service.BorrowScheduleService
}

func NewBorrowScheduleHandler(BorrowScheduleServices service.BorrowScheduleService) *BorrowScheduleHandler {
	return &BorrowScheduleHandler{
		BorrowScheduleServices: BorrowScheduleServices,
	}
}

func (h *BorrowScheduleHandler) Register(router *gin.Engine) {
	router.POST("/schedule-borrow", h.PostCreateBorrowSchedule)
	router.POST("/schedule-return", h.PostReturnBook)
}

func (h *BorrowScheduleHandler) PostCreateBorrowSchedule(c *gin.Context) {
	var form form.FormCreateSchedule

	if err := c.BindJSON(&form); err != nil {
		response.ErrorInvalidParameter(c)
		return
	}

	result, err := h.BorrowScheduleServices.CreateSchedule(form)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.OKWithHTTPCode(c, http.StatusCreated, "Borrow Schedule Created Successfully", result)
}

func (h *BorrowScheduleHandler) PostReturnBook(c *gin.Context) {
	var form form.FormReturnBook

	if err := c.BindJSON(&form); err != nil {
		response.ErrorInvalidParameter(c)
		return
	}

	result, err := h.BorrowScheduleServices.ReturnBook(form)
	if err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	response.OK(c, result)
}
