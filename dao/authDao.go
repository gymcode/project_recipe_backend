package dao

import (
	"database/sql"

	"github.com/gymcode/project_recipe_backend/database"
	"github.com/gymcode/project_recipe_backend/model"
	"gorm.io/gorm"
)

// ------------ user DAO -----------------

/**
	method: InsertUser
	desc: inserting a user into the user database
	params: user model	
*/
func InsertUser(user model.User) *gorm.DB {
	resp := database.DB.Create(&user)

	if resp.RowsAffected < 0 {
		panic("Could not insert into the database")
	}

	return resp
}


/**
	method: GetUserByMsisdn
	desc: gettting a user by the phone number
	params: user model and the msisdn 
*/
func GetUserByMsisdn(user model.User, msisdn string){
	database.DB.Where("msisdn = @msisdn", sql.Named("msisdn", msisdn)).Find(&user)
}



// ------------ OTP DAO -----------------

/**
	method: GetUserByMsisdn
	desc: gettting a user by the phone number
	params: user model and the msisdn 
*/
func InsertOtp(otp model.OTP) *gorm.DB{
	resp := database.DB.Create(&otp)

	if resp.RowsAffected < 0 {
		panic("Could not insert into the database")
	}

	return resp

}

/**
	method: GetUserByMsisdn
	desc: gettting a user by the phone number
	params: user model and the msisdn 
*/
func DeleteOtp(user model.User, msisdn string) *gorm.DB{
	return database.DB.Delete(&user, "msisdn = ?", msisdn)
}


