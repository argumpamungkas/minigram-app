package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// MYSQL
var (
	db  *gorm.DB
	err error
)

func ConnectDatabase() {
	// LOAD FILE .ENV
	err = godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// GET DB
	host := os.Getenv("DB_HOST")
	username := os.Getenv("DB_USERNAME")
	dbName := os.Getenv("DB_NAME")

	// MYSQL config
	dsn := fmt.Sprintf("%s:@tcp(%s)/%s?parseTime=true&loc=Local", username, host, dbName)

	// open dsn with gorm
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	// if there is an error opening the connection, handle it
	if err != nil {
		panic(err)
	}

	// err = db.Ping
	// if err != nil {
	// 	panic(err)
	// }

	fmt.Println("Success connect to DB")
}

func GetDb() *gorm.DB {
	return db
}
