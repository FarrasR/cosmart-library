package handler

import (
	"cosmart-library/entity/form"
	"cosmart-library/entity/response"
	"cosmart-library/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type BookHandler struct {
	BookService service.BookService
}

func NewBookHandler(BookService service.BookService) *BookHandler {
	return &BookHandler{
		BookService: BookService,
	}
}

func (h *BookHandler) Register(router *gin.Engine) {
	router.GET("/books", h.GetBooks)
	router.GET("/books/:id", h.GetBookById)
	router.POST("/books", h.PostBook)
}

func (h *BookHandler) GetBooks(c *gin.Context) {
	limit, err := strconv.Atoi(c.DefaultQuery("limit", "10"))
	if err != nil {
		response.ErrorInvalidParameter(c)
		return
	}
	offset, err := strconv.Atoi(c.DefaultQuery("offset", "0"))
	if err != nil {
		response.ErrorInvalidParameter(c)
		return
	}

	books, err := h.BookService.GetBooks(limit, offset)
	if err != nil {
		response.ErrorInvalidParameter(c)
		return
	}

	response.OK(c, books)
}

func (h *BookHandler) GetBookById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.ErrorInvalidParameter(c)
		return
	}

	book, err := h.BookService.GetBookById(id)

	if err != nil {
		response.ErrorInvalidParameter(c)
		return
	}
	response.OK(c, book)
}

func (h *BookHandler) PostBook(c *gin.Context) {
	var form form.FormCreateBook

	if err := c.BindJSON(&form); err != nil {
		response.ErrorInvalidParameter(c)
		return
	}

	book, err := h.BookService.CreateBook(form)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.OKWithHTTPCode(c, http.StatusCreated, "Book created successfully", book)
}
