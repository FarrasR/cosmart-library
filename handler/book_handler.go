package handler

import (
	"cosmart-library/entity/form"
	"cosmart-library/entity/request"
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
	qh := request.NewQueryHelper(c)

	limit := qh.GetInt("limit", 10)
	offset := qh.GetInt("offset", 0)
	genre := qh.GetString("genre", "")

	form := form.FormGetBooks{
		Limit:  limit,
		Offset: offset,
		Genre:  genre,
	}

	books, err := h.BookService.GetBooks(form)
	if err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
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
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	response.OK(c, book)
}

func (h *BookHandler) PostBook(c *gin.Context) {
	var form form.FormCreateBook

	if err := c.ShouldBindJSON(&form); err != nil {
		response.ErrorInvalidParameter(c)
		return
	}

	book, err := h.BookService.CreateBook(form)
	if err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}
	response.OKWithHTTPCode(c, http.StatusCreated, "Book created successfully", book)
}
