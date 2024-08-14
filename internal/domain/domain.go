package domain

type Ing struct {
	Amount int    `json:"amount"`
	Type   string `json:"type"`
}

type Recipe struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Ingredients []Ing  `json:"ingredients"`
	Temperature int    `json:"temperature"`
}

type ResponseRecipes struct {
	Recipes []Recipe `json:"recipes"`
}

type CountRecipes struct {
	Count int `json:"count" validate:"required"`
}
