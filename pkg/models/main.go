package models

import (
	"fmt"
	"os"

	_ "github.com/joho/godotenv/autoload"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetDbInfo() (User, Password, Host, Port, Database string) {
	User = os.Getenv("DB_USER")
	Password = os.Getenv("DB_PASSWORD")
	Host = os.Getenv("DB_HOST")
	Port = os.Getenv("DB_PORT")
	Database = os.Getenv("DB_DATABASE")
	return
}

func SetConnection() (db *gorm.DB, err error) {
	User, Password, Host, Port, Database := GetDbInfo()
	connectString := fmt.Sprintf("host=%s port=%s user=%s dbname=%s  password=%s sslmode=disable", Host, Port, User, Database, Password)
	db, err = gorm.Open(postgres.Open(connectString), &gorm.Config{})
	return
}
