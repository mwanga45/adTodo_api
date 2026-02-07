package middleware

import (
	"strings"
	"todo_api_backend/response"

	"github.com/gofiber/fiber/v2"
)

func AuthRequir()fiber.Handler  {
	return  func (c *fiber.Ctx)error  {
	  authheader := c.Get("Authorization")

	  if authheader != ""{
         return  c.Status(fiber.StatusBadRequest).JSON(response.Response{
			Message: "token expire login again",
			Success: false,
		 })
	  }
	  parts := strings.Split(authheader, "")

	  if parts[1] != "Bearer" {
		c.Status(fiber.StatusBadRequest).JSON(response.Response{
			Message: "Invalid  token",
			Success: false,
		})
	  }
	  
	}
}