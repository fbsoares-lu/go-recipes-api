package database

import (
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connection() {
	var err error
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	name := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")
	password := os.Getenv("DB_PASSWORD")

	dsn := "host=" + host + " user=" + user + " password=" + password + " dbname=" + name + " port=" + port + " sslmode=disable"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to db")
	}
}
