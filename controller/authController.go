package controller

import (
	"database/sql"
	"log"
	"time"

	"github.com/dgrijalva/jwt-go/v4"
	"github.com/gofiber/fiber/v2"
	"github.com/gymcode/project_recipe_backend/database"
	"github.com/gymcode/project_recipe_backend/model"
	utils "github.com/gymcode/project_recipe_backend/utils"
	"golang.org/x/crypto/bcrypt"
)

func SendMessage(c *fiber.Ctx) error {
	return c.SendString("Hello, World 👋!")
}

func Register(c *fiber.Ctx) error {
	user := new(model.User)

	err := c.BodyParser(user)
	log.Println("user model with the details :: ", user)

	if err != nil {
		log.Println(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(model.WrapFailureResponse{
			Code:    "01",
			Message: err.Error(),
			Error:   false,
		})
	}

	//password hashing
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 15)

	// msisdn validation base on countrycode
	msisdn := utils.CountryValidation(user.Msisdn, user.IsoCode)

	// check if the number alredy exists in the database before storing it
	var userData model.User
	database.DB.Where("msisdn = @msisdn", sql.Named("msisdn", msisdn)).Find(&userData)

	if userData.Msisdn == msisdn {
		return c.Status(fiber.StatusNotFound).JSON(model.WrapFailureResponse{
			Code:    "01",
			Message: "phone number already exists. please use a different number",
			Error:   true,
		})
	}

	// new user
	userInput := model.User{
		FirstName:  user.FirstName,
		OtherNames: user.OtherNames,
		Msisdn:     msisdn,
		Email:      user.Email,
		Password:   string(hashedPassword),
		CreatedAt:  time.Now().String(),
		Country:    user.Country,
		IsoCode:    user.IsoCode,
	}

	// insert user
	results := database.DB.Create(&userInput)

	// checking if it was inserted
	if results.RowsAffected < 0 {
		panic("Could not insert into the database")
	}

	// generate otp and send
	otp := utils.GenerateOtp(6)
	log.Println("otp generated for user :: %s", otp)

	hashedOtp, _ := bcrypt.GenerateFromPassword([]byte(otp), 15)

	otpInput := model.OTP{
		Msisdn:    msisdn,
		HashedOtp: string(hashedOtp),
		CreatedAt: time.Now().String(),
		UpdatedAt: time.Now().String(),
	}

	otpInputResponse := database.DB.Create(&otpInput)

	// checking if it was inserted
	if otpInputResponse.RowsAffected < 0 {
		panic("Could not insert into the database")
	}

	// send the otp to the user 


	log.Println("user input here :: ", userInput)
	return c.JSON(fiber.Map{
		"code":    "00",
		"message": "You have signed up successfully",
		"data":    userInput,
	})
}

func ResendOtp(c *fiber.Ctx)error {
	// get the msisdn from the parameter
	msisdn := c.Params("msisdn")
	if msisdn == "" {
		return c.Status(fiber.StatusBadRequest).JSON(model.WrapFailureResponse{
			Code:    "01",
			Message: "Msisdn must not be empty",
			Error:   true,
		})
	}

	// get the user by the msisdn 
	var userOtp model.OTP

	dbResponse := database.DB.Delete(&userOtp, "msisdn = ?", msisdn)

	if dbResponse.RowsAffected < 0 {
		panic("Could not insert into the database")
	}

	// generate otp and send
	otp := utils.GenerateOtp(6)
	log.Println("otp generated for user :: %s", otp)

	hashedOtp, _ := bcrypt.GenerateFromPassword([]byte(otp), 15)

	otpInput := model.OTP{
		Msisdn:    msisdn,
		HashedOtp: string(hashedOtp),
		CreatedAt: time.Now().String(),
		UpdatedAt: time.Now().String(),
	}

	otpInputResponse := database.DB.Create(&otpInput)

	// checking if it was inserted
	if otpInputResponse.RowsAffected < 0 {
		panic("Could not insert into the database")
	}


	// delete the existing otp for that user in the database 
	// generate a new otp 
	// hash the new otp 
	// store the new otp generated in the otp table 
	// send the otp to the msisdn
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
		return c.Status(fiber.StatusNotFound).JSON(model.WrapFailureResponse{
			Code:    "01",
			Message: "User with email does not exist",
			Error:   false,
		})
	}

	// compare passwords
	results := bcrypt.CompareHashAndPassword([]byte(userData.Password), []byte(user.Password))

	if results != nil {
		return c.Status(fiber.StatusForbidden).JSON(model.WrapFailureResponse{
			Code:    "01",
			Message: "Passwords do not match. Please try again",
			Error:   false,
		})
	}

	//  jwt signing user
	token, err := utils.JwtSign(userData.ID)

	// before creating cookie
	// check if there's already a cookie
	existingCookie := c.Cookies("token")
	if len(existingCookie) != 0 {
		return c.Status(fiber.StatusOK).JSON(
			model.WrapFailureResponse{
				Code:    "01",
				Message: "User is already logged in.",
				Error:   true,
			})
	}

	// create cookies
	cookie := new(fiber.Cookie)
	cookie.Name = "token"
	cookie.Value = token
	cookie.Expires = time.Now().Add(time.Hour * 24)
	cookie.HTTPOnly = true

	// set Cookie
	c.Cookie(cookie)

	if err != nil {
		panic("there was an issue with the jwt token")
	}

	log.Printf("jwt token %s", token)
	// log user in successfully
	return c.Status(fiber.StatusOK).JSON(
		model.WrapSuccessResponse{
			Code:    "00",
			Message: "User logged in successfully",
			Error:   false,
			Data:    userData,
		})
}

func User(c *fiber.Ctx) error {
	// get the cookie
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
				Error:   true,
			})
	}

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

	return c.Status(fiber.StatusOK).JSON(
		model.WrapSuccessResponse{
			Code:    "00",
			Message: "request received successfully",
			Error:   false,
			Data:    user,
		})
}

func SignOut(c *fiber.Ctx) error {
	// clearing the cookie
	c.Cookie(&fiber.Cookie{
		Name:     "token",
		Expires:  time.Now().Add(-(time.Hour * 2)),
		HTTPOnly: true,
		SameSite: "lax",
	})

	return c.Status(fiber.StatusOK).JSON(model.WrapSuccessResponse{
		Code:    "00",
		Message: "User logged out successfully",
		Error:   false,
	})
}

// func ConfirmOtp(c *fiber.Ctx) error {

// }
