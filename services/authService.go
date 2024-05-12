package authServices

import (
	"coffeshop/database"
	"coffeshop/helpers"
	"coffeshop/models"
	userRepository "coffeshop/repositories/UserRepository"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"net/mail"

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

func isEmailValid(email string) bool {
	_, err := mail.ParseAddress(email)

	return err == nil
}

func isEligiblePassword(password string, passwordConfirmation string) (bool, error) {
	if password != passwordConfirmation {
		return false, errors.New("passwordConfirmation not match")
	}

	return true, nil
}

func Register(ctx *gin.Context) (models.User, error) {
	var request userRepository.UserRegistrationRequestDTO
	if err := ctx.BindJSON(&request); err != nil {
		return models.User{}, err
	}

	if !isEmailValid(request.Email) {
		return models.User{}, errors.New("email invalid")
	}

	retrievedUser, err := userRepository.FindUserByEmail(request.Email)

	if err != nil {
		return models.User{}, err
	}

	if retrievedUser != (models.User{}) {
		return retrievedUser, errors.New("please use other credential")
	}

	_, err = isEligiblePassword(request.Password, request.PasswordConfirmation)

	if err != nil {
		return models.User{}, err
	}

	hiddenPassword, err := hidePassword(request.Password)

	if err != nil {
		return models.User{}, err
	}

	request.Password = hiddenPassword

	result, err := userRepository.InsertUser(request)

	if err != nil {
		return models.User{}, err
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

	if !isEmailValid(user.Email) {
		fmt.Println("error email")
		return "", errors.New("email invalid")
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

	token, _ := helpers.GenerateAllTokens(retrievedUser.Email, retrievedUser.Id)

	return token, nil
}
