package routes

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gymcode/project_recipe_backend/controller"
)

var BaseUrl string = "/api/v1" 


func Routes(app *fiber.App) {
	/*
	Routes for authentication 
	*/
	app.Get("/", controller.SendMessage)
	app.Post(fmt.Sprintf("%s/register", BaseUrl), controller.Register)
	app.Post(fmt.Sprintf("%s/login", BaseUrl), controller.Login)
	app.Post(fmt.Sprintf("%s/confirm-otp", BaseUrl), controller.ConfirmOtp)
	app.Get(fmt.Sprintf("%s/user", BaseUrl), controller.User)
	app.Get(fmt.Sprintf("%s/signout", BaseUrl), controller.SignOut)
	app.Post(fmt.Sprintf("%s/resend/:msisdn",  BaseUrl), controller.ResendOtp)

	/*
	Routes for Recipe Crud
	*/
	app.Get(fmt.Sprintf("%s/add", BaseUrl), controller.AddRecipe)

	/*
	Routes for Ingredients Crud
	*/


	/*
	Routes for Intructions Crud
	*/

}
