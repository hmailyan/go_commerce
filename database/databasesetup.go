package database

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DBSetup() *mongo.Client {
	// Initialize MongoDB client
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))

	if err != nil {
		panic(err)
	}

	ctx, cancle := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancle()

	err = client.Connect(ctx)
	if err != nil {
		panic(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		panic(err)
	}

	fmt.Println("Connected to MongoDB!")
	return client
}

var Client *mongo.Client = DBSetup()

func UserData(client *mongo.Client, collectionName string) *UserData {
	// collection := client.Database("ecommerce").Collection(collectionName)
	// return &UserData{
	// 	Collection: collection,
	// }
}

func ProductData(client *mongo.Client, collectionName string) *ProductData {

	// collection := client.Database("ecommerce").Collection(collectionName)
	// return &ProductData{
	// 	Collection: collection,
	// }
}
