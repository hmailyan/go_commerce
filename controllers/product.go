// filepath: /home/hrant/Desktop/go_commerce/controllers/product.go
package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/hmailyan/go_ecommerce/database"
	"github.com/hmailyan/go_ecommerce/models"
)

// SearchProduct searches products by name (case-insensitive regex).
// - ctx: request context
// - coll: mongo collection for products
// - query: search term for product_name
// - limit: if >0, limits the number of returned documents
func SearchProduct() gin.HandlerFunc {
	return func(c *gin.Context) {
		if database.DB == nil {
			database.SetupGORM()
		}
		var productList []models.Product
		if err := database.DB.Find(&productList).Error; err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "failed to retrieve products"})
			return
		}
		c.IndentedJSON(http.StatusOK, productList)
	}
}

// SearchProductByQuery performs a general query against the products collection.
// - ctx: request context
// - coll: mongo collection for products
// - q: a bson.M query (pass nil to match all)
// - findOpts: optional *options.FindOptions
func SearchProductByQuery() gin.HandlerFunc {
	return func(c *gin.Context) {
		queryParams := c.Query("name")

		if queryParams == "" {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Query parameter 'name' is required"})
			return
		}

		if database.DB == nil {
			database.SetupGORM()
		}

		var searchProducts []models.Product
		like := "%" + queryParams + "%"
		if err := database.DB.Where("product_name ILIKE ?", like).Find(&searchProducts).Error; err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Error occurred while searching for products"})
			return
		}
		c.IndentedJSON(http.StatusOK, searchProducts)
	}

}

func ProductViewerAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		var products models.Product

		if err := c.BindJSON(&products); err != nil {
			c.JSON(http.StatusNotAcceptable, gin.H{"error": err.Error()})
			return
		}

		// Ensure DB is initialized
		if database.DB == nil {
			database.SetupGORM()
		}

		if products.ID == "" {
			products.ID = uuid.NewString()
		}

		if err := database.DB.Create(&products).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Product was not inserted"})
			return
		}
		c.JSON(http.StatusOK, "Successfully added")
	}

}
