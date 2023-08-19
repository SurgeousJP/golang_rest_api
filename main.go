package main

import (
	"context"
	"golang_rest_api/cache"
	"golang_rest_api/controllers"
	"golang_rest_api/services"
	"log"
	"os"

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
	userService    services.UserService
	bookRedisCache cache.BookCache
	bookController controllers.BookController
	userController controllers.UserController
	ctx            context.Context
	bookCollection *mongo.Collection
	userCollection *mongo.Collection
	mongoClient    *mongo.Client
)

const (
	NUMBER_OF_SECONDS_IN_ONE_DAY = 86400
	DEFAULT_DATABASE_CODE        = 0
	DEFAULT_PORT                 = "9090"
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

	bookCollection = mongoClient.Database(os.Getenv("DB_NAME")).Collection("books")
	bookService = services.NewBookService(bookCollection, ctx)
	userCollection = mongoClient.Database(os.Getenv("DB_NAME")).Collection("users")
	userService = services.NewUserservice(userCollection, ctx)
	bookRedisCache = cache.NewRedisCache(
		os.Getenv("REDIS_ADDRESS"),
		DEFAULT_DATABASE_CODE,
		NUMBER_OF_SECONDS_IN_ONE_DAY,
		os.Getenv("REDIS_PASSWORD"),
		ctx,
	)

	bookController = controllers.NewBookController(bookService, bookRedisCache)
	userController = controllers.NewUserController(userService)
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
		port = DEFAULT_PORT
	}

	basePath := server.Group(serverGroup)

	bookController.RegisterBookRoutes(basePath)
	userController.RegisterUserRoutes(basePath)
	log.Fatal(server.Run(":" + port))
}
