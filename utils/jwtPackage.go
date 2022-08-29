package utils

import (
	"github.com/dgrijalva/jwt-go/v4"
)


func JwtSign()(sigedString string, err error){

	secret := []byte("pleaseprovideatokenforme")

	// creating the claims 
	claims := jwt.StandardClaims{
		ExpiresAt: jwt.NewTime(15000),
		Issuer: "Haute Cuisine",
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedString, err := token.SignedString(secret)
	return signedString, err
}