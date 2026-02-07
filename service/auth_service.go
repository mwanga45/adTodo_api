package service

import (
	"todo_api_backend/model"
	"todo_api_backend/repository"

	"github.com/gofiber/fiber/v2/log"
)

func Createuser(user *model.User)error{
  if user.Username == "" || user.Password == "" {
	log.Fatal("Please username and password required")
  }
  return  repository.CreateUser(user)

}
func Userlogin(user *model.User)error  {
		if user.Password == "" || user.Username == ""{
		  log.Fatal("Please username and password required")
	}
    _ , err := repository.Userlogin(user.Username)

	if err != nil{
		log.Fatalf(err.Error())
	}
	return  nil
}