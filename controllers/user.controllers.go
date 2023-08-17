package controllers

import (
	"golang_rest_api/services"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserService services.UserService
}

func NewUserController(userService services.UserService) UserController {
	return UserController{
		UserService: userService,
	}
}

func (uc *UserController) GetAllUsers(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"success": "access_granted"})
}

func (uc *UserController) RegisterUserRoutes(rg *gin.RouterGroup) {
	bookRoute := rg.Group("/user")

	bookRoute.GET("/getall", )
}
