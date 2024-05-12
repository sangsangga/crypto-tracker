package middlewares

import (
	"coffeshop/helpers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CookieTool() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get cookie
		token, err := c.Cookie("token")
		if err != nil {
			// Cookie verification failed
			c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden with no cookie"})
			c.Abort()
		}

		if token == "" {
			c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden with no cookie"})
			c.Abort()
		}

		_, err = helpers.ValidateToken(token)

		if err != nil {
			c.JSON(http.StatusForbidden, gin.H{
				"message": "Not Authorized",
			})
			c.Abort()
		}

		c.Next()
	}
}
