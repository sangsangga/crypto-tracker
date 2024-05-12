package routes

import (
	"coffeshop/controllers"
	"coffeshop/database"

	"github.com/gin-gonic/gin"
)

func CoinRoutes(r *gin.Engine) {
	// r.Use(middlewares.AuthenticateUser())
	r.GET("/coin/:userId", controllers.GetMyCoin())
	r.GET("/coin", controllers.GetAllCoins())
	r.POST("/coin/track-coin", controllers.TrackCoin())
	r.DELETE("/coin/untrack-coin", controllers.UntrackCoin(database.Client))
}
