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

func (s *bookService) GetBooks(limit int, offset int) ([]model.Book, error) {
	return s.BookRepository.Find(limit, offset)
}

func (s *bookService) CreateBook(form form.FormCreateBook) (model.Book, error) {
	book := model.Book{
		Title:   form.Title,
		Author:  form.Author,
		Edition: form.Edition,
	}

	return s.BookRepository.Create(book)
}
