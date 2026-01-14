package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/hmailyan/go_ecommerce/internal/app/http/middleware"
)

func RegisterCartRoutes(rg *gin.RouterGroup, deps *Dependencies) {
	cartsGroup := rg.Group("/cart")
	cartsGroup.Use(middleware.Auth())
	cartsGroup.POST("/add", deps.CartHandler.AddItem())
	cartsGroup.GET("/", deps.CartHandler.GetCart())
	cartsGroup.PUT("/remove", deps.CartHandler.RemoveItem())
	cartsGroup.POST("/clear", deps.CartHandler.Clear())

}
