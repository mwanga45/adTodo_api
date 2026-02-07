package utils

import (
	"errors"

	"github.com/golang-jwt/jwt/v5"
)

func ValidatetTk(tokenString string) (*JwtClaims, error) {
	token, err := jwt.ParseWithClaims(
		tokenString,
		&JwtClaims{},
		func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.New("unexpected signing method")
			}
			return secretKey, nil
		},
	)

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	claims, ok := token.Claims.(*JwtClaims)
	if !ok {
		return nil, errors.New("invalid token claims")
	}

	return claims, nil
}
