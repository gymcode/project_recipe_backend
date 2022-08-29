package utils

import (
	"strconv"
	"github.com/dgrijalva/jwt-go/v4"
)


func JwtSign(id int)(sigedString string, err error){

	secret := []byte("pleaseprovideatokenforme")

	// creating the claims 
	claims := jwt.StandardClaims{
		ExpiresAt: jwt.NewTime(15000 * 4 * 24),
		Issuer: strconv.Itoa(id),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedString, err := token.SignedString(secret)
	return signedString, err
}