package cache

import (
	"golang_rest_api/models"
)

type BookCache interface {
	Set(key *string, value *models.Book) error
	Get(key *string) (*models.Book, error)
}
