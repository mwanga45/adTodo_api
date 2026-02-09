package middleware

import (
	"strings"
	"todo_api_backend/response"
	"todo_api_backend/utils"

	"github.com/gofiber/fiber/v2"
)

func AuthRequir() fiber.Handler {
	return func(c *fiber.Ctx) error {
		authheader := c.Get("Authorization")

		if authheader == "" {
			return c.Status(fiber.StatusBadRequest).JSON(response.Response{
				Message: "token expire login again",
				Success: false,
			})
		}
		parts := strings.Split(authheader, " ")

		if parts[0] != "Bearer" {
			c.Status(fiber.StatusBadRequest).JSON(response.Response{
				Message: "Invalid  token",
				Success: false,
			})
		}
        if len(parts) != 2 || parts[1] != ""{
          return  c.Status(fiber.StatusBadRequest).JSON(response.Response{
			Message: "Invalid token",
			Success: false,
		  })
		}
		claims, err := utils.ValidatetTk(parts[1])
		if err != nil {
			return  c.Status(fiber.StatusBadRequest).JSON(response.Response{
				Message: err.Error(),
				Success: false,
			})
		}
		c.Locals(claims.Username)
		c.Locals(claims.Password)
		
		return c.Next()

	}
}
