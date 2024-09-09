package domain

const OwnerRoleAdmin = "admin"

type Ing struct {
	Amount int    `json:"amount"`
	Type   string `json:"type"`
}

type Recipe struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Ingredients []Ing  `json:"ingredients"`
	Temperature int    `json:"temperature"`
	CreatedBy   string `json:"created_by"`
}

type RecipeOwner struct {
	OwnerId string `json:"owner_id"`
}
