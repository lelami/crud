package mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MClient struct {
	client *mongo.Client
	dbname string
}

func NewMongoClient(url, dbname string) (*MClient, error) {

	var cl MClient

	client, err := mongo.Connect(context.TODO(), options.Client().
		ApplyURI(url))
	if err != nil {
		return nil, err
	}

	cl.client = client
	cl.dbname = dbname

	/*	// НЕРАБОЧАЯ ВЕРСИЯ!
		// Определение конвейера агрегации для создания представления
		pipeline := mongo.Pipeline{
			{
				{"$lookup", bson.D{
					{"from", "products"},
					{"localField", "ingredients._id"},
					{"foreignField", "id"},
					{"as", "ingredientView"},
				}},
			},
			{
				{"$addFields", bson.D{
					{"ingredients.name", "$ingredientView.name"},
					{"ingredients.calories", "$ingredientView.calories"},
					{"ingredients.protein", "$ingredientView.protein"},
					{"ingredients.fat", "$ingredientView.fat"},
					{"ingredients.carbohydrate", "$ingredientView.carbohydrate"},
				}},
			},
			{
				{"$project", bson.D{
					{"_id", 1},
					{"name", 1},
					{"temperature", 1},
					{"ingredients", 1},
				}},
			},
		}

		// Создание представления
		err = client.Database(dbname).RunCommand(
			context.TODO(),
			bson.D{
				{"create", "recipeView"},
				{"viewOn", "recipes"},
				{"pipeline", pipeline},
			},
		).Err()
		if err != nil {
			log.Fatal(err)
		}*/

	return &cl, nil
}
