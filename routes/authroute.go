package routes

import (
	"todo_api_backend/controller"

	"github.com/gofiber/fiber/v2"
)

func AuthenticationRoute(app *fiber.App)  {
	auth := app.Group("/auth")

	auth.Post("/register", controller.Createuser)

	
}