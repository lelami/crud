package mongo

import (
	"crud/internal/domain"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"testing"
)

func TestSetRecipes(t *testing.T) {
	type fields struct {
		client *mongo.Client
		dbname string
	}
	type args struct {
		recipe *domain.Recipe
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "тесто",
			args: args{recipe: &domain.Recipe{
				ID:   primitive.NewObjectID().Hex(),
				Name: "тесто",
				Ingredients: []domain.RecipeIng{{
					ProductID: "66f8465313e18e30fd25dbbb",
					Amount:    100,
				},
					{
						ProductID: "66f8465313e18e30fd25dbbc",
						Amount:    50,
					}},
				Temperature: 30,
			}},
			wantErr: false,
		},
		/*		{
					name: "5",
					args: args{recipe: &domain.Recipe{
						ID: primitive.NewObjectID().Hex(),
					}},
					wantErr: true,
				},
				{
					name: "6",
					args: args{recipe: &domain.Recipe{
						ID: primitive.NewObjectID().Hex(),
					}},
					wantErr: false,
				},*/
	}

	userDB, err := NewMongoClient("mongodb://admin:admin@localhost:27017/", "product")
	if err != nil {
		log.Fatalf("ERROR failed to initialize product database: %v", err)
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := userDB.SetRecipe(tt.args.recipe.ID, tt.args.recipe); (err != nil) != tt.wantErr {
				t.Errorf("SetUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}

}

func TestSetProduct(t *testing.T) {
	type fields struct {
		client *mongo.Client
		dbname string
	}
	type args struct {
		product *domain.Product
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "Water",
			args: args{product: &domain.Product{
				ID:           primitive.NewObjectID().Hex(),
				Name:         "Вода",
				Category:     "",
				Calories:     0,
				Protein:      0,
				Fat:          0,
				Carbohydrate: 0,
			}},
			wantErr: false,
		},
		{
			name: "Мука",
			args: args{product: &domain.Product{
				ID:           primitive.NewObjectID().Hex(),
				Name:         "Мука",
				Category:     "",
				Calories:     100,
				Protein:      0,
				Fat:          0,
				Carbohydrate: 10,
			}},
			wantErr: false,
		},
	}

	userDB, err := NewMongoClient("mongodb://admin:admin@localhost:27017/", "product")
	if err != nil {
		log.Fatalf("ERROR failed to initialize product database: %v", err)
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := userDB.SetProduct(tt.args.product.ID, tt.args.product); (err != nil) != tt.wantErr {
				t.Errorf("SetProduct() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestPipeline(t *testing.T) {
	type fields struct {
		client *mongo.Client
		dbname string
	}
	type args struct {
		product *domain.Product
	}
	/*	tests := []struct {
			name    string
			fields  fields
			args    args
			wantErr bool
		}{
		}*/

	db, err := NewMongoClient("mongodb://admin:admin@localhost:27017/", "product")
	if err != nil {
		log.Fatalf("ERROR failed to initialize product database: %v", err)
	}

	view, err := db.GetRecipeView("66f84b59215c7f795694ea91")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(view)
	/*	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := userDB.SetProduct(tt.args.product.ID, tt.args.product); (err != nil) != tt.wantErr {
				t.Errorf("SetUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}*/
}
