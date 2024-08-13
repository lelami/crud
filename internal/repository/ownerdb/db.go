package ownerdb

type DB interface {
	Get(id string) (string, error)
	Set(recipeId string, ownerId string) error
	Delete(id string) error
}
