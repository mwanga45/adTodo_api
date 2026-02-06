package utils

import (
	"time"
	"todo_api_backend/model"
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

func Generatetoken(user *model.User)(string,error){
	// generate an token
	secrete_key := godotenv.Load("SECRETE_KEY")
	claims := jwt.MapClaims{
		"username":user.Username,
		"userId":user.ID,
		"exp":time.Now().Add(time.Hour *24).Unix(),
	}
	method := jwt.SigningMethodES256

	token , err := jwt.NewWithClaims(method,claims).SignedString(secrete_key)

	if err != nil{
		return  "", err
	}
	return  token, nil

}

