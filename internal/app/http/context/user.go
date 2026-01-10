package context

import (
	"github.com/gin-gonic/gin"
	"github.com/hmailyan/go_ecommerce/internal/app/http/middleware"
)

func GetUserID(c *gin.Context) (string, bool) {
	id, ok := c.Get(middleware.ContextUserIDKey)
	if !ok {
		return "", false
	}
	userID, ok := id.(string)
	return userID, ok
}
