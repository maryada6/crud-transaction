package db

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"crud-transaction/config"
	"crud-transaction/models"
)

var db *gorm.DB

func GetDB() *gorm.DB {
	if db == nil {
		dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
			config.GetDatabaseHost(),
			config.GetDatabaseUser(),
			config.GetDatabasePassword(),
			config.GetDatabaseName(),
			config.GetDatabasePort())
		log.Println("Connecting to database")
		var err error
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Panic("Failed to connect to database")
		}

		err = db.AutoMigrate(&models.Transaction{})
		if err != nil {
			log.Panic("Failed to apply database migrations")
		}
		fmt.Println("Database connected successfully and migrations applied")
	}
	return db
}
