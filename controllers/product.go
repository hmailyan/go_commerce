// filepath: /home/hrant/Desktop/go_commerce/controllers/product.go
package controllers

import (
	"context"
	"errors"

	"go_commerce/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// SearchProduct searches products by name (case-insensitive regex).
// - ctx: request context
// - coll: mongo collection for products
// - query: search term for product_name
// - limit: if >0, limits the number of returned documents
func SearchProduct(ctx context.Context, coll *mongo.Collection, query string, limit int64) ([]models.Product, error) {
	if coll == nil {
		return nil, errors.New("nil collection")
	}
	if query == "" {
		return nil, errors.New("empty query")
	}

	filter := bson.M{"product_name": primitive.Regex{Pattern: query, Options: "i"}}
	findOpts := options.Find()
	if limit > 0 {
		findOpts.SetLimit(limit)
	}

	cur, err := coll.Find(ctx, filter, findOpts)
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	var products []models.Product
	if err := cur.All(ctx, &products); err != nil {
		return nil, err
	}
	return products, nil
}

// SearchProductByQuery performs a general query against the products collection.
// - ctx: request context
// - coll: mongo collection for products
// - q: a bson.M query (pass nil to match all)
// - findOpts: optional *options.FindOptions
func SearchProductByQuery(ctx context.Context, coll *mongo.Collection, q bson.M, findOpts ...*options.FindOptions) ([]models.Product, error) {
	if coll == nil {
		return nil, errors.New("nil collection")
	}
	if q == nil {
		q = bson.M{}
	}

	opts := options.MergeFindOptions(findOpts...)
	cur, err := coll.Find(ctx, q, opts)
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	var products []models.Product
	if err := cur.All(ctx, &products); err != nil {
		return nil, err
	}
	return products, nil
}