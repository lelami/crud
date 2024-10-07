package domain

// store
type RecipeIng struct {
	ProductID string `json:"id" bson:"_id"`
	Amount    int    `json:"amount" bson:"amount"`
	Type      string `json:"type" bson:"type"`
}

type Recipe struct {
	ID          string      `json:"id"`
	Name        string      `json:"name" bson:"name"`
	Ingredients []RecipeIng `json:"ingredients" bson:"ingredients"`
	Temperature int         `json:"temperature" bson:"temperature"`
}

type Product struct {
	ID           string `json:"id"`
	Name         string `json:"name" bson:"name"`
	Category     string `json:"category" bson:"category"`
	Calories     int    `json:"calories" bson:"calories"`
	Protein      int    `json:"protein" bson:"protein"`
	Fat          int    `json:"fat" bson:"fat"`
	Carbohydrate int    `json:"carbohydrate" bson:"carbohydrate"`
}

// view

type RecipeIngView struct {
	ProductID    string `json:"id" bson:"_id"`
	Amount       int    `json:"amount" bson:"amount"`
	Type         string `json:"type" bson:"type"`
	Name         string `json:"name" bson:"name"`
	Calories     int    `json:"calories" bson:"calories"`
	Protein      int    `json:"protein" bson:"protein"`
	Fat          int    `json:"fat" bson:"fat"`
	Carbohydrate int    `json:"carbohydrate" bson:"carbohydrate"`
}

type RecipeView struct {
	ID          string          `json:"id" bson:"id"`
	Name        string          `json:"name" bson:"name"`
	Ingredients []RecipeIngView `json:"ingredients" bson:"ingredients"`
	Temperature int             `json:"temperature" bson:"temperature"`
}
