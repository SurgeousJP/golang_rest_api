package main

import (
	"context"
	"fmt"
	"golang_rest_api/controllers"
	"golang_rest_api/services"
	"log"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	server *gin.Engine
	bookService services.BookService
	bookController controllers.BookController
	ctx context.Context
	bookCollection *mongo.Collection
	mongoClient *mongo.Client
	err error
)

func init(){
	ctx = context.TODO()

	mongoConn := options.
	Client().
	ApplyURI("mongodb+srv://baosurgeous:testDatabase@testcluster.dfxfjru.mongodb.net/?retryWrites=true&w=majority")

	mongoClient, err = mongo.Connect(ctx, mongoConn)
	if err != nil {
		log.Fatal(err)
	}
	if err := mongoClient.Ping(ctx, readpref.Primary()); err != nil {
		log.Fatal(err)
	}

	fmt.Println("mongodb connection established")

	bookCollection = mongoClient.Database("TestDB").Collection("books")
	bookService = services.NewBookService(bookCollection, ctx)
	bookController = controllers.New(bookService)
	server = gin.Default()
}

func main() {
	defer mongoClient.Disconnect(ctx)
	basePath := server.Group("/v1")
	bookController.RegisterBookRoutes(basePath)
	log.Fatal(server.Run(":9090"))
}