package service

import (
	"errors"
	"todo_api_backend/model"
	"todo_api_backend/repository"

	"golang.org/x/crypto/bcrypt"
)

func Createuser(user *model.User) error {
	if user.Username == "" || user.Password == "" {
	  return 	errors.New("Please username and password required")
	}
	return repository.CreateUser(user)

}
func Userlogin(user *model.User) error {
	if user.Password == "" || user.Username == "" {
		return errors.New("Please username and password required")
	}
	result, err := repository.Userlogin(user.Username)

	if err != nil {
		return errors.New("wrong password or username")
	}
	hashederror := bcrypt.CompareHashAndPassword([]byte(result.Password), []byte(user.Password))
	if hashederror != nil {
		return errors.New("wrong password or username")
	}

	return nil
}
