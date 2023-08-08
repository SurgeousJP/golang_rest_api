package services

import (
	"context"
	"errors"
	"golang_rest_api/models"

	"go.mongodb.org/mongo-driver/bson"
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
	_, err := b.bookCollection.InsertOne(b.contxt, book)
	return err
}

func (b *BookServiceImpl) GetBook(bookName *string) (*models.Book, error) {
	var book *models.Book
	query := bson.D{bson.E{Key: "name", Value: bookName}}
	err := b.bookCollection.FindOne(b.contxt, query).Decode(&book)
	return book, err
}

func (b *BookServiceImpl) GetAllBooks() ([]*models.Book, error) {
	var books []*models.Book
	cursor, err :=  b.bookCollection.Find(b.contxt, bson.D{{}})
	
	if err != nil {
		return nil, err
	}
	for cursor.Next(b.contxt) {
		var book models.Book
		err := cursor.Decode(&book)
		if err != nil {
			return nil, err
		}
		books = append(books, &book)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}
	cursor.Close(b.contxt)

	if len(books) == 0 {
		return nil, errors.New("documents not found")
	}

	return books, nil
}

func (b *BookServiceImpl) UpdateBook(book *models.Book) error {
	filter := bson.D{bson.E{Key: "name", Value: book.Name}}
	update := bson.D{
		bson.E{Key: "$set",
		Value: bson.D{
			bson.E{Key: "name", Value: book.Name},
			bson.E{Key: "book_img_url", Value: book.ImageURL},
			bson.E{Key: "author", Value: book.Author},
			bson.E{Key: "price", Value: book.Price},
			bson.E{Key: "supplier", Value: book.Supplier},
			bson.E{Key: "publisher", Value: book.Publisher},
			bson.E{Key: "book_layout", Value: book.Layout},
			bson.E{Key: "series", Value: book.Series},
			},
		},
	}
	result, _ := b.bookCollection.UpdateOne(b.contxt, filter, update)
	if result.MatchedCount != 1 {
		return errors.New("no matched document found for update")
	}
	return nil
}

func (b *BookServiceImpl) DeleteBook(bookName *string) error {
	filter := bson.D{bson.E{Key: "name", Value: bookName}}
	result, _ := b.bookCollection.DeleteOne(b.contxt, filter)
	if result.DeletedCount != 1 {
		return errors.New("no matched document found for delete")
	}
	return nil
}
