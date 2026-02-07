package utils

import "os"
import "github.com/golang-jwt/jwt/v5"

var secretKey = []byte(os.Getenv("SECRETE_KEY"))

type JwtClaims struct {
	Username string   `json:"username"`
	Password  string `json:"password"`
	jwt.RegisteredClaims
}
