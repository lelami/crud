package owner

import (
	"crud/internal/repository/ownerdb"
)

var owners ownerdb.DB

func Init(DB ownerdb.DB) {
	owners = DB
}

func Get(id string) (string, error) {
	return owners.Get(id)
}

func Delete(id string) error {
	return owners.Delete(id)
}

func AddOrUpd(recipeId string, ownerId string) error {
	return owners.Set(recipeId, ownerId)
}
