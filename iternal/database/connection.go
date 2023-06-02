package database

import (
	"go-backend/iternal/config"
	"go-backend/iternal/database/models"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Database struct {
	session *gorm.DB
}

var DB Database

func Connect() {
	db, err := gorm.Open(postgres.Open(config.Config.DB_DSN), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
		os.Exit(2)
	}

	log.Println("Connected")
	// db.Logger = logger.Default.LogMode(logger.Info)

	log.Println("Running migrations")
	db.AutoMigrate(&models.User{})

	DB = Database{
		session: db,
	}
}
