package routes

import (
	"coffeshop/controllers"
	"coffeshop/database"
	"coffeshop/middlewares"

	"github.com/gin-gonic/gin"
)

func CoinRoutes(r *gin.Engine) {
	r.GET("/coin/:userId", middlewares.AuthenticateUser(), controllers.GetMyCoin())
	r.GET("/coin", middlewares.AuthenticateUser(), controllers.GetAllCoins())
	r.POST("/coin/track-coin", middlewares.AuthenticateUser(), controllers.TrackCoin())
	r.DELETE("/coin/untrack-coin", middlewares.AuthenticateUser(), controllers.UntrackCoin(database.Client))
}
