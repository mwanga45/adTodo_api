package utils

import (

	"github.com/golang-jwt/jwt/v5"
)

func ValidatetTk(token string)error  {
	token, err := jwt.Parse(token,func(t *jwt.Token) (interface{}, error) {
		return 
	})
}