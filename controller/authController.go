package controller

import (
	"database/sql"
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
		Password:  string(hashedPassword),
	}

	// insert user
	results := database.DB.Create(&userInput)

	// checking if it was inserted
	if results.RowsAffected < 0 {
		panic("Could not insert into the database")
	}

	log.Println("user input here :: ", userInput)
	return c.JSON(fiber.Map{
		"code":    "00",
		"message": "You have signed up successfully",
		"data":    userInput,
	})
}

func Login(c *fiber.Ctx) error {
	user := new(model.User)

	err := c.BodyParser(user)
	if err != nil {
		return err
	}

	// take the user input we need
	var userData model.User

	// making the database query
	database.DB.Where("email = @email", sql.Named("email", user.Email)).First(&userData)

	if userData.ID == 0 {
		return c.Status(fiber.StatusNotFound).JSON(model.ApiResponse{
			Code:    "00",
			Message: "User with email does not exist",
			Error:   false,
			Data:    *user,
		})
	}

	// compare passwords
	results := bcrypt.CompareHashAndPassword([]byte(userData.Password), []byte(user.Password))

	if results != nil {
		return c.Status(fiber.StatusForbidden).JSON(model.ApiResponse{
			Code:    "01",
			Message: "Passwords do not match. Please try again",
			Error:   false,
			Data:    *user,
		})
	}

	// log user in successfully
	return c.Status(fiber.StatusOK).JSON(
		model.ApiResponse{
			Code:    "00",
			Message: "User logged in successfully",
			Error:   false,
			Data:    userData,
		})
}
