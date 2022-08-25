package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"

	"os"
)

var db *gorm.DB

func ConnectDB() {
	godotenv.Load(".env")
	dbUrl := os.Getenv("DB_URL")
	if dbUrl == "" {
		panic("URL not available")
	}
	d, err := gorm.Open("mysql", dbUrl)
	if err != nil {
		panic(err.Error())
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
