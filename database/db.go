package database

import (
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
	_ "log"
	"github.com/gymcode/project_recipe_backend/model"
)

var DB *gorm.DB

func Connect() {
	dsn := "root:H@ppy123@tcp(127.0.0.1:3306)/testDB"
	db_connection, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("an issue occured trying to establish a database connection")
	}

	DB = db_connection

	db_connection.AutoMigrate(&model.User{})
	// log.Println(connection)

}	