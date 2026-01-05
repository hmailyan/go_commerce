package routes

import (
	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(rg *gin.RouterGroup, deps *Dependencies) {
	users := rg.Group("/users")

	users.POST("/signup", deps.UserHandler.SignUp())
	users.GET("/me", deps.AuthMiddleware, deps.UserHandler.Me()) // add middleware for authentication
}
