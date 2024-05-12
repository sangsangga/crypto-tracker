package routes

import (
	"coffeshop/controllers"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(r *gin.Engine) {
	r.POST("/auth/register", controllers.Register())
	r.POST("/auth/login", controllers.Login())
	r.GET("/auth/logout", controllers.Logout())
}
