package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/hmailyan/go_ecommerce/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func AddAddress() gin.HandlerFunc {
	return func(c *gin.Context) {
		userQueryID := c.Query("id")

		if userQueryID == "" {
			c.JSON(http.StatusNotFound, gin.H{"error": "User ID is required"})
			c.Abort()
			return
		}

		_, err := primitive.ObjectIDFromHex(userQueryID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Server Error"})
			c.Abort()
			return
		}

		var address models.Address

		address.Address_ID = primitive.NewObjectID()

		if err := c.BindJSON(&address); err != nil {
			c.IndentedJSON(http.StatusNotAcceptable, err.Error())
		}

		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		// filter := bson.D{primitive.E{Key: "_id", Value: user_id}}
		// update := bson.D{{"$push", bson.D{primitive.E{Key: "addresses", Value: addresses}}}}
		match_filter := bson.D{{Key: "$match", Value: bson.D{primitive.E{Key: "_id", Value: address}}}}
		unwind := bson.D{{Key: "$unwind", Value: bson.D{primitive.E{Key: "path", Value: "$addresses"}}}}
		grouping := bson.D{{Key: "$group", Value: bson.D{primitive.E{Key: "_id", Value: "$address_id"}, {Key: "count", Value: bson.D{{Key: "$sum", Value: 1}}}}}}

		pointCursor, err := UserCollection.Aggregate(ctx, mongo.Pipeline{match_filter, unwind, grouping})

		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, "Wrong Information")
			return
		}

		var addressCount []bson.M
		if err = pointCursor.All(ctx, &addressCount); err != nil {
			panic(err)
		}

		var size int32
		for _, address_no := range addressCount {
			size = address_no["count"].(int32)
		}
		if size < 2 {
			filter := bson.D{primitive.E{Key: "_id", Value: address}}
			update := bson.D{{"$push", bson.D{primitive.E{Key: "addresses", Value: address}}}}

			_, err = UserCollection.UpdateOne(ctx, filter, update)
			if err != nil {
				panic(err)
			}
			c.IndentedJSON(http.StatusOK, "Successfully added address")
		} else {
			c.IndentedJSON(http.StatusBadRequest, "You can add maximum 2 addresses")
		}
		defer cancel()
		ctx.Done()
	}
}

func EditHomeAddress() gin.HandlerFunc {
	return func(c *gin.Context) {
		userQueryID := c.Query("id")
		if userQueryID == "" {
			c.JSON(http.StatusNotFound, gin.H{"error": "User ID is required"})
			c.Abort()
			return
		}

		user_id, err := primitive.ObjectIDFromHex(userQueryID)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Server Error"})
		}

		var editAddress models.Address
		if err := c.BindJSON(&editAddress); err != nil {
			c.IndentedJSON(http.StatusNotAcceptable, err.Error())
			return
		}
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		filter := bson.D{primitive.E{Key: "_id", Value: user_id}}
		update := bson.D{{Key: "$set", Value: bson.D{primitive.E{Key: "addresses.0.house_name", Value: editAddress.House}, primitive.E{Key: "addresses.0.street_name", Value: editAddress.Street}, primitive.E{Key: "addresses.0.city_name", Value: editAddress.City}, primitive.E{Key: "addresses.0.pin_code", Value: editAddress.Pincode}}}}

		_, err = UserCollection.UpdateOne(ctx, filter, update)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, "Wrong Information")
			return
		}
		defer cancel()
		ctx.Done()
		c.IndentedJSON(http.StatusAccepted, "successfuly")
	}
}

func EditWorkAddress() gin.HandlerFunc {
	return func(c *gin.Context) {
		userQueryID := c.Query("id")
		if userQueryID == "" {
			c.JSON(http.StatusNotFound, gin.H{"error": "User ID is required"})
			c.Abort()
			return
		}

		user_id, err := primitive.ObjectIDFromHex(userQueryID)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Server Error"})
		}

		var editAddress models.Address
		if err := c.BindJSON(&editAddress); err != nil {
			c.IndentedJSON(http.StatusNotAcceptable, err.Error())
			return
		}
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		filter := bson.D{primitive.E{Key: "_id", Value: user_id}}
		update := bson.D{{Key: "$set", Value: bson.D{primitive.E{Key: "addresses.1.house_name", Value: editAddress.House}, primitive.E{Key: "addresses.1.street_name", Value: editAddress.Street}, primitive.E{Key: "addresses.1.city_name", Value: editAddress.City}, primitive.E{Key: "addresses.1.pin_code", Value: editAddress.Pincode}}}}

		_, err = UserCollection.UpdateOne(ctx, filter, update)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, "Wrong Information")
			return
		}
		defer cancel()
		ctx.Done()
		c.IndentedJSON(http.StatusAccepted, "successfuly")
	}
}

func DeleteAddress() gin.HandlerFunc {
	return func(c *gin.Context) {
		userQueryID := c.Query("id")

		if userQueryID == "" {
			c.JSON(http.StatusNotFound, gin.H{"error": "User ID is required"})
			c.Abort()
			return
		}
		addresses := make([]models.Address, 0)
		user_id, err := primitive.ObjectIDFromHex(userQueryID)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Server Error"})
			c.Abort()
			return
		}

		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		filter := bson.D{primitive.E{Key: "_id", Value: user_id}}
		update := bson.D{{"$set", bson.D{primitive.E{Key: "addresses", Value: addresses}}}}

		_, err = UserCollection.UpdateOne(ctx, filter, update)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, "Wrong Information")
			return
		}
		defer cancel()
		ctx.Done()
		c.IndentedJSON(http.StatusAccepted, "Success")
	}
}
