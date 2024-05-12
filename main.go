package main

import (
	"coffeshop/database"
	"coffeshop/routes"
)

func main() {
	database.StartDatabase()

	r := routes.SetupRouter()

	// r.GET("/ping", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{"message": "pong"})
	// })

	r.Run()
}
