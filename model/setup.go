package model

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDb() {
	dsn := "apps@tcp(127.0.0.1:3306)/go_api"
	var err error
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}

	database.AutoMigrate(&News{})

	DB = database
}
