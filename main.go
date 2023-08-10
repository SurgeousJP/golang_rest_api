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
)

func init(){
	ctx = context.TODO()

	
	// You should load the DB_CONNECTION_STRING, SERVER_GROUP, PORT from your .env environment,
	// set it to fit your usage, you can check it from previous commits for more information

	// Load environment variables from the .env file (in local)
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
	if port == "" {
		port = "9090"
	}

	basePath := server.Group(serverGroup)

	bookController.RegisterBookRoutes(basePath)
	log.Fatal(server.Run(":" + port))
}