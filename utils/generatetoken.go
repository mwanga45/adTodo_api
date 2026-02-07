package utils

import (
	"fmt"
	"time"
	"todo_api_backend/model"

	"github.com/golang-jwt/jwt/v5"
)


func GenerateToken( user *model.User)(string,error){
    claims := JwtClaims{
		Username: user.Username,
		Password: user.Password,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt: jwt.NewNumericDate(time.Now()),
			Issuer: "MyApp",

		},
		
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedString,err := token.SignedString(secretKey)
	if err != nil{
		return  "", fmt.Errorf(err.Error())
	}
	return  signedString, nil
}