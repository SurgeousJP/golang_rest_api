package services

import "golang_rest_api/models"

type BookService interface {
	CreateBook(*models.Book) error
	CreateBooks([]*models.Book) error
	GetBook(*string) (*models.Book, error)
	GetAllBooks() ([]*models.Book, error)
	GetBooksInPage(int64, int64) ([]*models.Book, error) 
	UpdateBook(*models.Book) error
	DeleteBook(*string) error
}