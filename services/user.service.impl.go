package services

import (
	"context"
	"golang_rest_api/helper"
	"golang_rest_api/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserServiceImpl struct {
	userCollection *mongo.Collection
	contxt         context.Context
}

func NewUserservice(userCollection *mongo.Collection, ctx context.Context) UserService {
	return &UserServiceImpl{
		userCollection: userCollection,
		contxt:         ctx,
	}
}

func (u *UserServiceImpl) SignUp (user *models.User) error {
	user.ID = primitive.NewObjectID()
	user.Password, _ = helper.HashPassword(user.Password)
	_, err := u.userCollection.InsertOne(u.contxt, user)
	
	return err
}

func (u *UserServiceImpl) GetUser(username string) (*models.User, error){
	var user *models.User
	query := bson.D{bson.E{Key: "username", Value: username}}
	err := u.userCollection.FindOne(u.contxt, query).Decode(&user)
	return user, err
}
