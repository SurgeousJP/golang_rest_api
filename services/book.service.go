package services

import "golang_rest_api/models"

type BookService interface {
	CreateBook(*models.Book) error
	GetBook(*string) (*models.Book, error)
	GetAllBooks() ([]*models.Book, error)
	UpdateBook(*models.Book) error
	DeleteBook(*models.Book) error
}