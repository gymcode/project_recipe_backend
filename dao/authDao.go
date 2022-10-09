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

/**
	method: ActivateUserAccount
	desc: updating the user account by the activated status
	params: msisdn and the user model
*/

func ActivateUserAccount(user model.User, msisdn string) *gorm.DB{
	return database.DB.Model(&user).Where("msisdn = ?", msisdn).Update("activated", true)
}




// ------------ OTP DAO -----------------

/**
	method: InsertOtp
	desc: inserting the otp into the database
	params: OTP model
*/
func InsertOtp(otp model.OTP) *gorm.DB{
	resp := database.DB.Create(&otp)

	if resp.RowsAffected < 0 {
		panic("Could not insert into the database")
	}

	return resp

}

/**
	method: DeleteOtp
	desc: deleting an otp from the database
	params: user model and the msisdn 
*/
func DeleteOtp(otp model.OTP, msisdn string) *gorm.DB{
	return database.DB.Delete(&otp, "msisdn = ?", msisdn)
}

/**
	method: InsertOtp
	desc: inserting the otp into the database
	params: OTP model
*/
func GetOtpByMsisdn(otp model.OTP, msisdn string) *gorm.DB{
	return database.DB.Select("msisdn = @msisdn", sql.Named("msisdn", msisdn)).Find(&otp)
}
