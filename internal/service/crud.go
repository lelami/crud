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
func GetRecipes(index, size int) (*domain.ResponseRecipes, error) {
	return recipes.GetRecipes(index, size)
}
func Get(id string) (*domain.Recipe, error) {
	return recipes.Get(id)
}
func Count() domain.CountRecipes {
	return domain.CountRecipes{
		Count: recipes.Count(),
	}
}
func Delete(id string) error {
	return recipes.Delete(id)
}

func AddOrUpd(r *domain.Recipe) error {

	if r.ID == "" {
		r.ID = uuid.New().String()
	}

	return recipes.Set(r.ID, r)
}
