package controllers

import (
	authServices "coffeshop/services"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func Register() gin.HandlerFunc {

	return func(ctx *gin.Context) {
		user, err := authServices.Register(ctx)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			ctx.Abort()
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"data": user,
		})
	}
}

func Login() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		token, err := authServices.Login(ctx)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			ctx.Abort()
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"data": token,
		})

	}
}

func Logout() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		session := sessions.Default(ctx)
		session.Clear()

		ctx.JSON(http.StatusOK, gin.H{"data": "success"})
	}
}
