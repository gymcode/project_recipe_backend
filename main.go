package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gymcode/project_recipe_backend/database"
	"github.com/gymcode/project_recipe_backend/routes"
)

func main(){

	// establishing database conection
	database.Connect()

	app := fiber.New()

	routes.Routes(app)

    app.Listen(":3000")
}