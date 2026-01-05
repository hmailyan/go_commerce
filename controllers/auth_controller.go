// filepath: /home/hrant/Desktop/go_commerce/controllers/auth.go
package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hmailyan/go_ecommerce/database"
	"github.com/hmailyan/go_ecommerce/models"
	"github.com/hmailyan/go_ecommerce/services"
)

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
