package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hmailyan/go_ecommerce/internal/shared/utils"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization token required"})
			c.Abort()
			return
		}

		_, err := utils.NewTokenUtils().ValidateToken(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization token"})
			c.Abort()
			return
		}

		c.Next()
	}
}
