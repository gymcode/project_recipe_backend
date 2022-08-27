package controller

import (
	"log"
	"github.com/gofiber/fiber/v2"
	"github.com/gymcode/project_recipe_backend/database"
	"github.com/gymcode/project_recipe_backend/model"
	"golang.org/x/crypto/bcrypt"
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

	//password hashing
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 15)
	log.Println("hashed password:: ", hashedPassword) 

	// new user
	userInput := model.User{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Password: string(hashedPassword),
	}

	// insert user  
	results := database.DB.Create(&userInput)

	// checking if it was inserted
	if results.RowsAffected < 0 {
		panic("Could not insert into the database")
	}

	log.Println("user input here :: ", userInput)
	return c.JSON(fiber.Map{
		"code": "00",
		"message": "insertion was suceesfull",
		"data": userInput,
	})
}
