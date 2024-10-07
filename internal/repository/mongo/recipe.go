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

const CollRecipe = "recipes"

// GetRecipe retrieves a product by its ProductID.
func (c *MClient) GetRecipe(id string) (*domain.Recipe, error) {
	coll := c.client.Database(c.dbname).Collection(CollRecipe)

	objid, _ := primitive.ObjectIDFromHex(id)
	var recipe domain.Recipe
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

// SetRecipe inserts or updates a product by its ProductID.
func (c *MClient) SetRecipe(id string, recipe *domain.Recipe) error {
	coll := c.client.Database(c.dbname).Collection(CollRecipe)

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

// DeleteRecipe removes a product by its ProductID.
func (c *MClient) DeleteRecipe(id string) error {
	coll := c.client.Database(c.dbname).Collection(CollRecipe)
	objid, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": objid}
	_, err := coll.DeleteOne(context.TODO(), filter)
	if err != nil {
		return err
	}
	return nil
}

// GetRecipe retrieves a product by its ProductID.
func (c *MClient) GetRecipeView(id string) (*domain.RecipeView, error) {
	coll := c.client.Database(c.dbname).Collection("recipeView")

	objid, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": objid}

	var recipe domain.RecipeView
	err := coll.FindOne(context.TODO(), filter).Decode(&recipe)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, fmt.Errorf("product with ProductID %s not found", id)
		}
		return nil, err
	}
	return &recipe, nil
}
