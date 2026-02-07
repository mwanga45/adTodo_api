package main

import (
	"todo_api_backend/config"
	database "todo_api_backend/db"
	"todo_api_backend/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))
      config.LoadEnv()
	_= database.Databaseconnection()
	
	
     routes.AuthenticationRoute(app)


	app.Listen(":3000")
}
