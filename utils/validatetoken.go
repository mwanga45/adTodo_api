package utils

import (
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

func ValidatetTk(tokenString string) (*JwtClaims, error) {
	token, err := jwt.ParseWithClaims(
		tokenString,
		&JwtClaims{},
		func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method")
			}
			return secretKey, nil
		},
	)

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	claims, ok := token.Claims.(*JwtClaims)
	if !ok {
		return nil, fmt.Errorf("invalid token claims")
	}

	return claims, nil
}
