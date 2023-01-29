package service

import (
	"cosmart-library/entity/form"
	"cosmart-library/entity/model"
	"cosmart-library/repository"
)

type BookService interface {
	GetBookById(bookId int) (model.Book, error)
	GetBooks(form form.FormGetBooks) ([]model.Book, error)
	CreateBook(form form.FormCreateBook) (model.Book, error)
}

type bookService struct {
	BookRepository repository.BookRepository
}

func NewBookService(BookRepository repository.BookRepository) BookService {
	return &bookService{
		BookRepository: BookRepository,
	}
}

func (s *bookService) GetBookById(bookId int) (model.Book, error) {
	return s.BookRepository.FindOne(bookId)
}

func (s *bookService) GetBooks(form form.FormGetBooks) ([]model.Book, error) {
	return s.BookRepository.Find(form.Limit, form.Offset, form.Genre)
}

func (s *bookService) CreateBook(form form.FormCreateBook) (model.Book, error) {
	book := model.Book{
		Title:   form.Title,
		Author:  form.Author,
		Edition: form.Edition,
		Genre:   form.Genre,
	}

	return s.BookRepository.Create(book)
}
