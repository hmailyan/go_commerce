package controllers

import (
	"context"
	"net/http"
	"time"

	"go_commerce/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// SearchAddress searches addresses by city (case-insensitive regex).
// - ctx: request context
// - coll: mongo collection for addresses
// - query: search term for city
// - limit: if >0, limits the number of returned documents
// func AddAddress(ctx context.Context, coll *mongo.Collection, address models.Address) error {
// 	if coll == nil {
// 		return errors.New("nil collection")
// 	}

// 	_, err := coll.InsertOne(ctx, address)
// 	return err
// }
// func EditHomeAddress(ctx context.Context, coll *mongo.Collection, addressID primitive.ObjectID, updatedAddress models.Address) error {
// 	if coll == nil {
// 		return errors.New("nil collection")
// 	}

// 	filter := bson.M{"_id": addressID}
// 	update := bson.M{
// 		"$set": bson.M{
// 			"house":   updatedAddress.House,
// 			"street":  updatedAddress.Street,
// 			"city":    updatedAddress.City,
// 			"pincode": updatedAddress.Pincode,
// 		},
// 	}

// 	_, err := coll.UpdateOne(ctx, filter, update)
// 	return err
// }
// func EditWorkAddress(ctx context.Context, coll *mongo.Collection, addressID primitive.ObjectID, updatedAddress models.Address) error { // Ned to change
// 	if coll == nil {
// 		return errors.New("nil collection")
// 	}

// 	filter := bson.M{"_id": addressID}
// 	update := bson.M{
// 		"$set": bson.M{
// 			"house":   updatedAddress.House,
// 			"street":  updatedAddress.Street,
// 			"city":    updatedAddress.City,
// 			"pincode": updatedAddress.Pincode,
// 		},
// 	}

// 	_, err := coll.UpdateOne(ctx, filter, update)
// 	return err
// }

func DeleteAddress() gin.HandlerFunc {
	return func(c *gin.Context) {
		user_id := c.Query("id")

		if user_id == "" {
			c.JSON(http.StatusNotFound, gin.H{"error": "User ID is required"})
			c.Abort()
			return
		}
		addresses := make([]models.Address, 0)
		user_id, err := primitive.ObjectIDFromHex(user_id)

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
