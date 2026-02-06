package controller

import (
	"todo_api_backend/model"
	"todo_api_backend/response"
	"todo_api_backend/service"

	"github.com/gofiber/fiber/v2"
)


func CreateTodo(c *fiber.Ctx)error{
	var todo model.Todoinfo

	if err := c.BodyParser(&todo);err !=nil{
		return c.Status(fiber.StatusBadRequest).JSON(response.Response{
		 Message: "Badrequest",
		 Success: false,
		})
	}

    if err := service.CreateTodo(&todo);err != nil{
        return  c.Status(fiber.StatusInternalServerError).JSON(response.Response{
		  Message: err.Error(),
		  Success: false,

		})
	}
	return  c.Status(fiber.StatusCreated).JSON(response.Response{
		Message: "Successfuly add new todo list",
		Success: true,
	})

}