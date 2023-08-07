package services

import (
	"context"
	"golang_rest_api/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type BookServiceImpl struct {
	bookCollection *mongo.Collection
	contxt context.Context
}

func NewBookService (bookCollection *mongo.Collection, ctx context.Context) BookService {
	return &BookServiceImpl{
		bookCollection: bookCollection,
		contxt: ctx,
	}
}

func (b *BookServiceImpl) CreateBook(book *models.Book) error {
	return nil
}

func (b *BookServiceImpl) GetBook(bookName *string) (*models.Book, error) {
	return nil, nil
}

func (b *BookServiceImpl) GetAllBooks() ([]*models.Book, error) {
	return nil, nil
}

func (b *BookServiceImpl) UpdateBook(book *models.Book) error {
	return nil
}

func (b *BookServiceImpl) DeleteBook(book *models.Book) error {
	return nil
}
