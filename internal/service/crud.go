package service

import (
	"crud/internal/domain"
	"crud/internal/repository/recipedb"
	"github.com/google/uuid"
)

var recipes recipedb.DB

func Init(DB recipedb.DB) {
	recipes = DB
}

func Get(id string) (*domain.Recipe, error) {
	return recipes.GetRecipe(id)
}

func Delete(id string) error {
	return recipes.DeleteRecipe(id)
}

func AddOrUpd(r *domain.Recipe) error {

	if r.ID == "" {
		r.ID = uuid.New().String()
	}

	return recipes.SetRecipe(r.ID, r)
}
