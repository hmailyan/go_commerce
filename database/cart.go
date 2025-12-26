package database

import (
	"context"
	"time"

	"github.com/hmailyan/go_ecommerce/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type CartData struct {
	Collection *mongo.Collection
}

var (
	ErrCantFindProduct = errors.New("cannot find the product")
	ErrCantDecodeProduct = errors.New("cannot decode the product")
	ErrUserIdIsNotValid = errors.New("user id is not valid")
	ErrCantUpdateUser = errors.New("cannot update user cart")
	ErrCantRemoveItemCart = errors.New("cannot remove item from cart")
	ErrCantGetItem = errors.New("cannot get cart items")
	ErrCantByCartItem = errors.New("cannot buy cart item")
	
)

func CartData(client *mongo.Client, collectionName string) *CartData
 {
	collection := client.Database("ecommerce").Collection(collectionName)
	return &CartData{
		Collection: collection,
	}
}

// AddToCart adds a product to the user's cart.
func (cd *CartData) AddProductToCart(ctx context.Context, userID string, product models.ProductUser) error {
	filter := bson.M{"_id": userID}
	update := bson.M{
		"$push": bson.M{
			"usercart": product,
		},
		"$set": bson.M{
			"updated_at": time.Now(),
		},
	}

	_, err := cd.Collection.UpdateOne(ctx, filter, update)
	return err
}

// RemoveFromCart removes a product from the user's cart.
func (cd *CartData) RemoveCartItem(ctx context.Context, userID string, productID primitive.ObjectID) error {
	filter := bson.M{"_id": userID}
	update := bson.M{
		"$pull": bson.M{
			"usercart": bson.M{"_id": productID},
		},
		"$set": bson.M{
			"updated_at": time.Now(),
		},
	}

	_, err := cd.Collection.UpdateOne(ctx, filter, update)
	return err
}

// ClearCart clears all products from the user's cart.
// func (cd *CartData) ClearCart(ctx context.Context, userID string) error {
// 	filter := bson.M{"_id": userID}
// 	update := bson.M{
// 		"$set": bson.M{
// 			"usercart":   []models.ProductUser{},
// 			"updated_at": time.Now(),
// 		},
// 	}

// 	_, err := cd.Collection.UpdateOne(ctx, filter, update)
// 	return err
// }

// GetCart retrieves the user's cart.
// func (cd *CartData) GetCart(ctx context.Context, userID string) ([]models.ProductUser, error)
//  {
// 	filter := bson.M{"_id": userID}
// 	var user models.User
// 	err := cd.Collection.FindOne(ctx, filter).Decode(&user)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return user.UserCart, nil
// }

func ByItemFromCart(ctx context.Context, coll *mongo.Collection, userID string, productID primitive.ObjectID) error {
	filter := bson.M{"_id": userID}
	update := bson.M{
		"$pull": bson.M{
			"usercart": bson.M{"_id": productID},
		},
		"$set": bson.M{
			"updated_at": time.Now(),
		},
	}

	_, err := coll.UpdateOne(ctx, filter, update)
	return err
}