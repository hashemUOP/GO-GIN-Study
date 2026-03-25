package database

import (
	"fmt"
	"log"
	"login-app/keys"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// 1. Use parentheses (), not curly braces {}
// 2. Use = instead of := at the package level
var (
	host     = keys.Host
	user     = keys.User
	password = keys.Password
	dbname   = keys.Dbname
	port     = keys.Port
)

// DB is capitalized so you can use it in other files (like main.go)
var DB *gorm.DB

func ConnectDatabase() {
	var err error

	// Create the connection string (DSN)
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		host, user, password, dbname, port)

	// Open the connection
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	fmt.Println("Database connection successful!")
}
