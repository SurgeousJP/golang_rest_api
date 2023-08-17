package controllers

import (
	"fmt"
	"golang_rest_api/cache"
	"golang_rest_api/constants"
	"golang_rest_api/models"
	"golang_rest_api/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type BookController struct {
	BookService services.BookService
	BookCache   cache.BookCache
}

func NewBookController(bookService services.BookService, bookCache cache.BookCache) BookController {
	return BookController{
		BookService: bookService,
		BookCache:   bookCache,
	}
}

func (bc *BookController) CreateBook(ctx *gin.Context) {
	var book models.Book
	if err := ctx.ShouldBindJSON(&book); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if err := bc.BookService.CreateBook(&book); err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Successful"})

}

func (bc *BookController) CreateBooks(ctx *gin.Context) {
	var books []*models.Book
	if err := ctx.ShouldBindJSON(&books); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if err := bc.BookService.CreateBooks(books); err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Successful"})
}

func (bc *BookController) GetBook(ctx *gin.Context) {
	bookName := ctx.Param("name")

	book, _ := bc.BookCache.GetBook(&bookName)

	if book == nil {
		book, err := bc.BookService.GetBook(&bookName)

		if err != nil {
			ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, book)

		bc.BookCache.SetBook(&bookName, book)
		return
	}
	ctx.JSON(http.StatusOK, book)
}

func (bc *BookController) GetBooksInPage(ctx *gin.Context) {
	pageNumber := ctx.Param("pageNumber")
	pageNumberInt, err := strconv.Atoi(pageNumber)

	if err != nil || int64(pageNumberInt) <= 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Invalid page number"})
		return
	}

	booksPerPage := constants.BOOKS_PER_PAGE

	key := "all_books_page_" + fmt.Sprint(pageNumber)
	books, _ := bc.BookCache.GetAllBooks(&key)

	if books == nil {
		books, err := bc.BookService.GetBooksInPage(int64(pageNumberInt), booksPerPage)

		if err != nil {
			ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, books)

		bc.BookCache.SetAllBooks(&key, books)
		return
	}
	ctx.JSON(http.StatusOK, books)
}

func (bc *BookController) GetAllBooks(ctx *gin.Context) {
	// Try to get all books from the Redis cache
	key := "all_books"
	books, _ := bc.BookCache.GetAllBooks(&key)

	if books == nil {
		// Cache miss, fetch all books from the primary source
		books, err := bc.BookService.GetAllBooks()

		if err != nil {
			ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, books)

		// Cache the fetched books in Redis
		bc.BookCache.SetAllBooks(&key, books)
		return
	}

	ctx.JSON(http.StatusOK, books)
}

func (bc *BookController) UpdateBook(ctx *gin.Context) {
	var book models.Book

	if err := ctx.ShouldBindJSON(&book); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if err := bc.BookService.UpdateBook(&book); err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Successful"})
}

func (bc *BookController) DeleteBook(ctx *gin.Context) {
	bookName := ctx.Param("name")
	err := bc.BookService.DeleteBook(&bookName)

	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Successful"})
}

func (bc *BookController) RegisterBookRoutes(rg *gin.RouterGroup) {
	bookRoute := rg.Group("/book")
	// The URI must be diffent structure from each other !
	bookRoute.POST("/create", bc.CreateBook)

	bookRoute.POST("/createMany", bc.CreateBooks)

	bookRoute.GET("/get/:name", bc.GetBook)

	bookRoute.GET("/getall", bc.GetAllBooks)

	bookRoute.PATCH("/update", bc.UpdateBook)

	bookRoute.DELETE("/delete/:name", bc.DeleteBook)

	bookRoute.GET("/get/page/:pageNumber", bc.GetBooksInPage)
}
