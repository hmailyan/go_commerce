package controllers

import (
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"

	"context"

	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/hmailyan/go_ecommerce/database"
	"github.com/hmailyan/go_ecommerce/models"
)

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

func (app *Application) AddToCart() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Add to new func
		productQueryID := c.Query("id")
		if productQueryID == "" {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Product ID is required"})
			return
		}

		userQueryID := c.Query("UserID")
		if userQueryID == "" {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "User ID is required"})
			return
		}

		productID, err := primitive.ObjectIDFromHex(productQueryID)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid Product ID"})
			return
		}

		userID, err := primitive.ObjectIDFromHex(userQueryID)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid User ID"})
			return
		}

		var ctx, cancel = context.WithTimeout(context.Background(), 5.*time.Second)
		defer cancel()

		err = database.AddProductToCart(ctx, app.prodCollection, app.userCollection, productID, userID.Hex())
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
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Product ID is required"})
			return
		}

		userQueryID := c.Query("UserID")
		if userQueryID == "" {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "User ID is required"})
			return
		}

		productID, err := primitive.ObjectIDFromHex(productQueryID)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid Product ID"})
			return
		}

		userID, err := primitive.ObjectIDFromHex(userQueryID)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid User ID"})
			return
		}

		var ctx, cancel = context.WithTimeout(context.Background(), 5.*time.Second)
		defer cancel()

		err = database.RemoveCartItem(ctx, app.prodCollection, app.userCollection, productID, userID.Hex())
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
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "User ID is required"})
			return
		}

		var ctx, cancel = context.WithTimeout(context.Background(), 5.*time.Second)
		defer cancel()

		err := database.ByItemsFromCart(ctx, app.userCollection, userQueryID)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Failed to buy products from cart"})
			return
		}
		c.IndentedJSON(http.StatusOK, gin.H{"message": "Products bought from cart successfully"})
	}
}

func (app *Application) InstantBuy() gin.HandlerFunc {
	return func(c *gin.Context) {
		productQueryID := c.Query("id")
		if productQueryID == "" {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Product ID is required"})
			return
		}

		userQueryID := c.Query("UserID")
		if userQueryID == "" {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "User ID is required"})
			return
		}

		productID, err := primitive.ObjectIDFromHex(productQueryID)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid Product ID"})
			return
		}

		userID, err := primitive.ObjectIDFromHex(userQueryID)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid User ID"})
			return
		}

		var ctx, cancel = context.WithTimeout(context.Background(), 5.*time.Second)
		defer cancel()

		err = database.InstantBuy(ctx, app.prodCollection, app.userCollection, productID, userID.Hex())
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Failed to process instant buy"})
		}
		c.IndentedJSON(http.StatusOK, gin.H{"message": "Instant buy processed successfully"})
	}
}

func GetItemFromCart() gin.HandlerFunc {
	return func(c *gin.Context) {
		userQueryID := c.Query("id")

		if userQueryID == "" {
			c.JSON(http.StatusNotFound, gin.H{"error": "User ID is required"})
			c.Abort()
			return
		}

		user_id, err := primitive.ObjectIDFromHex(userQueryID)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Server Error"})
			c.Abort()
			return
		}

		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		var filledCart models.User
		err = UserCollection.FindOne(ctx, bson.D{primitive.E{Key: "_id", Value: user_id}}).Decode(&filledCart)

		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Not found"})
			c.Abort()
			return
		}

		filler_match := bson.D{{Key: "$match", Value: bson.D{primitive.E{Key: "_id", Value: user_id}}}}
		unwind_match := bson.D{{Key: "$unwind", Value: bson.D{primitive.E{Key: "path", Value: "$usercart"}}}}
		grouping := bson.D{{Key: "$group", Value: bson.D{primitive.E{Key: "_id", Value: "$_id"}, {Key: "total", Value: bson.D{{Key: "$sum", Value: "$usercart.price"}}}}}}

		pointCursor, err := UserCollection.Aggregate(ctx, mongo.Pipeline{filler_match, unwind_match, grouping})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Server Error"})
			return
		}

		var listing []bson.M
		if err = pointCursor.All(ctx, &listing); err != nil {
			log.Fatal(err)
			c.AbortWithStatus(http.StatusInternalServerError)
		}
		for _, json := range listing {
			c.IndentedJSON(http.StatusOK, json["total"])
			c.IndentedJSON(http.StatusOK, filledCart.UserCart)
		}
		ctx.Done()
	}
}
