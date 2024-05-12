package controllers

import (
	authServices "coffeshop/services"
	"log"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func hidePassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		log.Panic(err)
	}
	return string(bytes)

}

func ValidatePassword(userPassword string, inputPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(userPassword), []byte(inputPassword))

	return err == nil
}

func Register() gin.HandlerFunc {

	return func(ctx *gin.Context) {
		user, err := authServices.Register(ctx)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error() + "1"})
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
				"message": err.Error() + "2",
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
