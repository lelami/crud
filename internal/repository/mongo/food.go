package mongo

import (
	"context"
	"crud/internal/domain"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const CollProduct = "products"

// GetProduct retrieves a product by its ProductID.
func (c *MClient) GetProduct(id string) (*domain.Product, error) {
	coll := c.client.Database(c.dbname).Collection(CollProduct)

	var recipe domain.Product
	objid, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": objid}
	err := coll.FindOne(context.TODO(), filter).Decode(&recipe)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, fmt.Errorf("product with ProductID %s not found", id)
		}
		return nil, err
	}
	return &recipe, nil
}

// SetProduct inserts or updates a product by its ProductID.
func (c *MClient) SetProduct(id string, recipe *domain.Product) error {
	coll := c.client.Database(c.dbname).Collection(CollProduct)

	objid, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": objid}
	update := bson.M{"$set": recipe}

	opts := options.FindOneAndUpdate().SetUpsert(true)
	res := coll.FindOneAndUpdate(context.TODO(), filter, update, opts)
	if errors.Is(res.Err(), mongo.ErrNoDocuments) {
		return nil
	}
	return nil
}

// DeleteProduct removes a product by its ProductID.
func (c *MClient) DeleteProduct(id string) error {
	coll := c.client.Database(c.dbname).Collection(CollProduct)
	objid, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": objid}
	_, err := coll.DeleteOne(context.TODO(), filter)
	if err != nil {
		return err
	}
	return nil
}
