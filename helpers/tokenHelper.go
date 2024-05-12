package helpers

import (
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

type JWTDetails struct {
	Email  string
	UserId int64
	jwt.StandardClaims
}

func GenerateAllTokens(
	email string, userId int64,
) (signedToken string, err error) {
	claims := &JWTDetails{
		Email:  email,
		UserId: userId,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(12)).Unix(),
		},
	}

	secret := os.Getenv("SECRET_KEY")
	if secret == "" {
		secret = "secret"
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(secret))

	if err != nil {
		return "", err
	}

	return token, err

}

func ValidateToken(signedToken string) (claims *JWTDetails, msg error) {
	token, err := jwt.ParseWithClaims(signedToken, &JWTDetails{},
		func(t *jwt.Token) (interface{}, error) {
			secret := os.Getenv("SECRET_KEY")
			if secret == "" {
				secret = "secret"
			}
			return []byte(secret), nil
		})

	if err != nil {
		msg = err
		return
	}

	claims, ok := token.Claims.(*JWTDetails)

	if !ok {
		msg = err
		return
	}

	if claims.ExpiresAt < time.Now().Unix() {
		msg = err
		return
	}

	return claims, msg

}
