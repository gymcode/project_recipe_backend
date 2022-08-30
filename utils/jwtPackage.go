package utils

import (
	"strconv"
	"github.com/dgrijalva/jwt-go/v4"
)

var Secret []byte = []byte("pleaseprovideatokenforme")

func JwtSign(id int)(sigedString string, err error){

	// creating the claims 
	claims := jwt.StandardClaims{
		ExpiresAt: jwt.NewTime(1500),
		Issuer: strconv.Itoa(int(id)),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedString, err := token.SignedString(Secret)
	return signedString, err
}