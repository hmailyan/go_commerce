package controllers

import (
	"context"
	"errors"

	"go_commerce/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// SearchAddress searches addresses by city (case-insensitive regex).
// - ctx: request context
// - coll: mongo collection for addresses
// - query: search term for city
// - limit: if >0, limits the number of returned documents
func AddAddress(ctx context.Context, coll *mongo.Collection, address models.Address) error {
	if coll == nil {
		return errors.New("nil collection")
	}

	_, err := coll.InsertOne(ctx, address)
	return err
}
func EditHomeAddress(ctx context.Context, coll *mongo.Collection, addressID primitive.ObjectID, updatedAddress models.Address) error {
	if coll == nil {
		return errors.New("nil collection")
	}

	filter := bson.M{"_id": addressID}
	update := bson.M{
		"$set": bson.M{
			"house":   updatedAddress.House,
			"street":  updatedAddress.Street,
			"city":    updatedAddress.City,
			"pincode": updatedAddress.Pincode,
		},
	}

	_, err := coll.UpdateOne(ctx, filter, update)
	return err
}
func EditWorkAddress(ctx context.Context, coll *mongo.Collection, addressID primitive.ObjectID, updatedAddress models.Address) error { // Ned to change
	if coll == nil {
		return errors.New("nil collection")
	}

	filter := bson.M{"_id": addressID}
	update := bson.M{
		"$set": bson.M{
			"house":   updatedAddress.House,
			"street":  updatedAddress.Street,
			"city":    updatedAddress.City,
			"pincode": updatedAddress.Pincode,
		},
	}

	_, err := coll.UpdateOne(ctx, filter, update)
	return err
}

func DeleteAddress(ctx context.Context, coll *mongo.Collection, addressID primitive.ObjectID) error {
	if coll == nil {
		return errors.New("nil collection")
	}

	filter := bson.M{"_id": addressID}

	_, err := coll.DeleteOne(ctx, filter)
	return err
}
