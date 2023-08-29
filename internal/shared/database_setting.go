package shared

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SetConnection(config *DatabaseConfig) (db *gorm.DB, err error) {
	connectString := fmt.Sprintf("host=%s port=%s user=%s dbname=%s  password=%s sslmode=disable", config.Connection.Host, config.Connection.Port, config.Connection.User, config.Connection.Database, config.Connection.Password)
	db, err = gorm.Open(postgres.Open(connectString), &gorm.Config{})
	return
}
