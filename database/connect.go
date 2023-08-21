package database

import (
	"fmt"
	"os"

	"github.com/heshan-g/go-api/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var dbErr error

func Connect() {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		os.Getenv("PG_HOST"),
		os.Getenv("PG_USER"),
		os.Getenv("PG_PASSWORD"),
		os.Getenv("PG_DBNAME"),
		os.Getenv("PG_PORT"),
		os.Getenv("PG_SSLMODE"),
	)

	DB, dbErr = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if dbErr != nil {
		panic("Error connecting to database")
	}

	DB.AutoMigrate(&model.User{})
}
