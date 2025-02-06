package db

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"crud-transaction/config"
	"crud-transaction/models"
)

var DB *gorm.DB

func InitDB() {
	if DB == nil {
		dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
			config.GetDatabaseHost(),
			config.GetDatabaseUser(),
			config.GetDatabasePassword(),
			config.GetDatabaseName(),
			config.GetDatabasePort())
		log.Println("Connecting to database")
		var err error
		DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Panic("Failed to connect to database")
		}

		err = DB.AutoMigrate(&models.Transaction{})
		if err != nil {
			log.Panic("Failed to apply database migrations")
		}
		fmt.Println("Database connected successfully and migrations applied")
	}
}
