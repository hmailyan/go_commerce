package routes

import "github.com/gin-gonic/gin"

func RegisterProductRoutes(rg *gin.RouterGroup, deps *Dependencies) {
	productsGroup := rg.Group("/products")

	productsGroup.POST("/", deps.ProductHandler.Create())
	productsGroup.POST("/:id", deps.ProductHandler.CreateVariation())
	productsGroup.GET("/", deps.ProductHandler.List())
	productsGroup.GET("/:id", deps.ProductHandler.GetByID())
	productsGroup.GET("/search", deps.ProductHandler.SearchByQuery())
}
