package controller

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gymcode/project_recipe_backend/model"
)

func SendMessage(c *fiber.Ctx) error {
	return c.SendString("Hello, World 👋!")
}

func Register(c *fiber.Ctx) error {
	user := new(model.User)
	log.Println("user model with the details :: ", user)

	err := c.BodyParser(user)
	if err != nil {
		return err
	}

	// new user
	userInput := model.User{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Password: user.Password,
	}
	log.Println("user input here :: ", userInput)

	return c.JSON(user)
}
