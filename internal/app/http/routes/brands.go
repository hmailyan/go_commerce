package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/hmailyan/go_ecommerce/internal/app/http/middleware"
)

func RegisterBrandRoutes(rg *gin.RouterGroup, deps *Dependencies) {
	cartsGroup := rg.Group("/brand")
	cartsGroup.Use(middleware.Auth())
	cartsGroup.POST("/", deps.CartHandler.AddItem())
	// cartsGroup.GET("/", deps.CartHandler.GetCart())
	// cartsGroup.PUT("/remove", deps.CartHandler.RemoveItem())
	// cartsGroup.POST("/clear", deps.CartHandler.Clear())

}
