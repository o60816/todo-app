package main

import "todo-app/models"

func main() {
	models.InitializeDB()
	initializeRoutes()
}
