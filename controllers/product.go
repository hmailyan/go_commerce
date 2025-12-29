// filepath: /home/hrant/Desktop/go_commerce/controllers/product.go
package controllers

import (
	"context"
	"net/http"
	"time"

	"go_commerce/models"

	"github.com/gin-gonic/gin"
	"github.com/hmailyan/go_ecommerce/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var ProductCollection *mongo.Collection = database.ProductData(database.Client, "products")

// SearchProduct searches products by name (case-insensitive regex).
// - ctx: request context
// - coll: mongo collection for products
// - query: search term for product_name
// - limit: if >0, limits the number of returned documents
func SearchProduct() gin.HandlerFunc {
	return func(c *gin.Context) {
		var productList []models.Product

		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		cursor, err := ProductCollection.Find(ctx, bson.D{}).All(&productList)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "failed to retrieve products"})
		}
		err = cursor.All(ctx, &productList)
		if err != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
		}
		defer cursor.Close()

		if err = cursor.Err(); err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "invalid"})
		}
		defer cancel()

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
		var searchProducts []models.Product
		queryParams := c.Query("name")

		if queryParams == "" {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Query parameter 'name' is required"})
			return
		}

		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		searchQueryDB, err := ProductCollection.Find(ctx, bson.M{"product_name": bson.M{"$regex": queryParams}})
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Error occurred while searching for products"})
			return
		}

		err = searchQueryDB.All(ctx, &searchProducts)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Error occurred while searching for products"})
			return
		}
		defer searchQueryDB.Close(ctx)

		if err = searchQueryDB.Err(); err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Error occurred while searching for products"})
			return
		}
		defer cancel()

		c.IndentedJSON(http.StatusOK, searchProducts)
	}

}
