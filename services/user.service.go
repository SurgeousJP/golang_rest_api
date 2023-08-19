package services

import "golang_rest_api/models"

type UserService interface {
	SignUp(*models.User) error
	GetUser(string) (*models.User, error)
}
