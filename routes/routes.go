package routes

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gymcode/project_recipe_backend/controller"
)

var BaseUrl string = "/api/v1" 



func Routes(app *fiber.App) {
	app.Get("/", controller.SendMessage)
	app.Post(fmt.Sprintf("%s/register", BaseUrl), controller.Register)
}
