package models

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func ConnectDatabase() {
	LoadEnv()

	dsn := fmt.Sprintf(
		"host=%s user=%s  dbname=%s port=%s sslmode=disable",
		os.Getenv("HOST"),
		os.Getenv("USER"),
		os.Getenv("DB_NAME"),
		os.Getenv("PORT"),
	)
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Error while connecting to database", err)
		return
	}
	DB = database
	err = DB.AutoMigrate(&Person{})
	if err != nil {
		fmt.Println("Error During Migration:", err)
		return
	}
	fmt.Println("Successfully Connected to database !")
}
