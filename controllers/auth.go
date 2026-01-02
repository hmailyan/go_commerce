// filepath: /home/hrant/Desktop/go_commerce/controllers/auth.go
package controllers

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/hmailyan/go_ecommerce/database"
	"github.com/hmailyan/go_ecommerce/models"
	"github.com/hmailyan/go_ecommerce/services"
)

var validate = validator.New()

// SignUp registers a new user. Expects JSON matching models.User (password in plain text).
func SignUp() gin.HandlerFunc {

	return func(c *gin.Context) {
		var user models.User

		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		validationErr := validate.Struct(user)
		if validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
			return
		}

		// Ensure DB is ready
		if database.DB == nil {
			database.SetupGORM()
		}

		var count int64
		if err := database.DB.Model(&models.User{}).Where("email = ?", user.Email).Count(&count).Error; err != nil {
			log.Printf("error occurred while checking for the email: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occurred while checking for the email"})
			return
		}

		if count > 0 {
			c.JSON(http.StatusConflict, gin.H{"error": "this email already exists"})
			return
		}

		if err := database.DB.Model(&models.User{}).Where("phone = ?", user.Phone).Count(&count).Error; err != nil {
			log.Printf("error occurred while checking for the phone number: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occurred while checking for the phone number"})
			return
		}

		if count > 0 {
			c.JSON(http.StatusConflict, gin.H{"error": "this phone number already exists"})
			return
		}

		// Hash password and set meta
		password, _ := services.HashPassword(user.Password)
		user.Password = password
		user.Created_At = time.Now()
		user.Updated_At = time.Now()
		if user.ID == "" {
			user.ID = uuid.NewString()
		}

		token, refreshToken, _ := services.GenerateUserTokens(user)
		user.Token = token
		user.Refreshtoken = refreshToken

		if err := database.DB.Create(&user).Error; err != nil {
			log.Printf("user item was not created: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "user item was not created"})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"message": "user created successfully"})

	}
}

// Login authenticates a user by email and password.
func Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req models.User
		var foundUser models.User
		if err := c.BindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if database.DB == nil {
			database.SetupGORM()
		}

		if err := database.DB.Where("email = ?", req.Email).First(&foundUser).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "email or password is incorrect"})
			return
		}

		PasswordIsValid, msg := services.VerifyPassword(req.Password, foundUser.Password)

		if !PasswordIsValid {
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			fmt.Println(msg)
			return
		}
		token, refreshToken, _ := services.GenerateUserTokens(foundUser)

		services.UpdateAllTokens(token, refreshToken, foundUser.ID)

		c.JSON(http.StatusFound, foundUser)

	}
}
