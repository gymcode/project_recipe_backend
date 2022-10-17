package database

import (
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
	_ "log"
	"github.com/gymcode/project_recipe_backend/model"
)

var DB *gorm.DB

func Connect(db_config string) {
	dsn := db_config
	db_connection, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("an issue occured trying to establish a database connection")
	}

	DB = db_connection

	db_connection.AutoMigrate(
		&model.User{}, 
		&model.OTP{})
	// log.Println(connection)

}	