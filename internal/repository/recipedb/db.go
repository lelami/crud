package recipedb

import (
	"crud/internal/domain"
)

type DB interface {
	Get(id string) (*domain.Recipe, error)
	Set(id string, recipe *domain.Recipe) error
	Count() int
	GetRecipes(index, size int) (*domain.ResponseRecipes, error)
	Delete(id string) error
}
