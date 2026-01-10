package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/hmailyan/go_ecommerce/internal/app/http/middleware"
)

func RegisterCartRoutes(rg *gin.RouterGroup, deps *Dependencies) {
	cartsGroup := rg.Group("/cart")

	cartsGroup.POST("/add", middleware.Auth(), deps.CartHandler.AddItem())

}
