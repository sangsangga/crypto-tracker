package main

import (
	"coffeshop/database"
	"coffeshop/routes"
	"fmt"
	"os"
)

func main() {
	database.StartDatabase()

	r := routes.SetupRouter()

	port := os.Getenv("PORT")
	if port == "" {
		fmt.Println(">------")
		port = "8080"
	}

	r.Run(":" + port)
}
