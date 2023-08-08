package controllers

import (
	"golang_rest_api/models"
	"golang_rest_api/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type BookController struct {
	BookService services.BookService
}

func New(bookService services.BookService) BookController {
	return BookController{
		BookService: bookService,
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

func (bc *BookController) GetBook(ctx *gin.Context) {
	bookName := ctx.Param("name")
	book, err := bc.BookService.GetBook(&bookName)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, book)
}

func (bc *BookController) GetAllBooks(ctx *gin.Context)  {
	books, err := bc.BookService.GetAllBooks()

	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, books)
}

func (bc *BookController) UpdateBook(ctx *gin.Context){
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

func (bc *BookController) DeleteBook(ctx *gin.Context)  {
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
	bookRoute.POST("/create", bc.CreateBook)
	bookRoute.GET("/get/:name", bc.GetBook)
	bookRoute.GET("/getall", bc.GetAllBooks)
	bookRoute.PATCH("/update", bc.UpdateBook)
	bookRoute.DELETE("/delete/:name", bc.DeleteBook)
}