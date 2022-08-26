package database

import (
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
	_ "log"
)

func Connect() {
	dsn := "root:H@ppy123@tcp(127.0.0.1:3306)/testDB"
	_, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("an issue occured trying to establish a database connection")
	}

	// log.Println(connection)

	
}	