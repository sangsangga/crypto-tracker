package helpers

import (
	"log"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

type JWTDetails struct {
	Email string
	jwt.StandardClaims
}

func GenerateAllTokens(
	email string,
) (signedToken string, err error) {
	claims := &JWTDetails{
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(12)).Unix(),
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte("secret"))

	if err != nil {
		log.Panic(err)
		return
	}

	return token, err

}

func ValidateToken(signedToken string) (claims *JWTDetails, msg error) {
	token, err := jwt.ParseWithClaims(signedToken, &JWTDetails{},
		func(t *jwt.Token) (interface{}, error) {
			return []byte("secret"), nil
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
