package main

import (
	"context"
	"fmt"
	"golang_rest_api/controllers"
	"golang_rest_api/services"
	"log"

	"os"
	// "github.com/joho/godotenv"

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
	serverGroup string
	port string
)

func init(){
	ctx = context.TODO()

	// Load environment variables from the .env file
	// err := godotenv.Load()
	// if err != nil {
	// 	log.Fatal("Error loading .env file")
	// }

	// Retrieve the connection string from the environment
	connectionString := os.Getenv("DB_CONNECTION_STRING")

	mongoConn := options.	
	Client().
	ApplyURI(connectionString)

	mongoClient, err := mongo.Connect(ctx, mongoConn)
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
	server.SetTrustedProxies(nil)
}

func main() {
	defer mongoClient.Disconnect(ctx)
	serverGroup := os.Getenv("SERVER_GROUP")
	port := os.Getenv("PORT")
	basePath := server.Group(serverGroup)
	bookController.RegisterBookRoutes(basePath)
	log.Fatal(server.Run(":" + port))
}