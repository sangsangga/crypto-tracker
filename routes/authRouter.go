package routes

import (
	"coffeshop/controllers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(r *gin.Engine) {
	r.GET("/test", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "test123"})
	})
	r.POST("/auth/register", controllers.Register())
	r.POST("/auth/login", controllers.Login())
	r.GET("/auth/logout", controllers.Logout())
}
