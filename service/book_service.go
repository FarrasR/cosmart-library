package service

import (
	"cosmart-library/entity/form"
	"cosmart-library/entity/model"
	"cosmart-library/repository"
)

type BookService interface {
	GetBookById(bookId int) (model.Book, error)
	GetBooks(limit int, offset int) ([]model.Book, error)
	CreateBook(form form.FormCreateBook) (model.Book, error)
}

type bookService struct {
	bookRepository repository.BookRepository
}

func NewBookService(bookRepository repository.BookRepository) BookService {
	return &bookService{
		bookRepository: bookRepository,
	}
}

func (s *bookService) GetBookById(bookId int) (model.Book, error) {
	return s.bookRepository.FindOne(bookId)
}

func (s *bookService) GetBooks(limit int, offset int) ([]model.Book, error) {
	return s.bookRepository.Find(limit, offset)
}

func (s *bookService) CreateBook(form form.FormCreateBook) (model.Book, error) {
	book := model.Book{
		Title:   form.Title,
		Author:  form.Author,
		Edition: form.Edition,
	}

	return s.bookRepository.Create(book)
}
