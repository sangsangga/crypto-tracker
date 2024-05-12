package middlewares

import (
	"coffeshop/helpers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthenticateUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		clientToken := ctx.Request.Header.Get("token")

		if clientToken == "" {
			ctx.JSON(http.StatusForbidden, gin.H{
				"message": "Not Authorized",
			})
			ctx.Abort()
			return
		}

		claims, err := helpers.ValidateToken(clientToken)

		if err != nil {
			ctx.JSON(http.StatusForbidden, gin.H{
				"message": "Not Authorized",
			})
			ctx.Abort()
			return
		}

		ctx.Set("email", claims.Email)
		ctx.Next()

	}
}
