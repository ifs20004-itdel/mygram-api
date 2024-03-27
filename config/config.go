package config

import (
	"fmt"
	"mygramapi/models"

	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	root     = goDotEnvVariable("DB_USER")
	password = goDotEnvVariable("DB_PASSWORD")
	dbName   = goDotEnvVariable("DB_NAME")
	DB       *gorm.DB
)

func goDotEnvVariable(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	return os.Getenv(key)
}

func DBInit() {
	DbUrl := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local", root, password, dbName)
	db, err := gorm.Open(mysql.Open(DbUrl), &gorm.Config{})
	if err != nil {
		log.Fatal("DB connection error")
	}

	DB = db

	db.AutoMigrate(&models.User{}, &models.Photo{}, &models.Comment{}, &models.SocialMedia{})
	fmt.Println("DB connected!!!")
}

func GetDB() *gorm.DB {
	return DB
}
