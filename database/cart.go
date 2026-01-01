package database

import (
	"context"
	"time"

	"errors"

	"github.com/hmailyan/go_ecommerce/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type CartData struct {
	Collection *mongo.Collection
}

var (
	ErrCantFindProduct    = errors.New("cannot find the product")
	ErrCantDecodeProduct  = errors.New("cannot decode the product")
	ErrUserIdIsNotValid   = errors.New("user id is not valid")
	ErrCantUpdateUser     = errors.New("cannot update user cart")
	ErrCantRemoveItemCart = errors.New("cannot remove item from cart")
	ErrCantGetItem        = errors.New("cannot get cart items")
	ErrCantByCartItem     = errors.New("cannot buy cart item")
)

// AddToCart adds a product to the user's cart.
func AddProductToCart(ctx context.Context, prodCollection, userCollection *mongo.Collection, productID primitive.ObjectID, userID string) error {
	searchFromDb, err := prodCollection.FindOne(ctx, bson.M{"_id": productID})
	if err != nil {
		return ErrCantFindProduct
	}

	var productCart []models.ProductUser

	err = searchFromDb.All(ctx, &productCart)
	if err != nil {
		return ErrCantDecodeProduct
	}

	id, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return ErrUserIdIsNotValid
	}

	filter := bson.D{primitive.E{Key: "_id", Value: id}}
	update := bson.D{
		{Key: "$push", Value: bson.D{
			{Key: "usercart", Value: bson.D{
				{Key: "$each", Value: productCart},
			}},
		}},
	}

	_, err = userCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		return ErrCantUpdateUser
	}
	return nil
}

// RemoveFromCart removes a product from the user's cart.
func RemoveCartItem(ctx context.Context, prodCollection, userCollection *mongo.Collection, productID primitive.ObjectID, userID string) error {
	user_id, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return ErrUserIdIsNotValid
	}
	filter := bson.D{primitive.E{Key: "_id", Value: user_id}}
	update := bson.D{{Key: "$pull", Value: bson.D{primitive.E{Key: "usercart", Value: bson.M{"_id": productID}}}}}

	_, err = userCollection.UpdateMany(ctx, filter, update)
	if err != nil {
		return ErrCantRemoveItemCart
	}
	return nil
}

// ClearCart clears all products from the user's cart.
func ClearCart(ctx context.Context, user_id string, userCollection *mongo.Collection) error {

	userCartEmpty := make([]models.ProductUser, 0)
	filter := bson.D{primitive.E{Key: "_id", Value: user_id}}
	update := bson.M{"$set": bson.M{"usercart": userCartEmpty}}

	_, err := userCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	return nil
}

func ByItemsFromCart(ctx context.Context, userCollection *mongo.Collection, userID string) error {
	user_id, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return ErrUserIdIsNotValid
	}

	var getCartItems models.User
	var orderCart models.Order

	orderCart.Order_ID = primitive.NewObjectID()
	orderCart.Ordered_At = time.Now()
	orderCart.Order_Cart = make([]models.ProductUser, 0)
	orderCart.Payment_Method = "cod"

	unwind := bson.D{{Key: "$unwind", Value: bson.D{primitive.E{Key: "path", Value: "$usercart"}}}}
	grouping := bson.D{{Key: "$group", Value: bson.D{primitive.E{Key: "_id", Value: "$_id"}, {Key: "total", Value: bson.D{primitive.E{Key: "$sum", Value: "$usercart.price"}}}}}}
	currentResult, err := userCollection.Aggregate(ctx, mongo.Pipeline{unwind, grouping})
	ctx.Done()
	if err != nil {
		return ErrCantGetItem
	}

	var getUserCart []bson.M
	if err = currentResult.All(ctx, &getUserCart); err != nil {
		panic(err)
	}

	var totalPrice int32

	for _, user_item := range getUserCart {
		totalPrice = user_item["total"].(int32)
	}
	orderCart.Price = int(totalPrice)

	filter := bson.D{primitive.E{Key: "_id", Value: user_id}}
	update := bson.D{{Key: "$push", Value: bson.D{primitive.E{Key: "orders", Value: orderCart}}}}

	_, err = userCollection.UpdateMany(ctx, filter, update)
	if err != nil {
		return ErrCantByCartItem
	}

	err = userCollection.FindOne(ctx, bson.D{primitive.E{Key: "_id", Value: user_id}}).Decode(&getCartItems)
	if err != nil {
		return ErrCantGetItem
	}

	filter2 := bson.D{primitive.E{Key: "_id", Value: user_id}}
	update2 := bson.M{"$push": bson.M{"orders.$[].order_list": bson.M{"$each": getCartItems.UserCart}}}

	_, err = userCollection.UpdateOne(ctx, filter2, update2)
	if err != nil {
		panic(err)
	}

	err = ClearCart(ctx, userID, userCollection)
	if err != nil {
		panic(err)
	}
	return nil
}

func InstantBuy(ctx context.Context, prodCollection, userCollection *mongo.Collection, productID primitive.ObjectID, userID string) error {
	user_id, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return ErrUserIdIsNotValid
	}

	var productDetails models.ProductUser
	var orderDetails models.Order

	orderDetails.Order_ID = primitive.NewObjectID()
	orderDetails.Ordered_At = time.Now()
	orderDetails.Order_Cart = make([]models.ProductUser, 0)
	orderDetails.Payment_Method = "cod"

	err = prodCollection.FindOne(ctx, bson.D{primitive.E{Key: "_id", Value: productID}}).Decode(&productDetails)
	if err != nil {
		return ErrCantFindProduct
	}

	orderDetails.Price = int(productDetails.Price)

	filter := bson.D{primitive.E{Key: "_id", Value: user_id}}
	update := bson.D{{Key: "$push", Value: bson.D{primitive.E{Key: "orders", Value: orderDetails}}}}

	_, err = userCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		return ErrCantByCartItem
	}

	filter2 := bson.D{primitive.E{Key: "_id", Value: user_id}}
	update2 := bson.M{"$push": bson.M{"orders.$[].order_list": productDetails}}

	_, err = userCollection.UpdateOne(ctx, filter2, update2)
	if err != nil {
		panic(err)
	}
	return nil
}
