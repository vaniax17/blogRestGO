package db

import (
	"blogRest/src/models"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() {

	err := godotenv.Load(".env")

	if err != nil {
		panic("Error loading .env file")
	}

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbName := os.Getenv("DB_NAME")
	dbPassword := os.Getenv("DB_PASSWORD")

	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbName, dbPassword)

	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("Connected to database successfully")

	err = DB.AutoMigrate(&models.User{}, &models.Post{}, &models.Comment{})

	if err != nil {
		panic("failed to migrate database")
	}
	fmt.Println("Migrated database successfully")

}

func Close() {

	sqlDB, _ := DB.DB()

	err := sqlDB.Close()
	if err != nil {
		return
	}
	fmt.Println("Closed database successfully")
}
