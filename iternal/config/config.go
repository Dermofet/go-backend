package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var Config *config

func init() {
	Config, err := loadConfig()
	if err != nil {
		log.Fatal("load config error: ", err)
		os.Exit(2)
	} else {
		log.Print(".env file was parsed. Config:", Config)
	}
}

type config struct {
	DB_HOST     string
	DB_PORT     int
	DB_USERNAME string
	DB_PASSWORD string
	DB_NAME     string
	DB_DSN      string
}

func loadConfig() (*config, error) {
	err := godotenv.Load(".env")
	if err != nil {
		return nil, err
	}

	port, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		return nil, err
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		os.Getenv("DB_HOST"), os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"), port)

	return &config{
		DB_HOST:     os.Getenv("DB_HOST"),
		DB_PORT:     port,
		DB_USERNAME: os.Getenv("DB_USERNAME"),
		DB_PASSWORD: os.Getenv("DB_PASSWORD"),
		DB_NAME:     os.Getenv("DB_NAME"),
		DB_DSN:      dsn,
	}, nil
}
