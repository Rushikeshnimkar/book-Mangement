package models

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := "host=localhost port=5432 user=gitsplit dbname=gitsplit password=gitsplit sslmode=disable"
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		panic("failed to connect")

	}
	err = database.AutoMigrate(&Book{})
	if err != nil {
		return
	}
	DB = database

}
