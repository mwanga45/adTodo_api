package utils

import (
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

func ValidatetTk(tokenString string)error  {
	parsedtoken, err := jwt.Parse(tokenString,func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok{
			return  nil , fmt.Errorf("unexpected signing method")
		}
		return  secretKey, nil
	})

	if err != nil{
		return  err
	}

	if !parsedtoken.Valid{
		return fmt.Errorf("Invalid token")
	}
	return  nil
}