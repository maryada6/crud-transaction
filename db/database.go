package db

import (
	"fmt"
	"log"
	"sync"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"crud-transaction/config"
	"crud-transaction/models"
)

var (
	DB   *gorm.DB
	once sync.Once
)

func InitDB() {
	once.Do(func() {
		dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
			config.GetDatabaseHost(),
			config.GetDatabaseUser(),
			config.GetDatabasePassword(),
			config.GetDatabaseName(),
			config.GetDatabasePort())

		var err error
		DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Fatalf("Failed to connect to database: %v", err)
		}

		err = DB.AutoMigrate(&models.Transaction{})
		if err != nil {
			log.Fatalf("Failed to apply database migrations: %v", err)
		}
		fmt.Println("Database connected successfully and migrations applied")
	})
}
