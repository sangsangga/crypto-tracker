package main

import (
	"coffeshop/database"
	"coffeshop/routes"
	"os"
)

func main() {
	database.StartDatabase()

	r := routes.SetupRouter()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r.Run(":" + port)
}
