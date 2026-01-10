package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/hmailyan/go_ecommerce/internal/shared/utils"
)

const ContextUserIDKey = "user_id"

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization token required"})
			c.Abort()
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "invalid authorization format",
			})
			return
		}

		token := parts[1]

		userID, err := utils.NewTokenUtils().ValidateToken(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization token"})
			c.Abort()
			return
		}

		c.Set(ContextUserIDKey, userID)

		c.Next()
	}
}
