package database

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	// Get Environment
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Failed to get environment")
	}
	host := os.Getenv("HOST")
	username := os.Getenv("USERNAME")
	dbname := os.Getenv("DBNAME")
	port := os.Getenv("PORT")
	password := os.Getenv("PASSWORD")

	// Connect to DB
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, username, password, dbname, port)
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to DB!")
	}
}
