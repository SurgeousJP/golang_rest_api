package cache

import (
	"golang_rest_api/models"
)

type BookCache interface {
	SetBook(key *string, book *models.Book) error
	GetBook(key *string) (*models.Book, error)
	GetAllBooks(key *string) ([]*models.Book, error)
	SetAllBooks(key *string, books []*models.Book) error
}
