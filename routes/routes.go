package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gymcode/project_recipe_backend/controller"
)

func Routes(app *fiber.App) {

	app.Get("/", controller.SendMessage)

}
