package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/hmailyan/go_ecommerce/internal/app/http/middleware"
)

func RegisterUserRoutes(rg *gin.RouterGroup, deps *Dependencies) {
	users := rg.Group("/users")

	users.POST("/signup", deps.UserHandler.SignUp())
	users.GET("/verify", deps.UserHandler.VerifyEmail())
	users.POST("/login", deps.UserHandler.Login())
	users.GET("/me", middleware.Auth(), deps.UserHandler.Me()) // add middleware for authentication
}
