package main

import (
	"store_backend/database"
	"store_backend/environment"
	"store_backend/handlers"
	"store_backend/repositories"
	"store_backend/server"
)

func main() {
	env := environment.Initialize()
	db := database.Initialize(env)
	repos := repositories.Initialize(db)
	handlers := handlers.Initialize(repos)

	server := server.Initialize(handlers, env)
	server.Start()
}
