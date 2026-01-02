package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// maskURI hides credentials from logs (e.g. mongodb://user:pass@host -> mongodb://***@host)
func maskURI(uri string) string {
	if idx := strings.Index(uri, "@"); idx != -1 {
		return "mongodb://***@" + uri[idx+1:]
	}
	return uri
}

func DBSetup() *mongo.Client {
	// Determine MongoDB URI from environment or docker-compose credentials
	uri := os.Getenv("MONGO_URI")
	if uri == "" {
		user := os.Getenv("MONGO_INITDB_ROOT_USERNAME")
		pass := os.Getenv("MONGO_INITDB_ROOT_PASSWORD")
		host := os.Getenv("MONGO_HOST")
		if host == "" {
			host = "localhost"
		}
		port := os.Getenv("MONGO_PORT")
		if port == "" {
			port = "27017"
		}
		if user != "" && pass != "" {
			uri = fmt.Sprintf("mongodb://%s:%s@%s:%s", user, pass, host, port)
		} else {
			uri = fmt.Sprintf("mongodb://%s:%s", host, port)
		}
	}

	// Initialize MongoDB client
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))

	if err != nil {
		panic(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		panic(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		// Provide a clearer message when authentication is required or fails
		if strings.Contains(err.Error(), "requires authentication") || strings.Contains(err.Error(), "Authentication failed") {
			log.Fatalf("MongoDB connection failed: %v. Authentication is required. Set MONGO_URI (e.g. mongodb://user:pass@host:port) or set MONGO_INITDB_ROOT_USERNAME, MONGO_INITDB_ROOT_PASSWORD and MONGO_HOST/MONGO_PORT environment variables and restart the app.", err)
		}
		panic(err)
	}

	log.Printf("Connected to MongoDB at %s", maskURI(uri))
	return client
}

var Client *mongo.Client = DBSetup()

func UserData(client *mongo.Client, collectionName string) *mongo.Collection {
	collection := client.Database("ecommerce").Collection(collectionName)
	return collection
}

func ProductData(client *mongo.Client, collectionName string) *mongo.Collection {
	collection := client.Database("ecommerce").Collection(collectionName)
	return collection
}
