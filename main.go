package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gymcode/project_recipe_backend/database"
	"github.com/gymcode/project_recipe_backend/routes"
	"github.com/gymcode/project_recipe_backend/utils"
)

func main() {

	// load system conffigurations first
	config, err := utils.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	// establishing database conection
	database.Connect(config.Db_Config)

	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	// routes
	routes.Routes(app)

	app.Listen(":8000")
}
