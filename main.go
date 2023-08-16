package main

import (
	"context"
	"fmt"
	"golang_rest_api/cache"
	"golang_rest_api/controllers"
	"golang_rest_api/services"
	"log"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	// "github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	server         *gin.Engine
	bookService    services.BookService
	bookRedisCache cache.BookCache
	bookController controllers.BookController
	ctx            context.Context
	bookCollection *mongo.Collection
	mongoClient    *mongo.Client
)

func init() {
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
	bookRedisCache = cache.NewRedisCache(
		"redis-18440.c295.ap-southeast-1-1.ec2.cloud.redislabs.com:18440",
		0,
		time.Duration(604800*time.Second),
		os.Getenv("REDIS_PASSWORD"),
		ctx,
	)

	bookController = controllers.New(bookService, bookRedisCache)
	server = gin.Default()
	// Set up CORS (Cross-Origin Resource Sharing)
	server.Use(cors.Default())
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
