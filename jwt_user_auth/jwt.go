package jwtauth

import (
	"fmt"
	// "log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	// "github.com/joho/godotenv"
)


var (
	secret *[]byte
	apiKey *string
)

func loadEnvVariables () {
	// err := godotenv.Load()
	// if err != nil {
	// 	log.Fatal("Error loading .env file")
	// }
	temp_secret := []byte(os.Getenv("JWT_SECRET_KEY"))
	secret = &temp_secret
	temp_apiKey := os.Getenv("JWT_API_KEY")
	apiKey = &temp_apiKey
}

func GenerateJWTTokenString(userID string) (string, error) {
	/*
	HS256 (HMAC with SHA-256) is a symmetric keyed hashing algorithm that uses one secret key. 
	Symmetric means two parties share the secret key. 
	The key is used for both generating the signature and validating it.
	*/
	loadEnvVariables()

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["exp"] = time.Now().Add(time.Hour).Unix()
	claims["id"] = userID

	tokenStr, err := token.SignedString(*secret)

	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}

	return tokenStr, nil
}

func JWTAuthenticateMiddleware() gin.HandlerFunc {
	loadEnvVariables()

	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "not authorized, missing token"})
			c.Abort()
			return
		}

		token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
			_, ok := t.Method.(*jwt.SigningMethodHMAC)
			if !ok {
				return nil, fmt.Errorf("invalid token")
			}
			return []byte(*secret), nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		claims, _ := token.Claims.(jwt.MapClaims)
		userID := claims["id"].(string)
        c.Set("userID", userID)
		
		c.Next()
	}
}

func GetJwtToken(c *gin.Context) {
	loadEnvVariables()

	userID := c.Param("userID")

	accessKey := c.GetHeader("Access")
	if accessKey != *apiKey {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "not authorized"})
		return
	}

	token, err := GenerateJWTTokenString(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create token"})
		return
	}

	c.String(http.StatusOK, token)
}
