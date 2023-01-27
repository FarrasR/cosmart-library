package repository

import (
	"cosmart-library/database"
	"cosmart-library/entity/model"
)

type BookRepository interface {
	FindOne(id int) (model.Book, error)
	Find(offset int, limit int) ([]model.Book, error)
	Create(book model.Book) (model.Book, error)
}

type bookRepository struct {
	DatabaseInstance database.DatabaseInstance
}

func NewBookRepository(databaseInstance database.DatabaseInstance) BookRepository {
	return &bookRepository{
		DatabaseInstance: databaseInstance,
	}
}

func (r *bookRepository) FindOne(id int) (model.Book, error) {
	var book model.Book

	if result := r.DatabaseInstance.GetConn().First(&book, id); result.Error != nil {
		return model.Book{}, result.Error
	}

	return book, nil
}

func (r *bookRepository) Find(limit int, offset int) ([]model.Book, error) {
	var books []model.Book

	if result := r.DatabaseInstance.GetConn().Offset(offset).Limit(limit).Find(&books); result.Error != nil {
		return nil, result.Error
	}

	return books, nil
}

func (r *bookRepository) Create(book model.Book) (model.Book, error) {
	if result := r.DatabaseInstance.GetConn().Create(&book); result.Error != nil {
		return model.Book{}, result.Error
	}

	return book, nil
}
