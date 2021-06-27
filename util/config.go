package util

import (
	"log"
	"os"
)

var AppConfig Config

type Config struct {
	SentryDSN  string
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	AppPort    string
}

func init() {
	AppConfig = Config{
		SentryDSN:  os.Getenv("SENTRY_DSN"),
		DBHost:     os.Getenv("DB_HOST"),
		DBPort:     os.Getenv("DB_PORT"),
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBName:     os.Getenv("DB_NAME"),
	}
	if os.Getenv("PORT") != "" {
		AppConfig.AppPort = os.Getenv("PORT")
	} else {
		AppConfig.AppPort = "8080"
	}

	log.Println("Initialized configs")
}
