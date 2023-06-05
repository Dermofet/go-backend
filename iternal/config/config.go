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
	var err error
	Config, err = loadConfig()
	if err != nil {
		log.Fatal("Load config error: ", err)
	} else {
		log.Println(".env file was parsed")
	}
}

type config struct {
	BACKEND_HOST string
	BACKEND_PORT int

	DB_HOST     string
	DB_PORT     int
	DB_USERNAME string
	DB_PASSWORD string
	DB_NAME     string
	DB_DSN      string

	JWT_SECRET_KEY []byte
}

func loadConfig() (*config, error) {
	err := godotenv.Load(".env")
	if err != nil {
		return nil, err
	}

	dbPort, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		return nil, err
	}

	backendPort, err := strconv.Atoi(os.Getenv("BACKEND_PORT"))
	if err != nil {
		return nil, err
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		os.Getenv("DB_HOST"), os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"), dbPort)

	return &config{
		BACKEND_HOST: os.Getenv("BACKEND_HOST"),
		BACKEND_PORT: backendPort,

		DB_HOST:     os.Getenv("DB_HOST"),
		DB_PORT:     dbPort,
		DB_USERNAME: os.Getenv("DB_USERNAME"),
		DB_PASSWORD: os.Getenv("DB_PASSWORD"),
		DB_NAME:     os.Getenv("DB_NAME"),
		DB_DSN:      dsn,

		JWT_SECRET_KEY: []byte(os.Getenv("JWT_SECRET_KEY")),
	}, nil
}
