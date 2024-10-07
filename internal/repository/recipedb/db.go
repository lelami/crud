package recipedb

import (
	"crud/internal/domain"
)

type DB interface {
	GetRecipe(id string) (*domain.Recipe, error)
	SetRecipe(id string, recipe *domain.Recipe) error
	DeleteRecipe(id string) error
}
