package controllers

import (
	"golang_rest_api/helper"
	jwtauth "golang_rest_api/jwt_user_auth"

	"golang_rest_api/models"
	"golang_rest_api/services"
	"net/http"

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

func (uc *UserController) SignUp(ctx *gin.Context) {
	var newUser models.User
	if err := ctx.ShouldBindJSON(&newUser); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	err := uc.UserService.SignUp(&newUser)
	
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register user"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
}

func (uc *UserController) Login(ctx *gin.Context) {
	var credentials models.User
        if err := ctx.ShouldBindJSON(&credentials); err != nil {
            ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
            return
        }

        var user *models.User
        user, err := uc.UserService.GetUser(credentials.Username)
        if err != nil {
            ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
            return
        }

        if helper.CheckPassword((*user).Password, credentials.Password) {
            token, err := jwtauth.GenerateJWTTokenString((*user).Username)
            if err != nil {
                ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
                return
            }

            ctx.JSON(http.StatusOK, gin.H{"token": token})
            return
        }

        ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
    
}

func (uc *UserController) GetUser(ctx *gin.Context) {
	username := ctx.Param("username")

	user, err := uc.UserService.GetUser(username)

	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, user)
}


func (uc *UserController) protectedFunction(ctx *gin.Context) {
	userID, _ := ctx.Get("userID")
    ctx.JSON(http.StatusOK, gin.H{"message": "Access granted for user", "user_id": userID})
}

func (uc *UserController) RegisterUserRoutes(rg *gin.RouterGroup) {
	// Remember to use Authentication ?

	userRoute := rg.Group("/user")

	userRoute.GET("/get/:username", uc.GetUser)

	userRoute.GET("/testJWTValidation", jwtauth.JWTAuthenticateMiddleware(), uc.protectedFunction)

	userRoute.POST("/login", uc.Login)

	userRoute.POST("/signup", uc.SignUp)
}
