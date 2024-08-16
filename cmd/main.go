package main

import (
	_ "crud/docs"
	"crud/internal/app"
)

// @title Сервис CRUD Recipes
// @version 0.1.0
// @description Сервис для чтения, записи и удаления рецептов
// @host      localhost:8080
// @contact.name API Support

// @Accept   json
// @Produce  json
func main() {
	app.Run()
}
