package controller

import (
	"todo_api_backend/model"
	"todo_api_backend/response"
	"todo_api_backend/service"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)


func Createuser(c *fiber.Ctx)error{
	var user model.User

	if err := c.BodyParser(&user);err != nil{
		return c.Status(fiber.StatusBadRequest).JSON(response.Response{
		  Message: err.Error(),
		  Success: false,
		})
	}
	if user.Username == ""|| user.Password == "" {
		return c.Status(fiber.StatusBadRequest).JSON(response.Response{
			Message: "password and username required",
			Success: false,
		})
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil{
		return  c.Status(fiber.StatusInternalServerError).JSON(response.Response{
			Message: err.Error(),
			Success: false,
		})
	}
	user.Password = string(hashed)

    	if err := service.Createuser(&user);err != nil{
     return  c.Status(fiber.StatusInternalServerError).JSON(response.Response{
		Message: err.Error(),
		Success: false,
	 }) 
	} 
	return  c.Status(fiber.StatusCreated).JSON(response.Response{
		Message: "Success create new user",
		Success: false,
	})
	

}