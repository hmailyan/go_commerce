package controllers

import {
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

}

type Application struct {
	prodCollection *mongo.Collection
	userCollection *mongo.Collection
}

func NewApplication(prodCollection, userCollection *mongo.Collection) *Application {
	return &Application{
		prodCollection: prodCollection,
		userCollection: userCollection,
	}
}

func (app *Application) AddToCart() gin.Handler {
	return func(c *gin.Context) {
		// Add to new func
		productQueryID := c.Query("id")
		if productQueryID == "" {
			_ = c.AbortWithError(http.StatusBadRequest, gin.H{"error": "Product ID is required"})
			return
		}

		userQueryID := c.Query("UserID")
		if userQueryID == "" {
			_ = c.AbortWithError(http.StatusBadRequest, gin.H{"error": "User ID is required"})
			return
		}

		productID, err := primitive.ObjectIDFromHex(productQueryID)
		if err != nil {
			_ = c.AbortWithError(http.StatusBadRequest, gin.H{"error": "Invalid Product ID"})
			return
		}

		userID, err := primitive.ObjectIDFromHex(userQueryID)
		if err != nil {
			_ = c.AbortWithError(http.StatusBadRequest, gin.H{"error": "Invalid User ID"})
		}

		var ctx, cancel = context.WithTimeout(context.Background(), 5.*time.Second)
		defer cancel()

		err = database.AddProductToCart(ctx, app.prodCollection, app.userCollection, productID, userID)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Failed to add product to cart"})
		}
		c.IndentedJSON(http.StatusOK, gin.H{"message": "Product added to cart successfully"})

	}
}

func (app *Application) RemoveFromCart() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Add to new func

		productQueryID := c.Query("id")
		if productQueryID == "" {
			_ = c.AbortWithError(http.StatusBadRequest, gin.H{"error": "Product ID is required"})
			return
		}

		userQueryID := c.Query("UserID")
		if userQueryID == "" {
			_ = c.AbortWithError(http.StatusBadRequest, gin.H{"error": "User ID is required"})
			return
		}

		productID, err := primitive.ObjectIDFromHex(productQueryID)
		if err != nil {
			_ = c.AbortWithError(http.StatusBadRequest, gin.H{"error": "Invalid Product ID"})
			return
		}

		userID, err := primitive.ObjectIDFromHex(userQueryID)
		if err != nil {
			_ = c.AbortWithError(http.StatusBadRequest, gin.H{"error": "Invalid User ID"})
			return
		}

		var ctx, cancel = context.WithTimeout(context.Background(), 5.*time.Second)
		defer cancel()

		err = database.RemoveCartItem(ctx, app.prodCollection, app.userCollection, productID, userID)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Failed to remove product from cart"})
		}
		c.IndentedJSON(http.StatusOK, gin.H{"message": "Product removed from cart successfully"})
	}
}

func (app *Application) BuyFromCart() gin.HandlerFunc {
	return func(c *gin.Context) {
		userQueryID := c.Query("UserID")
		if userQueryID == "" {
			_ = c.AbortWithError(http.StatusBadRequest, gin.H{"error": "User ID is required"})
		}

		var ctx, cancel = context.WithTimeout(context.Background(), 5.*time.Second)
		defer cancel()

		err = database.ByItemFromCart(ctx, app.prodCollection, app.userCollection, userID)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Failed to buy products from cart"})
		}
		c.IndentedJSON(http.StatusOK, gin.H{"message": "Products bought from cart successfully"})
	}
}

func (app *Application) InsendBuy() gin.HandlerFunc {
	return func(c *gin.Context) {
		productQueryID := c.Query("id")
		if productQueryID == "" {
			_ = c.AbortWithError(http.StatusBadRequest, gin.H{"error": "Product ID is required"})
			return
		}

		userQueryID := c.Query("UserID")
		if userQueryID == "" {
			_ = c.AbortWithError(http.StatusBadRequest, gin.H{"error": "User ID is required"})
			return
		}

		productID, err := primitive.ObjectIDFromHex(productQueryID)
		if err != nil {
			_ = c.AbortWithError(http.StatusBadRequest, gin.H{"error": "Invalid Product ID"})
			return
		}

		userID, err := primitive.ObjectIDFromHex(userQueryID)
		if err != nil {
			_ = c.AbortWithError(http.StatusBadRequest, gin.H{"error": "Invalid User ID"})
			return
		}

		var ctx, cancel = context.WithTimeout(context.Background(), 5.*time.Second)
		defer cancel()

		err = database.InstandBuyer(ctx, app.prodCollection, app.userCollection, productID, userID)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Failed to process instant buy"})
		}
		c.IndentedJSON(http.StatusOK, gin.H{"message": "Instant buy processed successfully"})
	}
}