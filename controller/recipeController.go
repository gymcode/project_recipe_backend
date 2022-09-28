package controller

import (
	"database/sql"
	"log"

	"github.com/dgrijalva/jwt-go/v4"
	"github.com/gofiber/fiber/v2"
	"github.com/gymcode/project_recipe_backend/database"
	"github.com/gymcode/project_recipe_backend/model"
	utils "github.com/gymcode/project_recipe_backend/utils"
)

func AddRecipe(c *fiber.Ctx) error {
	// get the cookies
	cookie := c.Cookies("token")

	log.Println(cookie)
	// parsing with claims
	token, _ := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return utils.Secret, nil
	})

	if token == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(
			model.WrapFailureResponse{
				Code:    "01",
				Message: "Unauthorized :: you do not have access to get details",
				Error:   false,
			})
	}

	// get the user

	// convert it to the standard claim
	claim := token.Claims.(*jwt.StandardClaims)

	var user model.User

	database.DB.Where("id = @id", sql.Named("id", claim.Issuer)).Find(&user)

	if user.ID == 0 {
		return c.Status(fiber.StatusNotFound).JSON(model.WrapFailureResponse{
			Code:    "01",
			Message: "User with email does not exist",
			Error:   false,
		})
	}

	return c.JSON(fiber.Map{
		"message": "output",
	})
}
