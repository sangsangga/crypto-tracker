package authServices

import (
	"coffeshop/database"
	"coffeshop/helpers"
	"coffeshop/models"
	userRepository "coffeshop/repositories/UserRepository"
	"database/sql"
	"errors"
	"log"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func hidePassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		log.Panic(err)
		return "", err
	}

	return string(bytes), nil

}

func validatePassword(userPassword string, inputPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(userPassword), []byte(inputPassword))

	return err == nil
}

func Register(ctx *gin.Context) (models.User, error) {
	var user models.User
	if err := ctx.BindJSON(&user); err != nil {
		return user, err
	}

	retrievedUser, err := userRepository.FindUserByEmail(user.Email)

	if err != nil {
		return user, err
	}

	if retrievedUser != (models.User{}) {
		return retrievedUser, errors.New("please use other credential")
	}

	hiddenPassword, err := hidePassword(user.Password)

	if err != nil {
		return user, err
	}

	user.Password = hiddenPassword

	result, err := userRepository.InsertUser(user)

	if err != nil {
		return user, err
	}

	return result, nil
}

func Login(ctx *gin.Context) (string, error) {
	var user models.User
	var err error
	db := database.Client

	if err = ctx.BindJSON(&user); err != nil {
		return "", err
	}

	retrievedUser := models.User{}

	row := db.QueryRow("SELECT * FROM USERS WHERE email=?", user.Email)

	if err = row.Scan(&retrievedUser.Id, &retrievedUser.Email, &retrievedUser.Password); err == sql.ErrNoRows {
		return "", err
	}

	isValidPassword := validatePassword(retrievedUser.Password, user.Password)
	if !isValidPassword {
		return "", errors.New("invalid credential")
	}

	token, _ := helpers.GenerateAllTokens(retrievedUser.Email)

	return token, nil
}
