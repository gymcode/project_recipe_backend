package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gymcode/project_recipe_backend/database"
	"github.com/gymcode/project_recipe_backend/routes"
)

func main() {

	// establishing database conection
	database.Connect()

	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	// routes 
	routes.Routes(app)

	app.Listen(":3000")
}
